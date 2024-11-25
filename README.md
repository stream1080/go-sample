[![Build Status](https://github.com/stream1080/go-sample/actions/workflows/go.yml/badge.svg)](https://github.com/stream1080/go-sample/actions?query=branch%3Amaster) 
[![Go Report Card](https://goreportcard.com/badge/github.com/stream1080/go-sample)](https://goreportcard.com/report/github.com/stream1080/go-sample)
[![Go Reference](https://pkg.go.dev/badge/github.com/stream1080/go-sample.svg)](https://pkg.go.dev/github.com/stream1080/go-sample)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/stream1080/go-sample)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/stream1080/go-sample)
![license](https://img.shields.io/github/license/stream1080/go-sample)

# go-sample
基于 `gin` 的 `api` 服务端脚手架。 `gin` 框架在 `Web` 开发中是相当受欢迎的，但是 `gin` 也是一个轻量级 `Web` 框架，并不能像其他语言比如 `Java` 的 `Spring` 框架具有丰富的生态和标准，在实际开发中需要自己设计和添加一些额外的能力，来完善应用.

本项目布局为传统的MVC模式，参考了行业流行框架,适用于大部分业务api服务端开发。

## 项目特性
- 优雅停机实现，停机时清理资源；
- 使用 `go-env` 定义项目配置文件；
- 集成 `gorm` 和 `MySQL` 进行数据持久化；
- 提供了部分 `demo` 实现，可以按照 `demo` 在项目中直接使用；
- 整合 `redis` 换成，开箱即用，实现分布式缓存功能；
- 整合 `zap` 日志组件，完善日志输出；
- 集成 `jwt` 组件，提供 `demo` 代码，自定义授权失败成功等的响应格式，跟全局 `api` 响应格式统一；
- 增加 `md5`, `bcrypt` 和 `uuid` 生成工具包；
- 应用统一封装响应格式，参照行业内主流项目规范；
- 项目全局错误码封装；
- 添加 `Makefile` 文件，可以使用 `make` 命令进行编译，打包。

## 项目结构
```
├─config
├─controller
├─docs
├─global
├─middlewares
├─models
├─pkg
│  ├─consistenthash
│  ├─encrypt
│  ├─jwt
│  ├─lock
│  ├─response
│  ├─utils
│  └─uuid
└─router
```