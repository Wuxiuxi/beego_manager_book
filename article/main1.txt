package main

import (
	_ "article/models"
	_ "article/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
