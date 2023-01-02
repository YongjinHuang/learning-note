# Compiler

Most compilers break down into three primary stages: Parsing, Transformation, and Code Generation
1. *Parsing* is taking raw code and turning it into a more abstract representation of the code
1. *Transformation* takes this abstract representation and manipulates to do whatever the compiler wants it do
1. *Code Generation* takes the transformed representation of the code and turns it into new code

## Parsing

Parsing typically gets broken down into two phases: Lexical Analysis and Syntactic Analysis
1. *Lexical Analysis* takes the raw code and splits it apart into these things called tokens by a thing called a tokenizer(or lexer). Tokens are an array of tiny little objects that describe an isolated piece of the syntax. They could be numbers, labels, punctuation, operators, whatever 
1. *Syntactic Analysis* takes the tokens and reformats them into a representation that describes each part of the syntax and their relation to one another. This is known as an intermediate representation or Abstract Syntax Tree. An Abstract Syntax Tree, or AST for short, is a deeply nested object that represents code in a way that is both easy to work with and tells us a lot of information

For the following syntax

```lisp
(add 2 (subtract 4 2))
```
Tokens might look something like this:
```js
[
  { type: 'paren',  value: '('        },
  { type: 'name',   value: 'add'      },
  { type: 'number', value: '2'        },
  { type: 'paren',  value: '('        },
  { type: 'name',   value: 'subtract' },
  { type: 'number', value: '4'        },
  { type: 'number', value: '2'        },
  { type: 'paren',  value: ')'        },
  { type: 'paren',  value: ')'        },
]
```

And an Abstract Syntax Tree (AST) might look like this:

```js
{
  type: 'Program',
  body: [{
    type: 'CallExpression',
    name: 'add',
    params: [{
      type: 'NumberLiteral',
      value: '2',
    }, {
      type: 'CallExpression',
      name: 'subtract',
      params: [{
        type: 'NumberLiteral',
        value: '4',
      }, {
        type: 'NumberLiteral',
        value: '2',
      }]
    }]
  }]
}
```

## Transformation

The next type of stage for a compiler is transformation. Again, this just takes the AST from the last step and makes changes to it. It can manipulates the AST in the same language or it can traslate it into an entirely new language

Let's look at how we would transform an AST. You might notice that our AST has elements within it that look very similar. There're these objects with a `type` property. Each of these are known as an AST Node. These nodes have defined properties on them that describe one isolated part of the tree

We can have a node for a `NumberLiteral`
```js
{
  type: 'NumberLiteral',
  value: '2',
}
```
Or maybe a node for a `CallExpression`
```js
{
  type: 'CallExpression',
  name: 'substract',
  params: [...nested nodes go here...],
}
```
When transforming the AST we can manipulate nodes by adding/removing/replacing properties, we can add new nodes, remove nodes, or we could leave the existing AST alone and create an entirely new one based on it

### Traversal
In order to navigate through all of thease nodes, we need to be able to traverse through them. This traversal process goes to each node in the AST depth-first

```js
{
  type: 'Program',
  body: [{
    type: 'CallExpression',
    name: 'add',
    params: [{
      type: 'NumberLiteral',
      value: '2',
    }, {
      type: 'CallExpression',
      name: 'subtract',
      params: [{
        type: 'NumberLiteral',
        value: '4',
      }, {
        type: 'NumberLiteral',
        value: '2',
      }]
    }]
  }]
}
```
So for the above AST we would go:
+ Program {.mindmap}
    + CallExpression (add)
        + NumberLiteral (2)
        + CallExpression (subtract)
            + NumberLiteral (4)
            + NumberLiteral (2)

### Visitors

The basic idea is that we are going to create a visitor object that has methods that will accept different node types. When we traverse our AST, we will call the methods on this visitor whenever we *enter* a node of a matching type. In order to make this useful we will also pass the node and a reference to the parent node

```js
var visitor = {
  NumberLiteral(node, parent) {},
  CallExpression(node, parent) {},
}
```
As we traverse down, we're going to reach branches with dead ends. As we finish each branch of the tree we *exit* it. So going down the tree we *enter* node, and going back up we *exit*

+ Traversal{.mindmap}
    + Program (enter)
        + CallExpression (enter)
            + NumberLiteral (enter)
            + NumberLiteral (exit)
            + CallExpression (enter)
                + NumberLiteral (enter)
                + NumberLiteral (exit)
                + NumberLiteral (enter)
                + NumberLiteral (exit)
            + CallExpression (exit)
        + CallExpression (exit)
    + Program (exit)

In order to Support that, the final form our visitor will look like this:
```js
var visitor = {
  NumberLiteral: {
    enter(node, parent) {},
    exit(node, parent) {},
  }
}
```

## Code Generation
The final phase of a compiler is code generation. Sometimes compilers will do things that overlap with transformation, but for the most part code generation just means take our AST and string-ify code back out