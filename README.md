# 🌳 Personal Bonzai Commands C

This repository was created using `github.com/rwxrob/z` as a starting point. I have borrowed quite a bit of the code for the branch commands from this repository and others by rwxrob.

These days I prefer to maintain a single Go [stateful command tree monolith](https://rwxrob.github.io/zet/1729/) tool rather than a ton of shell scripts in whatever languages. In fact, [Bonzai](https://github.com/rwxrob/bonzai) was created specifically for this sort of thing. I just `curl` down a single binary to whatever system I'm on and I have all of my favorite functionality on *any* device with zero compatibility hassles and installation dependencies. Everything just works, *anywhere*.

## Install

Just download one of the [release binaries](https://github.com/tr00datp00nar/c/releases):

```bash
curl -L https://github.com/tr00datp00nar/c/releases/latest/download/c-linux-amd64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/c/releases/latest/download/c-darwin-amd64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/c/releases/latest/download/c-darwin-arm64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/c/releases/latest/download/c-windows-amd64 -o ~/.local/bin/tr00datp00nar
```

Or install directly with `go`:

```bash
go install github.com/tr00datp00nar/c@latest
```

I prefer to use `c` instead of setting up a multicall binary since the habits it builds into my muscle memory work on any operating system or device and it doesn't take too much space when using UNIX pipelines and
such:

```bash
echo $(c isosec) $(c y2j quotes.yaml | jq -r .mad )
```

## Tab Completion in Bash

To activate bash completion just use the `complete -C` option from your `.bashrc` or command line. There is no messy sourcing required. All the completion is done by the program itself.

```bash
complete -C c c
```

If you don't have bash or tab completion check out the shortcut commands instead.

## Tab Completion in Zsh

Zsh does a good job of learning your commands over time all by itself, but some of the custom completions may not work as well. Personally, I use the Oh-My-Zsh option below, but the creator of Bonzai and the original Z command tree (rwxrob) prefers the default Linux shell (Bash) over the default Mac shell (Zsh). (PRs to rwxrob's repository are welcome to integrate completion into Zsh without dumping a ton of shell code that has to be sourced.)

### Oh-My-Zsh

Oh-My-Zsh has an available plugin called [zsh-bash-completions-fallback](https://github.com/3v1n0/zsh-bash-completions-fallback). This plugin allows zsh to fallback to bash completions when it can't find the appropriate completions itself.

Once installed, you can use the same `complete -C c c` as you normally would in bash.


## Embedded Documentation

All documentation (like manual pages) has been embedded into the source code of the application. See the source or run the program with help to access it.

## Building

Releases are built using the following commands:

```bash
c go build
gh release create
gh release upload TAG build/*
```
