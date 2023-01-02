# Source Map

JavaScript 脚本变得越来越复杂，大部分源码都需要经过转换才能投入到生产环境，常见的源码转换主要有 3 种情况
1. 压缩，减小体积
1. 多个文件合并，减少 HTTP 请求数
1. 其他语言编译成 JavaScript

Source Map 就是一个信息文件，里面存储着位置信息，也就是说转换后的代码的每一个位置所对应的转换前的位置

```js
{
  version : 3,
  // file 文件名
  file: "out.js",
  // sourceRoot 源码根目录
  sourceRoot : "",
  // sources 源码文件
  sources: ["foo.js", "bar.js"],
  // names 转换前的变量名
  names: ["a", "b"],
  // mappings 位置映射
  mappings: "AAgBC,SAAQ,CAAEA;AAAEA",
  // sourcesContent 每个 sources 对应的源码内容
  sourcesContent: ['const a = 1; console.log(a)', 'const b = 2; console.log(b)']
}
```
各种调试工具一般都支持 sourcemap 的解析，只需要在文件末尾加上:
```js
//@ sourceMappingURL=/path/to/source.js.map
```

sourcemap 的用途:
+ 调试的时候会使用 sourcemap ，这样可以直接在源码打断点调试
+ 线上报错的时候会使用 sourcemap 来映射到源码，我们会把 sourcemap 单独上传 sentry 等错误收集平台
+ 生成的类型也能通过 sourcemap 关联到对应的源码中的定义

它的生成可以通过 source-map 包的 api，而 mapping 的位置来源可能是源码 parse 后的 AST 中的位置信息和打印代码时计算出的位置信息的关联。

## mappings 属性
map 文件的 mappings 属性是一个很长的字符串，它分为三层:
1. **行对应**，以分号表示，每个分号对应转换后源码的一行。所以第一个分号前的内容就对应源码的第一行，以此类推
1. **位置对应**，以逗号表示，每个逗号对应转换后源码的一个位置。所以第一个都好前的内容就对应该行源码的第一个位置，以此类推
1. **位置转换**，以 [VLQ](https://en.wikipedia.org/wiki/Variable-length_quantity) 编码表示，代表该位置对应的转换前的源码位置

## 位置对应的原理

每个位置使用 5 位，表示 5 个字段，从左边算起:
1. 第一位，表示这个位置在转换后代码的第几列
1. 第二位，表示这个位置属于 sources 属性中的哪一个文件
1. 第三位，表示这个位置属于转换前代码的第几行
1. 第四位，表示这个位置属于转换前代码的第几列
1. 第五位，表示这个位置属于 names 属性中的哪一个变量

> 注意，由于 VLQ 编码是变长的，所以每一位可以由多个字符构成

比如，如果某个位置是 AAAAA ，由于 A 在 VLQ 编码中表示 0 ，因此这个位置的 5 个位实际上都是 0.它的意思是该位置在转换后代码的第 0 列，对应 sources 属性中的第 0 个文件，属于转换前代码的第 0 行第 0 列，对应 names 属性中的第 0 个变量