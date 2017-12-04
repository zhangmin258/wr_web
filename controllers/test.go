package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"wr_web/cache"
	"wr_web/models"
)

type TestController struct {
	//beego.Controller
	H5Controller
}

// @router /gettask [get]
func (c *TestController) GetTest() {
	xsrfToken := c.XSRFToken()
	c.Data["xsrf_token"] = xsrfToken
	token := c.GetString("Token")
	c.Data["token"] = token
	packageId, _ := c.GetInt("PackageId")
	tplName := ""
	switch packageId {
	case 1:
		tplName = "package1/"
	case 2:
		tplName = "package2/"
	case 8:
		tplName = "package8/"
	case 9:
		tplName = "package9/"
	default:
		tplName = "package1/"
	}
	c.Data["PackageId"] = packageId
	c.TplName = tplName + "myTask.html"
}

//获取累计签到天数
// @router /gettestsigncount [post]
func (c *TestController) GetTestSignCount() {
	beego.Info("enter test/getsigncount")
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	signCount, err := models.GetSignCount(c.Uid, time.Now().Year(), int(time.Now().Month()))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取累计签到天数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsigncount")
		resultMap["err"] = "获取累计签到天数失败"
		return
	}
	resultMap["SignCount"] = signCount
	resultMap["ret"] = 200
}
