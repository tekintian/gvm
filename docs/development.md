# 开发指南

本文档介绍如何参与 gvm 的开发。

## 开发环境

### 前置要求

- Go 1.21 或更高版本
- Git
- Make（可选，用于构建）

### 克隆项目

```bash
# 克隆仓库
git clone https://github.com/tekintian/gvm.git
cd gvm

# 安装依赖
go mod download
```

### 项目结构

```
gvm/
├── build/              # 版本管理
│   ├── build.go       # 构建信息和版本
│   └── gen_version.go # 版本号生成器
├── cli/               # 命令行界面
│   ├── cli.go        # 命令定义
│   ├── install.go    # 安装命令
│   ├── use.go        # 切换版本命令
│   ├── update.go     # 更新命令
│   └── ...
├── collector/         # Go 版本收集
│   ├── collector.go  # 版本收集器
│   ├── stable.html   # 稳定版模板
│   └── archived.html # 归档版模板
├── pkg/              # 核心包
│   ├── checksum/    # 校验和计算
│   ├── http/        # HTTP 下载
│   ├── sdk/github/  # GitHub API
│   └── ...
├── version/          # 版本管理逻辑
│   ├── version.go   # 版本操作
│   └── ...
├── main.go          # 入口文件
├── Makefile         # 构建脚本
└── .github/         # CI/CD 配置
```

## 构建和运行

### 使用 Makefile

```bash
# 生成版本号
make gen-version

# 构建
make build

# 安装
make install

# 运行测试
make test

# 构建所有平台
make build-all
```

### 使用 Go 命令

```bash
# 生成版本号
go run build/gen_version.go

# 构建
go build

# 运行
./gvm --help

# 运行测试
go test -v ./...

# 运行特定包的测试
go test -v ./cli
```

## 开发流程

### 1. 分支管理

```bash
# 创建功能分支
git checkout -b feature/new-feature

# 或者修复分支
git checkout -b fix/bug-fix
```

### 2. 编码

```bash
# 编辑代码
# ...

# 运行测试
go test -v ./...

# 运行 linter
golangci-lint run
```

### 3. 提交

```bash
# 查看变更
git status

# 添加文件
git add .

# 提交
git commit -m "feat: add new feature"

# 推送到远程
git push origin feature/new-feature
```

### 4. Pull Request

1. 在 GitHub 上创建 Pull Request
2. 填写 PR 模板
3. 等待 Code Review
4. 根据反馈修改
5. 合并到主分支

## 测试

### 运行测试

```bash
# 运行所有测试
go test -v ./...

# 运行测试并生成覆盖率
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# 查看覆盖率
go tool cover -html=coverage.out
```

### 添加测试

```go
// 示例：在 cli/install_test.go 中
package cli

import (
    "testing"
)

func TestInstall(t *testing.T) {
    // 测试逻辑
}
```

## 调试

### 使用 VSCode

创建 `.vscode/launch.json`：

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug ls-remote",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}",
      "args": ["ls-remote"]
    },
    {
      "name": "Debug install",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}",
      "args": ["install", "1.21.0"]
    }
  ]
}
```

### 使用 Delve

```bash
# 安装 delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 调试
dlv debug -- ls-remote
```

## 代码规范

### 命名约定

- 包名：小写，无下划线
- 文件名：小写，下划线分隔
- 导出函数：大写开头（PascalCase）
- 私有函数：小写开头（camelCase）

### 注释

```go
// Package cli provides command-line interface for gvm.
package cli

// Install downloads and installs a Go version.
func Install(version string) error {
    // 实现逻辑
}
```

### 错误处理

```go
// 使用自定义错误类型
return fmt.Errorf("failed to download: %w", err)

// 使用 pkg/errors 包（如果有）
return errors.Wrap(err, "failed to parse version")
```

## 版本管理

### 开发中版本

开发时，版本号会自动基于 commit hash：

```bash
# 无 tag 时
go run build/gen_version.go
# 输出：Generated build/version.go with version: 0.0.0-e93d469

./gvm --version
# 输出：gvm version 0.0.0-e93d469
```

### 发布版本

创建 tag 后版本号会自动使用 tag：

```bash
git tag v1.6.0
go run build/gen_version.go
# 输出：Generated build/version.go with version: 1.6.0
```

详见 [版本管理文档](./version-management.md)。

## 贡献指南

### Pull Request 流程

1. Fork 项目
2. 创建功能分支
3. 实现功能并添加测试
4. 确保所有测试通过
5. 更新相关文档
6. 提交 Pull Request

### Commit Message 规范

遵循 [Conventional Commits](https://www.conventionalcommits.org/)：

```
<type>(<scope>): <subject>

<body>

<footer>
```

类型：
- `feat`: 新功能
- `fix`: 修复
- `docs`: 文档
- `style`: 代码格式
- `refactor`: 重构
- `test`: 测试
- `chore`: 构建/工具

示例：
```
feat(install): add support for installing beta versions

Closes #123
```

## 常见任务

### 添加新命令

1. 在 `cli/` 目录创建命令文件（如 `newcmd.go`）
2. 实现命令函数
3. 在 `cli/commands.go` 注册命令

```go
// cli/newcmd.go
package cli

import (
    "github.com/urfave/cli/v2"
)

func newCmd(c *cli.Context) error {
    // 实现逻辑
    return nil
}

// cli/commands.go
commands = append(commands, &cli.Command{
    Name:  "newcmd",
    Usage: "Description of new command",
    Action: newCmd,
})
```

### 添加新平台支持

1. 更新 `collector` 以支持新平台的版本收集
2. 更新 Makefile 添加构建目标
3. 更新 CI/CD 配置
4. 更新文档

### 修复 Bug

1. 创建修复分支：`git checkout -b fix/bug-name`
2. 添加测试用例验证 bug
3. 修复代码
4. 确保所有测试通过
5. 提交 PR

## 性能优化

### 性能分析

```bash
# CPU 性能分析
go test -cpuprofile=cpu.prof ./...

# 内存分析
go test -memprofile=mem.prof ./...

# 查看分析结果
go tool pprof cpu.prof
```

### 优化建议

- 避免不必要的内存分配
- 使用 sync.Pool 复用对象
- 减少系统调用
- 使用缓存减少重复计算

## 相关文档

- [版本管理](./version-management.md)
- [发布流程](./release.md)
- [项目 README](../README.md)
