package cache

import (
	"encoding/json"
	"strconv"

	"wr_web/models"
	"wr_web/utils"
)

//向redis 中添加数据
// func SetInterfaceToRedis(key string, object interface{}) bool {
// 	if conn, err := redis.Dial("tcp", utils.REDIS_URI); err == nil {
// 		data, _ := json.Marshal(object)
// 		conn.Do("lpush", key, data)
// 		return true
// 	}
// 	return false
// }

//===========================获取用户的基本信息缓存====================================
// 通过uid或者account获取用户的【注册信息】缓存
func GetUsersByAccountCache(account string) (m *models.Users, err error) {
	if utils.Re == nil && utils.Rc.IsExist(utils.KEY_USERS_IFNO_REGISTER+"_"+account) {
		if data, err1 := utils.Rc.RedisBytes(utils.KEY_USERS_IFNO_REGISTER + "_" + account); err1 == nil {
			err = json.Unmarshal(data, &m)
			if m != nil {
				return
			}
		}
	}
	return models.GetUsersByAccount(account)
}
func GetUsersByIdCache(uid int) (m *models.Users, err error) {
	if utils.Re == nil && uid != 0 && utils.Rc.IsExist(utils.KEY_USERS_IFNO_REGISTER+"_"+strconv.Itoa(uid)) {
		if data, err1 := utils.Rc.RedisBytes(utils.KEY_USERS_IFNO_REGISTER + "_" + strconv.Itoa(uid)); err1 == nil {
			err = json.Unmarshal(data, &m)
			if m != nil {
				return
			}
		}
	}
	return models.GetUsersById(uid)
}

//  通过code获取用户的【注册信息】缓存
func GetUserByCodeCache(code string) (m *models.Users, err error) {
	if utils.Re == nil && utils.Rc.IsExist(utils.KEY_USERS_IFNO_REGISTER+"_"+code) {
		if date, err1 := utils.Rc.RedisBytes(utils.KEY_USERS_IFNO_REGISTER + "_" + code); err1 == nil {
			err = json.Unmarshal(date, &m)
			if m != nil {
				return
			}
		}
	}
	return models.GetUsersByCode(code)
}

//  通过uid从邀请好友奖励金提现表查询用户
// func GetUserFromWithdrawById(id int)(id int,err error){
// 	if utils.Re == nil && utils.Rc.IsExist(key)
// }
