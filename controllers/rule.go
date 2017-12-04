package controllers

import (
	"github.com/astaxie/beego"
)

type RuleController struct {
	beego.Controller
}

//贷款稳下
// @router /vipterm [get]
func (c *RuleController) LoanSureDown() {
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
	c.TplName = tplName + "loanSureDown.html"
}

//新手必读
// @router /prsc [get]
func (c *RuleController) Prsc() {
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
	c.TplName = tplName + "PRSC.html"
}

//贷款攻略
// @router /loanstrategy [get]
func (c *RuleController) LoanStrategy() {
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
	c.TplName = tplName + "loanStrategy.html"
}

//邀请规则
// @router /inviterules [get]
func (c *RuleController) InviteRules() {
	xsrfToken := c.XSRFToken()
	c.Data["xsrf_token"] = xsrfToken
	token := c.GetString("Token")
	c.Data["token"] = token
	c.TplName = "terms/inviteRules.html"
}

//常见问题
// @router /faq [get]
func (c *RuleController) Faq() {
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
	c.TplName = tplName + "FAQ.html"
}

//贷款咨询服务协议
// @router /counselteams [get]
func (c *RuleController) CounselTeams() {
	xsrfToken := c.XSRFToken()
	c.Data["xsrf_token"] = xsrfToken
	token := c.GetString("Token")
	c.Data["token"] = token
	c.TplName = "terms/counselTeams.html"
}

//拒贷返现活动
//@router /activityrules [get]
func (c *RuleController) ActivityRules() {
	xsrfToken := c.XSRFToken()
	c.Data["xsrf_token"] = xsrfToken
	token := c.GetString("Token")
	c.Data["token"] = token
	c.TplName = "terms/activityRules.html"
}
