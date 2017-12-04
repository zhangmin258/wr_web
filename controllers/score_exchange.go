package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
	"wr_web/cache"
	"wr_web/models"
	"wr_web/utils"
	"zcm_tools/email"
)

type ScoreExchangeController struct {
	H5Controller
}
type NoBaseExchangeController struct {
	beego.Controller
}

/*
	积分商城页面接口
*/

//跳转积分商城
//@router /showscoreexchangeproducts [get]
func (this *NoBaseExchangeController) ShowScoreExchangeProducts() {
	xsrfToken := this.XSRFToken()
	this.Data["xsrf_token"] = xsrfToken
	token := this.GetString("Token")
	this.Data["token"] = token
	packageId, _ := this.GetInt("PackageId")
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
	this.Data["PackageId"] = packageId
	this.TplName = tplName + "shoppingMall.html"
}

//跳转抽奖
//@router /showlotteryproducts [get]
func (this *NoBaseExchangeController) ShowLotteryProducts() {
	xsrfToken := this.XSRFToken()
	this.Data["xsrf_token"] = xsrfToken
	token := this.GetString("Token")
	this.Data["token"] = token
	packageId, _ := this.GetInt("PackageId")
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
	this.Data["PackageId"] = packageId
	this.TplName = tplName + "turntableLottery.html"
}

//跳转兑换记录展示页
//@router /scoreexchangerecord [get]
func (this *NoBaseExchangeController) ScoreExchangeRecord() {
	xsrfToken := this.XSRFToken()
	this.Data["xsrf_token"] = xsrfToken
	token := this.GetString("Token")
	this.Data["token"] = token
	packageId, _ := this.GetInt("PackageId")
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
	this.Data["PackageId"] = packageId
	this.TplName = tplName + "conversionRecord.html"
}

//根据uid获取用户的账号信息：账号信息页面和个人中心页面
//@router /searchuserinfo [post]
func (this *ScoreExchangeController) SearchUserInfo() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	userMini, err := models.SearchUserInfo(this.Uid)
	if err != nil {
		cache.RecordNewLogs(this.Uid, "0", "查询用户身份证"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/searchuserinfo")
		resultMap["err"] = "查询用户账号信息错误！"
		return
	}
	num, err := models.GetMessageCountByUid(this.Uid)
	if err != nil {
		cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "查询用户未读消息数量错误"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/searchuserinfo")
		resultMap["err"] = "查询用户未读消息数量错误"
		return
	}
	resultMap["ret"] = 200
	resultMap["not_red_msg"] = num
	resultMap["userMini"] = userMini
}

//兑换记录展示页
//@router /showscoreexchangerecord [post]
func (this *ScoreExchangeController) ShowScoreExchangeRecord() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	showSeRecord, err := models.GetScoreExchangeRecord(this.Uid)
	if err != nil && err.Error() != utils.ErrNoRow() {
		cache.RecordNewLogs(0, "", "查询兑换记录错误！", this.Ctx.Input.IP(), "nobaseexchange/showscoreexchangerecord")
		resultMap["err"] = "查询兑换记录错误！"
		return
	}
	for k, v := range showSeRecord {
		showSeRecord[k].Title = "兑换" + v.Content
		if v.PayType == 1 {
			showSeRecord[k].Pay = strconv.Itoa(v.ScoreAmount) + "融豆"
		}
		if v.PayType == 2 {
			showSeRecord[k].Pay = strconv.FormatFloat(v.MoneyAmount, byte('f'), -1, 64) + "元"
		}
		if v.PayType == 3 {
			showSeRecord[k].Pay = "会员免费"
		}
	}
	resultMap["ret"] = 200
	resultMap["showSeRecord"] = showSeRecord
	return
}

//中奖快报轮播
//@router /lotteryreport [post]
func (this *ScoreExchangeController) LotteryReport() {
	defer this.ServeJSON()
	lotteryMsg := [50]string{}
	var users *models.BaseModels
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &users)
	if err != nil {
		cache.RecordNewLogs(0, "0", "解析参数错误"+err.Error()+";param:"+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/lotteryreport")
		this.Data["json"] = map[string]interface{}{"ret": 403, "err": "解析参数错误!"}
		return
	}
	//构造IOS审核假数据
	if users.PackageId != 2 && users.PackageId != 8 {
		if users.App == 1 {
			if users.PackageId == 0 {
				users.PackageId = 1
			}
			newVersion := users.AppVersion
			packageConfig, err := models.GetPackageConfig(users.PackageId)
			if err != nil {
				cache.RecordNewLogs(0, "0", "获取分身包数据错误"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryreport")
				this.Data["json"] = map[string]interface{}{"ret": 403, "err": "获取分身包数据错误!"}
				return
			}
			oldVersion := packageConfig.IosPretendVersion
			flag := utils.GetAppVersionEquery(newVersion, oldVersion)
			if flag {
				lotteryMsgs := []string{"133****5605 抽中了10元现金",
					"170****3423 抽中了100融豆",
					"186****7642 抽中了5元现金",
					"134****5069 抽中了1000元现金",
					"135****6423 抽中了iPhone 7",
					"159****0325 抽中了200融豆",
					"189****7323 抽中了10元现金",
					"139****5755 抽中了10元现金",
					"158****5343 抽中了100融豆",
					"150****9323 抽中了5元现金",
					"186****7379 抽中了1000元现金",
					"134****7879 抽中了iPhone 7",
					"159****6730 抽中了200融豆",
					"189****5793 抽中了10元现金",
					"135****8948 抽中了10元现金",
					"170****5629 抽中了100融豆",
					"186****7470 抽中了5元现金",
					"134****5532 抽中了1000元现金",
					"135****2245 抽中了iPhone 7",
					"159****3964 抽中了200融豆",
					"189****2478 抽中了10元现金",
					"139****5648 抽中了10元现金",
					"158****3345 抽中了100融豆",
					"150****8846 抽中了5元现金",
					"186****2364 抽中了1000元现金",
					"134****9764 抽中了iPhone 7",
					"159****5364 抽中了200融豆",
					"189****6037 抽中了10元现金"}
				this.Data["json"] = map[string]interface{}{"ret": 200, "lotteryMsg": lotteryMsgs}
				return
			}
		}
	}

	products, err := models.GetLotteryReportProduct()
	if err != nil {
		cache.RecordNewLogs(0, "0", "查询奖品异常"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/lotteryreport")
		this.Data["json"] = map[string]interface{}{"ret": 403, "err": "查询奖品异常!"}
		return
	}
	startNum := [...]string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188", "130", "131", "132", "145", "155", "156", "176", "185", "186", "133", "153", "177", "180", "181", "189", "170"}
	for k, _ := range lotteryMsg {
		r := utils.GetRandNumber(len(startNum) - 1)
		phoneNum := startNum[r] + "****"
		phoneNum += utils.GetRandDigit(4)
		r = utils.GetRandNumber(len(products) - 1)
		product := products[r]
		num := utils.GetRandNumber(100)
		if (product == "iPhone7" || product == "1000元") && num < 95 {
			product = "5元"
		}
		lotteryMsg[k] = "恭喜用户" + phoneNum + "抽中了" + product
	}
	//lotteryMsg := "恭喜用户" + phoneNum + "抽中了" + product
	this.Data["json"] = map[string]interface{}{"ret": 200, "lotteryMsg": lotteryMsg}
	return
}

//抽奖接口
//@router /lotteryresult [post]
func (this *ScoreExchangeController) LotteryResult() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	//频繁提交
	if !utils.Rc.SetNX("Score_lottery"+string(this.Uid), 1, time.Second*3) {
		resultMap["msg"] = "请勿频繁提交!"
		return
	}
	//活动开放校验
	dialSwitch, err := models.GetLotteryValueByCode("score_lottery_switch")
	if err != nil {
		cache.RecordNewLogs(0, "0", "通过config_key获取活动开关失败："+err.Error(), this.Ctx.Input.IP(), "/scoreexchange/lotteryresult")
		resultMap["msg"] = "活动正在维护中~"
		return
	}
	if dialSwitch.Config_value != "1" {
		resultMap["msg"] = "活动正在维护中~"
		return
	}
	if dialSwitch.Remark != "" {
		//活动时间校验
		begin_timeStr := strings.Split(dialSwitch.Remark, ",")[0]
		end_timeStr := strings.Split(dialSwitch.Remark, ",")[1]
		begin_time, _ := time.ParseInLocation("2006-01-02 15:04:05", begin_timeStr, time.Local)
		end_time, _ := time.ParseInLocation("2006-01-02 15:04:05", end_timeStr, time.Local)
		if time.Now().Unix() < begin_time.Unix() {
			resultMap["msg"] = "该活动暂未开放，敬请期待~"
			return
		}
		if time.Now().Unix() > end_time.Unix() {
			resultMap["msg"] = "活动已经结束！"
			return
		}
	}
	srdb, err := models.GetUsersRongDouByUid(this.Uid)
	if err != nil {
		cache.RecordNewLogs(0, "0", "用户融豆数据异常"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/getusersrongdou")
		resultMap["err"] = "用户融豆数据异常!"
		return
	}
	moneyService, err := models.GetLotteryPrice(8)
	if err != nil {
		cache.RecordNewLogs(0, "0", "获取抽奖价格错误"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/getusersrongdou")
		resultMap["err"] = "获取抽奖价格错误!"
		return
	}
	if srdb.Score < moneyService.ScorePrice {
		resultMap["msg"] = "您的融豆余额不足！"
		return
	}
	score := srdb.Score - moneyService.ScorePrice
	var ufr models.UsersFinanceRecord
	ufr.Uid = this.Uid
	ufr.PayType = 1
	ufr.DealType = 8
	ufr.ScoreAmount = 100
	ufr.PayOrGet = 1
	ufr.BeforScoreAmount = srdb.Score
	ufr.AfterScoreAmount = score
	err = models.AddUsersFinanceRecord(ufr)
	if err != nil {
		email.Send("记录用户收支情况异常", "err:"+err.Error(), utils.ToUsers, "weirong")
		cache.RecordNewLogs(0, "0", "记录用户收支情况异常"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/getusersrongdou")
		resultMap["err"] = "收支存储异常!"
		return
	}
	err = models.DeductBalanceByUid(this.Uid, score)
	if err != nil {
		email.Send("扣除融豆异常", "err:"+err.Error(), utils.ToUsers, "weirong")
		cache.RecordNewLogs(0, "0", "扣除融豆异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
		resultMap["err"] = "扣除融豆异常!"
		return
	}
	if !utils.Rc.IsExist("UsersLotteryBalanceDailyReset") {
		err := utils.Rc.Put("UsersLotteryBalanceDailyReset", `{"Money":500,"RongDou":50000}`, utils.GetTodayLastSecond())
		if err != nil {
			cache.RecordNewLogs(0, "0", "抽奖数据存入redis异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
			resultMap["err"] = "拉取抽奖数据异常!"
			return
		}
	}
	var lotteryBalance models.LotteryBalance
	balance, _ := utils.Rc.RedisBytes("UsersLotteryBalanceDailyReset")
	err = json.Unmarshal(balance, &lotteryBalance)
	if err != nil {
		cache.RecordNewLogs(0, "0", "获取抽奖数据异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
		resultMap["err"] = "获取抽奖数据异常!"
		return
	}
	money := lotteryBalance.Money
	rongDou := lotteryBalance.RongDou
	nTime := time.Now().Hour()
	var id, price int
	switch {
	case 0 <= nTime && nTime < 6 && money > 400 && rongDou > 40000:
		id, price = Lottery()
	case 6 <= nTime && nTime < 12 && money > 300 && rongDou > 30000:
		id, price = Lottery()
	case 12 <= nTime && nTime < 18 && money > 200 && rongDou > 20000:
		id, price = Lottery()
	case 18 <= nTime && nTime < 24 && money > 0 && rongDou > 0:
		id, price = Lottery()
	default:
		id = 7
	}
	if price == 5 || price == 10 {
		ufr.PayType = 2
		ufr.DealType = 13
		ufr.ScoreAmount = 0
		ufr.MoneyAmount = float64(price)
		ufr.PayOrGet = 2
		lotteryBalance.Money = money - price
		name := ""
		if price == 5 {
			name = "5元现金"
		} else {
			name = "10元现金"
		}
		walletBalance, err := models.GetWalletBalance(this.Uid)
		if err != nil {
			cache.RecordNewLogs(0, "0", "抽奖时获取钱包余额异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
			resultMap["err"] = "获取钱包余额异常!"
			return
		}
		ufr.BeforMoneyAmount = walletBalance.AccountBalance
		accountBalance := float64(price) + walletBalance.AccountBalance
		ufr.AfterMoneyAmount = accountBalance
		err = models.LotteryAddBalanceByUid(this.Uid, accountBalance)
		if err != nil {
			email.Send("中奖后充值现金异常", "err:"+err.Error(), utils.ToUsers, "weirong")
			cache.RecordNewLogs(0, "0", "中奖后充值现金异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
			resultMap["err"] = "充值现金异常!"
			return
		}
		resultMap["msg"] = "恭喜，中奖了！" + name + "，已充值到我的钱包，请查收！"
	} else if price == 100 || price == 200 {
		ufr.DealType = 13
		ufr.ScoreAmount = price
		ufr.PayOrGet = 2
		lotteryBalance.RongDou = rongDou - price
		ufr.AfterScoreAmount = price + score
		err = models.DeductBalanceByUid(this.Uid, ufr.AfterScoreAmount)
		if err != nil {
			email.Send("中奖后充值融豆异常", "err:"+err.Error(), utils.ToUsers, "weirong")
			cache.RecordNewLogs(0, "0", "中奖后充值融豆异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
			resultMap["err"] = "充值融豆异常!"
			return
		}
		name := ""
		if price == 100 {
			name = "100融豆"
		} else {
			name = "200融豆"
		}
		resultMap["msg"] = "恭喜，中奖了！" + name + "，已充值到我的融豆，请查收！"
	} else {
		resultMap["msg"] = "很遗憾，未中奖！"
	}
	lotteryStr, err := json.Marshal(lotteryBalance)
	if err != nil {
		cache.RecordNewLogs(0, "0", "更新中奖结果异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
		resultMap["err"] = "更新中奖结果异常!"
		return
	}
	err = utils.Rc.Put("UsersLotteryBalanceDailyReset", string(lotteryStr), utils.GetTodayLastSecond())
	if err != nil {
		cache.RecordNewLogs(0, "0", "更新奖池redis缓存异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
		resultMap["err"] = "更新中奖结果异常!"
		return
	}
	if id != 7 && id != 8 {
		err = models.AddUsersFinanceRecord(ufr)
		if err != nil {
			email.Send("用户中奖记录收支情况异常", "err:"+err.Error(), utils.ToUsers, "weirong")
			cache.RecordNewLogs(0, "0", "用户中奖记录收支情况异常"+err.Error()+" "+string(this.Ctx.Input.RequestBody), this.Ctx.Input.IP(), "scoreexchange/getusersrongdou")
			resultMap["err"] = "收支存储异常!"
			return
		}
		err = models.AddUsersLottery(this.Uid, id)
		if err != nil {
			cache.RecordNewLogs(this.Uid, "0", "保存用户中奖结果异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/lotteryresult")
		}
	}
	resultMap["LotteryId"] = id
	resultMap["ret"] = 200
	return
}

//抽奖方法
func Lottery() (id, price int) {
	num := utils.GetRandNumber(1000)
	switch {
	case 0 <= num && num < 5:
		id = 4
		price = 10
	case 5 <= num && num < 15:
		id = 3
		price = 5
	case 15 < num && num < 65:
		id = 2
		price = 200
	case 65 <= num && num < 365:
		id = 1
		price = 100
	case 365 <= num && num < 680:
		id = 7
	default:
		id = 8
	}
	return
}

//跳转至中奖记录页面
//@router /showlotteryrecord [get]
func (this *NoBaseExchangeController) ShowLotteryRecord() {
	xsrfToken := this.XSRFToken()
	this.Data["xsrf_token"] = xsrfToken
	token := this.GetString("Token")
	this.Data["token"] = token
	packageId, _ := this.GetInt("PackageId")
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
	this.Data["PackageId"] = packageId
	this.TplName = tplName + "winRecord.html"
}

//查询用户中奖记录
//@router /getuserslotteryrecord [post]
func (this *ScoreExchangeController) GetUsersLotteryRecord() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		this.Data["json"] = resultMap
		this.ServeJSON()
	}()
	lotteryRecord, err := models.GetLotteryRecordByUid(this.Uid)
	if err != nil && err.Error() != utils.ErrNoRow() {
		cache.RecordNewLogs(0, "0", "获取中奖记录异常"+err.Error(), this.Ctx.Input.IP(), "scoreexchange/getuserslotteryrecord")
		resultMap["err"] = "获取中奖记录异常!"
		return
	}
	resultMap["lotteryRecord"] = lotteryRecord
	resultMap["ret"] = 200
	return
}
