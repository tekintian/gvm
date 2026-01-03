package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
	"github.com/tekintian/gvm/build"
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
	app.Version = build.Version()
	app.Copyright = "This is a golang version manager app. Thanks for voidint."
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
`, build.ShortVersion)
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
