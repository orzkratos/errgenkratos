// Package erkgen provides public API for protoc plugin code generation.
// This package wraps the internal codegen functionality for external use.
//
// erkgen 提供用于 protoc 插件代码生成的公共 API。
// 此包包装了内部 codegen 功能以供外部使用。
package erkgen

import (
	"github.com/orzkratos/errgenkratos/internal/codegen"
	"google.golang.org/protobuf/compiler/protogen"
)

// Config holds generation configuration
// Config 保存生成配置
type Config struct {
	GeneratorName string // Generator name for comments // 注释中的生成器名称
}

// GenerateFile generates a _errors.pb.go file containing kratos errors definitions
// Wraps the internal codegen.GenerateFile function for public access
//
// GenerateFile 生成包含 kratos 错误定义的 _errors.pb.go 文件
// 包装内部 codegen.GenerateFile 函数以供公共访问
func GenerateFile(gen *protogen.Plugin, file *protogen.File, config Config) *protogen.GeneratedFile {
	return codegen.GenerateFile(gen, file, codegen.Config{
		GeneratorName: config.GeneratorName,
	})
}
