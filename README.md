# example-ping-service-server

## Build

```bash
$ gofmt -w .
$ golint .
$ go build .
$ docker build -t nokamoto13/example-ping-service-server:v0 .
$ docker push nokamoto13/example-ping-service-server:v0
```
