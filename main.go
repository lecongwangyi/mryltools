package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	d "mryltools/mysql"
)

var (
	db             *sql.DB
	driverName     = "mysql"
	dataSourceName = "root:Mingriyiliao8866@tcp(172.17.167.56:3306)/tm_clinic?charset=utf8mb4"
)

func init() {
	db, _ = sql.Open(driverName, dataSourceName)
	// if err != nil {

	// }
	//连接池
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(1)
	db.Ping()

	// return db

}

/**
 * 得到mysql连接
 */

func main() {
	fmt.Println("WO DE ", db)
	d.Insert(db, "", "")
}
