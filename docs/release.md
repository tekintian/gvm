# 发布流程

本文档介绍 gvm 的发布流程和 CI/CD 配置。

## 发布前检查清单

- [ ] 更新版本号（通过 git tag）
- [ ] 运行测试套件
- [ ] 更新 CHANGELOG.md
- [ ] 验证文档完整性
- [ ] 检查 CI/CD 配置

## 发布步骤

### 1. 准备发布

```bash
# 确保在 main/master 分支
git checkout main

# 拉取最新代码
git pull origin main

# 运行测试
make test

# 本地构建验证
make build
./gvm --version
```

### 2. 更新版本号

```bash
# 创建版本 tag（格式：vX.Y.Z）
git tag v1.6.0

# 推送 tag 到远程
git push origin v1.6.0
```

**注意事项**：
- 使用语义化版本规范（[Semantic Versioning](https://semver.org/)）
- Tag 格式：`v` + 主版本号`.`次版本号`.`修订号
- 示例：`v1.6.0`、`v2.0.0-rc1`

### 3. 自动发布

推送 tag 后，GitHub Actions 会自动执行：

#### 3.1 触发 Release Workflow

```yaml
# .github/workflows/release.yml
on:
  push:
    tags:
      - 'v*'
```

#### 3.2 执行步骤

1. **检出代码**：包含所有历史和 tags
2. **运行测试**：确保代码质量
3. **生成版本号**：从 git tag 自动生成
4. **构建多平台二进制**：
   - Linux: 386, amd64, arm, arm64, s390x
   - macOS: amd64, arm64
   - Windows: 386, amd64, arm, arm64
5. **生成校验和**：SHA256 和 MD5
6. **创建 GitHub Release**：自动发布
7. **上传所有文件**：二进制文件 + 校验和文件

#### 3.3 生成的文件

```
bin/
├── gvm1.6.0.linux-386
├── gvm1.6.0.linux-amd64
├── gvm1.6.0.linux-arm
├── gvm1.6.0.linux-arm64
├── gvm1.6.0.linux-s390x
├── gvm1.6.0.darwin-amd64
├── gvm1.6.0.darwin-arm64
├── gvm1.6.0.windows-386.exe
├── gvm1.6.0.windows-amd64.exe
├── gvm1.6.0.windows-arm.exe
├── gvm1.6.0.windows-arm64.exe
├── sha256sum.txt
└── md5sum.txt
```

### 4. 验证发布

访问 GitHub Release 页面：https://github.com/tekintian/gvm/releases

检查项：
- [ ] 版本号正确
- [ ] 所有平台二进制文件已上传
- [ ] 校验和文件完整
- [ ] Release 说明正确

### 5. 通知用户

- [ ] 更新 README.md
- [ ] 发布 GitHub Announcement
- [ ] 更新社交媒体
- [ ] 通知用户更新

## 平台兼容性

### macOS 兼容性

| 架构 | 最低系统版本 | 说明 |
|------|-------------|------|
| ARM64 | macOS 11.0+ | Apple Silicon (M1/M2/M3) |
| amd64 | macOS 10.13+ | Intel x86_64 |

编译配置：
```yaml
# macOS ARM64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w"

# macOS Intel
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w"
```

### Linux 兼容性

支持架构：386, amd64, arm, arm64, s390x

所有主流 Linux 发行版均支持。

### Windows 兼容性

| 架构 | 最低系统版本 |
|------|-------------|
| 386, amd64, arm, arm64 | Windows Vista+ |

## CI/CD 配置

### Release Workflow

文件：`.github/workflows/release.yml`

关键配置：
```yaml
name: Release

on:
  push:
    tags:
      - 'v*'
  workflow_dispatch:  # 支持手动触发

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0  # 确保检出所有 tags

      - name: Generate version from git tag
        run: |
          go run app_build/gen_version.go

      - name: Build
        run: |
          # 多平台编译...

      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          generate_release_notes: true
```

### CI Workflow

文件：`.github/workflows/ci.yml`

自动运行：
- 代码推送时
- Pull Request 时

测试矩阵：
- Ubuntu (Go 1.20, 1.21, 1.22)
- macOS ARM64 (Go 1.21)
- macOS Intel (Go 1.21)
- Windows (Go 1.20, 1.21, 1.22)

## 回滚发布

如果发现问题需要回滚：

```bash
# 1. 删除远程 tag
git push origin :refs/tags/v1.6.0

# 2. 删除本地 tag
git tag -d v1.6.0

# 3. 在 GitHub Release 页面删除发布
# 访问 https://github.com/tekintian/gvm/releases

# 4. 创建新的 patch 版本
git tag v1.6.1
git push origin v1.6.1
```

## 用户更新流程

### 自动更新

```bash
# 用户执行
gvm update

# 流程：
# 1. 检查 GitHub API 获取最新 release
# 2. 比较版本号
# 3. 下载对应平台的二进制文件
# 4. 验证 SHA256 校验和
# 5. 自动替换当前二进制文件
```

### 手动更新

用户也可以从 GitHub Releases 下载最新版本：

```bash
# 下载
wget https://github.com/tekintian/gvm/releases/download/v1.6.0/gvm1.6.0.linux-amd64

# 验证校验和
sha256sum -c sha256sum.txt

# 替换二进制文件
sudo mv gvm1.6.0.linux-amd64 /usr/local/bin/gvm
```

## Docker 发布

### 构建镜像

```bash
# 构建并推送
docker build -t tekintian/gvm:v1.6.0 .
docker push tekintian/gvm:v1.6.0

# 更新 latest 标签
docker tag tekintian/gvm:v1.6.0 tekintian/gvm:latest
docker push tekintian/gvm:latest
```

### Docker Compose

```yaml
# docker-compose.yml
version: '3.8'
services:
  gvm:
    image: tekintian/gvm:latest
    volumes:
      - gvm-data:/root/.gvm
```

## 相关文档

- [版本管理](./version-management.md)
- [CI/CD 配置](../.github/workflows/README.md)
- [Docker 支持](../docker-compose.yml)
