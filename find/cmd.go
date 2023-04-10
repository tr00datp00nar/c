package find

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/tr00datp00nar/c/find/web"
)

var Cmd = &Z.Cmd{
	Name:      `find`,
	Usage:     `[help]`,
	Version:   `v0.1.0`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	Summary:   help.S(_find),
	Commands: []*Z.Cmd{
		web.Cmd,
		help.Cmd,
	},
}
