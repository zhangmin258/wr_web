package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

// 会员
type VIP struct {
	Id           int
	Uid          int
	BeginTime    time.Time // 会员起始时间
	EndTime      time.Time // 会员结束时间
	ActivMonthNo int       // 总的开通月数
}

// 获取用户VIP信息
func GetVIPInfo(uid int) (vip *VIP, err error) {
	sql := `SELECT id,uid,begin_time,end_time FROM vip WHERE uid=? `
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&vip)
	return
}

// 初始化用户VIP信息
func InnitVIP(uid int) (err error) {
	beginTime := time.Now()
	endTime := beginTime.AddDate(0, 0, 31)
	sql := `INSERT INTO vip (uid, begin_time,end_time)VALUES(?,?,?)`
	_, err = orm.NewOrm().Raw(sql, uid, beginTime, endTime).Exec()
	return
}
