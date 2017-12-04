package controllers

import (
	"wr_web/models"
	"wr_web/utils"
	"strconv"
	"wr_web/cache"
)

/*
	VIP
*/

type VIPController struct {
	H5Controller
}

// 查询该用户是否为VIP，并返回VIP信息
// @router /getvipinfo [post]
func (this *VIPController) GetVIPInfo() {
	defer this.ServeJSON()
	vip,err:=models.GetVIPInfo(this.Uid)// 获取VIP信息
	if err!=nil && err.Error()==utils.ErrNoRow() {// 没有该用户数据，进行修复数据
		err:=models.InnitVIP(this.Uid)
		if err!=nil {
			this.Data["json"] = map[string]interface{}{"ret": 403,"msg":"初始化用户VIP数据失败!"}
			cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "初始化用户VIP数据失败"+err.Error(), this.Ctx.Input.IP(), "vip/getvipinfo")
			return
		}
		// 返回
		this.Data["json"] = map[string]interface{}{"ret": 200,"isvip":"N"}
		return
	}
	if err!=nil && err.Error()!=utils.ErrNoRow() {
		this.Data["json"] = map[string]interface{}{"ret": 403,"msg":"查询用户VIP数据失败!"}
		cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "查询用户VIP数据失败"+err.Error(), this.Ctx.Input.IP(), "vip/getvipinfo")
		return
	}
	// 返回
	this.Data["json"] = map[string]interface{}{"ret": 200,"isvip":"Y","endtime":vip.EndTime.Format(utils.FormatDate)}
	return
}
