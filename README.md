# gvm


**注意：**`master`分支可能处于开发之中并**非稳定版本**，请通过 tag 下载稳定版本的源代码，或通过[release](https://github.com/tekintian/gvm/releases)下载已编译的二进制可执行文件。

`g`是一个 Linux、macOS、Windows 下的命令行工具，可以提供一个便捷的多版本 [go](https://golang.org/) 环境的管理和切换。


## 特性

- 支持列出可供安装的 go 版本号
- 支持列出已安装的 go 版本号
- 支持在本地安装多个 go 版本
- 支持卸载已安装的 go 版本
- 支持在已安装的 go 版本之间自由切换
- 支持软件自我更新（>= 1.3.0）

## 安装

### 自动化安装

- Linux/macOS（适用于 bash、zsh）

  ```shell
  # 建议安装前清空`GOROOT`、`GOBIN`等环境变量
  $ curl -sSL https://raw.githubusercontent.com/tekintian/gvm/master/install.sh | bash
  $ echo "unalias gvm" >> ~/.bashrc # 可选。若其他程序（如'git'）使用了'gvm'作为别名。
  $ source ~/.bashrc # 或者 source ~/.zshrc
  ```

### 手动安装

- 下载[release](https://github.com/tekintian/gvm/releases)的二进制压缩包
- 将压缩包解压至`PATH`环境变量目录下（如`/usr/local/bin`）
- 编辑 shell 环境配置文件（如`~/.bashrc`、`~/.zshrc`...）

  ```shell
  $ cat>>~/.bashrc<<'EOF'
  export GOROOT="${HOME}/.gvm/go"
  export PATH="${HOME}/.gvm/go/bin:$PATH"
  export GVM_MIRROR=https://golang.google.cn/dl/
  EOF
  ```

- 启用环境变量
  ```shell
  $ source ~/.bashrc # 或source ~/.zshrc
  ```

## 使用

查询当前可供安装的`stable`状态的 go 版本

```shell
$ gvm ls-remote stable
  1.13.15
  1.14.7
```

安装目标 go 版本`1.14.7`

```shell
$ gvm install 1.14.7
Downloading 100% |███████████████| (119/119 MB, 9.939 MB/s) [12s:0s]
Computing checksum with SHA256
Checksums matched
Now using go1.14.7
```

查询已安装的 go 版本

```shell
$ gvm ls
  1.7.6
  1.11.13
  1.12.17
  1.13.15
  1.14.6
* 1.14.7
```

查询可供安装的所有 go 版本

```shell
$ gvm ls-remote
  1
  1.2.2
  1.3
  1.3.1
  ...    // 省略若干版本
  1.14.5
  1.14.6
* 1.14.7
  1.15rc1
```

切换到另一个已安装的 go 版本

```shell
$ gvm use 1.14.6
go version go1.14.6 darwin/amd64
```

卸载一个已安装的 go 版本

```shell
$ gvm uninstall 1.14.7
Uninstalled go1.14.7
```

更新 gvm 软件本身

```shell
$ gvm update
A new version of gvm(v1.2.2) is available
Downloading 100% |███████████████| (3.7/3.7 MB, 2.358 MB/s)
Computing checksum with SHA256
Checksums matched
Update completed
```

## FAQ

- 环境变量`GVM_MIRROR`有什么作用？

  由于中国大陆无法自由访问 Golang 官网，导致查询及下载 go 版本都变得困难，因此可以通过该环境变量指定一个或多个镜像站点（多个镜像站点之间使用英文逗号分隔），gvm 将从该站点查询、下载可用的 go 版本。已知的可用镜像站点如下：

  - Go 官方镜像站点：https://golang.google.cn/dl/
  - Go 语言中文网：https://studygolang.com/dl
  - 阿里云开源镜像站点：https://mirrors.aliyun.com/golang/

- 环境变量`GVM_EXPERIMENTAL`有什么作用？

  当该环境变量的值为`true`时，将**开启所有的实验特性**。

- 环境变量`GVM_HOME`有什么作用？

  按照惯例，gvm 默认会将`~/.gvm`目录作为其家目录。若想自定义家目录（Windows 用户需求强烈），可使用该环境变量切换到其他家目录。由于**该特性还属于实验特性**，需要先开启实验特性开关`GVM_EXPERIMENTAL=true`才能生效。特别注意，该方案并不十分完美，因此才将其归类为实验特性，详见[#18](https://github.com/tekintian/gvm/issues/18)。

- macOS 系统下安装 go 版本，gvm 抛出`[gvm] Installation package not found`字样的错误提示，是什么原因？

  Go 官方在**1.16**版本中才[加入了对 ARM 架构的 macOS 系统的支持](https://go.dev/doc/go1.16#darwin)。因此，ARM 架构的 macOS 系统下均无法安装 1.15 及以下的版本的 go 安装包。若尝试安装这些版本，gvm 会抛出`[gvm] Installation package not found`的错误信息。

- 是否支持网络代理？

  支持。可在`HTTP_PROXY`、`HTTPS_PROXY`、`http_proxy`、`https_proxy`等环境变量中设置网络代理地址。

- 支持哪些 Windows 版本？

  因为`g`的实现上依赖于`符号链接`，因此操作系统必须是`Windows Vista`及以上版本。

- Windows 版本安装以后不生效？

  这有可能是因为没有把下载安装的加入到 `$Path` 的缘故，需要手动将 `$Path` 纳入到用户的环境变量中。为了方便起见，可以使用项目中的 `path.ps1` 的 PowerShell 脚本运行然后重新启动计算机即可。

- 支持源代码编译安装吗？

  不支持

## 鸣谢

感谢[nvm](https://github.com/nvm-sh/nvm)、[n](https://github.com/tj/n)、[rvm](https://github.com/rvm/rvm)等工具提供的宝贵思路。


# goquery 使用
https://github.com/PuerkitoBio/goquery

~~~go
// 这个doc选择的内容为id="archive"元素下面的样式为.expanded下的div
	// <div id="archive"><div class="expanded""><div xxx></div></div></div>
	// divsArchive := c.doc.Find("#archive").ChildrenFiltered(".expanded").Find("div")
	// 查询id为archive下面的所有div的class为toggle的div元素
divsArchive := c.doc.Find("#archive").Find("div.toggle")

// 查询所有id为archive的元素下面的所有class为toggle的div元素
doc.Find("#archive").Find("div.toggle").Each(func(i int, div *goquery.Selection) {

}
~~~


