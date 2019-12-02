# DB Unshare

DB Unshare is a project created to unshare dashboards in the Sysdig Monitor front-end. 
You must be and administrator to modify other people's dashboards.

⚠️ Warning! This tool will become deprecated and this repository removed once the official [CLI](https://github.com/draios/sysdig-platform-cli) is able to unshare dashboards.

## Installation

Use Go to install it with: 

```
go get -v -u github.com/sysdiglabs/sdc-db-unshare
```

## Usage

For every request you need to provide your Sysdig API Token using the `-t` or `--token` option.
For the following examples, we will use the fictional token `ZZZZZZZZ-XXXX-YYYY-XXXX-AAAABBBBCCCC`.

If you need more help, check the help menu:

```
$ sdc-db-unshare -h                                                          
List and unshare dashboards from Sysdig Monitor

Usage:
  sdc-db-unshare [command]

Available Commands:
  dashboard   Manage dashboards
  help        Help about any command

Flags:
  -h, --help           help for sdc-db-unshare
  -t, --token string   API Token for Sysdig Platform communication

Use "sdc-db-unshare [command] --help" for more information about a command.
```

### List all dashboards

```
$ sdc-db-unshare -t ZZZZZZZZ-XXXX-YYYY-XXXX-AAAABBBBCCCC dashboard list
ID    DASHBOARD                                      USER                              
43256 Cisco examples                                 user1@mail.com    
43697 Kubernetes Namespace State 2                   user2@mail.com    
55355 Kubernetes Namespace Health                    user2@mail.com    
56714 Executive Dashboard for RBC                    user2@mail.com    
52261 Kubernetes Cluster and Node Capacity v2        user3@mail.com  
52262 Istio Overview                                 user3@mail.com  
```

### Get a dashboard info

```
$ sdc-db-unshare -t ZZZZZZZZ-XXXX-YYYY-XXXX-AAAABBBBCCCC dashboard get 52262
ID    NAME           AUTHOR         PUBLIC SHARED AUTOCREATED VERSION 
52262 Istio Overview user3@mail.com false  true   false       1
```

### Unshare a dashboard

```
$ sdc-db-unshare -t ZZZZZZZZ-XXXX-YYYY-XXXX-AAAABBBBCCCC dashboard unshare 52262 
```

The dashboard should now be unshared and the version should be increased by 1.

```
$ sdc-db-unshare -t ZZZZZZZZ-XXXX-YYYY-XXXX-AAAABBBBCCCC dashboard get 52262
ID    NAME           AUTHOR         PUBLIC SHARED AUTOCREATED VERSION 
52262 Istio Overview user3@mail.com false  false  false       2
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)
