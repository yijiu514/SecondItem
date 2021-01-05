package models

import (
	"newtest/pkg"
	"strings"

	"github.com/sirupsen/logrus"
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

// 初始化数据库
func init() {

	//从配置文件获取db信息
	_, db, err := pkg.TomlRead()
	if err != nil {
		logrus.WithError(err).Warn("toml read wrong")
	}

	//连接数据库
	path := strings.Join([]string{db.UserName, ":", db.PassWord, "@tcp(", db.IP, ":", db.Port, ")/", db.DBName, "?charset=utf8"}, "")
	DB, _ = sqlx.Open("mysql", path)
	err = DB.Ping()
	if err != nil {
		logrus.WithError(err).Warn("mysql  ping wrong")
		return
	}
	logrus.Info("mysql connet success")

	//创建数据库表
	_, err = DB.Exec(tableCreat)
	if err != nil {
		logrus.WithError(err).Warn("table creat wrong")
		return
	}
	logrus.Info("table create success")
}
