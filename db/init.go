package db

import (
	"database/sql"
	"fmt"

	//引入数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = ""
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "LocalOJ"
)

//DB is the interface for mysql operating
var DB *sql.DB

//InitDB Connect mysql
func InitDB() error {
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/?charset=utf8")
	//设置数据库最大连接数
	if err != nil {
		fmt.Printf("Connect database failed:%v\n", err)
		return err
	}
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err = DB.Ping(); err != nil {
		fmt.Printf("opon database failed:%v\n", err)
		return err
	}
	fmt.Println("connnect database successed")
	return nil
}

func init() {
	if err := InitDB(); err != nil {
		panic("InitDB error")
	}
}
