package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true",
		"chengchao",
		"over2012",
		"192.168.56.102",
		3306,
		"chaos_test",
		url.QueryEscape("Asia/Shanghai"))
	//dsn := "chengchao:over2012@tcp(192.168.56.102:3306)/chaos_test?charset=utf8&loc=%s&parseTime=true",
	var err error
	db, err = sql.Open("mysql", uri)
	if err != nil {
		log.Fatalln(err)
	}
}

// GetSqlDB just get DB
func GetSqlDB() *sql.DB {
	return db
}
