package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"wr_web/utils"
)

func init() {
	orm.RegisterDataBase("default", "mysql", utils.MYSQL_URL)
	orm.RegisterDataBase("wr_log", "mysql", utils.MYSQL_LOG_URL)
}
