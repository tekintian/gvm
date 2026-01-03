# gvm 文档

欢迎来到 gvm 文档中心。

## 快速导航

### 用户文档

- [README](../README.md) - 项目介绍和快速开始
- [中文 README](../README_CN.md) - 中文版文档

### 开发者文档

- [开发指南](./development.md) - 如何参与开发
- [版本管理](./version-management.md) - 版本号管理机制
- [发布流程](./release.md) - 发布新版本的步骤

### 架构文档

- [项目结构](#项目结构) - 代码组织方式
- [命令列表](#命令列表) - 所有可用命令
- [环境变量](#环境变量) - 配置选项

## 项目结构

```
gvm/
├── app_build/              # 构建和版本管理
│   ├── build.go       # 构建信息
│   └── gen_version.go # 版本号生成器
├── cli/               # 命令行界面
│   ├── cli.go        # 命令定义
│   ├── install.go    # 安装命令
│   ├── use.go        # 切换版本命令
│   ├── ls.go         # 列表命令
│   ├── uninstall.go  # 卸载命令
│   ├── update.go     # 更新命令
│   └── clean.go      # 清理命令
├── collector/         # Go 版本收集
│   ├── collector.go  # 版本收集器
│   ├── stable.html   # 稳定版 HTML
│   ├── unstable.html # 不稳定版 HTML
│   └── archived.html # 归档版 HTML
├── pkg/              # 核心包
│   ├── checksum/    # 校验和计算
│   ├── errors/      # 错误处理
│   ├── http/        # HTTP 下载
│   ├── sdk/github/  # GitHub API 集成
│   ├── sys/         # 系统信息
│   └── util/        # 工具函数
├── version/          # 版本管理逻辑
│   └── version.go   # 版本操作
├── main.go          # 入口文件
├── Makefile         # 构建脚本
├── go.mod           # Go 模块定义
├── go.sum           # 依赖锁定
├── Dockerfile       # Docker 镜像
├── docker-compose.yml # Docker Compose
└── .github/         # CI/CD 配置
    └── workflows/
        ├── ci.yml   # CI 工作流
        ├── release.yml # Release 工作流
        └── docker.yml # Docker 工作流
```

## 命令列表

| 命令 | 说明 | 示例 |
|------|------|------|
| `ls` | 列出已安装的版本 | `gvm ls` |
| `ls-remote` | 列出可用的版本 | `gvm ls-remote stable` |
| `install` | 安装指定版本 | `gvm install 1.21.0` |
| `use` | 切换到指定版本 | `gvm use 1.21.0` |
| `uninstall` | 卸载版本 | `gvm uninstall 1.21.0` |
| `update` | 更新 gvm 到最新版本 | `gvm update` |
| `clean` | 清理下载缓存 | `gvm clean` |

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `GVM_MIRROR` | Go 下载镜像站点 | - |
| `GVM_HOME` | 自定义 gvm 家目录 | `~/.gvm` |
| `GVM_EXPERIMENTAL` | 启用实验特性 | `false` |

### 镜像站点配置

```bash
# 官方镜像
export GVM_MIRROR=https://golang.google.cn/dl/

# Go 语言中文网
export GVM_MIRROR=https://studygolang.com/dl

# 阿里云镜像
export GVM_MIRROR=https://mirrors.aliyun.com/golang/
```

## 平台支持

### 支持的操作系统

| 操作系统 | 架构 | 最低版本 |
|---------|------|---------|
| Linux | 386, amd64, arm, arm64, s390x | 所有发行版 |
| macOS | amd64, arm64 | 10.13+ (Intel), 11.0+ (ARM) |
| Windows | 386, amd64, arm, arm64 | Vista+ |

### Docker 支持

```bash
# 使用 Docker
docker pull tekintian/gvm:latest
docker run --rm -v gvm-data:/root/.gvm tekintian/gvm ls-remote
```

## 开发资源

### 构建

```bash
# 使用 Makefile
make build

# 使用 Go
go run app_build/gen_version.go
go build
```

### 测试

```bash
# 运行所有测试
go test -v ./...

# 运行特定包的测试
go test -v ./cli

# 生成覆盖率报告
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Lint

```bash
# 运行 golangci-lint
golangci-lint run

# 或使用 CI
.github/workflows/ci.yml
```

## CI/CD

### GitHub Actions

- **CI 工作流**：每次提交和 PR 时运行测试
- **Release 工作流**：推送 tag 时自动发布
- **Docker 工作流**：构建并推送 Docker 镜像

### 发布流程

1. 创建 git tag：`git tag v1.6.0`
2. 推送 tag：`git push origin v1.6.0`
3. GitHub Actions 自动：
   - 运行测试
   - 生成版本号
   - 构建多平台二进制
   - 创建 GitHub Release
   - 上传所有文件

详见 [发布流程文档](./release.md)。

## 贡献

欢迎贡献！请查看 [开发指南](./development.md) 了解详细信息。

### 贡献步骤

1. Fork 项目
2. 创建功能分支
3. 提交变更
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License - 详见 [LICENSE](../LICENSE) 文件。

## 鸣谢

感谢这些优秀的项目提供的宝贵思路：

- [nvm](https://github.com/nvm-sh/nvm) - Node 版本管理器
- [n](https://github.com/tj/n) - Node 版本管理
- [rvm](https://github.com/rvm/rvm) - Ruby 版本管理器

## 相关链接

- [GitHub 仓库](https://github.com/tekintian/gvm)
- [Releases](https://github.com/tekintian/gvm/releases)
- [Issues](https://github.com/tekintian/gvm/issues)
- [Discussions](https://github.com/tekintian/gvm/discussions)
- [专业软件研发](https://dev.tekin.cn)

## 常见问题

详见 [FAQ](../README.md#faq)。

### 获取帮助

- 查看 [GitHub Issues](https://github.com/tekintian/gvm/issues)
- 使用 `gvm --help` 查看命令帮助
- 使用 `gvm <command> --help` 查看特定命令帮助

---

有问题？欢迎在 [GitHub Discussions](https://github.com/tekintian/gvm/discussions) 中讨论。
