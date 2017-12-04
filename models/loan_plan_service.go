package models

import (
	"github.com/astaxie/beego/orm"
)

type ChargeCall struct {
	PayToken  string // 付款凭证
	Uid       int    //用户编号
	ServiceId int    //商品id
}

//贷款稳下页面广告轮播
func GetAdvByConfigKey() (adv string, err error) {
	sql := `SELECT config_desc AS adv FROM config WHERE code="loan_adv_round"`
	err = orm.NewOrm().Raw(sql).QueryRow(&adv)
	return
}

//贷款稳下页面预约花费查询
func GetLoanPlanCost(id int) (showSeProducts *ShowSeProducts, err error) {
	sql := `SELECT * FROM score_exchange_product WHERE id=?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&showSeProducts)
	return
}

//添加预约记录
func AddLoanPlanRecord(uid, isValid, isFinished int, payOrder string) error {
	sql := `INSERT INTO support_loan(uid,pay_order,is_valid,create_time,is_finished)VALUES(?,?,?,NOW(),?)`
	o := orm.NewOrm()
	_, err := o.Raw(sql, uid, payOrder, isValid, isFinished).Exec()
	return err
}

//查询用户是否已经预约
func GetLoanPlanByUid(uid int) (count int, err error) {
	sql := `SELECT COUNT(1) FROM support_loan WHERE uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}
