# Kubernetes and Cloud Ready Boilerplate Go Server [![Docker Automated build](https://img.shields.io/docker/automated/onuryilmaz/boilerplate-go.svg?style=flat-square)](https://hub.docker.com/r/onuryilmaz/boilerplate-go/)
* REST Server included
* Cloud ready and all steps in Docker
* Kubernetes deployment and service included

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
make push
```

## Deploy and Test
```
kubectl create -f kubernetes

kubectl run -i --tty curl --image=tutum/curl -- sh 
$ curl -i boilerplate-go/ping
HTTP/1.1 200 OK
..
{"Status":"OK"}
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

