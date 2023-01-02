# css-secret
## Borders & Backgrounds
### Translucent borders
By default, backgrounds extend underneath the border area, which indicates the translucent border effect set previously will be covered. CSS3 allows us to change the default behavior of the background by setting the `background-clip: padding-box` to accomplish what we require

```html
<!-- --applet-- Translucent borders -->
<style>
  body {
      background: url('http://csssecrets.io/images/stone-art.jpg');
  }

  div {
      border: 10px solid hsla(0,0%,100%,.5);
      background: white;
      background-clip: padding-box;
      
      /* styling */
      max-width: 20em;
      padding: 2em;
      margin: 2em auto 0;
      font: 100%/1.5 sans-serif;
  }
</style>
<div>
  Can I haz semi-transparent borders?
  Pretty please?
</div>
```
### Multiple borders
It's little known that `box-shadow` accepts a fourth parameter (called **spread radius**), which makes the shadow larger or smaller by the amount you specify

A positive spread radius combined with 0 offsets and 0 blur creates a shadow that looks more like a solid border

```html
<!-- --applet-- Solid border with box-shadow -->
<style>
  div {
    width: 100px;
    height: 60px;
    margin: 25px;
    background: yellowgreen;
    box-shadow: 0 0 0 10px #655;
  }
</style>
<div></div>
```
The good thing about `box-shadow` is that we can have as many of them as we cant, comma separated. The only thing to keep in mind is that `box-shadow`s are overlaid one on top of the other, with the first one being the topmost. Therefore you need to adjust the spread radius accordingly 


```html
<!-- --applet-- Multiple borders -->
<style>
  div {
    width: 100px;
    height: 60px;
    margin: 25px;
    background: yellowgreen;
    box-shadow: 0 0 0 10px #655,
              0 0 0 15px deeppink,
              0 2px 5px 15px rgba(0,0,0,.6);
  }  
</style>
<div></div>
```

The shadow solution works quite well in most cases, but has a few caveats:
+ Shadows don't work exactly like borders, as they don't affect layout and are oblivious to the `box-sizing` property (emulate the extra space a border would occupy via padding or margins, depending on whether the shadow is `inset` or not)
+ Fake borders are on the outside of elements. These do not capture mouse events such as hovering or clicking. You can add the `inset` keyword to make the shadows be drawn on the inside of your element. Note that you will need to add extra padding to produce sufficient spacing

In some cases, if we only need 2 borders, we can use a regular border and the `outline` property for the outer one

```html
<!-- --applet-- Two borders -->
<style>
  div {
    width: 100px;
    height: 60px;
    margin: 25px;
    background: yellowgreen;
    border: 10px solid #655;
    outline: 5px solid deeppink;
  }  
</style>
<div></div>
```
Outlines created throuth the outline property do not follow the element's rounding, although that could change in the future

```html
<!-- --applet-- Rounding borders with outline -->
<style>
  div {
    width: 100px;
    height: 60px;
    margin: 25px;
    background: yellowgreen;
    border: 10px solid #655;
    outline: 5px solid deeppink;
    border-radius: 2px;
  }  
</style>
<div></div>
```

### Flexible background positioning
Fairly often, we want to position a background image with offsets from different corner than the top-left one, such as the bottom right.

The `background-position` property was extended to allow specifying offsets from any corner, by providing keywords before the offsets
```html
<!-- --applet-- Extended background-position solution -->
<style>
  div {
    /* Provide a decent fallback for browsers that don't support the extended background-position syntax */
    background: url(http://csssecrets.io/images/code-pirate.svg)
                no-repeat bottom right #58a;

    background-position: right 20px bottom 10px;
    
    /* Styling */
    max-width: 10em;
    min-height: 5em;
    padding: 10px;
    color: white;
    font: 100%/1 sans-serif;
  }
</style>
<div>Code Pirate</div>
```

There are 4 boxes in every elements

+ element's boxes{.mindmap}
    + margin box
    + border box
    + padding box
    + content box

By default, `background-position` refers to the **padding box**, so that borders don't end up obscuring background images. Therefore `top left` is by default the top-left outer corner of the padding box

We got a new property that we can use to change this behavior: `background-origin`. By default its value is `padding-box`

```html
<!-- --applet-- background-origin -->
<style>
  div {
    background: url(http://csssecrets.io/images/code-pirate.svg)
                no-repeat bottom right #58a;
    background-origin: content-box;
    
    /* Styling */
    max-width: 10em;
    min-height: 5em;
    padding: 10px;
    color: white;
    font: 100%/1 sans-serif;
  }
</style>
<div>Code Pirate</div>
```
If we think of it in terms of offsets from the top-left corner, we basically want an offset of 100% - 20px horizontally and 100% - 10px vertically. The `calc` function allows us to do eactly that sort of calculation
```html
<!-- --applet-- background-position-calc -->
<style>
  div {
    background: url(http://csssecrets.io/images/code-pirate.svg) no-repeat bottom right #58a;
    background-position: calc(100% - 20px) calc(100% - 10px);
    
    /* Styling */
    max-width: 10em;
    min-height: 5em;
    padding: 10px;
    color: white;
    font: 100%/1 sans-serif;
  }
</style>
<div>Code Pirate</div>
```

::: warning calc()
Don't forget to include white space around any - and + operators in `calc()`, otherwise it's a parsing error! The reason for this werid rule is forward compatibility
:::

### TODO Inner rounding

```html
<!-- --applet-- A container with an outline and rounding only on the inside -->
<style>
  .container {
    background: #655;
    padding: .8em;
  }
  .container > div {
      background: tan;
      border-radius: .8em;
      padding: 1em;
  }
</style>
<div class="container">
  <div>
    I have a nice subtle inner rounding,
    don't I look pretty?
  </div>
</div>
```

```html
<!-- --applet-- Inner rounding -->
<style>
  div {
    background: tan;
    border-radius: .8em;
    padding: 1em;   
  }
</style>

<div>
  I have a nice subtle inner rounding,
  don't I look pretty?
</div>
```

### Stripped backgrounds
To avoid having to adjust 2 numbers every time we want to change the stripe width, we can take advantage of the specification:

::: tip CSS Images Level3
If a color stop has a position that is less than the specified position of any color stop before it in the list, set its position to be equal to the largest specified position of any color stop before it
:::

This means that if we set the second color's position at 0, its position will be adjusted by the broswer to be equal to the position of the previous color stop, which is what we wanted anyway

```html
<!-- --applet-- Horizontal Stripes -->
<style>
  div {
    background: linear-gradient(#fb3 30%, #58a 0);
    background-size: 100% 20px;
    width: 200px;
    height: 100px;
  }
</style>
<div></div>
```
The code for vertical stripes is almost the same, with one main difference: an extra first argument that specifies the gradient direction
```html
<!-- --applet-- Vertical Stripes -->
<style>
  div {
    background: linear-gradient(to right, #fb3 50%, #58a 0);
    background-size: 80px 20px;
    width: 240px;
    height: 100px;
  }
</style>
<div></div>
```

### Complex background patterns

### Pseudo random backgrounds


## Shapes

### Flexible ellipses
::: tip Larger border-radius turns the element into a circle
When the sum of any 2 adjacent border radii exceeds the size of the border box, user agents must proportionally reduce the used values of all border radii until none of them overlap
:::

### Parallelograms
Parallelograms are a superset of rectangles: their sides are parallel but their corners are necessarily straight. In visual design, they're often useful to make the design appear more dynamic and convey a sense of movement

We can create the skewed rectangle shape with `skew()` transform, but this also results in the content being skewed, which makes it ugly and unreadable
```html
<!-- --applet-- Parallelograms -->
<style>
  .button {
    margin: 10px;
    padding: 10px;
    background: #58a;
    color: white;
    border-radius: 0px;
    display: inline-block;

    transform: skewX(-25deg);
  }
</style>
<a class="button">Click me</a>
```

**Apply an opposite skew() transform to the content, which will cancel out the outer transform**
```html
<!-- --applet-- Nested element solution -->
<style>
  .button {
    margin: 10px;
    padding: 10px;
    background: #58a;
    color: white;
    border-radius: 0px;
    display: inline-block;

    transform: skewX(-25deg);
  }
  .button > div {
    transform: skewX(25deg);
  }
</style>
<a class="button"><div>Click me</div></a>
```
::: warning Inline element
If you’re applying this effect to
an element that is inline by default, don’t forget to set its display property to something else, like inline-block or block, otherwise transforms will not apply. Same goes for the inner element.
:::


Another idea is to use a pseudo-element to apply all styling to backgrounds, borders, etc, and then transform that. Because our content is not contained in the pseudo-element, it's not affected by the transformation. These techniques can also be used with **any other transformation, in order to transform an element's shape without transoforming its content**

```html
<!-- --applet-- Pseudo element solution -->
<style>
.button {
	position: relative;
	display: inline-block;
	padding: .5em 1em;
	border: 0; margin: .5em;
	background: transparent;
	color: white;
	text-decoration: none;
}

.button::before {
	content: ''; /* To generate the box */
	position: absolute;
	top: 0; right: 0; bottom: 0; left: 0;
	z-index: -1;
	background: #58a;
	transform: skew(-25deg);
}}
</style>

<a href="#yolo" class="button">
  <div>Click me</div>
</a>
```

### Diamond images
```html
<!-- --applet-- diamond transform-based solution -->
<style>
.diamond {
	width: 250px;
	height: 250px;
	transform: rotate(45deg);
	overflow: hidden;
	margin: 100px;
}

.diamond img {
	max-width: 100%;
	transform: rotate(-45deg) scale(1.42);
}
</style>
<div class="diamond">
	<img src="http://csssecrets.io/images/adamcatlace.jpg" />
</div>
```

```html
<!-- --applet-- Clipping path solution -->
<style>
img {
	max-width: 250px;
	margin: 20px;
	-webkit-clip-path: polygon(50% 0, 100% 50%, 50% 100%, 0 50%);
	clip-path: polygon(50% 0, 100% 50%, 50% 100%, 0 50%);
	transition: 1s;
}

img:hover {
	-webkit-clip-path: polygon(0 0, 100% 0, 100% 100%, 0 100%);
	clip-path: polygon(0 0, 100% 0, 100% 100%, 0 100%);
}
</style>
<img src="http://csssecrets.io/images/adamcatlace.jpg" />
<img src="http://csssecrets.io/images/adam-sleeping.jpg" />
```

### Cutout corners
