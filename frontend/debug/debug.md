# debug

调试: 代码在某个平台运行，把运行时的状态通过某种方式暴露出来，传递给开发工具做 UI 展示和交互，辅助开发者排查问题、梳理流程、了解代码运行状态等

调试工具四要素:
1. frontend
1. backend
1. 调试协议
1. 信道

## Chrome DevTools 原理

Chrome DevTools 分为两部分，backend 和 frontend :
+ backend 和 Chrome 集成，负责把 Chrome 的网页运行时状态通过调试协议暴露出来
+ frontend 是独立的，负责对接调试协议，做 UI 的展示和交互

两者之间的调试协议叫做 Chrome DevTools Protocol ，简称 CDP

传输协议数据的方式叫做信道 (message channel)，有很多种，例如 Chrome DevTools 嵌入在 Chrome 里时，两者通过全局的函数通信；当 Chrome DevTools 远程调试某个目标的代码时，两者通过 WebSocket 通信

[chrome-devtools](./drawio/chrome-devtools.drawio){link-type="drawio"}


## VSCode Debugger 原理

VSCode Debugger 的原理和 Chrome DevTools 差不多，也是分为 frontend/backend/CDP 这几部分，只不过它多了一层适配器协议。因为 VSCode 并不是 JS 专用编辑器，自然不能和某一种语言的调试协议深度耦合，所以多了一层适配器层

好处:
1. VSCode Debugger 可以用同一套 UI 和逻辑来调试各种语言的代码，只要对接不同的 Debug Adapter 做协议转换
1. 其他编辑器也可以用这个 Debug Adapter Protocol 实现调试，这样就能直接复用 VSCode 的各种语言的 Debug Adapter 了

[vscode-debugger](./drawio/vscode-debugger.drawio){link-type="drawio"}


## Vue/React DevTools

