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
	if serviceCode != "" && uid != "" {
		msg, err := models.ReadFile("/Users/tzduan/Other/java/MockServer/userdata/" + uid + "/" + serviceCode + ".txt")
		// _, _ := models.ReadAll("/Users/tzduan/Other/java/MockServer/userdata/" + uid + "/" + serviceCode + ".txt")
		if err == nil {
			c.Ctx.WriteString(msg)
		} else {
			c.Ctx.WriteString("")
		}
	} else {
		c.Ctx.WriteString("Success")
	}

}
