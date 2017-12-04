package models

import "github.com/astaxie/beego/orm"

type Wallet struct {
	Id             int
	Uid            int     // 用户ID
	AccountBalance float64 // 钱包账户余额
}

// 获取钱包账户余额
func GetWalletBalance(uid int) (wallet *Wallet, err error) {
	sql := `SELECT id,uid,account_balance FROM wallet WHERE uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&wallet)
	return
}

// 初始化用户钱包数据
func InnitWallet(uid int) (err error) {
	sql := `INSERT INTO wallet (uid,account_balance) VALUES (?,0.00)`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

// 用户支付，扣除余额
func WalletPay(uid int, money float64) (err error) {
	sql := `UPDATE wallet SET account_balance=account_balance-? WHERE uid=?`
	_, err = orm.NewOrm().Raw(sql, money, uid).Exec()
	return
}
