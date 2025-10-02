[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/errgenkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/errgenkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/errgenkratos)](https://pkg.go.dev/github.com/orzkratos/errgenkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/errgenkratos/main.svg)](https://coveralls.io/github/orzkratos/errgenkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/errgenkratos.svg)](https://github.com/orzkratos/errgenkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/errgenkratos)](https://goreportcard.com/report/github.com/orzkratos/errgenkratos)

# errgenkratos

Kratos 错误处理代码生成插件，减少样板代码并支持嵌套枚举和自定义字段扩展。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🚀 **简洁代码生成**: 使用泛型函数减少重复代码
🎯 **单行函数体**: 每个生成的函数体内仅有一行代码
🔧 **可配置元数据**: 支持添加包含枚举值的自定义元数据字段
📦 **简单集成**: 完全兼容 protoc-gen-go-errors 的替代方案
⚡ **高性能**: 编译时泛型实现最小运行时开销
🔄 **嵌套枚举支持**: 从消息体内部枚举生成错误代码

## 安装

### 安装插件
```bash
go install github.com/orzkratos/errgenkratos/cmd/protoc-gen-orzkratos-errors@latest
```

## 基本使用

### 标准生成（仅顶级枚举）
```bash
protoc --orzkratos-errors_out=paths=source_relative:./your_output_dir your_proto_files.proto
```

### 开启嵌套枚举支持
```bash
protoc --orzkratos-errors_out=include_nested=true,paths=source_relative:./your_output_dir your_proto_files.proto
```

## 示例

查看 [examples](internal/examples/) DIR 了解详细用法：
- [example1](internal/examples/example1/) - 基础顶级枚举错误生成
- [example2](internal/examples/example2/) - 单文件嵌套枚举支持
- [example3](internal/examples/example3/) - 多文件项目与服务定义

## 生成代码示例

### 顶级枚举（标准）
```go
// 生成自：enum ErrorReason { USER_NOT_FOUND = 1 [(errors.code) = 404]; }

// 用户未找到
func IsUserNotFound(err error) bool {
    return errgenkratos.IsError(err, ErrorReason_USER_NOT_FOUND, 404)
}

// 用户未找到
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
    return errgenkratos.NewError(404, ErrorReason_USER_NOT_FOUND, format, args...)
}
```

### 嵌套枚举（新特性）
```go
// 生成自：message GetUserResponse { enum ErrorReason { USER_NOT_FOUND = 1 [(errors.code) = 404]; } }

// 用户未找到
func IsGetUserResponseUserNotFound(err error) bool {
    return errgenkratos.IsError(err, GetUserResponse_USER_NOT_FOUND, 404)
}

// 用户未找到
func ErrorGetUserResponseUserNotFound(format string, args ...interface{}) *errors.Error {
    return errgenkratos.NewError(404, GetUserResponse_USER_NOT_FOUND, format, args...)
}
```

### 设置自定义元数据字段

在业务代码中配置元数据字段：

```go
import "github.com/orzkratos/errgenkratos"

func init() {
    // 设置自定义元数据字段名
    errgenkratos.SetMetadataFieldName("codeNum")
}
```

## 环境要求

- Go 1.25.0+
- Protocol Buffers 编译器 (protoc)
- Kratos v2.8.0+

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/errgenkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/errgenkratos)