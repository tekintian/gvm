# GitHub Actions CI/CD

本项目使用 GitHub Actions 进行持续集成和持续部署。

## 工作流

### 1. CI (`.github/workflows/ci.yml`)

触发条件：
- 推送到 `master`, `main`, `develop` 分支
- 针对 `master`, `main`, `develop` 的 Pull Request

功能：
- 在多平台（Ubuntu, macOS, Windows）上运行测试
- 使用多个 Go 版本（1.20, 1.21, 1.22）测试兼容性
- 运行 golangci-lint 代码检查
- 编译验证二进制文件

### 2. Release (`.github/workflows/release.yml`)

触发条件：
- 推送标签（如 `v1.5.0`）
- 手动触发

功能：
- 跨平台编译 11 个二进制文件
- 生成 SHA256 和 MD5 校验和
- 自动创建 GitHub Release
- 上传构建产物

支持的平台和架构：
```
Linux:
  - linux-386
  - linux-amd64
  - linux-arm
  - linux-arm64
  - linux-s390x

macOS:
  - darwin-amd64
  - darwin-arm64

Windows:
  - windows-386.exe
  - windows-amd64.exe
  - windows-arm.exe
  - windows-arm64.exe
```

### 3. Docker Build (`.github/workflows/docker.yml`)

触发条件：
- 推送到 `master`, `main` 分支
- 推送标签（如 `v1.5.0`）
- 针对 `master`, `main` 的 Pull Request

功能：
- 构建多架构 Docker 镜像（amd64, arm64）
- 推送到 Docker Hub

## 使用方法

### 创建 Release

```bash
# 1. 创建并推送标签
git tag v1.5.0
git push origin v1.5.0

# 2. GitHub Actions 会自动：
#    - 编译所有平台的二进制文件
#    - 生成校验和
#    - 创建 GitHub Release
#    - 推送 Docker 镜像
```

### 手动触发 Release 工作流

1. 进入 GitHub 仓库的 Actions 页面
2. 选择 "Release" 工作流
3. 点击 "Run workflow"
4. 选择分支并确认

## Docker 使用

### 拉取镜像

```bash
docker pull tekintian/gvm:latest
```

### 使用 Docker 运行

```bash
# 列出远程 Go 版本
docker run --rm tekintian/gvm ls-remote

# 安装 Go 1.21.0
docker run --rm -v gvm-data:/root/.gvm tekintian/gvm install 1.21.0

# 查看已安装版本
docker run --rm -v gvm-data:/root/.gvm tekintian/gvm ls
```

### 使用 docker-compose

```bash
# 启动容器
docker-compose up -d

# 进入容器
docker-compose exec gvm sh

# 使用 gvm 命令
gvm ls-remote
gvm install 1.21.0
gvm use 1.21.0
```

## Secrets 配置

在 GitHub 仓库设置中添加以下 Secrets：

### Docker Hub（可选）

如需使用 Docker 工作流，请配置：
- `DOCKER_USERNAME`: Docker Hub 用户名
- `DOCKER_PASSWORD`: Docker Hub 密码或访问令牌

## 贡献指南

提交 PR 前，请确保：
1. 代码通过所有 CI 检查
2. 添加必要的测试用例
3. 遵循项目的代码规范
