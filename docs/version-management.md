# 版本管理

本文档介绍 gvm 的版本号管理机制。

## 概述

gvm 使用 `go generate` 机制自动从 git tag 生成版本号，这是 Go 社区推荐的最佳实践。

## 工作原理

### 版本号来源

版本号按以下优先级获取：

1. **Git Tag**（发布版本）
   - 如果存在 git tag（如 `v1.6.0`），则使用 tag 版本
   - 自动移除 `v` 前缀（`v1.6.0` → `1.6.0`）

2. **Git Commit**（开发中）
   - 如果没有 tag，使用 `0.0.0-<commit-hash>` 格式
   - 示例：`0.0.0-e93d469`

3. **默认值**（无 git）
   - 如果没有 git 环境，使用 `0.0.0-dev`

### 文件结构

```
build/
├── build.go           # 主文件，包含 //go:generate 指令
├── gen_version.go     # 版本号生成器
└── version.go         # 自动生成（gitignore）
```

## 使用方法

### 本地开发

#### 使用 Makefile（推荐）

```bash
# 生成版本号
make gen-version

# 构建（自动生成版本号）
make build

# 安装（自动生成版本号）
make install
```

#### 使用 Go 命令

```bash
# 手动生成版本号
go run build/gen_version.go

# 编译
go build
```

### 发布新版本

```bash
# 1. 标记新版本
git tag v1.6.0
git push origin v1.6.0

# 2. CI/CD 自动执行
# - 检出代码（包含 tag）
# - 生成版本号：go run build/gen_version.go
# - 编译多平台二进制文件
# - 生成校验和（sha256sum.txt）
# - 创建 GitHub Release
# - 上传二进制文件

# 3. 用户更新
gvm update
```

### 查看版本

```bash
# 查看当前版本
gvm --version

# 输出示例
gvm version 1.6.0
```

## CI/CD 集成

### GitHub Actions

在 `.github/workflows/release.yml` 中：

```yaml
- name: Generate version from git tag
  run: |
    go run build/gen_version.go

- name: Build
  run: |
    CGO_ENABLED=0 go build -ldflags="-s -w"
```

生成的二进制文件版本号将自动与 git tag 一致。

## 版本号格式

### 语义化版本

gvm 遵循 [语义化版本 2.0.0](https://semver.org/lang/zh-CN/) 规范：

- **主版本号**（Major）：不兼容的 API 修改
- **次版本号**（Minor）：向下兼容的功能性新增
- **修订号**（Patch）：向下兼容的问题修正

### 版本号示例

| Git Tag | 生成的版本号 | 说明 |
|---------|-------------|------|
| `v1.6.0` | `1.6.0` | 正式发布版本 |
| - | `0.0.0-e93d469` | 开发中版本 |
| - | `0.0.0-dev` | 无 git 环境 |

## 自更新机制

gvm 支持自动更新到最新版本：

### 更新流程

1. 检查 GitHub API 获取最新 release
2. 使用 `semver` 库比较版本号
3. 下载对应平台的二进制文件
4. 验证 SHA256 校验和
5. 使用 `go-update` 自动替换当前二进制文件

### 使用方法

```bash
# 检查并更新到最新版本
gvm update
```

## 注意事项

1. **首次构建**：需要先运行 `make gen-version` 或 `go run build/gen_version.go`
2. **生成的文件**：`build/version.go` 已加入 `.gitignore`，不会提交到 git
3. **版本一致性**：确保 git tag 与预期版本号一致
4. **CI/CD**：使用 `git checkout --fetch-depth=0` 确保检出所有 tags

## 常见问题

### Q: 为什么不直接在代码中硬编码版本号？

A: 使用 `go generate` 的优势：
- 自动从 git tag 获取，避免忘记更新
- 不污染 git 历史（生成的文件被忽略）
- 符合 Go 社区最佳实践
- CI/CD 自动化更简单

### Q: 如何验证版本号是否正确？

A: 
```bash
# 查看 git tag
git describe --tags --abbrev=0

# 查看生成的版本号
cat build/version.go

# 查看编译后的二进制版本
./gvm --version
```

### Q: 开发中如何避免版本号频繁变化？

A: 开发中版本号基于 commit hash，每次 commit 都会变化。这是正常的，确保每个构建都有唯一标识。

### Q: CI/CD 中如何确保使用正确的版本？

A: 在 release workflow 中：
- 使用 `fetch-depth: 0` 检出所有历史（包括 tags）
- 推送 tag 后触发 workflow
- 版本号自动从 tag 生成

## 相关文档

- [Go Generate 官方文档](https://go.dev/blog/generate)
- [语义化版本规范](https://semver.org/lang/zh-CN/)
- [发布流程](./release.md)
