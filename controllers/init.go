package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"net/http"
	"wr_web/utils"
	"net/url"
)

/*
	基础的Controller
	重写beeGo的JSON方法，最后返回的JSON数据会进行base64编码
	与APP进行JSON交互，需要另写方法
*/

type InitController struct {
	beego.Controller
	Uid int
}

func (c *InitController) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = true
		hasEncoding = false
	)
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	c.JSON(c.Data["json"], hasIndent, hasEncoding)
}

func (c *InitController) JSON(data interface{}, hasIndent bool, coding bool) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(c.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	ip := c.Ctx.Input.IP()
	requestBody, _ := url.QueryUnescape(string(c.Ctx.Input.RequestBody))
	apiLog.Println("请求地址：", c.Ctx.Input.URI(), "RequestBody：", requestBody, "ResponseBody", string(content), "IP：", ip)
	if coding {
		content = []byte(utils.StringsToJSON(string(content)))
	}
	// app接口都是页面接口，不需要处理加密问题。
	// H5接口返回json需要base64编码
	if utils.H5Encoded == "true" {
		return c.Ctx.Output.Body(utils.Base64Encrypt(content))
	}
	return c.Ctx.Output.Body(content)
}