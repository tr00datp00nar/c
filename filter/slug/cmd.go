package slug

import (
	"github.com/gosimple/slug"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var Cmd = &Z.Cmd{

	Name:        `slug`,
	Summary:     help.S(_slug),
	Description: help.D(_slug),
	Version:     `v0.1.0`,
	Copyright:   `Copyright 2021 Robert S Muhlestein`,
	License:     `Apache-2.0`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/slug.git`,
	Issues:      `github.com/rwxrob/slug/issues`,

	Commands: []*Z.Cmd{help.Cmd},

	Other: []Z.Section{
		{`Dependencies`, `
				This {{cmd .Name}} command lightly wraps the github.com/gosimple/slug
				module and therefore carries a technical and legal dependency on
				it. See it for its own licensing considerations and usage before
				deciding to use this command.
			`,
		},
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		str := Z.ArgsOrIn(args)
		term.Print(slug.Make(str))
		return nil
	},
}
