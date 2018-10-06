# gosca [![Go Report Card](https://goreportcard.com/badge/github.com/sjas/gosca?style=flat-square)](https://goreportcard.com/report/github.com/sjas/gosca)

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

- check wether git is present else fail
- check for a working direnv installation else fail
- make the destination path / module path settable from config file ( `~/.config/gosca`)
- prompt for destination folder if config is not present yet and write to config
- get a decent help page
- automatically 'go get' on creation if need be?

- parameter handling redone
- if single parameter starts with http{,s}:// string == create skeleton with github repo name and clone github repo
- if other string == create repo with that name, do nothing else
- if one parameter starts with http{,s}:// == create repo with other arguements' name
- if argc>0 == show usage
- if argc>2 == fail
- names can only consist of [A-Za-z_-]
- -h / --help
- -u / --usage
- flag to change dst folder path in config
- flag to show current dst folder path
