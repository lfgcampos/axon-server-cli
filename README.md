# Axon Server Command line interface

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
* [ ] register-application
* [ ] delete-application
* [ ] init-cluster
* [ ] cluster
* [ ] register-node
* [ ] unregister-node
* [x] contexts
* [ ] register-context
* [ ] delete-context
* [ ] add-node-to-context
* [ ] delete-node-from-context

The option `-S` with the url to the Axon Server is optional, if it is omitted it defaults to [http://localhost:8024](http://localhost:8024/).

## Access control

When running Axon Server with access control enabled, executing commands remotely requires an access token. 
This has to provided with the `-t` option. When you run a command on the Axon Server node itself, you don't have to provide 
a token.

For Axon Server Standard Edition the token is specified in the `axonserver.properties` file 
\(property name = `axoniq.axonserver.token`\). The token needs to be supplied using the `-t` option in any of the commands.

## Config

This specific cli accept a configuration file named `axonserver-cli.yaml` on the same directory or using
`-config` flag to override it. In this file you can set default values for flags such as `server` and `token`.

## Commands

This section describes all commands supported by the command line interface, grouped by the specific area.
Mind that the list above is marking the ones which are already done.
All the commands have the `-h` option, which will show all the info you need to know including all the flags you can set.

For example:

`./axonserver-cli-cobra.exe -h`
```
This CLI is used to perform actions on AxonServer

Usage:
  axonserver-cli-cobra [command]

Available Commands:
  application commands related to applications
  context     commands related to contexts
  help        Help about any command
  user        Commands related to users

Flags:
      --config string   config file (default is axonserver-cli.yaml)
  -h, --help            help for axonserver-cli-cobra
  -S, --server string   URL of AxonServer (default "http://localhost:8024")
  -t, --token string    Authentication Token
  -v, --version         version for axonserver-cli-cobra

Use "axonserver-cli-cobra [command] --help" for more information about a command.
```

or
`./axonserver-cli-cobra.exe user -h`
```
This is the command related to users

Usage:
  axonserver-cli-cobra user [command]

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

Use "axonserver-cli-cobra user [command] --help" for more information about a command.
```

or even deeper
`./axonserver-cli-cobra.exe user register -h`
```
register a user to be used on axonserver

Usage:
  axonserver-cli-cobra user register [flags]

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