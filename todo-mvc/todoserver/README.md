# todos 服务端

- 启动 mongo

```shell
docker run -d -p 27017:27017 --name example-mongo mongo:latest
```

- 启动 redis

```shell
docker run -itd --name redis_01 -p 6379:6379 redis:5.0.12
```

- 启动服务

```shell
go run main.go
```