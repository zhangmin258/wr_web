package models

import "github.com/astaxie/beego/orm"

// 初始化用户融豆表
func InitScore(uid int) (err error) {
	sql := `INSERT INTO score (uid,score) VALUES (?,0)`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

// 用户支付，扣除融豆
func ScorePay(uid int, score int) (err error) {
	sql := `UPDATE score SET score=score-? WHERE uid=?`
	_, err = orm.NewOrm().Raw(sql, score, uid).Exec()
	return
}
