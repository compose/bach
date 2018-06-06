# bach - Basic API (Compose) Handler

Bach is an application which uses the Compose API to provide the ability
to create, monitor and delete Compose databases.

To use, an environment variable - COMPOSEAPITOKEN must be set. This token value
can be obtained from the Compose console's Account view.

Where a command needs a parameter to identify a deployment, Bach can use either 
the deployment ID or the deployment name.

Bach is written in Go and uses the GoComposeAPI package for API requests.

The latest binary releases of Bach, for macOS, Linux and Windows, along with source snapshots, are available in the [releases](github.com/compose/bach/releases/latest) tab.

## Bach commands

```text
Bach is designed as simple route to accessing the Compose API

Usage:
  bach [command]

Available Commands:
  about       About Bach
  account     Show Account Details
  alerts      Show Alerts for deployment
  backups     Commands for backups
  cacert      Returns the self-signed cert for the deployment
  clusters    Show clusters
  create      Create a deployment
  databases   List databases
  datacenters Lists available datacenters
  deployments Show deployments attached to account
  deprovision Deprovision a deployment
  details     Show details for a deployment
  help        Help about any command
  list        List deployments attached to account
  recipe      Show details of a recipe
  recipes     Show Recipes related to deployment
  scale       Show scale information for a deployment
  teams       Commands for teams
  user        Commands for user management
  users       Show all users on account
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

## Bach subcommands

```text
bach backups [command]
  get         Show Backup details for deployment
  list        Show Backups for deployment
  restore     Restore a deployment
  start       Start backups for a deployment

bach user [command]
  add         Add user
  del         Delete user
  show        Show current user information

bach teams [command]
  create      Create named team
  get         get team details
  list        Show teams
  user        Commands for team users

bach teams user [command]
  add         add user to a team
  rem         removes user from a team

```
