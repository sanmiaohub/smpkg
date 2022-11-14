package smpkg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDBInst(user string, password string, host string, port string, dbname string) *sql.DB {
	dbFormat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, dbname, "utf8")

	var err error
	dbInst, err := sql.Open("mysql", dbFormat)
	if err != nil {
		panic(any("数据库链接异常:" + err.Error()))
	}
	return dbInst
}
