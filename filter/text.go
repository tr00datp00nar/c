package filter

import _ "embed"

//go:embed text/en/filter.md
var _filter string

//go:embed text/en/bon.md
var _bon string

//go:embed text/en/full.md
var _full string

//go:embed tmpl/fullcommand.tmpl
var fullcommandtmpl string

//go:embed text/en/deps.md
var _deps string

//go:embed tmpl/deps.tmpl
var depstmpl string

//go:embed text/en/branchreadme.md
var _branchreadme string

//go:embed tmpl/branchreadme.tmpl
var branchreadme string

//go:embed text/en/quick.md
var _quick string

//go:embed tmpl/quickcommand.tmpl
var quickcommandtmpl string

//go:embed text/en/gobadges.md
var _gobadges string

//go:embed tmpl/gobadges.tmpl
var gobadges string
