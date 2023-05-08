# :cloud: 多云平台接口对接管理系统

本项目是一款基于 Go 语言编写的多云平台接口对接管理系统，旨在为用户提供一个便捷的管理云平台的工具。

## 功能特点

- :heavy_check_mark: 支持多种云平台接口对接，包括 AWS、Azure、Google Cloud Platform 等。
- :heavy_check_mark: 提供简洁易用的管理界面，方便用户快速上手。
- :heavy_check_mark: 支持云平台资源的管理和监控，包括虚拟机、存储、网络等。
- :heavy_check_mark: 支持定时任务和自动化运维，提高工作效率。

## 系统架构

本系统采用前后端分离的架构，前端采用 Vue.js 框架，后端采用 Go 语言编写，数据库采用 MySQL。

系统结构如下图所示：

![系统结构图](./images/structure.png)

## 快速开始

1. 克隆代码库

```shell
git clone https://github.com/oyjs1989/CloudTools
```

2. 编译前端代码

```shell
cd frontend
npm install
npm run build
```

3. 启动后端服务

```shell
cd ../backend
go run main.go
```

4. 访问系统

在浏览器中输入 http://localhost:8080，即可访问系统。

## 贡献指南

本项目欢迎开发者参与贡献，您可以通过以下方式参与：

- 提交 bug 报告或功能需求
- 提交代码贡献
- 参与文档编写

具体贡献指南请参考 [CONTRIBUTING.md](./CONTRIBUTING.md)。

## 许可证

本项目采用 MIT 许可证，详情请参考 [LICENSE](./LICENSE) 文件。