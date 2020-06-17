# Golang 和 Web

## httprouter 的原理

- httprouter 是通过前缀树来进行匹配的，前缀树相比基础的字典树来说，在匹配很长的字符串上有更好的性能。
- 每一个注册的 url 都会通过 / 切分为 n 个树节点（httprouter 会有一些区别，会存在根分裂），然后挂到相应 method 树上去，所以业务中有几种不同的 method 接口，就会产生对应的前缀树。
- httprouter 对下级节点的查找进行了优化，它把当前节点的下级节点的首字母维护在本身，匹配时先进行索引的查找。

## 中间件

- 中间件可以帮助我们实现业务与非业务之间的剥离，减小代码冗余程度。
- 借助中间件可以实现代码间的解耦。
- 通过一个或多个函数对 handler 进行包装，返回一个包括了各个中间件的逻辑函数链。

## 处理客户端提交数据，如 formdata、json

- 接收表单方式提交的数据：

```go
func main() {
    router := gin.Default()
    router.POST("/post", func(c *gin.Context) {
        name := c.PostForm("name")
        message := c.PostForm("message")
        fmt.Printf（"name: %s; message: %s", name, message）
    })
}
```

- 接收提交的 JSON 数据时：

```go
type Login struct {
    User     string `json: "user" binding: "required"`
    Password string `json: "password" binding: "required"`
}
func main() {
    router := gin.Default()
    router.POST("/login", func(c *gin.Context) {
        var login Login
        if err := c.ShouldBind($login); err == nil {
            c.JSON(http.StatusOK, gin.H{
                "user":     login.User,
                "password"  login.Password,
            })
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })
    router.Run(":8080")
}
```

## Golang 中常见的坑

- Map 遍历顺序不固定，它是一种 Hash 表实现，每次便利的顺序可能不一样。
- 在函数调用中数组是值传递，无法通过修改数组类型的参数返回结果，必要时需要使用切片。
- Go 语言中对象的内存地址可能发生变化，因此指针不能从其它非指针类型的值生成。
