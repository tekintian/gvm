# gvm

<p align="center">
  <img src="https://img.shields.io/badge/go-v1.21+-00ADD8E?style=flat&logo=go" alt="Go 版本">
  <img src="https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-4E71EE?style=flat" alt="平台">
  <img src="https://img.shields.io/badge/license-MIT-green?style=flat" alt="许可证">
  <img src="https://img.shields.io/github/v/release/tekintian/gvm?color=brightgreen" alt="发布版本">
</p>

<div align="center">

**🚀 为开发者打造的轻量级 Go 版本管理工具**

</div>

<p align="center">
  <a href="#-安装">安装</a> •
  <a href="#-快速开始">快速开始</a> •
  <a href="#-特性">特性</a> •
  <a href="#-常见问题">常见问题</a> •
  <a href="https://github.com/tekintian/gvm/tree/master/docs">📚 文档</a>
</p>

---

**注意：** `master` 分支可能处于开发之中并**非稳定版本**，请通过 tag 下载稳定版本的源代码，或通过 [release](https://github.com/tekintian/gvm/releases) 下载已编译的二进制可执行文件。

`gvm` 是一个适用于 Linux、macOS 和 Windows 的轻量级命令行工具，提供便捷的多版本 [Go](https://golang.org/) 环境管理和切换。🎯

## ✨ 亮点

- 🚀 **极速安装** - 快速下载和安装 Go 版本
- 🔄 **无缝切换** - 一条命令切换 Go 版本
- 🌐 **跨平台支持** - 支持 Linux、macOS 和 Windows
- 🔒 **安全可靠** - 自动校验所有下载文件的 SHA256
- 💾 **节省空间** - 一次下载，多处使用
- 🛠️ **智能特性** - macOS 版本兼容性检测、PATH 配置验证等

## 📦 安装

### 一键安装（推荐）

**Linux / macOS**

```bash
curl -sSL https://raw.githubusercontent.com/tekintian/gvm/master/install.sh | bash
```

**Windows**

```powershell
# 从 releases 页面下载并运行安装程序
# 访问 https://github.com/tekintian/gvm/releases
```

### 手动安装

1. 从 [releases](https://github.com/tekintian/gvm/releases) 下载二进制压缩包
2. 将压缩包解压至 `PATH` 环境变量目录下（如 `/usr/local/bin`）
3. 添加到 shell 配置文件（`~/.bashrc` 或 `~/.zshrc`）：

```bash
cat >> ~/.bashrc <<'EOF'
export GOROOT="${HOME}/.gvm/go"
export PATH="${HOME}/.gvm/go/bin:$PATH"
export GVM_MIRROR=https://golang.google.cn/dl/
EOF
```

4. 重新加载 shell 配置：

```bash
source ~/.bashrc
```

## 🚀 快速开始

```bash
# 列出可用的稳定版本
gvm ls-remote stable

# 安装 Go 1.21.0
gvm install 1.21.0

# 列出已安装的版本
gvm ls

# 切换到 Go 1.21.0
gvm use 1.21.0

# 验证版本
go version
```

## 📚 使用说明

### 列出可用的 Go 版本

```bash
# 仅列出稳定版本
gvm ls-remote stable

# 列出所有可用版本（包括 beta、rc、已归档）
gvm ls-remote

# 列出特定通道
gvm ls-remote rc
gvm ls-remote beta
```

### 安装 Go 版本

```bash
# 安装指定版本
gvm install 1.21.0

# 使用简短版本安装（gvm 会自动匹配 1.21.0）
gvm install 1.21
```

### 管理已安装的版本

```bash
# 列出所有已安装的版本
gvm ls

# 切换到指定版本
gvm use 1.21.0

# 卸载版本
gvm uninstall 1.21.0
```

### 更新 gvm

```bash
# 检查并更新
gvm update
```

## ⚙️ 特性

- ✅ 支持列出可供安装的 Go 版本号
- ✅ 支持列出已安装的 Go 版本号
- ✅ 支持在本地安装多个 Go 版本
- ✅ 支持卸载已安装的 Go 版本
- ✅ 支持在已安装的 Go 版本之间自由切换
- ✅ 支持软件自我更新（>= 1.3.0）
- ✅ 跨平台支持（Linux、macOS、Windows）
- ✅ macOS 版本兼容性检测
- ✅ 重复版本过滤
- ✅ 完整版本号显示（如 1.24.0）
- ✅ 简短版本号匹配（如 `1.24` 可匹配 `1.24.0`）
- ✅ PATH 配置验证
- ✅ 网络代理支持
- ✅ 校验和验证

## 🌐 平台支持

| 平台 | 架构 | 最低系统版本 |
|------|-------|-------------|
| Linux | 386, amd64, arm, arm64, s390x | 所有发行版 |
| macOS | amd64, arm64 | 10.13+ (Intel), 11.0+ (ARM) |
| Windows | 386, amd64, arm, arm64 | Vista+ |

## 🔧 配置

### 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `GVM_MIRROR` | Go 下载镜像站点 | - |
| `GVM_HOME` | 自定义 gvm 家目录 | `~/.gvm` |
| `GVM_EXPERIMENTAL` | 启用实验特性 | `false` |

### 镜像站点

如果您在中国或需要更快的下载速度，可以使用以下镜像：

```bash
export GVM_MIRROR=https://golang.google.cn/dl/
# 或
export GVM_MIRROR=https://studygolang.com/dl
# 或
export GVM_MIRROR=https://mirrors.aliyun.com/golang/
```

## 🐳 Docker 支持

```bash
# 拉取镜像
docker pull tekintian/gvm:latest

# 在容器中运行 gvm
docker run --rm -v gvm-data:/root/.gvm tekintian/gvm ls-remote

# 使用 docker-compose
docker-compose up -d
docker-compose exec gvm gvm install 1.21.0
```

## ❓ 常见问题

### 为什么选择 gvm 而不是官方 Go 安装程序？

- **多版本管理** - 轻松安装和切换多个 Go 版本
- **快速切换** - 一条命令即可切换版本
- **自动配置** - 自动管理 PATH 和符号链接
- **跨平台一致** - 在 Linux、macOS 和 Windows 上体验一致

### 环境变量 `GVM_MIRROR` 有什么作用？

由于中国大陆无法自由访问 Go 官网，导致查询及下载 Go 版本都变得困难。因此可以通过该环境变量指定一个或多个镜像站点（多个镜像站点之间使用英文逗号分隔），gvm 将从该站点查询、下载可用的 Go 版本。已知的可用镜像站点如下：

- Go 官方镜像站点：https://golang.google.cn/dl/
- Go 语言中文网：https://studygolang.com/dl
- 阿里云开源镜像站点：https://mirrors.aliyun.com/golang/

### 环境变量 `GVM_EXPERIMENTAL` 有什么作用？

当该环境变量的值为 `true` 时，将**开启所有的实验特性**。

### 环境变量 `GVM_HOME` 有什么作用？

按照惯例，gvm 默认会将 `~/.gvm` 目录作为其家目录。若想自定义家目录（Windows 用户需求强烈），可使用该环境变量切换到其他家目录。由于**该特性还属于实验特性**，需要先开启实验特性开关 `GVM_EXPERIMENTAL=true` 才能生效。特别注意，该方案并不十分完美，因此才将其归类为实验特性，详见 [#18](https://github.com/tekintian/gvm/issues/18)。

### macOS 系统下安装 Go 版本，gvm 抛出 `[gvm] Installation package not found` 字样的错误提示，是什么原因？

Go 官方在 **1.16** 版本中才[加入了对 ARM 架构的 macOS 系统的支持](https://go.dev/doc/go1.16#darwin)。因此，ARM 架构的 macOS 系统下均无法安装 1.15 及以下版本的 Go 安装包。若尝试安装这些版本，gvm 会抛出 `[gvm] Installation package not found` 的错误信息。

### 是否支持网络代理？

支持。可在 `HTTP_PROXY`、`HTTPS_PROXY`、`http_proxy`、`https_proxy` 等环境变量中设置网络代理地址。

### 支持哪些 Windows 版本？

因为 `gvm` 的实现上依赖于`符号链接`，因此操作系统必须是 `Windows Vista` 及以上版本。

### Windows 版本安装以后不生效？

这有可能是因为没有把下载安装的加入到 `$PATH` 的缘故，需要手动将 `$PATH` 纳入到用户的环境变量中。为了方便起见，可以使用项目中的 `path.ps1` 的 PowerShell 脚本运行然后重新启动计算机即可。

### 支持源代码编译安装吗？

不支持。

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 鸣谢

感谢这些优秀的项目提供的宝贵思路：

- [nvm](https://github.com/nvm-sh/nvm) - Node 版本管理器
- [n](https://github.com/tj/n) - Node 版本管理
- [rvm](https://github.com/rvm/rvm) - Ruby 版本管理器

---

<div align="center">

**由 [TekinTian](https://github.com/tekintian) 用 ❤️ 打造**

[⬆ 返回顶部](#gvm)

</div>
