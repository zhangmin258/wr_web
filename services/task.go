package services

import (
	"github.com/astaxie/beego/toolbox"
	"time"
	"wr_web/utils"
)

func Task() {
	go InsertLotteryMoney()
}

//每天存入一条全新的奖池金额
func InsertLotteryMoney() {
	taskMission := toolbox.NewTask("DailyInsert", "0 0 0 * * 0-6", taskMission)
	toolbox.AddTask("DailyInsert", taskMission)
	toolbox.StartTask()
}

func taskMission() error {
	balance := `{"Money":500,"RongDou":50000}`
	utils.Rc.Put("UsersLotteryBalanceDailyReset", balance, time.Hour*24)
	return nil
}
