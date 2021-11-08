# table2struct
根据表名自动生成go语言中的结构体，主要为了方便go的数据库操作
## 使用方法
第一种方法：通过可执行程序demo.go
这种方法可以不初始化TableName，命令行会提示输入表名
```go
package main

import (
	"github.com/zhang-jianqiang/table2struct/database"
	"github.com/zhang-jianqiang/table2struct/service"
)

func main() {
	g := service.NewGenerator(&database.Model{
		Username: "root",
		Password: "",
		Url: "127.0.0.1:3306",
		Dbname: "test",
		PackageName: "dao",
		Dir: "./dao",
		TableName: "", // TableName为空时，通过命令行参数输入，如果使用test测试的话只能在这里传入TableName
	})
	g.CreateStruct()
}
```
第二种方法：通过test测试执行,test/example_test.go
这种方法必须指定表名，因为test测试无法传入命令行参数
```go
package test

import (
	"github.com/zhang-jianqiang/table2struct/database"
	"github.com/zhang-jianqiang/table2struct/service"
	"testing"
)

func TestCreateModel(t *testing.T) {
	g := service.NewGenerator(&database.Model{
		Username: "root",
		Password: "",
		Url: "127.0.0.1:3306",
		Dbname: "test",
		PackageName: "dao",
		Dir: "./../dao",
		TableName: "ai_lesson",
	})
	g.CreateStruct()
}
```

```shell
go test -v  -run TestCreateModel
```