package obsidian

import (
	"path/filepath"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `obsidian`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_obsidian),
	Description: help.D(_obsidian),

	Commands: []*Z.Cmd{
		help.Cmd,
		open,
	},
}

var open = &Z.Cmd{
	Name:        `open`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_open),
	Description: help.D(_open),

	Commands: []*Z.Cmd{
		help.Cmd,
		file,
		vault,
	},
}

var vault = &Z.Cmd{
	Name:        `vault`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_selectVault),
	Description: help.D(_selectVault),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		vault := selectVault()
		openVault(vault)
		return nil
	},
}

var file = &Z.Cmd{
	Name:        `file`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_file),
	Description: help.D(_file),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		vault := selectVault()
		file := selectFile(vault)
		openFile(filepath.Base(vault), file)
		return nil
	},
}
