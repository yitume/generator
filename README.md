# generator

generator 是一个基于 MySQL 表结构的 go CRUD 生成器。

## 使用

```bash
make build

./generator new \
    -m "root:root@tcp(127.0.0.1:3306)/information_schema" \ # 指定mysql地址
    -d shop \                                               # 指定数据库
    -o ./dist \                                             # 指定生成代码输出目录
    -t "%fa_%" \                                            # 指定表名
    -M "git.yitum.com/saas/shop-admin"                      # 指定go modules
    -s "ssh://USER_NAME:USER_PASSWORD@IP:PORT"              # 指定ssh隧道
    -T tmpl                                                 # 指定进行渲染的模板目录
```

## 注意

- MySQL 中类型为 Json 的字段会生成 TableName+FieldName + "Json" 的结构体
这个结构体需要自己在 model/mysql/addtion.json.go 中定义

- trans 中如果也引入了这个结构体，模板中会在结构体前新增 mysql 包名，需要手动执行
``goimports -w app`` 来 import model/mysql 包

- 对上述生成的 Json 结构体，如果是 ``[]int []string``，可以预定义 ``type IntJson []int `` 和 ``type StringJson []string`` 类型的 Json 结构体，通过 ``type XxxxYyyJson = IntsJson`` 
或者 ``type XxxxYyyJson = StringJson`` 的方式来使用

## 默认变量

- TABLE_NAME: 表名

## 其他

```golang
# 查看使用帮助
./generator new -h
./generator version -h
```

## LICENSE

[Apache License 2.0](./LICENSE)
