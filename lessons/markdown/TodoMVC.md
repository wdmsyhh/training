# TodoMVC 需求

## 功能需求

- 支持用户名密码登录，对数据进行用户级别隔离。
- 可以输入内容然后回车添加要做事项。
- 添加的内容在输入框下方列表显示。
- 已完成的可以在前面圆形单选框单选标记。
- 点击文本框左边可以全部标记。
- 点击列表下方 All 可以展示所有事项。
- 点击列表下方 Active 可以显示未标记的事项。
- 点击列表下方 Completed 可以显示已完成事项。
- 当有标记的事项时，列表下方右边显示 Clear completed 按钮。
- 点击 Clear completed 按钮可以删除已完成事项。
- 鼠标移动到某个事项右边显示删除按钮，点击可以删除该事项。

## 非功能性需求

- 后端使用指定 MVC 框架 + MongoDB + Redis，前端使用 Vue.js + vue-antd。
- 页面包含一个文本输入框。
- 文本输入框上方是 todos 字样。
- 要求相应及时。
