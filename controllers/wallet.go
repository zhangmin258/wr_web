package controllers

import (
	"strconv"
	"wr_web/cache"
	"wr_web/models"
	"wr_web/utils"
)

/*
	钱包
*/

type WalletController struct {
	H5Controller
}

// 获取用户的账户余额
func (this *WalletController) GetWalletBalance() {
	defer this.ServeJSON()
	wallet, err := models.GetWalletBalance(this.Uid)   // 根据用户ID查询余额
	if err != nil && err.Error() == utils.ErrNoRow() { // 没有用户数据,修复数据
		err := models.InnitWallet(this.Uid)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"ret": 403, "msg": "初始化用户钱包数据失败!"}
			cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "初始化用户钱包数据失败"+err.Error(), this.Ctx.Input.IP(), "wallet/getwalletbalance")
			return
		}
		this.Data["json"] = map[string]interface{}{"ret": 200, "balance": 0.00}
		return
	}
	if err != nil && err.Error() != utils.ErrNoRow() {
		this.Data["json"] = map[string]interface{}{"ret": 403, "msg": "查询用户余额失败!"}
		cache.RecordNewLogs(this.Uid, strconv.Itoa(this.Uid), "查询用户余额失败!"+err.Error(), this.Ctx.Input.IP(), "wallet/getwalletbalance")
		return
	}
	this.Data["json"] = map[string]interface{}{"ret": 200, "balance": wallet.AccountBalance}
	return
}

// 获取用户的收支明细
// @router /getwalletinfo [post]
func (this *WalletController) GetWalletInfo() {
	defer this.ServeJSON()

	return
}

// 该用户是否有设置密码
// @router /haswalletpwd [post]
func (this *WalletController) HasWalletPwd() {
	defer this.ServeJSON()

	return
}

// 设置钱包密码
// @router /setwalletpwd [post]
func (this *WalletController) SetWalletPwd() {
	defer this.ServeJSON()

	return
}
