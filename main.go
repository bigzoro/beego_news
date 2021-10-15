package main

import (
	_ "news/models"
	_ "news/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
