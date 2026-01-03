//go:generate go run gen_version.go

package app_build

import "strings"

// The value of variables come form `gb build -ldflags '-X "app_build.Build=xxxxx" -X "app_build.CommitID=xxxx"' `
var (
	// Build build time
	Build string
	// Branch current git branch
	Branch string
	// Commit git commit id
	Commit string
)

// Version 生成版本信息
func Version() string {
	var buf strings.Builder
	buf.WriteString(ShortVersion)

	if Build != "" {
		buf.WriteByte('\n')
		buf.WriteString("build: ")
		buf.WriteString(Build)
	}
	if Branch != "" {
		buf.WriteByte('\n')
		buf.WriteString("branch: ")
		buf.WriteString(Branch)
	}
	if Commit != "" {
		buf.WriteByte('\n')
		buf.WriteString("commit: ")
		buf.WriteString(Commit)
	}
	return buf.String()
}
