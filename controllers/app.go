package controllers

import (
	"zcm_tools/log"
)

/*
  APP接口的基础Controller
  APP接口的Controller都应当继承该Controller

  前置方法：对APP的请求进行Des3Base64解密
  提供Uid属性以接收通过Token参数查询出的用户ID
*/

var apiLog *log.Log

type APPController struct {
	InitController
	Uid int
}

func init() {
	apiLog = log.Init("20060102.api")
}

func (c *APPController) Prepare() {
	/*	token := c.GetString("Token")
		user, err := models.GetUsersByToken(token) //cache.GetUsersByTokenCache(m.Token)
		if err != nil && err.Error() == utils.ErrNoRow() {
			cache.RecordNewLogs(0, "0", "此账号未注册，请先行注册！"+err.Error()+" request:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "HTMLController")
			c.JSON(map[string]interface{}{"ret": 408, "err": "此账号未注册，请先行注册！"}, false, false)
			c.StopRun()
		}
		c.Uid = user.Id*/
	/*if c.Ctx.Input.Method() != "HEAD" && !c.Ctx.Input.IsUpload() {
		var m models.BaseModels
		if c.Ctx.Input.Method() == "POST" {
			if len(c.Ctx.Input.RequestBody) != 0 {
				if utils.APIEncoded=="true" {
					c.Ctx.Input.RequestBody = utils.DesBase64Decrypt(c.Ctx.Input.RequestBody)
				}
				err := json.Unmarshal(c.Ctx.Input.RequestBody, &m)
				if err != nil {
				    fmt.Println(err)
					cache.RecordNewLogs(0, "0", "base解析信息失败不存在"+err.Error()+" request:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "HTMLController")
					c.JSON(map[string]interface{}{"ret": 408, "err": "解析信息失败"}, false, false)
					c.StopRun()
				}
			}
			if m.IsEmulator {
				c.JSON(map[string]interface{}{"ret": 408, "err": "设备异常,请重新登录"}, false, false)
				c.StopRun()
			}
			if m.Token != "" && m.MobileType != "WrSystem" {
				//user, err := cache.GetUsersByIdCache(m.Uid)
				user, err := models.GetUsersByToken(m.Token) //cache.GetUsersByTokenCache(m.Token)
				if err != nil && err.Error() == utils.ErrNoRow() {
					cache.RecordNewLogs(0, "0", "此账号未注册，请先行注册！"+err.Error()+" request:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "HTMLController")
					c.JSON(map[string]interface{}{"ret": 408, "err": "此账号未注册，请先行注册！"}, false, false)
					c.StopRun()
				}
				if err != nil && err.Error() != utils.ErrNoRow() {
					cache.RecordNewLogs(0, "0", "查询用户信息发生错误！"+err.Error()+" request:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "HTMLController")
					c.JSON(map[string]interface{}{"ret": 408, "err": "服务器异常，请稍后再试！"}, false, false)
					c.StopRun()
				}
				u, err := cache.GetUsersByAccountCache(m.Account)
				if u == nil {
					cache.RecordNewLogs(0, "0", "用户信息为空"+err.Error(), c.Ctx.Input.IP(), "HTMLController")
					c.JSON(map[string]interface{}{"ret": 408, "err": "用户不存在"}, false, false)
					c.StopRun()
				}
				if user == nil {
					c.JSON(map[string]interface{}{"ret": 408, "err": "该账号在另一设备登录"}, false, false)
					c.StopRun()
				}
				c.Uid = user.Id
				if m.App == 3 {
					if user.WxToken != m.WxToken {
						c.JSON(map[string]interface{}{"ret": 408, "err": "该账号在另一设备登录"}, false, false)
						c.StopRun()
					}
				} else {
					if user.Token != m.Token {
						c.JSON(map[string]interface{}{"ret": 408, "err": "该账号在另一设备登录"}, false, false)
						c.StopRun()
					}
				}
			}
		} else {
			c.JSON(map[string]interface{}{"ret": 408, "err": "无效的请求"}, false, false)
			c.StopRun()
		}
	}*/
}
