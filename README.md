# Axon Server Command line interface

[![Release](https://img.shields.io/github/release/lfgcampos/axon-server-cli.svg?style=flat-square)](https://github.com/lfgcampos/axon-server-cli/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/lfgcampos/axon-server-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/lfgcampos/axon-server-cli?style=flat-square)](https://goreportcard.com/report/github.com/lfgcampos/axon-server-cli)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

The intention of this repo is to have a similar cli as in [axon-server-cli](https://github.com/AxonIQ/axon-server-se/tree/master/axonserver-cli) but written in go instead of java.
This is more of a learning exercise of go language but should be as much usable as the official cli.

The Axon Server command line interface allows updating the Axon Server configuration through scripts or from a command line.

For the [Axon Server Standard edition](https://axoniq.io/product-overview/axon-server) the only supported commands are:

* [ ] metrics
* [x] users 
* [x] register-user
* [x] delete-user

[Axon Server Enterprise edition](https://axoniq.io/product-overview/axon-server-enterprise) supports these additional commands:‌

* [x] applications
* [x] register-application
* [x] delete-application
* [x] init-cluster
* [x] cluster
* [ ] update-license
* [x] register-node
* [x] unregister-node
* [x] contexts
* [x] register-context
* [x] delete-context
* [x] add-node-to-context
* [x] delete-node-from-context

The option `-S` with the url to the Axon Server is optional, if it is omitted it defaults to [http://localhost:8024](http://localhost:8024/).

## Access control

When running Axon Server with access control enabled, executing commands remotely requires an access token. 
This has to provided with the `-t` option. When you run a command on the Axon Server node itself, you don't have to provide 
a token.

For Axon Server Standard Edition, the token is specified in the `axonserver.properties` file 
\(property name = `axoniq.axonserver.token`\). The token needs to be supplied using the `-t` option in any of the commands.

## Config

This specific cli accept a configuration file named `axonserver-cli.yaml` on the same directory or using
`-config` flag to override it. In this file you can set default values for flags such as `server` and `token`.

## Commands

This section describes some commands with examples supported by the command line interface.
Mind that the list above is marking the ones which are already done.
All commands have the `-h` option, which will show all the info you need to know including all the flags you can set.

For example:

`axon-server-cli -h`
```
This CLI is used to perform actions on AxonServer

Usage:
  axon-server-cli [command]

Available Commands:
  application commands related to applications
  context     commands related to contexts
  help        Help about any command
  user        Commands related to users

Flags:
      --config string   config file (default is axonserver-cli.yaml)
  -h, --help            help for axon-server-cli
  -S, --server string   URL of AxonServer (default "http://localhost:8024")
  -t, --token string    Authentication Token
  -v, --version         version for axon-server-cli

Use "axon-server-cli [command] --help" for more information about a command.
```

or
`axon-server-cli user -h`
```
This is the command related to users

Usage:
  axon-server-cli user [command]

Aliases:
  user, u

Available Commands:
  delete      Remove a user
  list        list all the users
  register    Register a user

Flags:
  -h, --help   help for user

Global Flags:
      --config string   config file (default is axonserver-cli.yaml)
  -S, --server string   URL of AxonServer (default "http://localhost:8024")
  -t, --token string    Authentication Token

Use "axon-server-cli user [command] --help" for more information about a command.
```

or even deeper
`axon-server-cli user register -h`
```
register a user to be used on axonserver

Usage:
  axon-server-cli user register [flags]

Aliases:
  register, r

Flags:
  -h, --help              help for register
  -p, --password string   user password
  -r, --roles strings     user roles
  -u, --username string   user username

Global Flags:
      --config string   config file (default is axonserver-cli.yaml)
  -S, --server string   URL of AxonServer (default "http://localhost:8024")
  -t, --token string    Authentication Token
```

## Release process

Release is automatic triggered by a git `tag` using [goreleaser](https://goreleaser.com).
For do that, you have to create a new tag using `git tag -a vx.y.z -m "my release message"`
and push it to the repo using `git push origin master --tags`.