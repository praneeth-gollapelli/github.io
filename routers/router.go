package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	beegoControllers "github.io/beego-util/controllers"
	"github.io/controllers"
)

//CORSFilter - CORS filter
var CORSFilter = cors.Allow(&cors.Options{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"content-type", "Authorization", "X-Requested-With", "Cache-Control", "Pragma"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
})

func init() {
	beego.InsertFilter("*", beego.BeforeExec, CORSFilter)
	beego.Router("*", &beegoControllers.BaseController{}, "options:CORS")

	beego.Router("/create", &controllers.Controller{}, "post:Create")
	beego.Router("/update/:id", &controllers.Controller{}, "put:Update")
	beego.Router("/delete/:id", &controllers.Controller{}, "delete:Destroy")
	beego.Router("/get", &controllers.Controller{}, "get:Find")
}
