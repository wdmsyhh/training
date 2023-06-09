# 语言基础

## 声明变量的几种方式

- 指定变量类型，如果没有初始化，则变量默认为零值：

    ```go
    var v_name v_type
    v_name = value
    ```

- 根据值自行判定变量类型：

    ```go
    var v_name = value
    ```

- := 左侧如果没有声明新的变量，就产生编译错误：

    ```go
    var intVal int
    intVal,intVal1 := 1,2
    ```

- 多变量声明：

    ```go
    var vname1, vname2, vname3 type
    vname1, vname2, vname3 = v1, v2, v3
    var vname1, vname2, vname3 = v1, v2, v3
    vname1, vname2, vname3 := v1, v2, v3
    var (
        vname1 v_type1
        vname2 v_type2
    )
    ```

## Golang 中的变量重声明

- 由于变量的类型在其初始化时就已经确定了，所以对它再次声明时赋予的类型必须与其原本的类型一致。
- 变量的重声明只可能发生在某一个代码块中。如果与当前的变量重名的是外层代码块的变量，那就是另一种含义了。
- 变量的重声明只有在使用短变量声明时才会发生，否则也无法通过编译。如果要在此处声明全新的变量，那么就应该使用包含关键字var的声明语句，但是这时就不能与同一代码块中的任何变量有重名了。
- 被“声明并赋值”的变量必须是多个，并且其中至少有一个是新的变量。这时我们才说对其中的旧变量进行重声明。

## 一个变量与其外层代码块中的变量重名

- 对于不同的代码块来说，其中的变量重名没什么大不了，照样可以通过编译，即便这些代码块有着直接的嵌套关系也是如此。
- 代码引用变量时总会最优先查找当前代码块中（不包含任何子代码块）的那个变量。
- 如果当前代码块中没有声明以此为名的变量，那么程序会沿着代码块的嵌套关系，一层一层的查找。
