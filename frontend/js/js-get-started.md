# YDKJSY

## Chapter 1 : What Is JavaScript?
A good start always depends on a solid first step

The truth is, the name JavaScript is an artifact of marketing shenanigans

Java is to JavaScript as ham is to hamster

TC39 is the technical steering committee that manages JS. Their primary task is managing the official specification for the language

JS's syntax and behavior are defined in the ES specification

For the most part, the JS defined in the specification and the JS that runs in browser-based JS engines is the same. But there are some differences that must be considered

Most of the cross-browser differences people complain about with "JS is so inconsistent" claims are actually due to differences in how those environment behaviors work, not in how the JS itself works

Using the Console or REPL(Read-Evaluate-Print-Loop) is not a pretty straightforward JS environment. Think of the console as a JS-Friendly environment

JavaScript is most definitely a multi-paradigm language. You can write procedural, class-oriented, or FP-style code, and you can make those decisions on a line-by-line basis instead of being forced into an all-or-nothing choice

Backwards compatibility means that once something is accepted as valid JS, there will not be a future change to the language that causes that code to become invalid JS

Being forwards compatibility means that including a new addition to the language in a program would not cause the program to break if it were run in an older JS engine. **JS is not forwards-compatibility**

Typically, forwards-compatibility problems related to syntax are solved by using a transpiler

Developers should focus on writing the clean, new syntax forms, and let the tools take care of producing a forward-compatible version of that code that is suitable to deploy and run on the oldest-supported JS engine environments

Transpilation and polyfilling are two highly effective techniques for addressing that gap between code that uses the latest stable features in the language and the old environments a site or application needs to still support

**JS is a compiled language**, meaning the tools (including the JS engine) process and verify a program (reporting any errors) before it executes


WASM is similar to ASM.js in that its original intent was to provide a path for non-JS programs to be converted to a form that could run in the JS engine. Unlike ASM.js, WASM chose to additionally get around some of
the inherent delays in JS parsing/compilation before a program can execute,
by representing the program in a form that is entirely unlike JS

WASM will not replace JS. WASM significantly auguments what the web (including JS) can accomplish. That's a great thing, entirely orthogonal to whether some people will use it as an escape hatch from having to write JS

Taken together, strict mode is largely the de facto default even though  technically it’s not actually the default

## Chapter 2 : Surveying JS

Regardless of which code organization pattern and loading mechanism is used for a file (standalone or module), you should still think of each file as its own program, which may then cooperate with other programs to perform the functions of your overall application

The most fundamental unit of information in a program is a value. Values are data. They're how the program maintains state

The better approach is to use " or ' for strings unless you need interpolation; reverse ` only for strings that will include interpolated expressions

For distinguishing values, the typeof operator tells you its built-in type, if primitive, or "object" otherwise

> typeof null unfortunately returns "object" instead of the expected "null". Also typeof returns the specific "function" for functions, but not the expected "array" for arrays

Think of variables as just containers for values

```js
var name = "Kyle";
var age
```

The `var` keyword declares a variable to be used in that part of the program, and optionally allows initial value assignment

```js
let name = "Kyle";
let age
```
The `let` keyword has some differences to `var`, with the most obvious being that `let` allows a more limited access to the variable than `var`. This's called `block scoping` as opposed to regular or function scoping

A third declartion form is `const`. It's like `let` but has an additional limitation that it must be given a value at the moment it's declared, and can not be re-assigned a different value later

It's ill-advised to use `const` with object values, because those values can still be changed even though the variable can't be re-assigned


We must be aware of the nuanced differences between an equality comparison and an euivalence comparison

`===`'s equality comparison is often described as "checking both the value and the type". But JS does not define `===` as structural equality for object values. Instead, `===` uses identity equality for object values

The `===` operator is designed to lie in two cases of special values: `NaN` and `-0`. It's best to avoid using `===` for them

```js
NaN === NaN; // false
0 === -0; // true
```

For `NaN` comparisons use the `Number.isNaN(..)` utility, which does not lie. For `-0` comparison use the `Object.is(..)` utility, which also does not lie

Coercion means a value of one type being converted to its respective representation in another type

If the comparison is between the same value type, both `==` and `===` do exactly the same thing, no difference whatsoever. If the value types being compared are different, the `==` differs from `===` in that it allows coercion before the comparison

The `==` operator should be described as coercive equality

Two major patterns for organizing code (data and behavior) are used broadly accross the JS ecosystem: classes and modules

The key hallmarks of a classic module are an outer function (that runs at least once), which returns an "instance" of the module with one or more functions exposed that can operate on the module instance's internal data

A class in a program is a definition of a "type" of custom data structure that includes both data and behaviors that operate on that data


## Chapter 3 : Digging to the Roots of JS


Closure is when a function remembers and continues to access variables from outside its scope, even when the function is executed in a different scope


## Chapter 4 : The Bigger Picture

JS is lexically scoped, though many clain it isn't, because of two particular characteristics of its model that are not present in other lexically scoped languages
+ `hoisting` : When all variables declared anywhere in a scope are treated as if they're declared at the beginning of the scope
+ `var-declared variables` : var-declared are function scoped, even if they appear inside a block