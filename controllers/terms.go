package controllers

import "github.com/astaxie/beego"

/*存放协议相关页面*/
type TermsController struct {
	beego.Controller
}

// VIP协议
// @router /vipterm [get]
func (c *TermsController) VipTerm() {
	c.TplName = "terms/PaidServiceTeams.html"
}
