package cli

import (
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/k0kubun/go-ansi"
	"github.com/tekintian/gvm/collector"
	"github.com/tekintian/gvm/version"
	"github.com/urfave/cli/v2"
)

const (
	stableChannel   = "stable"
	unstableChannel = "unstable"
	archivedChannel = "archived"
	betaChannel     = "beta"
	rcChannel       = "rc"
)

func listRemote(ctx *cli.Context) (err error) {
	channel := ctx.Args().First()
	if channel != "" && channel != stableChannel && channel != unstableChannel && channel != archivedChannel && channel != betaChannel && channel != rcChannel {
		return cli.ShowSubcommandHelp(ctx)
	}

	c, err := collector.NewCollector(strings.Split(os.Getenv(mirrorEnv), ",")...)
	if err != nil {
		return cli.Exit(errstring(err), 1)
	}

	var vs []*version.Version
	switch channel {
	case stableChannel:
		vs, err = c.StableVersions()
	case unstableChannel:
		vs, err = c.UnstableVersions()
	case archivedChannel:
		vs, err = c.ArchivedVersions()
	case betaChannel:
		vs, err = c.FilterVersions(betaChannel)
	case rcChannel:
		vs, err = c.FilterVersions(rcChannel)
	default:
		vs, err = c.AllVersions()
	}
	if err != nil {
		return cli.Exit(errstring(err), 1)
	}

	items := make([]*semver.Version, 0, len(vs))
	for i := range vs {
		vname := vs[i].Name
		var idx int
		if strings.Contains(vname, "alpha") {
			idx = strings.Index(vname, "alpha")

		} else if strings.Contains(vname, "beta") {
			idx = strings.Index(vname, "beta")

		} else if strings.Contains(vname, "rc") {
			idx = strings.Index(vname, "rc")
		}
		if idx > 0 {
			vname = vname[:idx] + "-" + vname[idx:]
		}
		v, err := semver.NewVersion(vname)
		if err != nil || v == nil {
			continue
		}
		items = append(items, v)
	}

	render(inuse(goroot), items, ansi.NewAnsiStdout())
	return nil
}
