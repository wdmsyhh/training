# MongoDB Homework

## mongo 命令

- 按 (age >= 70 && in_array("A", tags)) || (age >= 30 && age < 50 && (in_array("B", tags) || in_array("C", tags))) 的匹配逻辑正确查找客户，输出其 id，年龄，标签信息：

    ```shell
    db.member.find({
        "$or": [
            {"age": {"$gte": 70}, "tags": "A"},
            {
                "age": {"$gte": 30, "$lt": 50},
                "$or": [{"tags": "B"},{"tags": "C"}]
            }
        ]
    },{
        _id: 1,
        age: 1,
        tags: 1
    }).pretty()
    ```

- 使用 $inc 操作符，将 in_array("B", tags) 的客户的年龄加 10：

    ```shell
    db.member.update({
        "tags": "B"
    },{
        "$inc": {"age": 10}
    },{
        multi: true
    })
    ```

- 假设 member 的 age 字段只对应于客户创建时的年龄，写一个迁移脚本为每个客户添加一个叫 birthYear 的字段，并更新成正确的出生年份：

    ```shell
    var cursor = db.member.find()
    cursor.forEach(function(doc) {
        db.member.update({
            "_id": doc._id
        },{
            "$set": {"birthYear": doc.createdAt.getFullYear() - doc.age}
        })
    })
    ```

- 使用 aggregate 统计下各 tag 中客户的平均年龄和人数：

    ```shell
    db.member.aggregate([
        {"$unwind": "$tags"},
        {"$group": {
            "_id": "$tags",
            "avgAge": {"$avg": "$age"},
            "count": {"$sum": 1}
        }}
    ])
    ```
