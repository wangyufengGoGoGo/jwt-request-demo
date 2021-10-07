# 该项目布局来自于[golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```markdown
/cmd 项目主要应用程序。每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，/cmd/myapp)。
/internal 私有应用程序和库代码。（代码不是可重用的或不希望其他人重用它）
/pkg 外部应用可以使用的库代码。（代码可以导入并在其他项目中使用）
/api OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件。
/web 特定于 Web 应用程序的组件:静态 Web 资产、服务器端模板和 SPAs。
/configs 配置文件模板或默认配置。
/init 系统初始化或进程管理配置。
/scripts 执行各种构建、安装、分析等操作的脚本。
/build 打包和持续集成。
/deployments 部署。
/test 额外的测试应用程序和测试数据。
/docs 设计和用户文档。
/tools 项目支持的工具。
/examples 你的应用程序或公共库示例。
/third_party 外部辅助工具，分叉代码和其他第三方工具(例如 Swagger UI)。
/githooks git hooks代码审查。
/asserts 与存储库一起使用的其他资产(图像、徽标等)。
/website 如果你不使用 Github 页面，则在这里放置项目的网站数据。
```