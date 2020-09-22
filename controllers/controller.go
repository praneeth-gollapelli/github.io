package controllers

import (
	"github.io/beego-util/controllers"
	"github.io/services"
	"github.io/util"
)

//Controller ...
type Controller struct {
	controllers.BaseController
}

//Create
func (dc *Controller) Create() {
	err := services.Create(dc.ParseReqBody())
	if err != nil {
		dc.Data["json"] = util.MakeFailureResponse(err)
		dc.ServeJSON()
		return
	}
	dc.Data["json"] = util.MakeSuccessResponse()
	dc.ServeJSON()
	return
}

//Update
func (dc *Controller) Update() {
	query := make(map[string]interface{})
	query["_id"] = dc.Ctx.Input.Param(":id")
	err := services.Update(query, dc.ParseReqBody())
	if err != nil {
		dc.Data["json"] = util.MakeFailureResponse(err)
		dc.ServeJSON()
		return
	}
	dc.Data["json"] = util.MakeSuccessResponse()
	dc.ServeJSON()
	return
}

//Destroy
func (dc *Controller) Destroy() {
	query := make(map[string]interface{})
	query["_id"] = dc.Ctx.Input.Param(":id")
	err := services.Destroy(query)
	if err != nil {
		dc.Data["json"] = util.MakeFailureResponse(err)
		dc.ServeJSON()
		return
	}
	dc.Data["json"] = util.MakeSuccessResponse()
	dc.ServeJSON()
	return
}

//Find
func (dc *Controller) Find() {
	result := services.Find(dc.GetAllParams())
	dc.Data["json"] = util.MakeCustomSuccessResponse(result)
	dc.ServeJSON()
	return
}
