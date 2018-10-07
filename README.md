# gosca [![Go Report Card](https://goreportcard.com/badge/github.com/sjas/gosca?style=flat-square)](https://goreportcard.com/report/github.com/sjas/gosca)

the missing golang project skeleton generator tool

## why? 

fixes the need to use golang's forced global $GOPATH without having to fuzz around with the environment every damn time

## depends on?

- `direnv` package installed or in `$PATH` to properly work (for dynamically fixing `$GOPATH` upon entering/leaving the created project folder)
- `git`

## what/how?

- creates standalone binary that can be run for you scaffolding needs
- checks wether git and direnv binaries are present and usable
- prompts for its own config settings if these are not yet defined via spf13/viper
- will create the the defacto go project structure into a dedicated folder within your desired workspace
- including a dummy main.go
- sets your github account as your own namespace
- automatically creates `direnv`'s `.envrc`
- runs `git init` right where it needs to be ran
- hands out project path to copy-paste after running

## misses? todos?

- automatically 'go get' on creation if need be? (== include godep and initialize that, too)

- parameter handling redone (== use 'spf13/cobra')
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
