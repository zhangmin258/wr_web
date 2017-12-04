package utils

import (
	"github.com/astaxie/beego"
	"zcm_tools/cache"
)

var (
	RunMode       string //运行模式
	MYSQL_URL     string //数据库连接
	MYSQL_LOG_URL string
	BEEGO_CACHE   string       //缓存地址
	Rc            *cache.Cache //redis缓存
	Re            error        //redis错误
	Enablexsrf    string       // XSRF校验开关
	H5Encoded     string       // H5接口base64编码开关
	APIEncoded    string       // API接口base64编码开关
)

func init() {
	RunMode = beego.AppConfig.String("run_mode")
	H5Encoded = beego.AppConfig.String("h5_encoded")
	APIEncoded = beego.AppConfig.String("api_encoded")
	config, err := beego.AppConfig.GetSection(RunMode)
	if err != nil {
		panic("配置文件读取错误 " + err.Error())
	}
	Enablexsrf = beego.AppConfig.String("enable_xsrf")
	beego.Info("┌───────────────────")
	beego.Info("│XSRF校验:" + Enablexsrf)
	beego.Info("│APP接口加解密:" + APIEncoded)
	beego.Info("│H5接口编码:" + H5Encoded)
	beego.Info("│模式:" + RunMode)
	beego.Info("└───────────────────")
	MYSQL_URL = config["mysql_url"]
	MYSQL_LOG_URL = config["mysql_log_url"]
	BEEGO_CACHE = config["beego_cache"]
	Rc, Re = cache.NewCache(BEEGO_CACHE) //初始化缓存
}
