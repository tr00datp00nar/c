package main

import (
	_ "embed"
	"log"

	// Core Bonzai Library
	Z "github.com/rwxrob/bonzai/z"

	// Branches by rwxrob
	"github.com/rwxrob/conf"
	"github.com/rwxrob/good" // Common Go tools
	"github.com/rwxrob/grep"
	"github.com/rwxrob/help" // All-in-one Help package
	"github.com/rwxrob/keg"
	"github.com/rwxrob/uniq" // Generate unique strings and numbers
	"github.com/rwxrob/vars" // Manage vars like env vars in bonzai

	// My local branches
	"github.com/tr00datp00nar/c/filter"
	"github.com/tr00datp00nar/c/zet"

	// My branches that are live on Github

	"github.com/tr00datp00nar/bon"
	"github.com/tr00datp00nar/file"
	"github.com/tr00datp00nar/find"
	"github.com/tr00datp00nar/get"
	"github.com/tr00datp00nar/mal"
	"github.com/tr00datp00nar/music"

	"github.com/tr00datp00nar/obsdn"
	"github.com/tr00datp00nar/rank"
	"github.com/tr00datp00nar/set"
	"github.com/tr00datp00nar/worldbuilding"
)

func init() {
	Z.Dynamic["uname"] = func() string { return Z.Out("uname", "-a") }
	Z.Dynamic["ls"] = func() string { return Z.Out("ls", "-l", "-h") }
}

func main() {

	// remove log prefixes
	log.SetFlags(0)

	// provide panic trace
	Z.AllowPanic = true

	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:        `c`,
	Usage:       `COMMAND [args]`,
	Summary:     help.S(_c),
	Description: help.D(_c),
	Copyright:   `Copyright 2023 Micah Nadler`,
	Version:     `v0.1.1`,
	License:     `Apache-2.0`,
	Source:      `git@github.com:tr00datp00nar/c.git`,
	Issues:      `github.com/tr00datp00nar/c/issues`,

	Commands: []*Z.Cmd{
		bon.Cmd,
		conf.Cmd,
		file.Cmd,
		filter.Cmd,
		find.Cmd,
		get.Cmd,
		good.Cmd,
		grep.Cmd,
		help.Cmd,
		keg.Cmd,
		mal.Cmd,
		music.Cmd,
		obsdn.Cmd,
		rank.Cmd,
		set.Cmd,
		uniq.Cmd,
		vars.Cmd,
		worldbuilding.Cmd,
		zet.Cmd,
	},

	Shortcuts: Z.ArgMap{
		`bond`:     {`bon`, `deps`},
		`bonf`:     {`bon`, `full`},
		`bonq`:     {`bon`, `quick`},
		`bonbr`:    {`bon`, `branchreadme`},
		`netcheck`: {`get`, `net`, `check`},
		`lan`:      {`get`, `net`, `lan`},
		`wan`:      {`get`, `net`, `wan`},
		`mac`:      {`get`, `net`, `mac`},
		`dns`:      {`get`, `net`, `dns`},
		`router`:   {`get`, `net`, `router`},
		`wifipass`: {`get`, `net`, `wifipass`},
		`manga`:    {`get`, `manga`},
		`newworld`: {`worldbuilding`, `new-world`},
		`nvim`:     {`get`, `nvim`},
		`ob`:       {`obsidian`},
		`ofile`:    {`obsidian`, `open`, `file`},
		`ovault`:   {`obsidian`, `open`, `vault`},
		`sw`:       {`find`, `web`},
		`swb`:      {`find`, `web`, `brave`},
		`swyt`:     {`find`, `web`, `youtube`},
		`animerec`: {`mal`, `anime`, `rec`},
		`waifus`:   {`set`, `wallpaper`, `waifus`},
		`z`:        {`zet`, `new`},
		`zo`:       {`zet`, `obsidian`},
	},
}
