package utility

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// ConnectSQL Connect to database
func ConnectSQL() (*DB, error) {
	loading, _ := ReadConfigFile([]string{"database.default.link"} /*, "manifest/config/"*/)
	// connect to database
	dbData, err := sql.Open("mysql", loading["database.default.link"])
	if err != nil {
		log.Println("<==== navigation:数据库连接异常 ====>", err)
	} else {
		log.Println("<++++ navigation:数据库连接成功 ++++>")
	}

	// Check database connection
	if err = dbData.Ping(); err != nil {
		log.Println("<==== navigation:数据库检查失败 ====>", err)
	} else {
		log.Println("<++++ navigation:数据库检查通过 ++++>")
	}

	dbConn.SQL = dbData

	return dbConn, err
}
