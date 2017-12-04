package services

import (
	"strconv"
	"time"
	"wr_web/models"
	"wr_web/utils"
	"zcm_tools/uuid"
)

// 根据支付方式和金额给用户扣款,并返回支付凭证
func Pay(uid int, moneyService *models.ShowSeProducts, payType int) (payToken string, err error) {

	var financeLog models.UsersFinanceRecord
	financeLog.Uid = uid            // 用户uid
	payToken = uuid.NewUUID().Hex() // 付款凭证
	financeLog.PayToken = payToken
	financeLog.DealType = 1 // 交易类型
	financeLog.PayOrGet = 1 // 1 付款
	financeLog.CreateTime = time.Now()
	financeLog.PayType = payType
	switch payType { // 支付方式：1融豆 2现金余额 3会员免费
	case 1:
		err = models.ScorePay(uid, moneyService.ScorePrice)
		if err != nil {
			return
		}
		financeLog.ScoreAmount = moneyService.ScorePrice
		financeLog.MoneyAmount = 0
	case 2:
		err = models.WalletPay(uid, moneyService.MoneyPrice)
		if err != nil {
			return
		}
		financeLog.ScoreAmount = 0
		financeLog.MoneyAmount = moneyService.MoneyPrice
	case 3:
		financeLog.ScoreAmount = 0
		financeLog.MoneyAmount = 0
	}
	err = utils.Rc.Put(utils.CACHE_KEY_SERVICE_PAY_TOKEN+strconv.Itoa(uid)+strconv.Itoa(moneyService.ServiceId), payToken, 20*time.Minute) // 将用户支付token放入redis
	err = models.AddFinanceLog(financeLog)
	return
}
