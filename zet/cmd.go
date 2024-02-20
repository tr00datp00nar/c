package zet

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `zet`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_zet),
	Description: help.D(_zet),

	Commands: []*Z.Cmd{
		help.Cmd,
		newZet,
		zetObsidian},
}

var newZet = &Z.Cmd{
	Name:        `newZet`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_newZet),
	Description: help.D(_newZet),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		NewZet()
		return nil
	},
}

var zetObsidian = &Z.Cmd{
	Name:        `zetObsidian`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_zetObsidian),
	Description: help.D(_zetObsidian),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		ZetObsidian()
		return nil
	},
}
