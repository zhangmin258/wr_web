package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Sign struct {
	Id     int
	Uid    int
	Year   int
	Month  int
	Day    int
	IsSign int
	Count  int
}

type SignRewardProduct struct {
	Id           int
	SignCount    int //签到次数
	RewardAmount int //签到奖励融豆金额
	IsReceive    int //是否领取 0：未领取 1:已领取 2:未完成
}

type Score struct {
	Id    int
	Uid   int
	Score int
}

type SignRewardRecord struct {
	Id                  int
	Uid                 int
	SignRewardProductId int //签到奖励项的ID
	Year                int
	Month               int
	CreateTime          string
}

type Mission struct {
	Id             int
	Title          string //任务标题
	Content        string //任务内容
	ImgUrl         string //图片链接
	MissionAward   int    //任务奖励（融豆数量）
	MissionType    int    //任务类型：1.日常任务 2.成长任务
	Status         int    //状态 0：未领取 1：已领取 2：未完成
	ProgressStatus int    //进度状态  1:1/1  0:0/1
}

type MissionDailyRecord struct {
	Id           int
	Uid          int
	MissionId    int       //任务ID
	IsReceive    int       //是否已领取：0，未领取 1，已领取
	CompleteDate time.Time //完成时间
}

type MissionGrowRecord struct {
	Id           int
	Uid          int
	MissionId    int       //任务ID
	IsReceive    int       //是否已领取：0，未领取 1，已领取
	CompleteDate time.Time //完成时间
}

//签到
func AddSign(uid, year, month, day int) (err error) {
	sql := `INSERT INTO sign (uid,year,month,day) VALUES (?,?,?,?)`
	_, err = orm.NewOrm().Raw(sql, uid, year, month, day).Exec()
	return
}

//是否签到
func IsSign(uid, year, month, day int) (count int, err error) {
	sql := `SELECT count(1) FROM sign WHERE uid=? AND year=? AND month=? AND day=?`
	err = orm.NewOrm().Raw(sql, uid, year, month, day).QueryRow(&count)
	return
}

//查询融豆
func GetScore(uid int) (score int, err error) {
	sql := `SELECT score FROM score WHERE uid = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&score)
	return
}

//插入融豆数据
func AddScore(uid int) (err error) {
	sql := `INSERT INTO score (uid,score) VALUES (?,0)`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

//更新融豆
func UpdateScore(score, uid int) (err error) {
	sql := `UPDATE score SET score=? WHERE uid=?`
	_, err = orm.NewOrm().Raw(sql, score, uid).Exec()
	return
}

//查询累计签到天数
func GetSignCount(uid, year, month int) (count int, err error) {
	sql := `SELECT count(1) AS count FROM sign where uid=? AND year=? AND month=?  `
	err = orm.NewOrm().Raw(sql, uid, year, month).QueryRow(&count)
	return
}

//签到日历--返回签到时间
func GetDays(uid, year, month int) (days []int, err error) {
	sql := `SELECT day FROM sign WHERE uid=? AND year=? AND month=?`
	_, err = orm.NewOrm().Raw(sql, uid, year, month).QueryRows(&days)
	return
}

//获取额外奖励
func GetExtraReward(id int) (srp SignRewardProduct, err error) {
	sql := `SELECT id,sign_count,reward_amount FROM sign_reward_product WHERE id = ?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&srp)
	return
}

//添加奖励记录
func AddRewardRecord(uid, sign_reward_product_id, year, month int) (err error) {
	sql := `INSERT INTO sign_reward_record (uid,sign_reward_product_id,year,month,create_time) VALUES (?,?,?,?,NOW())`
	_, err = orm.NewOrm().Raw(sql, uid, sign_reward_product_id, year, month).Exec()
	return
}

//获取每日任务信息
func GetDailyMession() (dailyMission []Mission, err error) {
	sql := `SELECT id,title,content,img_url,mission_award,mission_type,is_used,2 status FROM mission WHERE mission_type=1 AND is_used=0`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&dailyMission)
	return
}

//获取成长任务信息
func GetGrowMession() (growMission []Mission, err error) {
	sql := `SELECT id,title,content,img_url,mission_award,mission_type,is_used,2 status FROM mission WHERE mission_type=0 AND is_used=0`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&growMission)
	return
}

//完成添加每日任务
func AddDailySession(uid, mission_id int) (err error) {
	sql := `INSERT INTO mission_daily_record (uid,mission_id,is_receive,complete_date) VALUES (?,?,0,NOW())`
	_, err = orm.NewOrm().Raw(sql, uid, mission_id).Exec()
	return
}

//完成添加成长任务
func AddGrowSession(uid, mission_id int) (err error) {
	sql := `INSERT INTO mission_grow_record (uid,mission_id,is_receive,complete_date) VALUES (?,?,0,NOW())`
	_, err = orm.NewOrm().Raw(sql, uid, mission_id).Exec()
	return
}

//获取每日任务记录(是否领取)
func GetDailyRecord(uid, mission_id int, complete_date string) (is_receive int, err error) {
	sql := `SELECT is_receive FROM mission_daily_record WHERE uid=? AND mission_id=? AND complete_date=?`
	err = orm.NewOrm().Raw(sql, uid, mission_id, complete_date).QueryRow(&is_receive)
	return
}

//查询每日任务是否完成
func GetDailyCount(uid, mission_id int, complete_date string) (count int, err error) {
	sql := `SELECT count(1) FROM mission_daily_record WHERE uid=? AND mission_id =? AND complete_date=?`
	err = orm.NewOrm().Raw(sql, uid, mission_id, complete_date).QueryRow(&count)
	return
}

//获取成长任务记录(是否领取)
func GetGrowRecord(uid, mission_id int) (is_receive int, err error) {
	sql := `SELECT is_receive FROM mission_grow_record WHERE uid=? AND mission_id=?`
	err = orm.NewOrm().Raw(sql, uid, mission_id).QueryRow(&is_receive)
	return
}

//查询成长任务是否完成
func GetGrowCount(uid, mission_id int) (count int, err error) {
	sql := `SELECT count(1) FROM mission_grow_record WHERE uid=? AND mission_id =?`
	err = orm.NewOrm().Raw(sql, uid, mission_id).QueryRow(&count)
	return
}

//领取更新每日任务领取状态
func UpadteDailyReceive(uid, mission_id int, complete_date string) (err error) {
	sql := `UPDATE mission_daily_record SET is_receive=1 WHERE uid=? AND mission_id=? AND complete_date=?`
	_, err = orm.NewOrm().Raw(sql, uid, mission_id, complete_date).Exec()
	return
}

//领取更新成长任务领取状态
func UpadteGrowReceive(uid, mission_id int) (err error) {
	sql := `UPDATE mission_grow_record SET is_receive=1 WHERE uid=? AND mission_id=?`
	_, err = orm.NewOrm().Raw(sql, uid, mission_id).Exec()
	return
}

//获取奖励数据
func GetSignReward() (srp []SignRewardProduct, err error) {
	sql := `SELECT id,sign_count,reward_amount FROM sign_reward_product`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&srp)
	return
}

//获取领取次数判断是否已领取
func GetReceiveCount(uid, sign_reward_product_id, month, year int) (count int, err error) {
	sql := `SELECT count(1) FROM sign_reward_record WHERE uid=? AND sign_reward_product_id=? AND month=? AND year=?`
	err = orm.NewOrm().Raw(sql, uid, sign_reward_product_id, month, year).QueryRow(&count)
	return
}

//成长任务提示
func GetNoReceiveCount(uid int) (count int, err error) {
	sql := `SELECT count(1) FROM mission_grow_record WHERE is_receive=0 AND uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

//获取所有任务总数
func GetMissionCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM mission WHERE is_used=0`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//获取每日任务完成数
func GetDailyMissionCount(uid int, complete_date string) (count int, err error) {
	sql := `SELECT count(1) FROM mission_daily_record WHERE uid =? AND complete_date=?`
	err = orm.NewOrm().Raw(sql, uid, complete_date).QueryRow(&count)
	return
}

//获取成长任务完成数
func GetGrowMissionCount(uid int) (count int, err error) {
	sql := `SELECT count(1) FROM mission_grow_record WHERE uid =? `
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

//获取已签到额外奖励id
func GetReceiveExtraId(uid, month, year int) (sign_reward_product_id []int, err error) {
	sql := `SELECT sign_reward_product_id FROM sign_reward_record WHERE uid=? AND month=? AND year=?`
	_, err = orm.NewOrm().Raw(sql, uid, month, year).QueryRows(&sign_reward_product_id)
	return
}

//获取已完成成长任务所有信息
func GetAlredayGrowMission(uid int) (mgr []MissionGrowRecord, err error) {
	sql := `SELECT mission_id,is_receive FROM mission_grow_record WHERE uid=?`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&mgr)
	return
}

//获取已完成每日任务所有信息
func GetAlredayDailyMission(uid int, complete_date string) (mdr []MissionDailyRecord, err error) {
	sql := `SELECT mission_id,is_receive FROM mission_daily_record WHERE uid = ? AND complete_date=?`
	_, err = orm.NewOrm().Raw(sql, uid, complete_date).QueryRows(&mdr)
	return
}

//查询任务奖励(融豆数量)
func GetMissionAward(id int) (mission_award int, err error) {
	sql := `SELECT mission_award FROM mission WHERE id=?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&mission_award)
	return
}

//记录用户收支记录 1：表示支付方式为融豆 2：代表收款 0：业务状态为成功
func AddUserFinanceRecord(uid, deal_type, score_amount, befor_score_amount, after_score_amount int) (err error) {
	sql := `INSERT INTO users_finance_record (uid,pay_type,deal_type,score_amount,pay_or_get,create_time,service_states,befor_score_amount,after_score_amount)VALUES(?,1,?,?,2,NOW(),0,?,?)`
	_, err = orm.NewOrm().Raw(sql, uid, deal_type, score_amount, befor_score_amount, after_score_amount).Exec()
	return
}

//获取手机运营商认证状态
func GetStatus(uid int) (is_mobile int, err error) {
	sql := `SELECT is_mobile_operators FROM users_auth WHERE uid= ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_mobile)
	return
}

//获取支付宝认证状态
func GetAlipayStatus(uid int) (is_alipay int, err error) {
	sql := `SELECT is_alipay_bqs FROM users_auth WHERE uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_alipay)
	return
}

//获取京东商城认证状态
func GetJdStatus(uid int) (is_jd int, err error) {
	sql := `SELECT is_jd_bqs FROM users_auth WHERE uid= ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_jd)
	return
}

//获取淘宝商城认证状态
func GetTbStatus(uid int) (is_tb int, err error) {
	sql := `SELECT is_tb_bqs FROM users_auth WHERE uid= ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_tb)
	return
}

//获取芝麻认证状态
func GetZmStatus(uid int) (is_zm int, err error) {
	sql := `SELECT is_zm_auth FROM users_auth WHERE uid= ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_zm)
	return
}

//获取身份认证状态
func GetIsRealNameStatus(uid int) (is_real int, err error) {
	sql := `SELECT is_real_name FROM users_auth WHERE uid= ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_real)
	return
}

//获取人脸识别状态
func GetIsFaceStatus(uid int) (is_face int, err error) {
	sql := `SELECT is_face_record FROM users_auth WHERE uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&is_face)
	return
}

//获取当前签到提醒状态
func GetSignRemindStatus(uid int) (status int, err error) {
	sql := `SELECT is_sign_prompt FROM users WHERE id = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&status)
	return
}

//更新当前签到提醒状态
func UpdateSignRemindStatus(status, uid int) (err error) {
	sql := `UPDATE users SET is_sign_prompt = ? WHERE id = ?`
	_, err = orm.NewOrm().Raw(sql, status, uid).Exec()
	return
}

//查看任务是否使用
func GetIsUsed(id int) (is_used int, err error) {
	sql := `SELECT is_used FROM mission WHERE id =?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&is_used)
	return
}

//在users表中更新最近签到时间
func UpdateLastSignTime(uid int) (err error) {
	sql := `UPDATE users SET last_sign_date = NOW() WHERE id = ?`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

//判断人脸任务是否完成
func GetFaceIsFinish(uid int) (count int, err error) {
	sql := `SELECT count(1) FROM mission_grow_record WHERE uid = ? AND mission_id = 12`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

type MissionImage struct {
	MissionId int    //
	ImgUrl    string //
}

//按包名获取图片url
func GetUrlByPackageId(packageId int) (missionImage []MissionImage, err error) {
	sql := `SELECT mission_id,img_url FROM mission_image WHERE package_id=?`
	_, err = orm.NewOrm().Raw(sql, packageId).QueryRows(&missionImage)
	return
}
