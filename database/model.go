package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Model struct {
	Username string
	Password string
	Url string
	Dbname string
	PackageName string
	Dir string
	TableName string
}

type DescResult struct {
	Field string
	Type string
	Null string
	Default string
	Extra string
}

var db *gorm.DB

func InitDb(m *Model) {
	if m.Username == "" {
		log.Fatalln("Username 不能为空")
	}
	// if m.Password == "" {
	// 	log.Fatalln("Password 不能为空")
	// }
	if m.Url == "" {
		log.Fatalln("Url 不能为空")
	}
	if m.Dbname == "" {
		log.Fatalln("Dbname 不能为空")
	}
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Url, m.Dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败，err:%+v\n", err)
	}
}

func (m *Model)GetTablesInfo(tableName string) []DescResult {
	var result []DescResult
	sql := "desc " + tableName
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		log.Fatalf("数据库执行失败，err: %+v\n", err)
	}
	return result
}




