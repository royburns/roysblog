package main

import (
	"github.com/astaxie/beego"
	_ "github.com/royburns/roysblog/models"
	_ "github.com/royburns/roysblog/routers"
)

func main() {
	beego.Run()
}
