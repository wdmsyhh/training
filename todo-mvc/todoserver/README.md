# todos 服务端

- 启动 mongo

```shell
docker run -itd --name todo-mongo -p 27017:27017 mongo:4.2.10 --auth

docker exec -it todo-mongo mongo admin

# 创建一个名为 admin，密码为 123456 的用户。
db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});

# 尝试使用上面创建的用户信息进行连接。
db.auth('admin', '123456')

# 创建数据库
use todos

# 创建 jwt 秘钥数据
db.account.insertOne({"secretKey":"abcdefg", "enable": true})

# 查询
db.account.find({"enable": true})
```

- 启动 redis

```shell
docker run -itd --name todo-redis -p 6379:6379 redis:5.0.12 --requirepass 123456

docker exec -it todo-redis bash

redis-cli

auth 123456

set key1 value1

get key1 value1
```

- 启动服务

```shell
go run main.go
```
