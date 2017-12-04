package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"wr_web/cache"
	"wr_web/models"
	"wr_web/utils"
)

type SignController struct {
	H5Controller
}
type SignHomeController struct {
	APPController
}

type SignH5Controller struct {
	beego.Controller
}

//我的任务页面
// @router /gettask [get]
func (c *SignH5Controller) GetTask() {
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
	c.TplName = tplName + "myTask.html"
}

//签到日历页面
// @router /getsignin [get]
func (c *SignH5Controller) GetSignIn() {
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
	c.TplName = tplName + "calendarSignIn.html"
}

//签到主页
// @router /getsignhome [get]
func (c *SignHomeController) GetSignHome() {
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
	c.TplName = tplName + "signIn.html"
}

//签到
// @router /addsign [post]
func (c *SignController) AddSign() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()
	count, err := models.IsSign(c.Uid, year, month, day)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取签到数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/addsign")
		resultMap["err"] = "获取签到数据失败"
		return
	}
	if count == 1 {
		resultMap["msg"] = "已签到"
		resultMap["ret"] = 304
		return
	}
	err = models.AddSign(c.Uid, year, month, day)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "新增签到数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/addsign")
		resultMap["err"] = "新增签到数据失败"
		return
	}
	beforeCount, err := models.GetScore(c.Uid)
	afterCount := beforeCount + 5
	err = models.UpdateScore(afterCount, c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/addsign")
		resultMap["err"] = "更新融豆失败"
		return
	}
	//10：交易类型-签到奖励  5：融豆5个
	err = models.AddUserFinanceRecord(c.Uid, 10, 5, beforeCount, afterCount)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "记录用户收支记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/addsign")
		resultMap["err"] = "记录用户收支记录失败"
		return
	}
	err = models.UpdateLastSignTime(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新最近签到时间失败"+err.Error(), c.Ctx.Input.IP(), "sign/addsign")
		resultMap["err"] = "更新最近签到时间失败"
		return
	}
	resultMap["msg"] = "新增签到数据成功"
	resultMap["ret"] = 200
}

//获取签到页面数据（是否签到、签到提醒的当前状态、奖励数据）
// @router /getsign [post]
func (c *SignController) GetSign() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()
	signCount, err := models.IsSign(c.Uid, year, month, day)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取签到数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsign")
		resultMap["err"] = "获取签到数据失败"
		return
	}
	if signCount == 1 {
		//已签到
		resultMap["status"] = 1
	} else {
		resultMap["status"] = 0
	}

	status, err := models.GetSignRemindStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取当前签到提醒状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsign")
		resultMap["err"] = "获取当前签到提醒状态失败"
		return
	}
	if status == 1 {
		resultMap["RemindStatus"] = 1 //开启状态
	} else {
		resultMap["RemindStatus"] = 0 //关闭状态
	}

	signReward, err := models.GetSignReward()
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取奖励数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsign")
		resultMap["err"] = "获取奖励数据失败"
		return
	}
	sr_product_id, err := models.GetReceiveExtraId(c.Uid, month, year)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取已签到额外奖励id失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsign")
		resultMap["err"] = "获取已签到额外奖励id失败"
		return
	}
	signCount, err = models.GetSignCount(c.Uid, year, month)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取累计签到天数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsign")
		resultMap["err"] = "获取累计签到天数失败"
		return
	}
	idMap := make(map[int]int)
	srIdMap := make(map[int]int)
	for _, v := range sr_product_id {
		srIdMap[v] = v
	}
	for k, v := range signReward {
		idMap[v.Id] = v.Id
		if v.SignCount > signCount {
			signReward[k].IsReceive = 2
		}
		if _, ok := srIdMap[v.Id]; ok {
			signReward[k].IsReceive = 1
		}
	}
	resultMap["SignReward"] = signReward
	resultMap["ret"] = 200
}

//查询融豆数量
// @router /getscore [post]
func (c *SignController) GetScore() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	score, err := models.GetScore(c.Uid)
	if err != nil {
		if err.Error() == utils.ErrNoRow() {
			err := models.AddScore(c.Uid)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加融豆数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/getscore")
				resultMap["err"] = "添加融豆数据失败"
				return
			}
		} else {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询融豆数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/getscore")
			resultMap["err"] = "查询融豆数据失败"
			return
		}
	}
	resultMap["Score"] = score
	resultMap["ret"] = 200
}

//获取累计签到天数
// @router /getsigncount [post]
func (c *SignController) GetSignCount() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	signCount, err := models.GetSignCount(c.Uid, time.Now().Year(), int(time.Now().Month()))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取累计签到天数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getsigncount")
		resultMap["err"] = "获取累计签到天数失败"
		return
	}
	resultMap["SignCount"] = signCount
	resultMap["ret"] = 200
}

//获取可完成任务数量
// @router /getmissioncount [post]
func (c *SignController) GetMissionCount() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()

	//完成注册任务
	isUsed, err := models.GetIsUsed(5)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取注册任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取注册任务是否使用失败"
		return
	}
	if isUsed == 0 {
		countZc, err := models.GetGrowCount(c.Uid, 5)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取注册任务记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取注册任务记录数失败"
			return
		}
		if countZc == 0 {
			err := models.AddGrowSession(c.Uid, 5)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加注册任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "添加注册任务记录失败"
				return
			}
		}
	}

	//身份认证任务
	isRealName, err := models.GetIsRealNameStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取身份认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取身份认证状态失败"
		return
	}
	if isRealName == 2 {
		isUsed, err := models.GetIsUsed(6)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取实名认证是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取实名认证是否使用失败"
			return
		}
		if isUsed == 0 {
			countReal, err := models.GetGrowCount(c.Uid, 6)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取实名认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取实名认证记录数失败"
				return
			}
			if countReal == 0 {
				err := models.AddGrowSession(c.Uid, 6)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加实名认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加实名认证任务记录失败"
					return
				}
			}
		}
	}

	//芝麻信用认证任务
	isZm, err := models.GetZmStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取芝麻信用认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取芝麻信用认证状态失败"
		return
	}
	if isZm == 2 {
		isUsed, err := models.GetIsUsed(7)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取芝麻信用认证任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取芝麻信用认证任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countZm, err := models.GetGrowCount(c.Uid, 7)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取芝麻信用认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取芝麻信用认证记录数失败"
				return
			}
			if countZm == 0 {
				err := models.AddGrowSession(c.Uid, 7)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加芝麻信用认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加芝麻信用认证任务记录失败"
					return
				}
			}
		}
	}

	//移动运营认证任务
	isMobile, err := models.GetStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取手机运营商认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取手机运营商认证状态失败"
		return
	}

	if isMobile == 2 {
		isUsed, err := models.GetIsUsed(8)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取手机运营商认证任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取手机运营商认证任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countMobile, err := models.GetGrowCount(c.Uid, 8)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取手机运营商认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取手机运营商认证记录数失败"
				return
			}
			if countMobile == 0 {
				err := models.AddGrowSession(c.Uid, 8)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加收集运营商认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加手机运营商信用认证任务记录失败"
					return
				}
			}
		}
	}

	//支付宝认证任务
	isAlipay, err := models.GetAlipayStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取支付宝认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取支付宝认证状态失败"
		return
	}
	if isAlipay == 2 {
		isUsed, err := models.GetIsUsed(9)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取支付宝认证任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取支付宝认证任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countAlipay, err := models.GetGrowCount(c.Uid, 9)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取支付宝认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取支付宝认证记录数失败"
				return
			}
			if countAlipay == 0 {
				err := models.AddGrowSession(c.Uid, 9)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加支付宝认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加支付宝认证任务记录失败"
					return
				}
			}
		}
	}

	//京东商城认证任务
	isJd, err := models.GetJdStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取京东商城认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取京东商城认证状态失败"
		return
	}
	if isJd == 2 {
		isUsed, err := models.GetIsUsed(10)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取京东商城认证任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取京东商城认证任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countJd, err := models.GetGrowCount(c.Uid, 10)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取京东商城认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取京东商城认证记录数失败"
				return
			}
			if countJd == 0 {
				err := models.AddGrowSession(c.Uid, 10)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加京东商城认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加京东商城认证任务记录失败"
					return
				}
			}
		}
	}

	//淘宝商城认证任务
	isTb, err := models.GetTbStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取淘宝商城认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取淘宝商城认证状态失败"
		return
	}

	if isTb == 2 {
		isUsed, err := models.GetIsUsed(11)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取淘宝商城认证任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取淘宝商城认证任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countTb, err := models.GetGrowCount(c.Uid, 11)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取淘宝商城认证记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取淘宝商城认证记录数失败"
				return
			}
			if countTb == 0 {
				err := models.AddGrowSession(c.Uid, 11)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加淘宝商城认证任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加淘宝商城认证任务记录失败"
					return
				}
			}
		}
	}

	//人脸识别认证任务
	isFace, err := models.GetIsFaceStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取人脸识别状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取人脸识别状态失败"
		return
	}
	if isFace == 2 {
		isUsed, err := models.GetIsUsed(12)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取人脸识别任务是否使用失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取人脸识别任务是否使用失败"
			return
		}
		if isUsed == 0 {
			countFace, err := models.GetGrowCount(c.Uid, 12)
			if err != nil {
				cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取人脸识别记录数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
				resultMap["err"] = "获取人脸识别记录数失败"
				return
			}
			if countFace == 0 {
				err := models.AddGrowSession(c.Uid, 12)
				if err != nil {
					cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加人脸识别任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
					resultMap["err"] = "添加人脸识别任务记录失败"
					return
				}
			}
		}
	}

	missionCount, err := models.GetMissionCount()
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取所有任务总数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmissioncount")
		resultMap["err"] = "获取所有任务总数失败"
		return
	}
	dailyCount, err := models.GetDailyMissionCount(c.Uid, time.Now().Format("2006-01-02"))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取每日任务完成数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmissioncount")
		resultMap["err"] = "获取每日任务完成数失败"
		return
	}
	growCount, err := models.GetGrowMissionCount(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取成长任务完成数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmissioncount")
		resultMap["err"] = "获取成长任务完成数失败"
		return
	}
	count := missionCount - dailyCount - growCount
	resultMap["Count"] = count
	resultMap["ret"] = 200
}

//签到日历
// @router /getdays [post]
func (c *SignController) GetDays() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	days, err := models.GetDays(c.Uid, time.Now().Year(), int(time.Now().Month()))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "新增签到天数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdays")
		resultMap["err"] = "新增签到天数失败"
		return
	}
	resultMap["Days"] = days
	resultMap["ret"] = 200
}

//领取额外奖励
// @router /getextrareward [post]
func (c *SignController) GetExtraReward() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	year := time.Now().Year()
	month := int(time.Now().Month())
	var reward models.SignRewardProduct
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reward)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "解析参数失败"+err.Error()+";param:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "sign/getextrareward")
		resultMap["err"] = "解析参数失败"
		return
	}
	extraReward, err := models.GetExtraReward(reward.Id)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取额外奖励数据失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
		resultMap["err"] = "获取额外奖励数据失败"
		return
	}
	signCount, err := models.GetSignCount(c.Uid, year, month)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取累计签到天数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
		resultMap["err"] = "获取累计签到天数失败"
		return
	}

	count, err := models.GetReceiveCount(c.Uid, reward.Id, month, year)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取签到奖励失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
		resultMap["err"] = "获取签到奖励失败"
		return
	}
	if count == 1 {
		resultMap["msg"] = "已领取"
		resultMap["ret"] = 304
		return
	}
	if extraReward.SignCount <= signCount {
		err = models.AddRewardRecord(c.Uid, extraReward.Id, year, month)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "添加奖励记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
			resultMap["err"] = "添加奖励记录失败"
			return
		}
		beforeCount, err := models.GetScore(c.Uid)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
			resultMap["err"] = "查询融豆失败"
			return
		}
		afterCount := beforeCount + extraReward.RewardAmount
		err = models.UpdateScore(afterCount, c.Uid)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
			resultMap["err"] = "更新融豆失败"
			return
		}
		//10：交易类型-签到奖励
		err = models.AddUserFinanceRecord(c.Uid, 10, extraReward.RewardAmount, beforeCount, afterCount)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "记录用户收支记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getextrareward")
			resultMap["err"] = "记录用户收支记录失败"
			return
		}
	} else {
		resultMap["msg"] = "累计天数不足"
		return
	}

	resultMap["msg"] = "领取成功"
	resultMap["ret"] = 200
}

//获取任务信息
// @router /getmission [post]
func (c *SignController) GetMission() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	var users *models.BaseModels
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &users)
	if err != nil {
		cache.RecordNewLogs(0, "0", "解析参数错误"+err.Error()+";param:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "解析参数错误!"
		return
	}
	dailyMission, err := models.GetDailyMession()
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取每日任务失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取每日任务失败"
		return
	}
	dailyAlreday, err := models.GetAlredayDailyMission(c.Uid, time.Now().Format("2006-01-02"))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取已完成每日任务所有信息失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取已完成每日任务所有信息失败"
		return
	}
	dailyMap := make(map[int]int)
	missionIdMap := make(map[int]int)
	isReceiveMap := make(map[int]int)
	for _, v := range dailyAlreday {
		missionIdMap[v.MissionId] = v.MissionId
		isReceiveMap[v.MissionId] = v.IsReceive
	}
	missionImage, err := models.GetUrlByPackageId(users.PackageId)
	if err != nil && err.Error() != utils.ErrNoRow() {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取任务列表logo异常"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取每日任务信息失败"
		return
	}
	if users.PackageId == 0 || len(missionImage) == 0 {
		missionImage, err = models.GetUrlByPackageId(1)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取任务列表logo异常"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
			resultMap["err"] = "获取每日任务信息失败"
			return
		}
	}
	imgMap := make(map[int]string)
	for _, v := range missionImage {
		imgMap[v.MissionId] = v.ImgUrl
	}

	for k, v := range dailyMission {
		dailyMap[v.Id] = v.Id
		if _, ok := missionIdMap[v.Id]; ok {
			dailyMission[k].ProgressStatus = 1
			if isReceiveMap[v.Id] == 1 {
				dailyMission[k].Status = 1
			} else {
				dailyMission[k].Status = 0
			}
		}
		dailyMission[k].ImgUrl = imgMap[v.Id]
	}

	growMission, err := models.GetGrowMession()
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取成长任务失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取成长任务失败"
		return
	}
	growAlreday, err := models.GetAlredayGrowMission(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取已完成成长任务所有信息失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取已完成成长任务所有信息失败"
		return
	}

	//移动运营认证任务
	isMobile, err := models.GetStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取手机运营商认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取手机运营商认证状态失败"
		return
	}
	if isMobile == 3 {
		growMission[3].Status = 3
	}

	//支付宝认证任务
	isAlipay, err := models.GetAlipayStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取支付宝认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取支付宝认证状态失败"
		return
	}
	if isAlipay == 3 {
		growMission[4].Status = 3
	}

	//京东商城认证任务
	isJd, err := models.GetJdStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取京东商城认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取京东商城认证状态失败"
		return
	}
	if isJd == 3 {
		growMission[5].Status = 3
	}

	//淘宝商城认证任务
	isTb, err := models.GetTbStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取淘宝商城认证状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取淘宝商城认证状态失败"
		return
	}
	if isTb == 3 {
		growMission[6].Status = 3
	}

	growMap := make(map[int]int)
	missionMap := make(map[int]int)
	receiveMap := make(map[int]int)

	for _, v := range growAlreday {
		missionMap[v.MissionId] = v.MissionId
		receiveMap[v.MissionId] = v.IsReceive
	}
	for k, v := range growMission {
		growMap[v.Id] = v.Id
		if _, ok := missionMap[v.Id]; ok {
			growMission[k].ProgressStatus = 1
			if receiveMap[v.Id] == 1 {
				growMission[k].Status = 1
			} else {
				growMission[k].Status = 0
			}
		}
		growMission[k].ImgUrl = imgMap[v.Id]
	}

	count, err := models.GetNoReceiveCount(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取未领取成长任务数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getmission")
		resultMap["err"] = "获取未领取成长任务数失败"
		return
	}
	if count > 0 {
		//成长任务提示
		resultMap["PointStatus"] = 1
	} else {
		resultMap["PointStatus"] = 0
	}
	resultMap["DailyMission"] = dailyMission
	resultMap["GrowMission"] = growMission
	resultMap["ret"] = 200
	return
}

//每日任务立即领取
// @router /getdailyreward [post]
func (c *SignController) GetDailyReward() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	var dailyMission models.MissionDailyRecord
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dailyMission)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "解析参数失败"+err.Error()+";param:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "解析参数失败"
		return
	}
	isReceive, err := models.GetDailyRecord(c.Uid, dailyMission.MissionId, time.Now().Format("2006-01-02"))
	if err != nil {
		if err.Error() == utils.ErrNoRow() {
			resultMap["msg"] = "未完成"
			return
		} else {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取每日任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
			resultMap["err"] = "获取每日任务记录失败"
			return
		}
	}
	if isReceive == 1 {
		resultMap["msg"] = "已领取"
		resultMap["ret"] = 304
		return
	}
	err = models.UpadteDailyReceive(c.Uid, dailyMission.MissionId, time.Now().Format("2006-01-02"))
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新每日任务领取状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "更新每日任务领取状态失败"
		return
	}
	missionAward, err := models.GetMissionAward(dailyMission.MissionId)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询任务奖励(融豆数量)失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "查询任务奖励(融豆数量)失败"
		return
	}
	beforeCount, err := models.GetScore(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "查询融豆失败"
		return
	}
	afterCount := beforeCount + missionAward

	err = models.UpdateScore(afterCount, c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "更新融豆失败"
		return
	}
	//9：交易类型-任务奖励
	err = models.AddUserFinanceRecord(c.Uid, 9, missionAward, beforeCount, afterCount)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "记录用户收支记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getdailyreward")
		resultMap["err"] = "记录用户收支记录失败"
		return
	}
	resultMap["msg"] = "领取成功"
	resultMap["ret"] = 200
}

//成长任务立即领取
// @router /getgrowreward [post]
func (c *SignController) GetGrowReward() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	var dailyMission models.MissionDailyRecord
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dailyMission)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "解析参数失败"+err.Error()+";param:"+string(c.Ctx.Input.RequestBody), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "解析参数失败"
		return
	}
	isReceive, err := models.GetGrowRecord(c.Uid, dailyMission.MissionId)
	if err != nil {
		if err.Error() == utils.ErrNoRow() {
			resultMap["msg"] = "未完成"
			return
		} else {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取成长任务记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
			resultMap["err"] = "获取成长任务记录失败"
			return
		}
	}
	if isReceive == 1 {
		resultMap["msg"] = "已领取"
		resultMap["ret"] = 304
		return
	}
	err = models.UpadteGrowReceive(c.Uid, dailyMission.MissionId)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新成长任务领取状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "更新成长任务领取状态失败"
		return
	}
	missionAward, err := models.GetMissionAward(dailyMission.MissionId)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询任务奖励(融豆数量)失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "查询任务奖励(融豆数量)失败"
		return
	}
	beforeCount, err := models.GetScore(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "查询融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "查询融豆失败"
		return
	}
	afterCount := beforeCount + missionAward

	err = models.UpdateScore(afterCount, c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新融豆失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "更新融豆失败"
		return
	}
	//9：交易类型-任务奖励
	err = models.AddUserFinanceRecord(c.Uid, 9, missionAward, beforeCount, afterCount)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "记录用户收支记录失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "记录用户收支记录失败"
		return
	}

	count, err := models.GetNoReceiveCount(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取未领取成长任务数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getgrowreward")
		resultMap["err"] = "获取未领取成长任务数失败"
		return
	}
	if count > 0 {
		//成长任务提示
		resultMap["PointStatus"] = 1
	} else {
		resultMap["PointStatus"] = 0
	}
	resultMap["msg"] = "领取成功"
	resultMap["ret"] = 200
}

//签到提醒状态
// @router /getremindstatus [post]
func (c *SignController) GetRemindStatus() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	status, err := models.GetSignRemindStatus(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取当前签到提醒状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getremindstatus")
		resultMap["err"] = "获取当前签到提醒状态失败"
		return
	}
	if status == 0 {
		err := models.UpdateSignRemindStatus(1, c.Uid)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新当前签到提醒状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getremindstatus")
			resultMap["err"] = "更新当前签到提醒状态失败"
			return
		}
		resultMap["status"] = 1 //提醒
		resultMap["ret"] = 200
		return
	}
	if status == 1 {
		err := models.UpdateSignRemindStatus(0, c.Uid)
		if err != nil {
			cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "更新当前签到提醒状态失败"+err.Error(), c.Ctx.Input.IP(), "sign/getremindstatus")
			resultMap["err"] = "更新当前签到提醒状态失败"
			return
		}
		resultMap["status"] = 0 //不提醒
		resultMap["ret"] = 200
		return
	}
}

//移动运营去完成判断人脸认证是否完成
// @router /getfacecount [post]
func (c *SignController) GetFaceCount() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	count, err := models.GetFaceIsFinish(c.Uid)
	if err != nil {
		cache.RecordNewLogs(c.Uid, strconv.Itoa(c.Uid), "获取人脸认证数失败"+err.Error(), c.Ctx.Input.IP(), "sign/getfacecount")
		resultMap["err"] = "获取人脸认证数失败"
		return
	}
	if count == 1 {
		resultMap["status"] = 1 //已完成
	}
	if count == 0 {
		resultMap["status"] = 0 //未完成
	}
	resultMap["ret"] = 200
	return
}
