package utils

//常量定义
const (
	key            = `h*.d;cy7x_12dkx?#j39fdl!` //api数据加密、解密key
	FormatDate     = "2006-01-02"               //日期格式
	FormatDateTime = "2006-01-02 15:04:05"      //完整时间格式
)
const (
	CACHE_KEY_Operation = "Operation_token_" //防跨域
	CACHE_KEY_Ip        = "Operation_Ip_"
)

//缓存key
const (
	KEY_USERS_IFNO_REGISTER     = "users:info:registerwr"        //_uid或_account //  用户注册信息，加入uid作为key放入redis中
	CACHE_KEY_SERVICE_PAY_TOKEN = "CACHE_KEY_SERVICE_PAY_TOKEN_" //支付Token
	WR_CACHE_KEY_SCORE_SERVICE  = "WR_CACHE_KEY_SCORE_SERVICE"   // 增值服务详情
	WR_CACHE_IOS_VERSION_DATA   = "WR_CACHE_IOS_VERSION_DATA"    //ios审核假数据
)

//缓存配置
const (
	CACHE_KEY_LOGS            = "WrCacheErrKey_"
	CACHE_KEY_WRWEB_OPERATION = "CACHE_KEY_WRWEB_OPERATION_"
)

const (
	Regular = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(17[0-9]))\\d{8}$"
)

// send to users
const (
	ToUsers = "lwc@zcmlc.com;jgl@zcmlc.com"
)
