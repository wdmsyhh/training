# Redis Homework

## 写出下面的 redis 命令

- redis 中插入十条 student{id, name, age} 的数据：

    ```shell
    hmset student:1 id 1 name name1 age 11
    hmset student:2 id 2 name name1 age 12
    hmset student:3 id 3 name name1 age 13
    hmset student:4 id 4 name name1 age 14
    hmset student:5 id 5 name name1 age 15
    hmset student:6 id 6 name name1 age 16
    hmset student:7 id 7 name name1 age 17
    hmset student:8 id 8 name name1 age 18
    hmset student:9 id 9 name name1 age 19
    hmset student:10 id 10 name name1 age 20
    ```

- redis 中记录 student 的投票次数，并执行加 1 和加 3 的操作：

    ```shell
    # 设置学生 id 对应的投票次数 
    zadd vote 10 1
    zadd vote 10 2
    zadd vote 10 3
    zadd vote 10 4
    zadd vote 10 5
    zadd vote 10 6
    zadd vote 10 7
    zadd vote 10 8
    zadd vote 10 9
    zadd vote 10 10

    # id 为 2 的会加 １
    zincrby vote 1 2
    # id 为 2 的会加 ３
    zincrby vote ３ 2
    ```

## 使用后端框架操作 redis 实现下面的功能

- redis 中插入十条 student{id, name, age} 的数据。其中 student 需要定义成 model，id、name、age 都需要随机生成。
- redis 中记录 student 的投票次数(初始值随机生成)，并按从低到高的顺序取出来。
