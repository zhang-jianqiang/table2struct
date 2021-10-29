package main

import (
	"github.com/zhang-jianqiang/table2struct/database"
	"github.com/zhang-jianqiang/table2struct/service"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("获取当前目录失败")
	}
	g := service.NewGenerator(&database.Model{
		Username: "root",
		Password: "root",
		Url: "127.0.0.1:3306",
		Dbname: "test",
		PackageName: "dao",
		Dir: pwd + "/dao",
	});
	g.CreateStruct()
}