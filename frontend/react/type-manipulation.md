# Type Manipulation

## Keyof Type Operator
The `keyof` operator takes an object type and produces a string or numeric literal union of its keys

```ts
type Point = { x: number; y: number }
type P = keyof Point;

type Arrayish = { [n: number]: unknown };
type A = keyof Arrayish; // number

type Mapish = { [n: string]: unknown };
type B = keyof Mapish; // string | number
```
::: warning Notice
B is `string | number`. This's because Javascript object keys are always coerced to a string, so `obj[0]` is always the same as `obj["0"]`
:::

## Typeof Type Operator
TypeScript adds a `typeof` operator you can use in a type context to refer to the type of a variable or property
```ts
function f() {
  return { x: 10 };
}
type P = ReturnType<typeof f>;
```

## Indexed Access Types
We can use an indexed access type to look up a specific property on another type

The indexing type is itself a type, so we can use unions, `keyof`, or other types entirely

**You can only use types when indexing, meaning you can't use a `const` to make a variable reference**
```ts
type Person = {
  age: number;
  name: string;
}
type Age = Person["age"]; // number
type I1 = Person["age" | "name"]
type I2 = Person[keyof Person]

const MyArr = [
  { name: 'a', age: 1},
  { name: 'b', age: 2},
  { name: 'c', age: 3},
];
type P = typeof MyArr[number];
```

## Conditional Types
Conditional types take a form that looks a little like conditional expressions
```ts
SomeType extends OtherType ? TrueType : FalseType
```
The power of conditional types comes from using them with generics
```ts
interface IdLabel {
  id: number;
}
interface NameLabel {
  name: string;
}
type NameOrId<T extends number | string> = T extends number ? IdLabel : NameLabel;
function createLabel<T extends number | string>(idOrName: T): NameOrId<T> {
  // ...
}
```
When conditional types act on a generic type, they become distributive when given a union type
```ts
type ToArray<T> = T extends any ? T[] : never;
type StrArrOrNumArr = ToArray<string | number>; // string[] | number[]
```
Typically, distributivity is the desired behavior. To avoid that behavior, you can surround each side of the `extends` keyword with square brackets
```ts
type ToArrayNonDist<T> = [T] extends [any] ? T[] : never;
type StrArrOrNumArr = ToArrayNonDist<string | number>; // (string | number)[]
```

Conditional types provide us with a way to infer from types we compare against in the true branch using the `infer` keyword. For example, we could have inferred the element type in `Flatten` instead of fetching it out manually with an indexed access type
```ts
// type Flatten<T> = T extends any[] ? T[number] : T;
type Flatten<T> = T extends Array<infer I> ? I : T;
```
***When inferring from a type with multiple call signatures(such as the type of an overloaded function), inferences are made from the last signature***
```ts
type GetReturnType<T> = T extends (...args: never[]) => infer Return ? Return : never;
type Num = GetReturnType<() => number>;
type Str = GetReturnType<() => string>;

declare function stringOrNum(x: string): number;
declare function stringOrNum(x: number): string;
declare function stringOrNum(x: string | number): string | number;
 
type T1 = ReturnType<typeof stringOrNum>;
```

## Mapped Types
A mapped type is a generic type which uses a union of `PropertyKey` (frequently created via a `keyof`) to iterate through keys to create a type
```ts
type OptionsFlags<T> = {
  [P in keyof T]: boolean;
}
type FeatureFlags = {
  darkMode: () => void;
}
type FeatureOptions = OptionsFlags<FeatureFlags>;

```
There are two additional modifiers which can be applied during mapping: `readonly` and `?` which affect mutability and optionality respectively. You can remove or add these modifiers by prefixing with `-` and `+`. If you don't add a prefix, then `+` is assumed

```ts
type CreateMutable<T> = {
  - readonly [P in keyof T]: T[P];
}
type LockedAccount {
  readonly id: string;
  readonly name: string;
}
type UnlockedAccount = CreateMutable<LockedAccount>;

type Concrete<T> = {
  [P in keyof T]-?: T[P];
}
type MaybeUser = {
  id: string;
  name?: string;
  age?: number;
}
type User = Concrete<MaybeUser>;
```
We can remap keys in mapped types with an `as` clause in a mapped type
+ We can leverage features like `template literal types` to create new property names from prior ones
  ```ts
  type Getter<T> = {
    [P in keyof T as `get${Capitalize<string & P>}`]: () => T[P];
  }
  interface Person {
    name: string;
    age: number;
    location: string;
  }
  type LazyPerson = Getter<Person>;
  // type LazyPerson = {
  //   getName: () => string;
  //   getAge: () => number;
  //   getLocation: () => string;
  // }
  ```
+ We can filter out keys by producing `never` via a conditional type
  ```ts
  type RemoveKindField<T> = {
    [P in keyof T as Exclude<P, "kind">]: T[P];
  }
  ```