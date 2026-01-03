package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func use(ctx *cli.Context) (err error) {
	vname := ctx.Args().First()
	if vname == "" {
		return cli.ShowSubcommandHelp(ctx)
	}
	targetV := filepath.Join(versionsDir, vname)

	if finfo, err := os.Stat(targetV); err != nil || !finfo.IsDir() {
		return cli.Exit(fmt.Sprintf("[gvm] The %q version does not exist, please install it first.", vname), 1)
	}

	_ = os.Remove(goroot)

	if err = mkSymlink(targetV, goroot); err != nil {
		return cli.Exit(errstring(err), 1)
	}
	if output, err := exec.Command(filepath.Join(goroot, "bin", "go"), "version").Output(); err == nil {
		fmt.Print(string(output))
	}

	// 检查当前 go 命令是否由 gvm 管理
	checkCurrentGoBinary()

	return nil
}

// checkCurrentGoBinary 检查当前使用的 go 命令是否由 gvm 管理
func checkCurrentGoBinary() {
	// 查找当前 go 可执行文件路径
	goPath, err := exec.LookPath("go")
	if err != nil {
		return
	}

	// 解析真实路径（处理符号链接）
	realGoPath, err := filepath.EvalSymlinks(goPath)
	if err != nil {
		return
	}

	// 检查是否在 gvm 管理的目录下
	gvmGoPath := filepath.Join(goroot, "bin", "go")
	gvmGoPath, err = filepath.EvalSymlinks(gvmGoPath)
	if err != nil {
		return
	}

	// 如果当前 go 不是 gvm 管理的版本
	if filepath.Clean(realGoPath) != filepath.Clean(gvmGoPath) {
		fmt.Printf("\n[gvm] Warning: Current 'go' command is not managed by gvm.\n")
		fmt.Printf("       Current go: %s\n", goPath)
		fmt.Printf("       GVM go:     %s\n\n", filepath.Join(goroot, "bin", "go"))
		fmt.Printf("Quick fix:\n")
		fmt.Printf("  1. Add to shell profile: export PATH=\"%s:$PATH\"\n", filepath.Join(goroot, "bin"))
		fmt.Printf("  2. Then run: source ~/.bashrc (or ~/.zshrc)\n\n")
	}
}
