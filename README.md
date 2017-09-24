# Boilerplate Go Package
* REST Server included
* Cloud ready and all steps in Docker

[![asciicast](https://asciinema.org/a/mw1363jxJNmqdycXx39Ol6LpJ.png)](https://asciinema.org/a/mw1363jxJNmqdycXx39Ol6LpJ)

## Requirements
* Docker (> version 17.05)
* GNU make

## Test
```
make test
```

## Build
```
make build
```
## Run
```
make run
```

## Push
```
make push DOCKER_REGISTRY=$REGISTRY
```

## Example Flow
```
$ make build
$ make run

# on another shell
$ curl -i localhost:8080/ping
HTTP/1.1 200 OK
..
{"Status":"OK"}
```

## Dependency Management
* [govendor](https://github.com/kardianos/govendor) is used for dependency management.
* Fixed versions can be checked from [vendor.json](vendor/vendor.json)

