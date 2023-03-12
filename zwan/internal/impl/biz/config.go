package biz

import (
	"fmt"
	"os"

	"github.com/jmzwcn/db"
)

var (
	mysqlServiceName = "mysql_emart"
	mysqlServicePort = "3306"
)

func init() {
	if msn, ok := os.LookupEnv("MYSQL_SERVICE_NAME"); ok {
		mysqlServiceName = msn
	}
	if msp, ok := os.LookupEnv("MYSQL_SERVICE_PORT"); ok {
		mysqlServicePort = msp
	}
	db.Open(fmt.Sprintf("root:123456@tcp(%s:%s)/emart", mysqlServiceName, mysqlServicePort))
}
