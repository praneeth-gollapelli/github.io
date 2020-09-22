package main

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.io/routers"
)

func main() {
	log.Println("Api up and running.......")
	beego.Run()
}
