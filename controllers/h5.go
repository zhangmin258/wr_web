package controllers

import (
	"encoding/base64"
	"wr_web/cache"
	"wr_web/models"
	"wr_web/utils"
	//"encoding/json"
	"encoding/json"
)

/*
  H5接口的基础Controller
  H5接口都应当继承该Controller

  前置方法：进行XSRF校验，对请求体进行base64解码
*/
type H5Controller struct {
	InitController
}

//前置方法
func (this *H5Controller) Prepare() {
	if utils.Enablexsrf == "true" {
		token := this.Ctx.Input.Query("_xsrf")
		if token == "" {
			token = this.Ctx.Request.Header.Get("X-Xsrftoken")
		}
		if token == "" {
			token = this.Ctx.Request.Header.Get("X-Csrftoken")
		}
		//处理恶意请求
		if token == "" {
			this.Data["json"] = map[string]interface{}{"ret": 403, "err": "网络异常！"}
			this.ServeJSON()
			this.StopRun()
		}
		/*token_count := utils.Check5Time(utils.CACHE_KEY_WRWEB_OPERATION + token)
		if token_count < 0 {
			this.Data["json"] = map[string]interface{}{"ret": 403, "err": "网络异常！"}
			this.ServeJSON()
			this.StopRun()
		}*/
	}
	if utils.H5Encoded == "true" {
		// 请求体base64解码
		if this.Ctx.Input.Method() == "POST" && len(this.Ctx.Input.RequestBody) != 0 {
			var err error
			this.Ctx.Input.RequestBody, err = base64.StdEncoding.DecodeString(string(this.Ctx.Input.RequestBody))
			if err != nil {
				this.Data["json"] = map[string]interface{}{"ret": 403, "err": "网络异常！"}
				cache.RecordNewLogs(0, "0", "解析请求体失败！"+err.Error(), this.Ctx.Input.IP(), "h5")
				this.ServeJSON()
				this.StopRun()
			}
		}
	}

	// 接收token，查找uid并赋值
	var m models.BaseModels
	if this.Ctx.Input.Method() == "POST" {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &m)
		if err != nil {
			cache.RecordNewLogs(0, "0", "base解析信息失败不存在"+err.Error()+" request:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "h5")
			this.JSON(map[string]interface{}{"ret": 408, "err": "解析信息失败"}, false, false)
			this.StopRun()
		}

		if m.Token != "" {
			user, err := models.GetUsersByToken(m.Token)         // 根据token去查询用户信息
			if err != nil && err.Error() == utils.ErrNoRow() { // 没有该token对应的用户
				_, err := cache.GetUsersByAccountCache(m.Account)  // 根据account去查询用户信息
				if err != nil && err.Error() == utils.ErrNoRow() { // 没有account对应的用户信息
					cache.RecordNewLogs(0, "0", "此账号未注册，请先行注册！"+err.Error()+" request:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "h5")
					this.JSON(map[string]interface{}{"ret": 408, "err": "此账号未注册，请先行注册！"}, false, false)
					this.StopRun()
				}
				if err != nil && err.Error() != utils.ErrNoRow() {
					cache.RecordNewLogs(0, "0", "根据account查询用户信息发生错误！"+err.Error()+" request:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "h5")
					this.JSON(map[string]interface{}{"ret": 408, "err": "服务器异常，请稍后再试！"}, false, false)
					this.StopRun()
				}
				this.JSON(map[string]interface{}{"ret": 408, "err": "该账号在另一设备登录!"}, false, false)
				this.StopRun()
			}
			if err != nil && err.Error() != utils.ErrNoRow() {
				cache.RecordNewLogs(0, "0", "根据token查询用户信息发生错误！"+err.Error()+" request:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "h5")
				this.JSON(map[string]interface{}{"ret": 408, "err": "服务器异常，请稍后再试！"}, false, false)
				this.StopRun()
			}
			this.Uid = user.Id
		}
	} else {
		this.JSON(map[string]interface{}{"ret": 408, "err": "无效的请求"}, false, false)
		this.StopRun()
	}
}
