```
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' ./main.go
$ docker build -t hbchen/micro-mesh-example-api:v0.0.1 .
$ docker build -t hbchen/micro-mesh-example-srv-1:v0.0.1 .
$ docker build -t hbchen/micro-mesh-example-srv-2:v0.0.1 .
```