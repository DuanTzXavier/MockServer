package controllers

import (
	"MockServer/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	uid := c.GetString("uid")
	serviceCode := c.GetString("servicecode")
	println("uid: " + uid)
	println("ServiceCode: " + serviceCode)
	c.Data["Json"] = uid
	c.Ctx.WriteString(models.GetUIDList("hahah"))
}
