package cli

import "github.com/urfave/cli/v2"

var (
	commands = []*cli.Command{
		{
			Name:      "ls",
			Usage:     "List installed versions",
			UsageText: "gvm ls",
			Action:    list,
		},
		{
			Name:      "ls-remote",
			Usage:     "List remote versions available for install",
			UsageText: "gvm ls-remote [stable|archived|unstable|rc|beta]",
			Action:    listRemote,
		},
		{
			Name:      "use",
			Usage:     "Switch to specified version",
			UsageText: "gvm use <version>",
			Action:    use,
		},
		{
			Name:      "install",
			Usage:     "Download and install a version",
			UsageText: "gvm install <version>",
			Action:    install,
		},
		{
			Name:      "uninstall",
			Usage:     "Uninstall a version",
			UsageText: "gvm uninstall <version>",
			Action:    uninstall,
		},
		{
			Name:      "update",
			Usage:     "Fetch the newest version of gvm",
			UsageText: "gvm update",
			Action:    update,
		},
		{
			Name:      "clean",
			Usage:     "Remove files from the package download directory",
			UsageText: "gvm clean",
			Action:    clean,
		},
	}
)
