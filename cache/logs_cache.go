package cache

import (
	"time"
	"wr_web/models"
	"wr_web/utils"
)

//record more information
func RecordNewLogs(uid int, account, content, operateIp, model string) bool {

	log := &models.Logs{Uid: uid, Account: account, OperateTime: time.Now(), Content: content, OperateIp: operateIp,
		Model:               model}
	if utils.Re == nil {
		utils.Rc.LPush(utils.CACHE_KEY_LOGS, log)
		return true
	}
	return false
}
