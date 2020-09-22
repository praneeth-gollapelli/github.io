package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

//BaseController handles Data
type BaseController struct {
	beego.Controller
}

//ParseReqBody - RequestBody to JSON
func (dc BaseController) ParseReqBody() map[string]interface{} {
	data := make(map[string]interface{})
	json.Unmarshal(dc.Ctx.Input.RequestBody, &data)
	return data
}

//ParseReqBodyAsXML - RequestBody to xml string
func (dc BaseController) ParseReqBodyAsXML() string {
	return string(dc.Ctx.Input.RequestBody[:])
}

//GetAllParams - Get All Query Params as a map
func (dc BaseController) GetAllParams() map[string]interface{} {
	if dc.Ctx.Request.Form == nil {
		dc.Ctx.Request.ParseForm()
	}
	params := make(map[string]interface{}, 0)
	for k, v := range dc.Ctx.Request.Form {
		params[k] = v[0]
	}
	return params
}

//CORS - Handle Options requests
func (dc *BaseController) CORS() {
	dc.ServeJSON()
}

//GetQueryParam --
func (dc BaseController) GetQueryParam(key string) string {
	return dc.Ctx.Input.Query(fmt.Sprintf(":%s", key))
}
