package smpkg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDBInst(user string, password string, host string, port string, dbname string) (dbInst *sql.DB, err error) {
	dbFormat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, dbname, "utf8")
	return sql.Open("mysql", dbFormat)
}
