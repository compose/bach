# bach - Basic API (Compose) Handler

Bach is a Go-based application which uses the Compose API to provide the ability
to create, monitor and delete Compose databases.

To use, an environment variable - COMPOSEAPITOKEN must be set. This token value
can be obtained from the Compose console's Account view.

Further details to follow.

```
$ ./bach --help-long                                         master ✚ ✱
usage: bach [<flags>] <command> [<args> ...]

A Compose CLI application

Flags:
      --help                  Show context-sensitive help (also try --help-long and
                              --help-man).
  -r, --raw                   Output raw JSON responses
  -j, --json                  Output post-processed JSON results
  -f, --fullca                Show all of CA Certificates
  -t, --token="yourAPItoken"  Set API token

Commands:
  help [<command>...]
    Show help.


  account
    Show account details


  deployments
    Show deployments


  recipe [<recid>]
    Show recipe


  recipes <deployment id>
    Show deployment recipes


  versions <deployment id>
    Show version and upgrades


  details <deployment id>
    Show deployment information


  scale <deployment id>
    Get scale of deployment


  clusters
    Show available clusters


  user
    Show current associated user


  datacenters
    Show available datacenters


  databases
    Show available database types


  create [<flags>] [<name>] [<type>]
    Create deployment

    --cluster=CLUSTER        Cluster ID
    --datacenter=DATACENTER  Datacenter location

  set scale <deployment id> <units>
    Set scale of deployment


  watch [<flags>] <recipe id>
    Watch recipe

    --refresh=10  Refresh rate in seconds


```
