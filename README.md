# Golang  REST API Boilerplate
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

Simple golang REST API boilerplate with service discovery and remote config supported by [Consul](https://www.consul.io/).

## Prerequisite
Things that you need to do before running this boilerplate:
1. Make sure, your consul service already running.
2. Provide your configuration in consul KV. For more details see this [post]((https://nodejs.org/en/)). 
3. If you're not using consul, please provide local configuration file by rename `.app-config.example.yaml` to `.app-config.yaml`.

## How to run
Run command below:
```shell script
$ go build -o boilerplate .
$ ./boilerplate serveHttp --consul localhost:8500 
```
In this case, consul service run at port `8500` by default.

## License
Copyright © 2020, [Bareksa Portal Investasi](https://bareksa.com).
Released under the [MIT License](LICENSE).