package cgo

import (
	"database/sql"
	"strings"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

//数据库的配置
const (
	username = "root"
	password = "root"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "cgo"
	driverName = "mysql"
)

//DB数据库连接池
var DB *sql.DB

func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{username,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8&parseTime=True&loc=Local"},"")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB,_ = sql.Open(driverName,path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping();err != nil {
		log.Panic(err)
	}
	log.Println("database connect success")
}


