package models

import (
	"fmt"
	"strings"

	//引入mysql数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User user表基本信息
type User struct {
	ID           int
	Email        string
	Role         string
	Creatat      int64
	Lockat       int64
	Password     string
	Passwordsalt string
	Sessionsalt  string
}

var tableCreat = `CREATE TABLE if not exists User (
					id             INTEGER,
					email          VARCHAR(40),
					role           VARCHAR(40),
					creatat       INT(64),
					lockat        INT(64),
					password       VARCHAR(40),
					passwordsalt  VARCHAR(40),
					sessionsalt   VARCHAR(40),
					PRIMARY KEY (email)
					)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

// DB 数据库指针
var DB *sqlx.DB

// TomlConfig 定义配置文件结构体
type TomlConfig struct {
	WebPort int64
	DB      Database `toml:"database"`
}

// Database 定义数据库信息结构体
type Database struct {
	UserName string
	PassWord string
	IP       string
	Port     string
	DBName   string
}

//DBInit 初始化数据库
func DBInit(db Database) (err error) {
	//连接数据库
	path := strings.Join([]string{db.UserName, ":", db.PassWord, "@tcp(", db.IP, ":", db.Port, ")/", db.DBName, "?charset=utf8"}, "")
	DB, _ = sqlx.Open("mysql", path)
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("DB ping wrong %m", err)
	}

	//创建数据库表
	_, err = DB.Exec(tableCreat)
	if err != nil {
		return fmt.Errorf("table creat wrong %x", err)
	}

	return nil
}
