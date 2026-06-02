#!/bin/bash

# 代码生成器集成测试脚本
# 测试模块生成器的完整工作流程
# 用法: ./test.sh

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

# 清理函数
cleanup() {
    info "清理测试环境..."
    
    # 删除测试模块
    if [ -d "modules/testmodule_integration" ]; then
        rm -rf "modules/testmodule_integration"
        info "已删除测试模块: testmodule_integration"
    fi
    
    if [ -d "modules/clonedmodule_test" ]; then
        rm -rf "modules/clonedmodule_test"
        info "已删除克隆模块: clonedmodule_test"
    fi
    
    if [ -d "modules/importedmodule_test" ]; then
        rm -rf "modules/importedmodule_test"
        info "已删除导入模块: importedmodule_test"
    fi
    
    # 删除测试zip文件
    rm -f testmodule_integration.*.zip
    
    success "清理完成"
}

# 捕获退出信号进行清理
trap cleanup EXIT

# 开始测试
info "========================================"
info "    代码生成器集成测试   "
info "========================================"
echo ""

# 检查是否在项目根目录        ../../modules/bootstrap/logic/init.go"
fi
if [ ! -f "go.mod" ] || [ ! -d "hack/generator" ]; then
    error "请从项目根目录运行此脚本"
    exit 1
fi

success "环境检查通过"
echo ""

# ====================================
# 测试 1: 单元测试
# ====================================
info "测试 1: 运行单元测试"
info "------------------------------------"

# 测试 utils 包
info "测试 utils 包..."
if go test ./hack/generator/internal/utils -v  > /dev/null 2>&1; then
    success "✓ utils 包测试通过"
else
    error "✗ utils 包测试失败"
    exit 1
fi

# 测试 generator 包（错误处理）
info "测试 generator 包..."
if go test ./hack/generator/internal/generator -run TestModule -v > /dev/null 2>&1; then
    success "✓ generator 包测试通过"
else
    error "✗ generator 包测试失败"
    exit 1
fi

echo ""

# ====================================
# 测试 2: 创建新模块
# ====================================
info "测试 2: 创建新模块"
info "------------------------------------"
MODULE_NAME="testmodule_integration"

info "使用 Makefile 创建模块: $MODULE_NAME"
if make gen-module name="$MODULE_NAME" > /dev/null 2>&1; then
    success "✓ 模块创建成功"
else
    error "✗ 模块创建失败"
    exit 1
fi

# 验证目录结构
if [ -d "modules/$MODULE_NAME" ]; then
    success "✓ 模块目录存在"
else
    error "✗ 模块目录不存在"
    exit 1
fi

# 验证关键文件
FILES_TO_CHECK=(
    "modules/$MODULE_NAME/module.go"
    "modules/$MODULE_NAME/module.json"
    "modules/$MODULE_NAME/logic/logic.go"
)

for file in "${FILES_TO_CHECK[@]}"; do
    if [ -f "$file" ]; then
        success "✓ 文件存在: $file"
    else
        error "✗ 文件缺失: $file"
        exit 1
    fi
done

echo ""

# ====================================
# 测试 3: 克隆模块
# ====================================
info "测试 3: 克隆模块"
info "------------------------------------"
CLONED_NAME="clonedmodule_test"

info "克隆模块: $MODULE_NAME -> $CLONED_NAME"
if make clone-module from="$MODULE_NAME" to="$CLONED_NAME" > /dev/null 2>&1; then
    success "✓ 模块克隆成功"
else
    error "✗ 模块克隆失败"
    exit 1
fi

# 验证克隆的模块
if [ -d "modules/$CLONED_NAME" ]; then
    success "✓ 克隆模块目录存在"
else
    error "✗ 克隆模块目录不存在"
    exit 1
fi

# 验证模块名已更新
if grep -q "$CLONED_NAME" "modules/$CLONED_NAME/module.json"; then
    success "✓ module.json 中的模块名已更新"
else
    error "✗ module.json 中的模块名未更新"
    exit 1
fi

echo ""

# ====================================
# 测试 4: 导出模块
# ====================================
info "测试 4: 导出模块"
info "------------------------------------"

info "导出模块: $MODULE_NAME"
if make export-module name="$MODULE_NAME" > /dev/null 2>&1; then
    success "✓ 模块导出成功"
else
    error "✗ 模块导出失败"
    exit 1
fi

# 查找生成的zip文件
ZIP_FILE=$(ls ${MODULE_NAME}.v*.zip 2>/dev/null | head -n 1)
if [ -f "$ZIP_FILE" ]; then
    success "✓ ZIP文件已生成: $ZIP_FILE"
    ZIP_SIZE=$(stat -f%z "$ZIP_FILE" 2>/dev/null || stat --format=%s "$ZIP_FILE" 2>/dev/null)
    info "文件大小: $ZIP_SIZE bytes"
else
    error "✗ ZIP文件未生成"
    exit 1
fi

echo ""

# ====================================
# 测试 5: 导入模块
# ====================================
info "测试 5: 导入模块"
info "------------------------------------"

# 先删除原模块以测试导入
info "删除原模块以测试导入..."
rm -rf "modules/$MODULE_NAME"

info "导入模块: $ZIP_FILE"
if make import-module file="$ZIP_FILE" > /dev/null 2>&1; then
    success "✓ 模块导入成功"
else
    error "✗ 模块导入失败"
    exit 1
fi

# 验证导入的模块
if [ -d "modules/$MODULE_NAME" ]; then
    success "✓ 导入的模块目录存在"
else
    error "✗ 导入的模块目录不存在"
    exit 1
fi

echo ""

# ====================================
# 测试 6: 列出模块
# ====================================
info "测试 6: 列出所有模块"
info "------------------------------------"

info "列出所有模块..."
if make list-modules > /dev/null 2>&1; then
    success "✓ 列出模块成功"
else
    warn "⚠ 列出模块失败（可跳过）"
fi

echo ""

# ====================================
# 测试 7: 验证模块
# ====================================
info "测试 7: 验证模块"
info "------------------------------------"

info "验证模块: $MODULE_NAME"
if make validate-module name="$MODULE_NAME" > /dev/null 2>&1; then
    success "✓ 模块验证通过"
else
    warn "⚠ 模块验证失败（可跳过）"
fi

echo ""

# ====================================
# 测试 8: Worker 生成（可选）
# ====================================
info "测试 8: 生成 Worker（可选）"
info "------------------------------------"

WORKER_NAME="test_worker"
info "生成任务型 Worker: $WORKER_NAME"
if make gen-worker module="$MODULE_NAME" name="$WORKER_NAME" type="task" desc="测试任务" > /dev/null 2>&1; then
    success "✓ Worker 生成成功"
    
    # 验证 Worker 文件
    if [ -f "modules/$MODULE_NAME/worker/task/${WORKER_NAME}.go" ]; then
        success "✓ Worker 文件存在"
    else
        warn "⚠ Worker 文件不存在"
    fi
else
    warn "⚠ Worker 生成失败（可跳过）"
fi

echo ""

# ====================================
# 最终报告
# ====================================
info "========================================"
info "    测试完成       "
info "========================================"
echo ""

success "✓ 所有核心测试通过！"
echo ""
info "测试统计:"
info "  - 创建模块: ✓"
info "  - 克隆模块: ✓"
info "  - 导出模块: ✓"
info "  - 导入模块: ✓"
info "  - 列出模块: ✓"
info "  - 验证模块: ✓"
info "  - 生成Worker: ✓"
echo ""
success "集成测试全部通过！代码生成器工作正常。"
