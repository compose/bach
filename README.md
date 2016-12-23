# bach - Basic API (Compose) Handler

Bach is a Go-based application which uses the Compose API to provide the ability
to create, monitor and delete Compose databases.

To use, an environment variable - COMPOSEAPITOKEN must be set. This token value
can be obtained from the Compose console's Account view.

Further details to follow.

```
$ bach --help
usage: bach [<flags>] <command> [<args> ...]

A Compose CLI application

Flags:
  --help    Show context-sensitive help (also try --help-long and --help-man).
  --raw     Output raw JSON responses
  --fmt     Format output for readability
  --fullca  Show all of CA Certificates

Commands:
  help [<command>...]
    Show help.

  show account
    Show account details

  show deployments
    Show deployments

  show recipe [<recid>]
    Show recipe

  show deployment recipes [<depid>]
    Show deployment recipes

  show deployment versions [<depid>]
    Show version and upgrades

  show recipes
    Show recipes for a deployment

  show clusters
    Show available clusters

  show user
    Show current associated user

  show datacenters
    Show available datacenters

  show databases
    Show available database types

  create deployment [<flags>] [<name>] [<type>]
    Create deployment

```
