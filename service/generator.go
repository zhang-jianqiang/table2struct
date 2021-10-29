package service

import (
	"fmt"
	"github.com/zhang-jianqiang/table2struct/database"
	"github.com/zhang-jianqiang/table2struct/pkg"
	"log"
	"strings"
)

type generator struct {
}

var (
	m *database.Model
	tableName string
)

func NewGenerator(model *database.Model) *generator {
	// 初始化数据库连接
	 m = model
	 database.InitDb(m)
	 return &generator{}
}

func (g *generator) CreateStruct() {
	fmt.Print("请输入表名并回车：")
	fmt.Scanln(&tableName)
	if tableName == "" {
		log.Fatalf("表名不能为空")
	}
	// 获取表信息
	tableDesc := m.GetTablesInfo(tableName)
	structStr := createStructTemplate(tableDesc)
	createStructFile(structStr)
}

func createStructTemplate(tableDesc []database.DescResult) string {
	str := fmt.Sprintf(`
type %s struct {`, pkg.StrToPascalCase(tableName));
	for _, v := range tableDesc {
		str += fmt.Sprintf(`
	%s %s`, pkg.StrToPascalCase(v.Field), getType(v.Type));
	}
	str += `
}`
	// 拼接包名和import
	if strings.Contains(str, "time.Time") {
		// 包含import
		str = fmt.Sprintf(`package %s

import "time"
`, m.PackageName) + str;
	} else {
		str = fmt.Sprintf(`package %s
`, m.PackageName) + str;
	}
	return str
}

func createStructFile(structStr string) {
	err := pkg.OutputFile(m.Dir+"/"+pkg.StrToPascalCase(tableName)+".go", structStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("文件生成成功 " + m.Dir+"/"+pkg.StrToPascalCase(tableName)+".go")
}

func getType(str string) string {
	goType := "string"
	switch {
	case strings.Contains(str, "bigint"):
		goType = "int64"
	case strings.Contains(str, "int"):
		goType = "int"
	case strings.Contains(str, "char"), strings.Contains(str, "text"), strings.Contains(str, "longtext"):
		goType = "string"
	case strings.Contains(str, "datetime"), strings.Contains(str, "timestamp"), strings.Contains(str, "date"), strings.Contains(str, "date"):
		goType = "time.Time"
	}
	return goType
}
