package filter

import (
	_ "embed"

	"github.com/tr00datp00nar/c/filter/slug"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `filter`,
	Aliases:     []string{`filt`},
	Summary:     help.S(_filter),
	Description: help.D(_filter),
	Copyright:   `Copyright 2022 Robert S Muhlestein`,
	Version:     `v0.1.0`,
	License:     `Apache-2.0`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/filter.git`,
	Issues:      `github.com/rwxrob/filter/issues`,

	Commands: []*Z.Cmd{
		slug.Cmd,
		help.Cmd,
	},
}
