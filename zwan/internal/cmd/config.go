package main

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/emart.io/cross/zwan/internal/impl/biz/db"
	_ "github.com/glebarez/go-sqlite"
)

const (
	port     = "8443"
	certFile = "fullchain.cer"
	keyFile  = "iyou.city.key"
)

// var (
// 	mysqlServiceName = "mysql_emart"
// 	mysqlServicePort = "3306"
// )

func init() {
	// if msn, ok := os.LookupEnv("MYSQL_SERVICE_NAME"); ok {
	// 	mysqlServiceName = msn
	// }
	// if msp, ok := os.LookupEnv("MYSQL_SERVICE_PORT"); ok {
	// 	mysqlServicePort = msp
	// }
	//db.Open(fmt.Sprintf("root:123456@tcp(%s:%s)/emart", mysqlServiceName, mysqlServicePort))
	db.Open("sqlite", "emart.db")
}
