package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Logs struct {
	Id          int       `xorm:"int pk autoincr 'id'"`
	Uid         int       `xorm:"int 'uid'"`
	Account     string    `xorm:"varchar(1000) 'account'"`
	OperateTime time.Time `xorm:"timestamp null 'operate_time'"`
	Content     string    `xorm:"text 'content'"`
	OperateIp   string    `xorm:"varchar(1000) null 'operate_ip'"`
	Model       string    `xorm:"varchar(1000) null 'model'"`
}

//变更借款状态的记录
type LoanLog struct {
	Uid        int       //用户uid
	LoanId     int       //借款id
	CreateTime time.Time //创建时间
	Status     string    //借款状态
}

// 记录用户收支操作
func AddFinanceLog(financeLog UsersFinanceRecord) (err error) {
	sql := `INSERT INTO users_finance_record (uid, pay_token,pay_type, deal_type, score_amount, money_amount, pay_or_get, create_time) VALUES (?,?,?,?,?,?,?,?)`
	_, err = orm.NewOrm().Raw(sql, financeLog.Uid, financeLog.PayToken, financeLog.PayType, financeLog.DealType, financeLog.ScoreAmount, financeLog.MoneyAmount, financeLog.PayOrGet, financeLog.CreateTime).Exec()
	return
}
