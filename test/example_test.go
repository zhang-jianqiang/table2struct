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
