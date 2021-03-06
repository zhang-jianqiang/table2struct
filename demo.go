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
