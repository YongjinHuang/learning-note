# js-types

## Chapter 1: Primitive Values

Here, we'll look at the core value types of JS, specifically the non-object types called *primitives*

JS doesn't apply types to variables or properties -- what I call, "container types" -- but rather, values themselves have types -- what I call, "value types"

The language provides seven built-in, primitive (non-object) value types:
+ `undefined`
+ `null`
+ `boolean`
+ `number`
+ `bigint`
+ `symbol`
+ `string`

Any value's value-type can be inspected via the `typeof` operator, which always returns a `string` value representing the underlying JS value-type:

```js
typeof true
typeof 42
typeof 42n
typeof Symbol("42")
```

JS variables themselves don't have type. They hold any arbitrary value, which itself has a value-type

## Chapter 2: Primitive Behaviors

All primitive values are immutable, meaning nothing in a JS program can reach into the contents of the value and modify it in any way

Even a string value, which looks like merely an array of characters, and array contents are typically mutable, is immutable

```js
// --run--
var a = 'Hello.';
a[5] = '!';
console.log(a);
```
::: warning Warning
In non-strict mode, assigning to a read-only property (like `greeting[5] = ..`) silently fails. In strict-mode, the disallowed assignment will throw an exception.
:::

Additionally, properties can not add to any primitive values

```js
// --run--
var b = 'Hello.'
b.hei = true
b.hei
```
Property access is not allowed in any way on nullish primitive values `null` and `undefined`. But properties can be accessed on all other primitive values

Non-nullish primitive values also have a couple of standard built-in methods that can be accessed

```js
// --run--
var c = 'Hello.'
c.toString();
c.valueOf();
```
Additionally, most of the primitive value-types define their own methods with specific Behaviors inherent to that type

**Any assignment of a primitive value from one variable/container to another is a *value copy***

::: details Note
Inside the JS engine, it _may_ be the case that only one `42` value exists in memory, and the engine points both `myAge` and `yourAge` variables at the shared value. Since primitive values are immutable, there's no danger in a JS engine doing so. But what's important to us as JS developers is, in our programs, `myAge` and `yourAge` act as if they have their own copy of that value, rather than sharing it
:::

## Chapter 3: Object Values


## Chpater 4: Coercing Values