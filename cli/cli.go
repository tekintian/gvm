package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
	app_build "github.com/tekintian/gvm/app_build"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	ghomeDir     string
	downloadsDir string
	versionsDir  string
	goroot       string
)

// Run 运行g命令行
func Run() {
	app := cli.NewApp()
	app.Name = "gvm"
	app.Usage = "Golang Version Manager"
	app.Version = app_build.Version()
	app.Copyright = "This is a golang version manager app. More info? visit https://dev.tekin.cn "
	app.Authors = []*cli.Author{{Name: "TekinTian", Email: "tekintian@gmail.com"}}

	app.Before = func(ctx *cli.Context) (err error) {
		ghomeDir = ghome()
		goroot = filepath.Join(ghomeDir, "go")
		downloadsDir = filepath.Join(ghomeDir, "downloads")
		if err = os.MkdirAll(downloadsDir, 0755); err != nil {
			return err
		}
		versionsDir = filepath.Join(ghomeDir, "versions")
		return os.MkdirAll(versionsDir, 0755)
	}
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func init() {
	cli.AppHelpTemplate = fmt.Sprintf(`NAME:
	{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

 USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Commands}} command{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

 VERSION:
	%s{{end}}{{end}}{{if .Description}}

 DESCRIPTION:
	{{.Description}}{{end}}{{if len .Authors}}

 AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
	{{range $index, $author := .Authors}}{{if $index}}
	{{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

 COMMANDS:{{range .VisibleCategories}}{{if .Name}}

	{{.Name}}:{{end}}{{range .VisibleCommands}}
	  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

 GLOBAL OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

 COPYRIGHT:
	{{.Copyright}}{{end}}
`, app_build.ShortVersion)
}

const (
	experimentalEnv = "GVM_EXPERIMENTAL"
	homeEnv         = "GVM_HOME"
	mirrorEnv       = "GVM_MIRROR"
)

// ghome 返回gvm根目录
func ghome() (dir string) {
	if experimental := os.Getenv(experimentalEnv); experimental == "true" {
		if dir = os.Getenv(homeEnv); dir != "" {
			return dir
		}
	}
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".gvm")
}

// inuse 返回当前的go版本号
func inuse(goroot string) (version string) {
	p, _ := os.Readlink(goroot)
	return filepath.Base(p)
}

// getMacOSVersion 获取 macOS 版本号
func getMacOSVersion() (major, minor int, err error) {
	if runtime.GOOS != "darwin" {
		return 0, 0, nil
	}
	cmd := exec.Command("sw_vers", "-productVersion")
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	version := strings.TrimSpace(string(output))
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("invalid macOS version: %s", version)
	}
	major, _ = strconv.Atoi(parts[0])
	minor, _ = strconv.Atoi(parts[1])
	return major, minor, nil
}

// checkGoCompatibility 检查 Go 版本兼容性
// 返回不兼容的原因，如果兼容则返回空字符串
func checkGoCompatibility(goVersion string) string {
	// 检查 macOS 兼容性
	if runtime.GOOS == "darwin" {
		major, minor, err := getMacOSVersion()
		if err != nil {
			return ""
		}
		// 获取 Go 版本号
		v, err := semver.NewVersion(goVersion)
		if err != nil {
			return ""
		}

		// macOS 10.11-10.12 不支持 Go 1.13+
		if major == 10 && minor <= 12 && minor >= 11 {
			if v.Major() == 1 && v.Minor() >= 13 {
				return fmt.Sprintf("requires macOS 10.13+, you are on %d.%d", major, minor)
			}
		}
		// macOS 10.10 及以下不支持 Go 1.7+
		if major == 10 && minor <= 10 {
			if v.Major() == 1 && v.Minor() >= 7 {
				return fmt.Sprintf("requires macOS 10.11+, you are on %d.%d", major, minor)
			}
		}
		// macOS 10.9 及以下不支持 Go 1.5+
		if major == 10 && minor <= 9 {
			if v.Major() == 1 && v.Minor() >= 5 {
				return fmt.Sprintf("requires macOS 10.10+, you are on %d.%d", major, minor)
			}
		}
		// macOS 10.6-10.8 不支持 Go 1.2+
		if major == 10 && minor <= 8 && minor >= 6 {
			if v.Major() == 1 && v.Minor() >= 2 {
				return fmt.Sprintf("requires macOS 10.9+, you are on %d.%d", major, minor)
			}
		}
		// macOS 10.5 及以下不支持 Go 1.0+
		if major == 10 && minor <= 5 {
			return fmt.Sprintf("requires macOS 10.6+, you are on %d.%d", major, minor)
		}
	}
	return ""
}

// render 渲染go版本列表
func render(curV string, items []*semver.Version, out io.Writer) {
	sort.Sort(semver.Collection(items))

	// 使用 map 去重
	seen := make(map[string]bool)
	for i := range items {
		fields := strings.SplitN(items[i].String(), "-", 2)
		v := fields[0]
		if len(fields) > 1 {
			v += "-" + fields[1]
		}
		// 如果已经显示过该版本，跳过
		if seen[v] {
			continue
		}
		seen[v] = true
		if v == curV {
			color.New(color.FgGreen).Fprintf(out, "* %s\n", v)
		} else {
			fmt.Fprintf(out, "  %s\n", v)
		}
		// 检查兼容性并显示提示
		if reason := checkGoCompatibility(v); reason != "" {
			color.New(color.FgYellow).Fprintf(out, "    [!] %s\n", reason)
		}
	}
}

// errstring 返回统一格式的错误信息
func errstring(err error) string {
	if err == nil {
		return ""
	}
	return wrapstring(err.Error())
}

func wrapstring(str string) string {
	if str == "" {
		return str
	}
	words := strings.Fields(str)
	if len(words) > 0 {
		words[0] = cases.Title(language.SimplifiedChinese).String(words[0])
	}
	return fmt.Sprintf("[gvm] %s", strings.Join(words, " "))
}
