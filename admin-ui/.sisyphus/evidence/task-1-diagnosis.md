# 诊断报告：Vue SFC "Failed to resolve extends base type" 错误

## 日期
2026-04-19

## 1. 当前错误状态

- **dev server**: 成功启动（VITE v8.0.0 ready, http://localhost:5999/）
- **错误数量**: 3 个 "Failed to resolve extends base type" 错误
- **错误级别**: Pre-transform error（编译时警告，不影响运行，但类型信息丢失）

### 出错文件

| 文件 | 出错行 | extends 基类型 |
|------|--------|---------------|
| `packages/effects/layouts/src/basic/menu/extra-menu.vue` | L12 `interface Props extends MenuProps` | `MenuProps` from `@vben-core/menu-ui` |
| `packages/effects/layouts/src/basic/menu/menu.vue` | L8 `interface Props extends MenuProps` | `MenuProps` from `@vben-core/menu-ui` |
| `packages/effects/layouts/src/basic/menu/mixed-menu.vue` | L13 `interface Props extends NormalMenuProps` | `NormalMenuProps` from `@vben-core/menu-ui` |

## 2. 版本对比表

| 依赖 | 本地实际安装版本 | 本地 catalog 声明 | 上游 catalog 声明 | 差异 |
|------|-----------------|------------------|------------------|------|
| TypeScript | 5.8.2 / 5.9.3（两个版本） | ^5.9.3 | ^6.0.2 | **大版本差异（5.x vs 6.x）** |
| Vue | 3.5.30 | ^3.5.30 | ^3.5.32 | 微小差异 |
| @vue/compiler-sfc | 3.5.30 | (跟随 vue) | (跟随 vue) | 相同 |
| @vitejs/plugin-vue | 5.2.4 / 6.0.5（两个版本） | ^6.0.5 | ^6.0.6 | 微小差异 |
| Vite | 5.4.21 / 8.0.0（两个版本） | ^8.0.0 | ^8.0.8 | 微小差异 |

**注意**：本地同时安装了两个 TypeScript 版本（5.8.2 和 5.9.3），这是因为 pnpm 的严格依赖隔离。

## 3. 配置对比表

### 3.1 tsconfig 配置

| 配置项 | 本地 `internal/tsconfig/base.json` | 上游 `internal/tsconfig/base.json` | 差异 |
|--------|-----------------------------------|-----------------------------------|------|
| 所有 compilerOptions | **完全一致** | **完全一致** | 无差异 |

| 配置项 | 本地 `internal/tsconfig/web.json` | 上游 `internal/tsconfig/web.json` | 差异 |
|--------|----------------------------------|----------------------------------|------|
| 所有 compilerOptions | **完全一致** | **完全一致** | 无差异 |

**结论：tsconfig 配置无差异。**

### 3.2 pnpm-workspace.yaml

| 配置项 | 本地 | 上游 | 差异 |
|--------|------|------|------|
| packages | 相同 | 相同 | 无 |
| overrides | 相同 | 相同 | 无 |
| catalog 依赖数量 | ~130+ | ~150+ | 上游更多（新增包） |
| **关键依赖版本** | | | |

### 3.3 .npmrc

| 配置项 | 本地 | 上游 | 差异 |
|--------|------|------|------|
| 所有设置 | **完全一致** | **完全一致** | 无差异 |

### 3.4 vite Vue 插件配置

| 配置项 | 本地 `plugins/index.ts` | 上游 `plugins/index.ts` | 差异 |
|--------|------------------------|------------------------|------|
| viteVue 配置 | `{ script: { defineModel: true } }` | `{ script: { defineModel: true } }` | 无差异 |

## 4. 根因分析

### 根因：**pnpm workspace 包链接问题**（配置/环境差异）

**不是单纯的版本差异。** 根因是一个隐蔽的 pnpm 模块解析问题。

### 详细根因链

```
1. @vben/tsconfig 包没有被链接到任何 node_modules/@vben/tsconfig 目录
   ↓
2. TypeScript parseJsonConfigFileContent 无法解析 extends: "@vben/tsconfig/web.json"
   → 报错: "File '@vben/tsconfig/web.json' not found."
   → moduleResolution 配置丢失，变为 undefined
   ↓
3. TS resolveModuleName 使用默认的 Node10 模式（而非 Bundler）
   ↓
4. Node10 模式不识别 package.json exports 的 "types" 条件
   → 解析 @vben-core/menu-ui 到 dist/index.mjs（而非 src/index.ts）
   ↓
5. @vue/compiler-sfc 的 parseFile 对 .mjs 文件返回空 AST（无类型信息）
   ↓
6. MenuProps / NormalMenuProps 类型无法被找到
   ↓
7. interface Props extends MenuProps → "Failed to resolve extends base type"
```

### 关键证据

1. **@vben/tsconfig 包不存在于 node_modules 中**：
   - 搜索了从 `packages/effects/layouts` 到根目录的所有 node_modules 路径
   - `@vben/tsconfig` 只存在于 `internal/tsconfig/` 源码目录，没有符号链接

2. **TS 无法解析 tsconfig extends 链**：
   ```
   TS 报错: "File '@vben/tsconfig/web.json' not found."
   → moduleResolution = undefined（应该为 Bundler = 100）
   ```

3. **模块解析结果对比**：
   - Node10 模式 → `dist/index.mjs`（无类型信息）
   - Bundler 模式 → `src/index.ts`（有类型信息）✓

4. **上游也有相同配置**（.npmrc、tsconfig、package.json exports），但上游可能因为 TypeScript 6.0 的行为差异或其他环境因素而能正常工作。

### 为什么只有 3 个文件受影响

只有这 3 个 .vue 文件使用了 `interface Props extends ExternalType` 模式，其中基类型来自**非相对路径的包导入**（`@vben-core/menu-ui`）。其他 .vue 文件：
- 使用内联类型字面量（不需要解析外部模块）
- 或只使用相对路径导入（compiler-sfc 直接解析文件路径，不经过 TS module resolution）

## 5. 修复方案建议

### 方案 A：确保 @vben/tsconfig 可被链接到 node_modules（推荐）

在根 `package.json` 的 `devDependencies` 中已有 `@vben/tsconfig`，但它没有被链接。
可能的原因是 root package.json 被视为 workspace root，其 devDependencies 不一定被链接。

**解决方法**：在每个使用 `extends: "@vben/tsconfig/..."` 的包的 `package.json` 中添加：
```json
{
  "devDependencies": {
    "@vben/tsconfig": "workspace:*"
  }
}
```

受影响的包包括：
- `packages/effects/layouts`
- `packages/@core/ui-kit/menu-ui`
- 以及其他所有使用 `@vben/tsconfig` 的包

### 方案 B：调整 tsconfig extends 为相对路径

将各包的 tsconfig.json 中的：
```json
{ "extends": "@vben/tsconfig/web.json" }
```
改为相对路径：
```json
{ "extends": "../../../internal/tsconfig/web.json" }
```
（路径根据包的层级深度调整）

### 方案 C：升级 TypeScript 到 6.x

TypeScript 6.x 可能在模块解析上有改进，使得即使 tsconfig extends 解析失败，也能通过其他方式正确处理 package.json exports。但这是最不确定的方案，需要验证。

## 6. 上游差异的关键发现

- **TypeScript 版本差异最大**（5.9.3 vs 6.0.2），但这不是直接原因
- **所有配置文件完全一致**（tsconfig、.npmrc、vite plugin 配置）
- **关键差异在于包链接状态**：@vben/tsconfig 在本地环境中没有被正确链接到 node_modules
- 上游可能因为以下原因之一而不受影响：
  1. TypeScript 6.0 的 parseJsonConfigFileContent 对 extends 解析有改进
  2. 上游的 pnpm 版本或配置不同，导致 workspace 包被正确链接
  3. 上游在安装时有不同的 hoisting 行为

## 7. 影响评估

- **运行时影响**：无（dev server 正常启动，页面可正常加载）
- **类型安全影响**：受影响组件的 extends 基类型属性变为 fallthrough attrs，可能导致 props 类型检查失效
- **HMR 影响**：Pre-transform error 可能在热更新时产生额外警告
- **构建影响**：构建可能通过但类型信息不完整
