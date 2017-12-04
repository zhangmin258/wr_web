package models

type BaseModels struct {
	Uid            int
	Account        string      //账号--手机号码
	App            int         //来自哪个平台 1 ios,2 android,3 wx,4pc，
	AppVersion     string      // app版本
	MobileType     string      //手机型号
	MobileVersion  string      //手机版本号
	Token          string      // 登录成功后返回一个token,作为下次post 请求的
	IsEmulator     bool        //是否虚拟机
	Location       string      //定位市区
	Address        string      //详细地址
	IpLocation     interface{} //手机定位(根据ip获取)
	IpAddress      string      //详细地址
	LongiAndlati   string      //经纬度
	DeviceUniqueID string      //手机设备号
	UserCode       string
	PackageId      int //分身包ID
	SingleSource   string
	WxToken        string //微融生成的微信公众号token
	Page           int    //分页信息
}
