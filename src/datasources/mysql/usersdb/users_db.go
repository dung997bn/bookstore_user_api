package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	//_ immport
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	//Client type
	Client *sql.DB
	// username = os.Getenv(mysqlUsersUsername)
	// password = os.Getenv(mysqlUsersPassword)
	// host     = os.Getenv(mysqlUsersHost)
	// schema   = os.Getenv(mysqlUsersSchema)
	username = "root"
	password = "cntt"
	host     = "127.0.0.1:3306"
	schema   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	fmt.Println(fmt.Sprintf("about to connect to %s", dataSourceName))
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
