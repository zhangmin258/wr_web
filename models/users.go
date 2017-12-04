package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"wr_web/utils"
)

type Users struct {
	Id                  int
	Code                string
	Account             string `orm:"column(account);size(40)"`
	LoginPassword       string
	State               int       `orm:"column(state);null"`
	CreateTime          time.Time `orm:"column(create_time);type(timestamp);null;auto_now"`
	MobileType          string    `orm:"column(mobile_type);size(100)"`
	MobileVersion       string    `orm:"column(mobile_version);size(100)"`
	LoginTime           time.Time `orm:"column(login_time);type(timestamp);auto_now"`
	HisDevicecode       string    `orm:"column(his_devicecode);null"`        // 注册后第一次登陆后的客户端唯一标识
	MobileTypeRecent    string    `orm:"column(mobile_type_recent);null"`    //手机型号1.2.1添加 手机登录后的手机类型
	MobileVersionRecent string    `orm:"column(mobile_version_recent);null"` //手机版本号
	OperationTime       time.Time `orm:"column(operation_time);null"`        //最近操作的时间
	App                 int       `orm:"column(app);null"`                   //来自哪个平台1.2.1添加 1 ios,2 android,3 wx,4pc，5vd 6 wap
	Token               string    `orm:"column(token);null"`                 //手机型号1.2.3 每次登录产生的唯一参数
	Source              string    `orm:"column(source);null"`                // 渠道
	AppVersion          string    //app 版本号
	AppVersionRecent    string    //最近操作app版本
	UserImg             string    `orm:"column(user_img);null"` //用户头像
	Isemulator          bool      //是否模拟器
	Jpushid             string    //手机推送标识
	LoanControl         int       `orm:"column(loan_control);null"` //借款控制: 0进行中有一笔不能申请第二笔,1:不做限制
	Ip                  string    //ip地址
	HisDevicecodeRecent string    `orm:"column(his_devicecode_recent);null"` //最近手机设备标识
	Location            string    //开启定位，手机定位位置
	Address             string    //开启定位，详细地址
	IpLocation          string    //ip定位位置
	IpAddress           string    //ip详细地址
	Stage               string    //注册记录
	LoanState           string    //借款申请状态
	LongiAndlati        string    //经纬度
	IsIndexPopup        int       //首页是否弹窗0:弹窗1：不弹窗
	IsAuthPopup         int       //认证之后是否弹窗0:弹窗1：不弹窗
	RegisterSource      int       //1:主包，>1：分包
	WxToken             string    //微融生成的微信公众号token
	SlientLogin         int       //0：常规登录 ，1：静默登录
}

//根据token(平台的唯一标识)数据查询用户注册信息；如果以后使用这个查询,请使用缓存
func GetUsersByToken(token string) (v *Users, err error) {
	o := orm.NewOrm()
	if err = o.Raw(`SELECT * FROM users WHERE token = ? `, token).QueryRow(&v); err == nil {
		if data, err2 := json.Marshal(v); err2 == nil && utils.Re == nil {
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+token, data, 24*time.Hour)
		}
		return
	}
	return nil, err
}

//根据手机号（账号）数据查询用户注册信息；如果以后使用这个查询,请使用缓存
func GetUsersByAccount(account string) (v *Users, err error) {
	o := orm.NewOrm()
	if err = o.Raw(`SELECT * FROM users WHERE account=? `, account).QueryRow(&v); err == nil {
		if data, err2 := json.Marshal(v); err2 == nil && utils.Re == nil {
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+account, data, 2*60*time.Second) // utils.GetTodayLastSecond()获取今天结束剩余时间
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+strconv.Itoa(v.Id), data, 60*time.Second)
		}
		return v, nil
	}
	return nil, err
}

//根据uid(平台的唯一标识)数据查询用户注册信息；如果以后使用这个查询,请使用缓存
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	if err = o.Raw(`SELECT * FROM users WHERE id = ?`, id).QueryRow(&v); err == nil {
		if data, err2 := json.Marshal(v); err2 == nil && utils.Re == nil {
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+v.Account, data, 60*time.Second)
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+strconv.Itoa(v.Id), data, 60*time.Second)
		}
		return
	}
	return nil, err
}

//  根据用户编码code查询用户注册信息；如果以后使用这个查询，用使用缓存
func GetUsersByCode(code string) (v *Users, err error) {
	o := orm.NewOrm()
	if err = o.Raw(` SELECT * FROM users WHERE code = ? `, code).QueryRow(&v); err == nil {
		if data, err2 := json.Marshal(v); err2 == nil && utils.Re == nil {
			utils.Rc.Put(utils.KEY_USERS_IFNO_REGISTER+"_"+code, data, 60*time.Second)
		}
		return
	}
	return nil, err
}
