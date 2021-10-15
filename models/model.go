package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	UserName string
	Password string
}

type Article struct {
	Id      int
	Title   string
	Content string
	Img     string
	Type    string
	Time    time.Time
	Count   int
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/news")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}
