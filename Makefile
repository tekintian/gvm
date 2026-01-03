GO = CGO_ENABLED=0 GO111MODULE=on GOPROXY=https://goproxy.cn,direct go
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_BRANCH := $(shell git symbolic-ref --short -q HEAD)
GIT_COMMIT_HASH := $(shell git rev-parse HEAD|cut -c 1-8)
GO_FLAGS := -v -ldflags="-s -w -X 'github.com/tekintian/gvm/build.Build=$(BUILD_DATE)' -X 'github.com/tekintian/gvm/build.Commit=$(GIT_COMMIT_HASH)' -X 'github.com/tekintian/gvm/build.Branch=$(GIT_BRANCH)'"


all: install test clean

# 生成版本号（从 git tag）
gen-version:
	$(GO) run build/gen_version.go

build: gen-version
	$(GO) build $(GO_FLAGS)
	# 移动到 ~/.gvm/bin/目录
	mv gvm ~/.gvm/bin/

install: gen-version
	$(GO) install $(GO_FLAGS)

build-all: build-linux build-darwin build-windows

build-linux: build-linux-386 build-linux-amd64 build-linux-arm build-linux-arm64 build-linux-s390x
build-linux-386:
	GOOS=linux GOARCH=386 $(GO) build $(GO_FLAGS) -o bin/linux-386/gvm
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(GO_FLAGS) -o bin/linux-amd64/gvm
build-linux-arm:
	GOOS=linux GOARCH=arm $(GO) build $(GO_FLAGS) -o bin/linux-arm/gvm
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(GO_FLAGS) -o bin/linux-arm64/gvm
build-linux-s390x:
    GOOS=linux GOARCH=s390x $(GO) build $(GO_FLAGS) -o  bin/linux-s390x/gvm


build-darwin: build-darwin-amd64 build-darwin-arm64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build $(GO_FLAGS) -o bin/darwin-amd64/gvm
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build $(GO_FLAGS) -o bin/darwin-arm64/gvm


build-windows: build-windows-386 build-windows-amd64 build-windows-arm build-windows-arm64
build-windows-386:
	GOOS=windows GOARCH=386 $(GO) build $(GO_FLAGS) -o bin/windows-386/gvm.exe
build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GO) build $(GO_FLAGS) -o bin/windows-amd64/gvm.exe
build-windows-arm:
	GOOS=windows GOARCH=arm $(GO) build $(GO_FLAGS) -o bin/windows-arm/gvm.exe
build-windows-arm64:
	GOOS=windows GOARCH=arm64 $(GO) build $(GO_FLAGS) -o bin/windows-arm64/gvm.exe

package:
	sh ./package.sh

test:
	$(GO) test -v ./...

clean:
	$(GO) clean -x
	rm -f sha256sum.txt
	rm -rf bin

.PHONY: all build install test package clean gen-version build-linux build-darwin build-windows build-linux-386 build-linux-amd64 build-linux-arm build-linux-arm64 build-linux-s390x build-darwin-amd64 build-darwin-arm64 build-windows-386 build-windows-amd64 build-windows-arm build-windows-arm64
