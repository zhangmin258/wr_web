package models

import (
	//"zcm_tools/orm"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"wr_web/utils"
)

type Config struct {
	Id           int
	Config_key   string
	Config_value string
	Config_desc  string
	Remark       string
	ConfigUrl    string
}

//项目配置
func GetConfig(key string) (cf *Config, err error) {
	o := orm.NewOrm()
	sql := `SELECT * FROM config where  config_key= ? `
	err = o.Raw(sql, key).QueryRow(&cf)
	return cf, err
}

type PackageConfig struct {
	Id                  int
	PkgName             string //分身包名
	IosVersion          string //ios版本号
	IosPretendVersion   string //ios伪装版本号
	AndroidVersion      string //android版本号
	IosCoerceUpdate     int    //ios是否强制更新
	AndroidCoerceUpdate int    //android
	IosUpdateMsg        string //ios更新标语
	AndroidUpdateMsg    string //android更新标语
	IosUpdateUrl        string //ios更新链接
	AndroidUpdateUrl    string //android更新链接
}

func GetPackageConfig(id int) (p PackageConfig, err error) {
	o := orm.NewOrm()
	if err = o.Raw(`SELECT * FROM package_config WHERE id=?`, id).QueryRow(&p); err == nil {
		if data, err1 := json.Marshal(p); err1 == nil && utils.Re == nil {
			utils.Rc.Put(utils.WR_CACHE_IOS_VERSION_DATA+"_"+strconv.Itoa(id), data, 60*time.Second)
			utils.Rc.Put(utils.WR_CACHE_IOS_VERSION_DATA+"_"+strconv.Itoa(p.Id), data, 24*time.Second)
		}
		return
	}
	return
}
