package printinorder

type Foo struct {
	firstDone  chan struct{}
	secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.firstDone <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.firstDone
	printSecond()
	f.secondDone <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.secondDone
	printThird()
}
