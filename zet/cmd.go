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
		edit,
		new,
		update},
}

var new = &Z.Cmd{
	Name:        `new`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_new),
	Description: help.D(_new),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		newZet()
		return nil
	},
}

var update = &Z.Cmd{
	Name:        `update`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_update),
	Description: help.D(_update),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		dir := ZetDir
		updateIndexList(dir)
		return nil
	},
}

var edit = &Z.Cmd{
	Name:        `edit`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_edit),
	Description: help.D(_edit),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		filePicker()
		return nil
	},
}
