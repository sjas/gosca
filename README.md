# gosca

[![Go Report Card](https://goreportcard.com/badge/github.com/sjas/gosca?style=flat-square)](https://goreportcard.com/report/github.com/sjas/gosca)

go project scaffolding tool.

## why? 

fixes the need to use golang's forced global $GOPATH without having to fuzz around with the environment all the damn time

## depends on?

- `direnv` package installed or in `$PATH` to properly work (for dynamically fixing `$GOPATH` upon entering/leaving the created project folder)
- `git`

## what?

- creates standalone binary
- will create the the usual go project structure into "/home/wrk" (change as you wish) including a dummy main.go
- assumes `github.com/sjas` as module path (change as you wish)
- automatically creates `direnv`'s `.envrc`
- runs `git init` where it needs to be run, in `~/wrk/PROJECT/src/github.com/sjas/PROJECT` (fixed automatically upon changes above)
- hands out project path to copy-paste after running

## misses? todos?

- some error reporting/handling
- make the destination path / module path settable from config file (not hardcoded, is no use for me since the path structure is a fixed convention for all my setups)
