# most-adequate-guide
## Pure Functions

A function is a special relationship between values: Each of its input values given back exactly one outeput value

A ***pure function*** is a function that, given the same input, will always return the same output and does not have any observable side effect

```js
const xs = [1,2,3,4,5];

// pure
xs.slice(0,3); // [1,2,3]

xs.slice(0,3); // [1,2,3]

xs.slice(0,3); // [1,2,3]


// impure
xs.splice(0,3); // [1,2,3]

xs.splice(0,3); // [4,5]

xs.splice(0,3); // []
```

A ***side effect*** is a change of system state or observable interaction with the outside world that occurs during the calculation of a result

+ Side effects{.mindmap}
    + changing the file system
    + inserting a record into a database
    + making an http call
    + mutations
    + printing to the screen / logging
    + obtaining user input
    + querying the DOM
    + accessing system state

The case for purity
+ The Case for Purity{.mindmap}
    + Cacheable
    + Portable / Self-documenting
    + Testable
    + Reasonable
    + Parallel Code

## Currying
The concept of currying is simple: you can call a function with fewer arguments than it accepts. It returns a function that takes the remaining arguments

```js
const add = x => y => x + y;
const increment = add(1);
const addTen = add(10);
```

Giving a function fewer arguments than it expects is typically called ***partial application***


## Compose
Composition connects our functions together like a series of pipes. Data will flow from right to left through our application as it must

```js
// --run--
const compose = (f, g) => x => f(g(x));
const toUpperCase = x => x.toUpperCase();
const exclaim = x => `${x}!`;
const shout = compose(exclaim, toUpperCase);
// without compose, the above would read
// const shout = exclaim(toUpperCase(x))
console.log(shout('send in the clowns'));
```

Composition is associative, meaning it doesn't matter how you group two of them

Pointfree style means never having to say your data
```js
// not pointfree because we mention the data: w
const snakeCase = w => w.toLowerCase().replace(/\s+/ig, '_');
// pointfree
const snakeCase = compose(replace(/\s+/ig, '_'), toLowerCase)
```

A common mistake is to compose something like `map`, a function of two arguments, without first partially applying it

```js
// wrong - we end up giving angry an array and we partially applied map with who knows what.
const latin = compose(map, angry, reverse);

latin(['frog', 'eyes']); // error

// right - each function expects 1 argument.
const latin = compose(map(angry), reverse);

latin(['frog', 'eyes']); // ['EYES!', 'FROG!'])
```
If you are having trouble debugging a composition, we can use this helpful, but impure trace function to see what's going on

The trace function allows us to view the data at a certain point for debugging purposes
```js
const trace = curry((tag, x) => {
  console.log(tag, x);
  return x;
});

const dasherize = compose(
  intercalate('-'),
  toLower,
  split(' '),
  replace(/\s{2,}/ig, ' '),
);

dasherize('The world is a vampire');
// TypeError: Cannot read property 'apply' of undefined
const dasherize = compose(
  intercalate('-'),
  toLower,
  trace('after split'),
  split(' '),
  replace(/\s{2,}/ig, ' '),
);
```
## Functor
A ***functor*** is a type that implements `map` and obeys some laws. In category theory, functors take the objects and morphisms of of a category and map them to a different category
```js
// map :: Functor f => (a -> b) -> f a -> f b
const map = curry((f, anyFunctor) => anyFunctor.map(f));
// identity
map(id) === id;

// composition
compose(map(f), map(g)) === map(compose(f, g));
```

![Categories mapped](https://2816854455-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-MT09zmSclnRGt38Vn0l%2Fsync%2Fba3d9f467fe60dffed05d7503e2f0d41d20876e0.png?generation=1612779926245350&alt=media)

![functor diagram](https://2816854455-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-MT09zmSclnRGt38Vn0l%2Fsync%2F31fa2ac5ef79006d1b75bd097d0f0209dcd0de02.png?generation=1612779927448867&alt=media)

```js
// topRoute :: String -> Maybe String
const topRoute = compose(Maybe.of, reverse);

// bottomRoute :: String -> Maybe String
const bottomRoute = compose(map(reverse), Maybe.of);

topRoute('hi'); // Just('ih')
bottomRoute('hi'); // Just('ih')
```
![functor diagram 2](https://2816854455-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-MT09zmSclnRGt38Vn0l%2Fsync%2F65421639ee6470954bd6fba3e9c45c86939515f8.png?generation=1612779925219492&alt=media)

## Monads

A ***pointed functor*** is a functor with an `of` method

Monads are pointed functors that can flatten

Any functor which defines a `join` method, has an `of` method, and obeys a few laws is a monad

```js
// join :: Monad m => m (m a) -> m a
const join = mma => mma.join();

// firstAddressStreet :: User -> Maybe Street
const firstAddressStreet = compose(
  join,
  map(safeProp('street')),
  join,
  map(safeHead), safeProp('addresses'),
);

firstAddressStreet({
  addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: 'WC2N' }],
});
// Maybe({name: 'Mulburry', number: 8402})
```

```js
// chain :: Monad m => (a -> m b) -> m a -> m b
const chain = curry((f, m) => m.map(f).join());

// or

// chain :: Monad m => (a -> m b) -> m a -> m b
const chain = f => compose(join, map(f));
// map/join
const firstAddressStreet = compose(
  join,
  map(safeProp('street')),
  join,
  map(safeHead),
  safeProp('addresses'),
);

// chain
const firstAddressStreet = compose(
  chain(safeProp('street')),
  chain(safeHead),
  safeProp('addresses'),
);

// map/join
const applyPreferences = compose(
  join,
  map(setStyle('#main')),
  join,
  map(log),
  map(JSON.parse),
  getItem,
);

// chain
const applyPreferences = compose(
  chain(setStyle('#main')),
  chain(log),
  map(JSON.parse),
  getItem,
);
```

```js
// associativity
compose(join, map(join)) === compose(join, join);
```
![monad associativity law](https://2816854455-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-MT09zmSclnRGt38Vn0l%2Fsync%2F5626436501de6f2d4db2ffd8e1f8dbee35074094.png?generation=1612779914164248&alt=media)

```js
// identity for all (M a)
compose(join, of) === compose(join, map(of)) === id;
```
![monad identity law](https://2816854455-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-MT09zmSclnRGt38Vn0l%2Fsync%2Fc4ef025f6e0c0a5565e1cc3889b3709347b76ba7.png?generation=1612779913153395&alt=media)

## Applicative Functors

`ap` is a function that can apply the function contents of one functor to the value contents of another

An applicative functor is a pointed functor with an `ap` method

```js
F.of(x).map(f) === F.of(f).ap(F.of(x));
```