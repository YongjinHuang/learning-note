# js-object


## Chapter 1: Object Foundations
> JS definitely has objects, but that doesn't mean that all values are objects

Object are the foundation for the second of JS's three pillars: the prototype


## Chapter 2: How Objects Work

The characteristics that define the underlying behavior of objects are collectively referred to in formal terms as the [metaobject protocol (MOP)](https://github.com/getify/You-Dont-Know-JS/blob/2nd-ed/objects-classes/ch2.md#user-content-fn-mop-e24d26e7bbf7e3abc8f3a7997d01b0c2). The MOP is useful not only for understanding how objects will behave, but also for overriding the default behaviors of objects to bend the language to fit our program's needs more fully

Each property on an object is internally described by what's known as a *property descriptor*. This is, itself, an object(aka metaobject) with several properties on it, dictating how the target property behaves
```js
// --run--
var myObj = {
    favoriteNumber: 42,
    isDeveloper: true,
    firstName: "Kyle"
};

console.log(Object.getOwnPropertyDescriptor(myObj,"favoriteNumber"));

var anotherObj = {};

Object.defineProperty(anotherObj,"fave",{
    value: 42,
    enumerable: true,     // default if omitted
    writable: true,       // default if omitted
    configurable: true    // default if omitted
});

console.log(anotherObj.fave);          // 42

var anotherObj = {};

Object.defineProperties(anotherObj,{
    "fave": {
        // a property descriptor
    },
    "superFave": {
        // another property descriptor
    }
});
```

A special kind of property, known as an *accessor property*(aka a getter/setter) can be defined. For these a property like this, its descriptor does not define a fixed value property, but would instead look something like this:

```js
{
  get() { .. },    // function to invoke when retrieving the value
  set(v) { .. },   // function to invoke when assigning the value
  // .. enumerable, etc
}
```

A getter looks like a property access(`obj.prop`), but under the covers it invokes the `get()` method as defined; it's sort of like if you had called `obj.prop()`; A setter looks like a property assignment(`obj.prop = value`), but it invokes the `set(...)` method as defined; it's sort of like if you had called `obj.prop(value)`

```js
// --run--
var anotherObj = {};

Object.defineProperty(anotherObj,"fave",{
    get() { console.log("Getting 'fave' value!"); return 123; },
    set(v) { console.log(`Ignoring ${v} assignment.`); }
});

console.log(anotherObj.fave);
// Getting 'fave' value!
// 123

anotherObj.fave = 42;
// Ignoring 42 assignment.

console.log(anotherObj.fave);
// Getting 'fave' value!
// 123
```

TODO
Besides `value` or `get()` / `set(...)` the other 3 attributes of a property descriptor are as shown above:
+ `enumerable`
+ `writable`
+ `configurable`

## Chapter 4: This Works

Here's the most important thing to understand `this`: the determination of what value (usually object) `this` points at is not made at author time, but rather determined at runtime. 

That means you can not simply look at a `this`-aware function (even a method in a `class` definition) and know for sure what `this` will hold while that function runs. Instead you have to find each place the function is invoked, and look at how it's invoked (not even where matters). That's the only way to fully answer what `this` will point to

::: tip Typical question
When the function is invoked a certain way, what `this` will be assigned for that invocation
:::

### This Aware

