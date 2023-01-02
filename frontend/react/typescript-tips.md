# [typescript-tips](https://basarat.gitbook.io/typescript/main-1)

## String Based Enums

Sometimes you need a collection of strings collected under a common key. Use string literal types to create string based enums by combining with union types

## Nominal Typing

## Stateful Functions
A common feature in other programming languages is usage of the `static` keyword to increase the lifetime(not scope) of a function variable to live beyond function invocations. Here's a C sample that achieves this:
```c
// --run--
void called() {
  static count = 0;
  count++;
  printf("Called: %d", count);
}
int main() {
  called();
  called();
  return 0;
}
```
Since Typescript doesn't have function statics you can achieve the same thing using various abstractions that wrap over a local variable e.g. using a `class`:

```ts
// --run--
const { called } = new class {
  count = 0;
  called = () => {
    this.count++;
    console.log(`Called: ${this.count}`);
  }
}

called();
called();
```

## Currying
Just use a chain of fat arrow functions
```ts
// A curried function
let add = (x: number) => (y: number) => x+y;
add(1)(2);
// Partially applied
let add9 = add(9);
// fully apply the function
add9(10);
```
## Type Instantiation

## Lazy Object Literal Initialization
Quite commonly in Javascript code bases you would initialize object literals in the following manner:
```js
let foo = {};
foo.bar = 12;
foo.bas = 'Hello';
```
As soon as you move the code to TypeScript you will start to get Errors like the following
```ts
let foo = {};
foo.bar = 123; // Error: Property 'bar' does not exist on type '{}'
foo.bas = "Hello World"; // Error: Property 'bas' does not exist on type '{}'
```
This is because from the state let foo = {}, TypeScript infers the type of foo (left hand side of initializing assignment) to be the type of the right hand side {} (i.e. an object with no properties). So, it error if you try to assign to a property it doesn't know about

### Ideal Fix
The *proper* way to initialize an object in TypeScript is to do it in the assignment:
```ts
let foo = {
  bar: 123,
  bas: 'Hello',
};
```
This is also great for code review and code maintainability purposes

> The quick fix and middle ground lazy initialization patterns described below suffer from *mistakenly forgetting to initialize a property*
### Quick Fix
If you have a large Javascript code base that you are migrating to TypeScript the ideal fix might not be a viable solution for you. In that case you can carefully use a *type assertion* to silence the compiler
```ts
let foo = {} as any;
foo.bar = 123;
foo.bas = 'Hello';
```

### Middle Ground
Of course using the *any* assertion can be very bad as it sort of defeats the safety of TypeScript. The middle ground fix is to create an `interface` to ensure:
+ Good Docs
+ Safe assignment

This is shown below
```ts
interface Foo {
    bar: number
    bas: string
}

let foo = {} as Foo;
foo.bar = 123;
foo.bas = "Hello World";
```

## Classes are Useful

## Avoid Export Default

## Limit Property Setters
Prefer explict set/get functions(e.g. `setA` and `getA` functions) over setters/getters

Consider the following code:
```ts
foo.bar = {
  a: 123,
  b: 252
};
```
In the presence of setter/getters:

```ts
class Foo {
  a: number;
  b: number;
  set bar(value:{a: number, b: number}) {
    this.a = a;
    this.b = b;
  }
}
```
This is not a good use of property setters. The person reading the first code sample has no context about all the things that will change. Whereas someone calling `foo.setBar(value)` might have an idea that something might change on `foo`

> Bonus points: Find references works better if you have different functions. In TypeScript tools if you find references for a getter or a setter you get *both* whereas with explict function calls you only get references to the relevant function

## outFile caution

## JQuery 
> Note: you need to install the `jquery.d.ts` file for these tips

Just create `jquery-foo.d.ts` with:
```ts
interface JQuery {
  foo: any;
}
```
And now you can use `$('something').foo({whateverYouWant:'hello jquery plugin'})`
## Static Constructors

## Singleton Pattern

## Function Parameters

## Build Toggles

## Barrel

## Create Arrays
Creating an empty array is super easy:
```ts
const foo: string[] = [];
```
If you want to create an array pre-filled with some content using the ES6 `Array.prototype.fill`:
```ts
const foo: string[] = new Array(3).fill('');
console.log(foo);
```
If you want to create an array of a predefined length with calls you can use the spread operator:
```ts
const someNumbers = [...new Array(3)].map((_, i) => i * 10);
console.log(someNumbers);
```


## Typesafe Event Emitter
Conventionally in Node.js and traditional JavaScript you have a single event emitter. This event emitter internally tracks listener for different event types e.g.

```ts
const emitter = new EventEmitter();
```