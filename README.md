# gvm

<p align="center">
  <img src="https://img.shields.io/badge/go-v1.21+-00ADD8E?style=flat&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-4E71EE?style=flat" alt="Platform">
  <img src="https://img.shields.io/badge/license-MIT-green?style=flat" alt="License">
  <img src="https://img.shields.io/github/v/release/tekintian/gvm?color=brightgreen" alt="Release">
</p>

<div align="center">

**🚀 The Fast and Simple Go Version Manager for Developers**

</div>

<p align="center">
  <a href="#-installation">Installation</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-features">Features</a> •
  <a href="#-faq">FAQ</a> •
  <a href="https://github.com/tekintian/gvm/tree/master/docs">📚 Documentation</a>
</p>

---

**Note:** The `master` branch may be under development and is **not a stable version**. Please download the source code of a stable version via a tag, or download the compiled binary from [releases](https://github.com/tekintian/gvm/releases).

`gvm` is a lightweight command-line tool for Linux, macOS, and Windows that provides convenient management and switching of multiple versions of [Go](https://golang.org/). 🎯

## ✨ Highlights

- 🚀 **Blazing Fast** - Quick installation and switching between Go versions
- 🔄 **Seamless Switching** - Switch between Go versions with a single command
- 🌐 **Cross-Platform** - Works on Linux, macOS, and Windows
- 🔒 **Secure** - Automatic checksum verification for all downloads
- 💾 **Space Efficient** - Download once, install multiple versions
- 🛠️ **Smart Features** - macOS version compatibility detection, PATH validation, and more

## 📦 Installation

### One-Line Installation (Recommended)

**Linux / macOS**

```bash
curl -sSL https://raw.githubusercontent.com/tekintian/gvm/master/install.sh | bash
```

**Windows**

```powershell
# Download from releases and run the installer
# See https://github.com/tekintian/gvm/releases
```

### Manual Installation

1. Download the binary from [releases](https://github.com/tekintian/gvm/releases)
2. Extract to a directory in your `PATH` (e.g., `/usr/local/bin`)
3. Add to your shell profile (`~/.bashrc` or `~/.zshrc`):

```bash
cat >> ~/.bashrc <<'EOF'
export GOROOT="${HOME}/.gvm/go"
export PATH="${HOME}/.gvm/go/bin:$PATH"
export GVM_MIRROR=https://golang.google.cn/dl/
EOF
```

4. Reload your shell:

```bash
source ~/.bashrc
```

## 🚀 Quick Start

```bash
# List available stable versions
gvm ls-remote stable

# Install Go 1.21.0
gvm install 1.21.0

# List installed versions
gvm ls

# Switch to Go 1.21.0
gvm use 1.21.0

# Verify the version
go version
```

## 📚 Usage

### List Available Go Versions

```bash
# List stable versions only
gvm ls-remote stable

# List all available versions (including beta, rc, archived)
gvm ls-remote

# List specific channel
gvm ls-remote rc
gvm ls-remote beta
```

### Install Go Versions

```bash
# Install specific version
gvm install 1.21.0

# Install with short version (gvm will auto-match 1.21.0)
gvm install 1.21
```

### Manage Installed Versions

```bash
# List all installed versions
gvm ls

# Switch to a version
gvm use 1.21.0

# Uninstall a version
gvm uninstall 1.21.0
```

### Update gvm

```bash
# Check for updates
gvm update
```

## ⚙️ Features

- ✅ List available Go versions for installation
- ✅ List installed Go versions
- ✅ Install multiple Go versions locally
- ✅ Uninstall installed Go versions
- ✅ Switch freely between installed Go versions
- ✅ Self-update support (>= 1.3.0)
- ✅ Cross-platform support (Linux, macOS, Windows)
- ✅ macOS version compatibility detection
- ✅ Duplicate version filtering
- ✅ Complete version display (e.g., 1.24.0)
- ✅ Short version matching (e.g., `1.24` matches `1.24.0`)
- ✅ PATH configuration validation
- ✅ Network proxy support
- ✅ Checksum verification

## 🌐 Platform Support

| Platform | Architectures | Minimum OS Version |
|----------|---------------|-------------------|
| Linux | 386, amd64, arm, arm64, s390x | All distributions |
| macOS | amd64, arm64 | 10.13+ (Intel), 11.0+ (ARM) |
| Windows | 386, amd64, arm, arm64 | Vista+ |

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|----------|
| `GVM_MIRROR` | Mirror site for downloading Go | - |
| `GVM_HOME` | Custom home directory for gvm | `~/.gvm` |
| `GVM_EXPERIMENTAL` | Enable experimental features | `false` |

### Mirror Sites

If you're in China or need faster downloads, use one of these mirrors:

```bash
export GVM_MIRROR=https://golang.google.cn/dl/
# or
export GVM_MIRROR=https://studygolang.com/dl
# or
export GVM_MIRROR=https://mirrors.aliyun.com/golang/
```

## 🐳 Docker Support

```bash
# Pull the image
docker pull tekintian/gvm:latest

# Run gvm in container
docker run --rm -v gvm-data:/root/.gvm tekintian/gvm ls-remote

# Using docker-compose
docker-compose up -d
docker-compose exec gvm gvm install 1.21.0
```

## ❓ FAQ

### Why use gvm instead of official Go installer?

- **Multiple versions**: Install and switch between multiple Go versions easily
- **Fast switching**: Switch versions with a single command
- **No manual setup**: Automatic PATH configuration and symlink management
- **Cross-platform**: Works consistently across Linux, macOS, and Windows

### What is the purpose of the `GVM_MIRROR` environment variable?

Due to access restrictions to the official Go website in some regions, querying and downloading Go versions can be difficult. This environment variable allows you to specify one or more mirror sites (separated by commas). gvm will query and download available Go versions from these sites.

### What is the purpose of the `GVM_EXPERIMENTAL` environment variable?

When this environment variable is set to `true`, **all experimental features are enabled**.

### What is the purpose of the `GVM_HOME` environment variable?

By convention, gvm uses the `~/.gvm` directory as its home directory. If you want to customize the home directory, you can use this environment variable to switch to another directory. Since **this feature is still experimental**, you need to enable `GVM_EXPERIMENTAL=true` for it to take effect.

### When installing Go on macOS, gvm throws `[gvm] Installation package not found`?

The official Go website only [added support for ARM architecture macOS systems in version 1.16](https://go.dev/doc/go1.16#darwin). Therefore, Go installation packages version 1.15 and below cannot be installed on ARM architecture macOS systems.

### Does it support network proxy?

Yes. You can set network proxy addresses in `HTTP_PROXY`, `HTTPS_PROXY`, `http_proxy`, `https_proxy` environment variables.

### Which Windows versions are supported?

Since the implementation of `gvm` depends on `symbolic links`, the operating system must be `Windows Vista` or higher.

### Windows version doesn't work after installation?

This may be because the downloaded installation was not added to `$PATH`. You need to manually include `$PATH` in the user's environment variables. For convenience, run the PowerShell script `path.ps1` provided in the project and restart your computer.

### Does it support building from source?

No, it does not support building from source.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

Thanks to these amazing projects for providing valuable ideas:

- [nvm](https://github.com/nvm-sh/nvm) - Node Version Manager
- [n](https://github.com/tj/n) - Node version management
- [rvm](https://github.com/rvm/rvm) - Ruby Version Manager

---

<div align="center">

**Made with ❤️ by [TekinTian](https://github.com/tekintian)**

[⬆ Back to Top](#gvm)

</div>

