package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"wr_web/cache"
	"wr_web/models"
	//"wr_web/services"
	"wr_web/utils"
	"zcm_tools/email"
)

type LoanPlanServiceController struct {
	beego.Controller
}

type LoanStabController struct {
	H5Controller
}

/*
	贷款稳下页面
*/

//广告轮播
//@router /getadvofloan [post]
func (this *LoanStabController) GetAdvOfLoan() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	adv, err := models.GetAdvByConfigKey()
	if err != nil {
		cache.RecordNewLogs(this.Uid, "0", "查询轮播广告异常"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "loanstab/getadvofloan")
		resultMap["err"] = "查询轮播广告失败"
		return
	}
	resultMap["ret"] = 200
	resultMap["adv"] = adv
	return
}

//立即申请贷款稳下
//@router /loanplanservice [post]
func (this *LoanStabController) LoanPlanService() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	var chargeCall models.ChargeCall
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &chargeCall)
	if err != nil {
		cache.RecordNewLogs(0, "0", "解析参数错误"+err.Error()+";param:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "loanstab/loanplanservice")
		resultMap["err"] = "解析参数错误！"
		return
	}
	// 判断是否有付款TOKEN
	if !utils.CheckPay(this.Uid, chargeCall.ServiceId, chargeCall.PayToken) {
		resultMap["err"] = "尚未支付!"
		cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "未找到该业务对应的付款信息!"+err.Error(), this.Ctx.Input.IP(), "loanstab/loanplanservice")
		return
	}
	err = models.AddLoanPlanRecord(this.Uid, 1, 1, chargeCall.PayToken)
	if err != nil {
		email.Send("记录用户贷款预约信息异常", "Uid:"+strconv.Itoa(this.Uid)+";err:"+err.Error(), utils.ToUsers, "weirong")
		cache.RecordNewLogs(this.Uid, "0", "记录用户贷款预约信息异常"+err.Error(), this.Ctx.Input.IP(), "loanstab/loanplanservice")
		resultMap["err"] = "信息记录异常!"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "专属贷款顾问跟您电话联系，请保持通讯通畅！"
	return
}
