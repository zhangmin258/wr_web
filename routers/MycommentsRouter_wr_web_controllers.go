package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.MyGlobalControllerRouter["wr_web/controllers:LoanStabController/GetAdvOfLoan"] =beego.ControllerComments{
			"GetAdvOfLoan",
			`/getadvofloan`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:LoanStabController/LoanPlanService"] =beego.ControllerComments{
			"LoanPlanService",
			`/loanplanservice`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:NoBaseExchangeController/ScoreExchangeRecord"] =beego.ControllerComments{
			"ScoreExchangeRecord",
			`/scoreexchangerecord`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:NoBaseExchangeController/ShowLotteryProducts"] =beego.ControllerComments{
			"ShowLotteryProducts",
			`/showlotteryproducts`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:NoBaseExchangeController/ShowLotteryRecord"] =beego.ControllerComments{
			"ShowLotteryRecord",
			`/showlotteryrecord`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:NoBaseExchangeController/ShowScoreExchangeProducts"] =beego.ControllerComments{
			"ShowScoreExchangeProducts",
			`/showscoreexchangeproducts`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/ActivityRules"] =beego.ControllerComments{
			"ActivityRules",
			`/activityrules`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/CounselTeams"] =beego.ControllerComments{
			"CounselTeams",
			`/counselteams`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/Faq"] =beego.ControllerComments{
			"Faq",
			`/faq`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/InviteRules"] =beego.ControllerComments{
			"InviteRules",
			`/inviterules`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/LoanStrategy"] =beego.ControllerComments{
			"LoanStrategy",
			`/loanstrategy`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/LoanSureDown"] =beego.ControllerComments{
			"LoanSureDown",
			`/vipterm`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:RuleController/Prsc"] =beego.ControllerComments{
			"Prsc",
			`/prsc`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:ScoreExchangeController/GetUsersLotteryRecord"] =beego.ControllerComments{
			"GetUsersLotteryRecord",
			`/getuserslotteryrecord`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:ScoreExchangeController/LotteryReport"] =beego.ControllerComments{
			"LotteryReport",
			`/lotteryreport`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:ScoreExchangeController/LotteryResult"] =beego.ControllerComments{
			"LotteryResult",
			`/lotteryresult`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:ScoreExchangeController/SearchUserInfo"] =beego.ControllerComments{
			"SearchUserInfo",
			`/searchuserinfo`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:ScoreExchangeController/ShowScoreExchangeRecord"] =beego.ControllerComments{
			"ShowScoreExchangeRecord",
			`/showscoreexchangerecord`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/AddSign"] =beego.ControllerComments{
			"AddSign",
			`/addsign`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetDailyReward"] =beego.ControllerComments{
			"GetDailyReward",
			`/getdailyreward`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetDays"] =beego.ControllerComments{
			"GetDays",
			`/getdays`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetExtraReward"] =beego.ControllerComments{
			"GetExtraReward",
			`/getextrareward`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetFaceCount"] =beego.ControllerComments{
			"GetFaceCount",
			`/getfacecount`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetGrowReward"] =beego.ControllerComments{
			"GetGrowReward",
			`/getgrowreward`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetMission"] =beego.ControllerComments{
			"GetMission",
			`/getmission`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetMissionCount"] =beego.ControllerComments{
			"GetMissionCount",
			`/getmissioncount`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetRemindStatus"] =beego.ControllerComments{
			"GetRemindStatus",
			`/getremindstatus`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetScore"] =beego.ControllerComments{
			"GetScore",
			`/getscore`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetSign"] =beego.ControllerComments{
			"GetSign",
			`/getsign`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignController/GetSignCount"] =beego.ControllerComments{
			"GetSignCount",
			`/getsigncount`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignH5Controller/GetSignIn"] =beego.ControllerComments{
			"GetSignIn",
			`/getsignin`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignH5Controller/GetTask"] =beego.ControllerComments{
			"GetTask",
			`/gettask`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:SignHomeController/GetSignHome"] =beego.ControllerComments{
			"GetSignHome",
			`/getsignhome`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:TermsController/VipTerm"] =beego.ControllerComments{
			"VipTerm",
			`/vipterm`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:TestController/GetTest"] =beego.ControllerComments{
			"GetTest",
			`/gettask`,
			[]string{"get"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:TestController/GetTestSignCount"] =beego.ControllerComments{
			"GetTestSignCount",
			`/gettestsigncount`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:VIPController/GetVIPInfo"] =beego.ControllerComments{
			"GetVIPInfo",
			`/getvipinfo`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:WalletController/GetWalletInfo"] =beego.ControllerComments{
			"GetWalletInfo",
			`/getwalletinfo`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:WalletController/HasWalletPwd"] =beego.ControllerComments{
			"HasWalletPwd",
			`/haswalletpwd`,
			[]string{"post"},
			nil}

	beego.MyGlobalControllerRouter["wr_web/controllers:WalletController/SetWalletPwd"] =beego.ControllerComments{
			"SetWalletPwd",
			`/setwalletpwd`,
			[]string{"post"},
			nil}

}
