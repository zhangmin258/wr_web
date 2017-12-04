package routers

import (
	"github.com/astaxie/beego"
	"wr_web/controllers"
)

func init() {
	beego.MyAutoRouter(&controllers.WalletController{})
	beego.MyAutoRouter(&controllers.VIPController{})
	beego.MyAutoRouter(&controllers.SignController{})
	beego.MyAutoRouter(&controllers.SignHomeController{})
	beego.MyAutoRouter(&controllers.SignH5Controller{})
	beego.MyAutoRouter(&controllers.ScoreExchangeController{})   //积分商城页面
	beego.MyAutoRouter(&controllers.NoBaseExchangeController{})  //跳转积分商城页面
	beego.MyAutoRouter(&controllers.TermsController{})           //跳转存放协议相关页面
	beego.MyAutoRouter(&controllers.LoanPlanServiceController{}) //跳转贷款稳下页面
	beego.MyAutoRouter(&controllers.LoanStabController{})        //贷款稳下页面
	beego.MyAutoRouter(&controllers.RuleController{})            //贷款稳下页面
}
