# every-layout
## Rudiments

### Box
As Rachel Andrew has reminded us, everything in web design is a box, or the absence of a box. Layout is inevitably, therefore, the arrangement of boxes


#### The box model
The box model is the formula upon which layout boxes are based, and comprises content, padding, border and margin. CSS lets us alter these values to change the overall size and shape of elements' display

![box-model](./FILES/every-layout.md/3304e921.svg)

#### The `display` property

Block elements assume all of the available space in one dimension. Typically this's the horizontal dimension, because the `writing-mode` is set to `horizontal-tb` (horizontal: with a top to bottom flow direction)

![writing-mode](./FILES/every-layout.md/0285336f.svg)

Inline elements (with the `display` value `inline`) behave differently. They're laid out in line with the current context, writing mode, and direction. They're only as wide as their content, and are placed adjacently wherever there is space to do so. Block elements follow flow direction, and inline elements follow writing direction

![inline](./FILES/every-layout.md/cb8a4595.svg)

> Thinking typographically, it could be said that block elements are like paragraphs, and inline elements are like words

### Composition
The idea (Composition Over Inheritance) is that combining simple independent parts (objects; classes; functions) gives you more flexibility, and leads to more efficiency, than connecting everything through inheritance to a shared origin 

#### Composition and layout

#### Layout primitives
The purpose of Every Layout is to identify and document what each of these smaller layouts are

![dialog-layout-primitives](./FILES/every-layout.md/fa017e44.svg)

### Units



### Global and local styling

### Modular scale

### Axioms



## Layouts

Each of the following layout primitives is documented with a code generator, and includes a custom element implementation. You can use them tother, in composition, to create robust and responsive layouts with the need for `@media` breakpoints

![The Stack](./FILES/every-layout.md/358b6913.svg)
<center>**The Stack**</center>


![The Box](./FILES/every-layout.md/74a1ff60.svg)
<center>**The Box**</center>

![The Center](./FILES/every-layout.md/e77cc3b7.svg)
<center>**The Center**</center>

![The Cluster](./FILES/every-layout.md/940c5e40.svg)
<center>**The Cluster**</center>

![The Sidebar](./FILES/every-layout.md/b410ec20.svg)
<center>**The Sidebar**</center>

![The Switcher](./FILES/every-layout.md/9de2de6a.svg)
<center>**The Switcher**</center>

![The Cover](./FILES/every-layout.md/2b4875bc.svg)
<center>**The Cover**</center>


![The Grid](./FILES/every-layout.md/7bec036d.svg)
<center>**The Grid**</center>

![The Frame](./FILES/every-layout.md/7d703f0d.svg)
<center>**The Frame**</center>

![The Reel](./FILES/every-layout.md/721e93c2.svg)
<center>**The Reel**</center>

![The Imposter](./FILES/every-layout.md/38134dc1.svg)
<center>**The Imposter**</center>

![The Icon](./FILES/every-layout.md/b38a01c6.svg)
<center>**The Icon**</center>
