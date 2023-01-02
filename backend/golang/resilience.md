# Go Resiliency

近期看了一个国外老哥写的小项目 [go-resiliency](https://github.com/eapache/go-resiliency)，这个小项目实现了客户端执行*作业*(可以理解为是执行业务逻辑的函数调用)的**弹性处理机制**(超时、信号量、重试、熔断、批量执行等)，从中我学到了很多关于并发相关的编程技巧，这里将其整理成博客分享



## 超时
对作业的执行设定超时时间是非常常见的操作
+ 构造函数接收超时时间
```go
// Deadline implements the deadline/timeout resiliency pattern.
type Deadline struct {
  timeout time.Duration
}

  // New constructs a new Deadline with the given timeout.
func New(timeout time.Duration) *Deadline {
  return &Deadline{
    timeout: timeout,
  }
}
```
+ 提供入参为回调函数(作业)的 `Run` 接口，用于执行待执行的作业
```go
dl := New(1 * time.Second)
err := dl.Run(func(stopper <-chan struct{}) error {
  // do something possibly slow
  // check stopper function and give up if timed out
  return nil
})
switch err {
case ErrTimedOut:
  // execution took too long, oops
default:
  // some other error
}
```
每个作业在执行前都要设定一个定时器，超时的情况需要及时返回接口调用方已超时，未超时的情况下则返回作业的执行结果

这里涉及两个协程:
+ `Deadline` : 触发作业执行以及接收作业执行结果
+ `Work` : 执行作业具体逻辑并将执行结果发送给 `Deadline` 协程

协程间通过 `channel` 进行交流

[Deadline](./drawio/deadline.drawio){link-type="drawio"}

```go
func (d *Deadline) Run(work func(<-chan struct{}) error) error {
  result := make(chan error)
  stopper := make(chan struct{})
  go func() {
    value := work(stopper)
    select {
    case result <- value:
    case <-stopper:
    }
  }()
  select {
  case ret := <-result:
    return ret
  case <-time.After(d.timeout):
    close(stopper)
    return ErrTimedOut
  }
}
```

## 信号量
信号量可以看作是一种计数器，常被用来管理资源的数量，一个被协程持有的信号量可以被任何协程释放，其使用方式和 `lock` 类似: `Require` 和 `Release`

```go
sem := New(3, 1*time.Second)
for i := 0; i < 10; i++ {
  go func() {
    if err := sem.Acquire(); err != nil {
      return //could not acquire semaphore
    }
    defer sem.Release()
    // do something semaphore-guarded
  }()
}
```
::: tip 互斥锁与信号量的区别
[与互斥锁相比](https://www.zhihu.com/question/47704079)，**互斥锁在信号量的基础上增加了所有权的概念**。一个被协程锁住的互斥锁只能持锁的协程解锁，简单来说，互斥锁管理的是资源的使用权
:::

信号量的构造函数需要指定数量 `tickets` 和 `Require` 的超时时间，每个 `Semaphore` 实例拥有一个长度为 `tickets` 的 `buffered channel`: `sem`

```go
// Semaphore implements the semaphore resiliency pattern
type Semaphore struct {
  sem     chan struct{}
  timeout time.Duration
}

// New constructs a new Semaphore with the given ticket-count
// and timeout.
func New(tickets int, timeout time.Duration) *Semaphore {
  return &Semaphore{
    sem:     make(chan struct{}, tickets),
    timeout: timeout,
  }
}
```

对于信号量的 `Require` 和 `Release` 操作
+ `Require` : 向 `sem` 发送消息，同时启动定时器监听是否超时，当 `sem` 队列满的时候将会阻塞
```go
// Acquire tries to acquire a ticket from the semaphore. If it can, it returns nil.
// If it cannot after "timeout" amount of time, it returns ErrNoTickets. It is
// safe to call Acquire concurrently on a single Semaphore.
func (s *Semaphore) Acquire() error {
  select {
  case s.sem <- struct{}{}:
    return nil
  case <-time.After(s.timeout):
    return ErrNoTickets
  }
}
```
+ `Release` : 从 `sem` 消费消息
```go
// Release releases an acquired ticket back to the semaphore. It is safe to call
// Release concurrently on a single Semaphore. It is an error to call Release on
// a Semaphore from which you have not first acquired a ticket.
func (s *Semaphore) Release() {
  <-s.sem
}
```

[Semaphore](./drawio/semaphore.drawio){link-type="drawio"}

## 重试
作业可能存在不稳定的情况(如第三方服务偶先异常等)，客户端往往需要重试机制确保作业的顺利执行
```go
r := New([]time.Duration{
  100 * time.Microsecond,
  200 * time.Microsecond,
  400 * time.Microsecond,
}, nil)

err := r.Run(func() error {
  // do some work
  return nil
})

if err != nil {
  // handle the case where the work failed three times
}
```

重试往往需要考虑以下问题:
1. 需要重试多少次
2. 不同重试次数间执行作业的时间间隔如何设置(防止因重试加大作业执行的压力)
3. 作业返回什么错误才需要重试

因此重试机制的构造函数里，第一个参数 `backoff` 用于指定作业的[重试节奏](#重试节奏)，第二个参数 `class` 用于指定作业的[重试触发条件](#重试触发条件)
```go
// New constructs a Retrier with the given backoff pattern and classifier. The length of the backoff pattern
// indicates how many times an action will be retried, and the value at each index indicates the amount of time
// waited before each subsequent retry. The classifier is used to determine which errors should be retried and
// which should cause the retrier to fail fast. The DefaultClassifier is used if nil is passed.
func New(backoff []time.Duration, class Classifier) *Retrier {
  if class == nil {
    class = DefaultClassifier{}
  }
  return &Retrier{
    backoff: backoff,
    class:   class,
    rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
  }
}
```


### 重试节奏
从构造函数可以看出，重试节奏由一个指定时长的数组表示，比如
```go
[]time.Duration{
  10*time.Millisecond,
  20*time.Millisecond,
}
```
就表示在执行第一次重试前等待 10 毫秒，在执行第二次重试前等待 20 毫秒，数组长度即表示最多重试多少次

最简单的重试节奏是每次重试前等待相同的时长
```go
// ConstantBackoff generates a simple back-off strategy of retrying 'n' times, and waiting 'amount' time after each one.
func ConstantBackoff(n int, amount time.Duration) []time.Duration {
  ret := make([]time.Duration, n)
  for i := range ret {
    ret[i] = amount
  }
  return ret
}
```
也可以只指定第一次重试前等待的时长，后续重试时长是前一次重试等待时长的 2 倍([指数级增长](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/))，为了不让等待时长太久，也可以指定最大等待时长
```go
// ExponentialBackoff generates a simple back-off strategy of retrying 'n' times, and doubling the amount of
// time waited after each one.
func ExponentialBackoff(n int, initialAmount time.Duration) []time.Duration {
  ret := make([]time.Duration, n)
  next := initialAmount
  for i := range ret {
    ret[i] = next
    next *= 2
  }
  return ret
}

// LimitedExponentialBackoff generates a simple back-off strategy of retrying 'n' times, and doubling the amount of
// time waited after each one.
// If back-off reaches `limitAmount` , thereafter back-off will be filled with `limitAmount` .
func LimitedExponentialBackoff(n int, initialAmount time.Duration, limitAmount time.Duration) []time.Duration {
  ret := make([]time.Duration, n)
  next := initialAmount
  for i := range ret {
    if next < limitAmount {
      ret[i] = next
      next *= 2
    } else {
      ret[i] = limitAmount
    }
  }
  return ret
}

```

**实际计算作业重试前等待时间，可以适量加一些[抖动因子(`jitter`)](https://aws.amazon.com/builders-library/timeouts-retries-and-backoff-with-jitter/)以防止相同时间对作业的执行产生压力**

```go
func (r *Retrier) calcSleep(i int) time.Duration {
  // lock unsafe rand prng
  r.randMu.Lock()
  defer r.randMu.Unlock()
  // take a random float in the range (-r.jitter, +r.jitter) and multiply it by the base amount
  return r.backoff[i] + time.Duration(((r.rand.Float64()*2)-1)*r.jitter*float64(r.backoff[i]))
}
```


### 重试触发条件
重试的触发条件使用接口 `Classifier` 来定义，我们可以根据业务需要实现接口来自定义具体的触发条件
```go
// Action is the type returned by a Classifier to indicate how the Retrier should proceed.
type Action int

const (
	Succeed Action = iota // Succeed indicates the Retrier should treat this value as a success.
	Fail                  // Fail indicates the Retrier should treat this value as a hard failure and not retry.
	Retry                 // Retry indicates the Retrier should treat this value as a soft failure and retry.
)
// Classifier is the interface implemented by anything that can classify Errors for a Retrier.
type Classifier interface {
  Classify(error) Action
}
```
比如最简单的重试触发条件: 如果作业执行出错就重试
```go
// DefaultClassifier classifies errors in the simplest way possible. If
// the error is nil, it returns Succeed, otherwise it returns Retry.
type DefaultClassifier struct{}

// Classify implements the Classifier interface.
func (c DefaultClassifier) Classify(err error) Action {
  if err == nil {
    return Succeed
  }
  return Retry
}
```
也可以提前指定作业返回错误的白名单列表，如果作业执行出错且错误在列表里，才重试
```go
// WhitelistClassifier classifies errors based on a whitelist. If the error is nil, it
// returns Succeed; if the error is in the whitelist, it returns Retry; otherwise, it returns Fail.
type WhitelistClassifier []error

// Classify implements the Classifier interface.
func (list WhitelistClassifier) Classify(err error) Action {
  if err == nil {
    return Succeed
  }
  for _, pass := range list {
    if errors.Is(err, pass) {
      return Retry
    }
  }
  return Fail
}
```

### 重试流程

结合[节奏](#重试节奏)和[触发条件](#重试触发条件)，利用重试机制执行作业的流程就很清晰了

[Retrier](./drawio/retrier.drawio){link-type="drawio"}

1. 执行作业
1. 利用自定义的 `Classifier` 判断作业的执行结果，如果成功或失败则立即返回作业的执行结果，否则触发重试
1. 每次重试前先判断当前重试次数是否已达上限，如果是则返回作业的执行结果，不是则先计算下一次重试的等待时间并 `sleep`，等待结束后累加重试次数，再回到第 1 步执行作业


```go
// Run executes the given work function by executing RunCtx without context.Context.
func (r *Retrier) Run(work func() error) error {
  return r.RunCtx(context.Background(), func(ctx context.Context) error {
    // never use ctx
    return work()
  })
}

// RunCtx executes the given work function, then classifies its return value based on the classifier used
// to construct the Retrier. If the result is Succeed or Fail, the return value of the work function is
// returned to the caller. If the result is Retry, then Run sleeps according to the its backoff policy
// before retrying. If the total number of retries is exceeded then the return value of the work function
// is returned to the caller regardless.
func (r *Retrier) RunCtx(ctx context.Context, work func(ctx context.Context) error) error {
  retries := 0
  for {
    ret := work(ctx)
    switch r.class.Classify(ret) {
    case Succeed, Fail:
      return ret
    case Retry:
      if retries >= len(r.backoff) {
        return ret
      }
      timeout := time.After(r.calcSleep(retries))
      if err := r.sleep(ctx, timeout); err != nil {
        return err
      }
      retries++
    }
  }
}
```

## 熔断
熔断是一种服务自我保护的机制，尤其是针对一个接口可能调用多个微服务的场景: 如果其中一个被调用的微服务因各种原因超时或返回失败，调用方对微服务的调用会积累更多的时间及资源，进而导致调用方级联地故障。为了避免出现系统雪崩的场景，我们可以使用熔断器，**当探测到作业执行失败达到阈值时，即打开熔断器，使得作业无需被执行就立即返回失败，而调用方可以基于熔断器打开的失败进行兜底操作**


```go
breaker := New(3, 1, 5*time.Second)

for {
  result := breaker.Run(func() error {
    // communicate with some external service and
    // return an error if the communication failed
    return nil
  })

  switch result {
  case nil:
    // success!
  case ErrBreakerOpen:
    // our function wasn't run because the breaker was open
  default:
    // some other error
  }
}
```
![circuit-breaker-open](./FILES/resilience.md/290f3fa6.png)

当熔断器被打开时，返回到调用方的只有 `ErrBreakerOpen` 错误，而其他情况会先执行作业，再基于作业执行结果进行统计和状态转移，毕竟**不能让熔断器一直返回熔断错误，要给它一定的弹性去恢复到原来正常运转的情况**


```go
func (b *Breaker) Run(work func() error) error {
  state := atomic.LoadUint32(&b.state)

  if state == open {
    return ErrBreakerOpen
  }

  return b.doWork(state, work)
}

func (b *Breaker) doWork(state uint32, work func() error) error {
  var panicValue interface{}

  result := func() error {
    defer func() {
      panicValue = recover()
    }()
    return work()
  }()

  if result == nil && panicValue == nil && state == closed {
    // short-circuit the normal, success path without contending
    // on the lock
    return nil
  }

  // oh well, I guess we have to contend on the lock
  b.processResult(result, panicValue)

  if panicValue != nil {
    // as close as Go lets us come to a "rethrow" although unfortunately
    // we lose the original panicing location
    panic(panicValue)
  }

  return result
}
```


这里实现的熔断器比较简单，主要提供三个参数:
+ `errorThreshold` : 作业执行错误数阈值
+ `successThreshold` : 作业执行成功数阈值
+ `timeout` : 熔断器打开后，过多长时间才进入半打开状态

```go
// New constructs a new circuit-breaker that starts closed.
// From closed, the breaker opens if "errorThreshold" errors are seen
// without an error-free period of at least "timeout". From open, the
// breaker half-closes after "timeout". From half-open, the breaker closes
// after "successThreshold" consecutive successes, or opens on a single error.
func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker {
  return &Breaker{
    errorThreshold:   errorThreshold,
    successThreshold: successThreshold,
    timeout:          timeout,
  }
}
```

到这里是不是有点一头雾水，为什么要统计作业执行结果情况？半打开状态又是什么东西？是的，熔断器的概念比较多，后续将以[3 个状态](#3-个状态)、[2 个动作](#2-个动作)、[状态流转](#状态流转)进行讲解



### 3 个状态

[Circuit Breaker](./drawio/circuit-breaker.drawio){link-type="drawio"}

+ `closed` : 关闭状态，此状态会统计作业执行的错误数是否达到阈值，如果是则将切换为 `open` 状态
+ `open` : 打开状态，此状态不会执行作业，快速返回错误；同时会开启一个定时器，当时间一到将状态改为 `halfOpen` 状态
+ `halfOpen` : 半打开状态，此状态下如果作业还出现错误则立马回到 `open` 状态；此外在半打开状态下会统计作业执行的成功数是否达到阈值，如果是则将切换为 `closed` 状态

因此每次切换状态都会把统计数据清零

```go
const (
  closed uint32 = iota
  open
  halfOpen
)
// ...
func (b *Breaker) changeState(newState uint32) {
  b.errors = 0
  b.successes = 0
  atomic.StoreUint32(&b.state, newState)
}

```

### 2 个动作
#### Open Breaker
+ `openBreaker` : 打开熔断器，状态变为 `open`，过了设定的超时时间后会自动变为 `halfOpen` 状态
+ 触发条件
  1. 熔断器处于 `closed` 状态，且作业执行累积的异常次数已达到阈值 `errorThreshold`
  1. 熔断器处于 `halfOpen` 状态，且最新作业的执行结果异常
```go
func (b *Breaker) openBreaker() {
  b.changeState(open)
  go b.timer()
}

func (b *Breaker) timer() {
  time.Sleep(b.timeout)
  b.lock.Lock()
  defer b.lock.Unlock()
  b.changeState(halfOpen)
}
```


#### Close Breaker

+ `closeBreaker` : 关闭熔断器，状态变为 `closed`
+ 触发条件: 熔断器处于 `halfOpen` 状态，且作业执行累积的成功次数已达到阈值
```go
func (b *Breaker) closeBreaker() {
  b.changeState(closed)
}
```

### 状态流转

每次作业执行完成后都会对执行结果进行判断和统计
1. 如果执行结果没有异常，只有当熔断器状态为 `halfOpen` 时才进行作业成功次数的统计，如果达到阈值则关闭熔断器
1. 如果执行结果返回异常
    + 只有当熔断器状态为 `closed` 时才进行作业异常次数的统计，如果达到阈值则[打开熔断器](#Open-Breaker)
    + 当熔断器状态为 `halfOpen` 时，立即[打开熔断器](#Open-Breaker)

```go
func (b *Breaker) processResult(result error, panicValue interface{}) {
  b.lock.Lock()
  defer b.lock.Unlock()

  if result == nil && panicValue == nil {
    if b.state == halfOpen {
      b.successes++
      if b.successes == b.successThreshold {
        b.closeBreaker()
      }
    }
  } else {
    if b.errors > 0 {
      expiry := b.lastError.Add(b.timeout)
      if time.Now().After(expiry) {
        b.errors = 0
      }
    }

    switch b.state {
    case closed:
      b.errors++
      if b.errors == b.errorThreshold {
        b.openBreaker()
      } else {
        b.lastError = time.Now()
      }
    case halfOpen:
      b.openBreaker()
    }
  }
}
```

这样，熔断器就能基于作业的执行情况，动态地调整应对策略，从而使得服务免于雪崩效应

## 批量处理

对于像监控上报的场景，如果消息总是单条上报，对客户端或服务端都会造成不小的压力，因此更常见的策略是**收集一批次的消息，等待指定的时间再统一进行上报**。这里就涉及到对作业的批量处理了


```go
b := New(10*time.Millisecond, func(params []interface{}) error {
  // do something with the batch of parameters
  return nil
})

b.Prefilter(func(param interface{}) error {
  // do some sort of sanity check on the parameter, and return an error if it fails
  return nil
})

for i := 0; i < 10; i++ {
  go b.Run(i)
}
```
[Request Batch](./drawio/request-batch.drawio){link-type="drawio"}


批量处理 `Batcher` 的构造函数需要两个参数:
+ `timeout` : 批次从创建到执行前的[收集时间](#收集)
+ `doWork` : 批次作业执行的具体逻辑

```go
// Batcher implements the batching resiliency pattern
type Batcher struct {
  timeout   time.Duration
  prefilter func(interface{}) error

  lock   sync.Mutex
  submit chan *work
  doWork func([]interface{}) error
}

// New constructs a new batcher that will batch all calls to Run that occur within
// `timeout` time before calling doWork just once for the entire batch. The doWork
// function must be safe to run concurrently with itself as this may occur, especially
// when the timeout is small.
func New(timeout time.Duration, doWork func([]interface{}) error) *Batcher {
  return &Batcher{
    timeout: timeout,
    doWork:  doWork,
  }
}
```

简单地说，`Batcher` 的执行逻辑是先收集待执行的作业单元 `work`，再组合作业单元为批次 `batch`，然后以批次为作业单位去执行作业

### 收集
每一个作业单元在 `Batcher` 里都使用 `work` 表示

```go
type work struct {
  param  interface{}
  future chan error
}
```

每次 `Run` 一个作业单元时，`Batcher` 并没有立即执行作业单元，而是将 `work` 发送至作业单元队列 `submit` 里

此处的作业单元队列可以理解为一个特定的批次 `batch`，**当某个 `batch` 的所有 `work` 都被执行完成后，这个批次将会被清空，等待下一个批次的收集**


```go
// Run runs the work function with the given parameter, possibly
// including it in a batch with other calls to Run that occur within the
// specified timeout. It is safe to call Run concurrently on the same batcher.
func (b *Batcher) Run(param interface{}) error {
  if b.prefilter != nil {
    if err := b.prefilter(param); err != nil {
      return err
    }
  }

  if b.timeout == 0 {
    return b.doWork([]interface{}{param})
  }

  w := &work{
    param:  param,
    future: make(chan error, 1),
  }

  b.submitWork(w)

  return <-w.future
}

func (b *Batcher) submitWork(w *work) {
  b.lock.Lock()
  defer b.lock.Unlock()

  if b.submit == nil {
    b.submit = make(chan *work, 4)
    go b.batch()
  }

  b.submit <- w
}
```
并不是所有收集的作业单元都需要被执行，所以提供了个 `Prefilter` 函数，在提交作业单元到 `batch` 前可校验作业单元参数的合法性

```go
// Prefilter specifies an optional function that can be used to run initial checks on parameters
// passed to Run before being added to the batch. If the prefilter returns a non-nil error,
// that error is returned immediately from Run and the batcher is not invoked. A prefilter
// cannot safely be specified for a batcher if Run has already been invoked. The filter function
// specified must be concurrency-safe.
func (b *Batcher) Prefilter(filter func(interface{}) error) {
  b.prefilter = filter
}
```

### 触发
对于每一个批次都会有一个定时器，等时间一到就会关闭当前批次接收 `work` 的通道

```go
func (b *Batcher) timer() {
  time.Sleep(b.timeout)

  b.lock.Lock()
  defer b.lock.Unlock()

  close(b.submit)
  b.submit = nil
}
```

这恰好使得批次不阻塞于从 `submit` 变量中获取 `work`，而是继续对收集好的 `work` 执行 `doWork` 逻辑，等执行结束再将此批次的作业执行结果发到每个 `work` 的 `future` 通道中，使得调用方能得知批次作业的执行结果

```go
func (b *Batcher) batch() {
  var params []interface{}
  var futures []chan error
  input := b.submit

  go b.timer()

  for work := range input {
    params = append(params, work.param)
    futures = append(futures, work.future)
  }

  ret := b.doWork(params)

  for _, future := range futures {
    future <- ret
    close(future)
  }
}
```

综上所述，批量执行作业的流程如下图所示

[Batcher](./drawio/batcher.drawio){link-type="drawio"}

## 总结
虽然写了一年的 `golang` 语言，但是从不敢说自己真的深入理解它

近期比较深入地学习，才发现自己的水平根本还达不到标准，后续应该要从这些短小精悍的项目开始去夯实自己的语言基础
