package generator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestModuleCreator_Create 测试创建新模块
// 注意：此测试依赖项目模板，跳过以避免文件系统依赖
func TestModuleCreator_Create(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// TestModuleCreator_CreateExisting 测试创建已存在的模块
func TestModuleCreator_CreateExisting(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// TestModuleCloner_Clone 测试克隆模块
func TestModuleCloner_Clone(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// TestModuleCloner_CloneNonExistent 测试克隆不存在的模块
func TestModuleCloner_CloneNonExistent(t *testing.T) {
	ctx := context.Background()
	cloner := NewModuleCloner(ctx)

	// 克隆不存在的模块应该失败
	err := cloner.Clone("nonexistent999999", "target999999")
	assert.Error(t, err, "克隆不存在的模块应该返回错误")
	assert.Contains(t, err.Error(), "不存在", "错误消息应该提及源模块不存在")

	t.Logf("✓ 克隆不存在模块测试通过，错误: %v", err)
}

// TestModuleCloner_CloneToExisting 测试克隆到已存在的目标
func TestModuleCloner_CloneToExisting(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// TestModuleExporterImporter_RoundTrip 测试模块导出导入完整流程
func TestModuleExporterImporter_RoundTrip(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// TestModuleExporter_ExportNonExistent 测试导出不存在的模块
func TestModuleExporter_ExportNonExistent(t *testing.T) {
	ctx := context.Background()
	exporter := NewModuleExporter(ctx)

	// 导出不存在的模块应该失败
	_, err := exporter.Export("nonexistent999999")
	assert.Error(t, err, "导出不存在的模块应该返回错误")

	t.Logf("✓ 导出不存在模块测试通过，错误: %v", err)
}

// TestModuleImporter_ImportNonExistent 测试导入不存在的zip文件
func TestModuleImporter_ImportNonExistent(t *testing.T) {
	ctx := context.Background()
	importer := NewModuleImporter(ctx)

	// 导入不存在的zip文件应该失败
	_, err := importer.Import("nonexistent999999.zip")
	assert.Error(t, err, "导入不存在的文件应该返回错误")
	assert.Contains(t, err.Error(), "不存在", "错误消息应该提及文件不存在")

	t.Logf("✓ 导入不存在文件测试通过，错误: %v", err)
}

// TestModuleImporter_ImportToExisting 测试导入已存在的模块
func TestModuleImporter_ImportToExisting(t *testing.T) {
	t.Skip("跳过：需要完整的模板环境，推荐使用集成测试")
}

// BenchmarkModuleCreator_Create 性能测试：创建模块
func BenchmarkModuleCreator_Create(b *testing.B) {
	b.Skip("跳过：需要完整的模板环境")
}

// BenchmarkModuleExporter_Export 性能测试：导出模块
func BenchmarkModuleExporter_Export(b *testing.B) {
	b.Skip("跳过：需要完整的模板环境")
}
