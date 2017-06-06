# bach - Basic API (Compose) Handler

Bach is a Go-based application which uses the Compose API to provide the ability
to create, monitor and delete Compose databases.

To use, an environment variable - `COMPOSEAPITOKEN` must be set. This token value
can be obtained from the Compose console's Account view.

Further details to follow.

## Installing

You can install `bach` using Homebrew (for macOS), downloading the binaries from Github, or building from source.

### Homebrew

```
brew install bach
```

### Download binaries from Github

Use the [releases](https://github.com/compose/bach/releases) page.

### Building from source

To be explained.

## Usage

```
Bach is designed as simple route to accessing the Compose API

Usage:
  bach [command]

Available Commands:
  about       About Bach
  account     Show Account Details
  alerts      Show Alerts for deployment
  backups     Commands for backups
  clusters    Show clusters
  create      Create a deployment
  databases   List databases
  datacenters Lists available datacenters
  deployments Show deployments attached to account
  deprovision Deprovision a deployment
  details     Show details for a deployment
  help        Help about any command
  recipe      Show details of a recipe
  recipes     Show Recipes related to deployment
  scale       Show scale information for a deployment
  user        Show user information
  versions    Show versions for deployment database
  watch       Watch a recipe status

Flags:
      --caescaped      Display full CAs as escaped strings
      --fullca         Show all of CA Certificates
  -h, --help           help for bach
      --json           Output post-processed JSON results
      --nodecodeca     Do not Decode base64 CA Certificates
      --raw            Output raw JSON responses (disables --watch)
      --token string   Your API Token (default "Your API Token")
      --wait           Automatically silently wait for a resulting recipe to finish
      --watch          Automatically watch a resulting recipe

Use "bach [command] --help" for more information about a command.
```
