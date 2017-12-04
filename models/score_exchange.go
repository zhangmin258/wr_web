package models

import (
	"encoding/json"
	"strconv"
	"time"
	"wr_web/utils"

	"github.com/astaxie/beego/orm"
)

type ShowSeProducts struct {
	ServiceId        int     `orm:"column(id)"`
	Content          string  //兑换商品名称
	ScorePrice       int     //兑换商品融豆价格
	MoneyPrice       float64 //兑换商品人民币价格
	IsShowInExchange int     //是否在兑换页面展示
	IsVipFree        int     //VIP是否免费
	ImgUrl           string  //logo图片url
}

type MoneyService struct {
	Id               int
	Content          string
	ScorePrice       int
	MoneyPrice       float64
	IsVipFree        int    //vip是否免费：0免费 1不免费
	DescribeTitle1   string //描述标题
	DescribeContent1 string //描述内容
	DescribeTitle2   string
	DescribeContent2 string
}

type ShowSeRecord struct {
	ServiceId   int       `orm:"column(id)"`
	ImgUrl      string    //logo图片url
	Content     string    //名称
	Title       string    //标题
	PayType     int       //付款方式
	CreateTime  time.Time //兑换时间
	ScoreAmount int       //融豆
	MoneyAmount float64   //人民币
	Pay         string    //付款
}

type ShowLotteryProducts struct {
	LotteryId       int    `orm:"column(id)"`
	Content         string //奖品名称
	Price           int    //抽奖价格
	ProType         int    //奖品类型：1.融豆；2.现金；3.实物；4.谢谢惠顾
	ScorePay        int    //当奖品为融豆时，奖品价值多少融豆
	MoneyPay        int    //当奖品为现金时，奖品价值多少现金
	IsShowInLottery int    //奖品是否需要展示给用户看到
}

type SeProducts struct {
	BaseModels
	Price int //积分兑换的价格
}

type LotteryBalance struct {
	Money   int //奖池剩余金额
	RongDou int //奖池剩余融豆
}

type UsersPayPwdInfo struct {
	IsVIP          int       //是否VIP
	AccountBalance float64   //钱包余额
	Score          int       //融豆余额
	IsPayPwd       int       //是否设置支付密码
	IsBinding      int       //是否绑卡
	VIPEndTime     time.Time //VIP到期时间
}

type SeRongDouBalance struct {
	Score    int //
	IsFreeze int //
}
type UsersVipInfo struct {
	Uid       int       //用户id
	BeginTime time.Time //vip开始时间
	EndTime   time.Time //vip结束时间
}

type UsersFinanceRecord struct {
	Uid              int       //用户ID
	PayToken         string    //付款凭证（用户ID+商品ID）
	PayType          int       //支付方式 1融豆；2现金余额；3会员免费
	DealType         int       //交易类型
	ScoreAmount      int       //融豆金额
	MoneyAmount      float64   //现金金额
	PayOrGet         int       //1：收款   2：付款
	CreateTime       time.Time //创建时间
	BeforScoreAmount int       //交易前融豆余额
	AfterScoreAmount int       //交易后融豆余额
	BeforMoneyAmount float64   //交易前钱包余额
	AfterMoneyAmount float64   //交易后钱包余额
}

//中奖记录
type LotteryRecord struct {
	Content    string    //奖品名称
	CreateTime time.Time //中奖时间
}

//用户账号信息
type UsersMini struct {
	Uid     int    //用户id
	UserImg string //用户头像
	Account string //账号/手机号
	IdNo    string //身份证号码
	IdName  string //IDnum
}

//查询可兑换的商品并展示到前台
func GetShowSeProducts() (showSeProducts []ShowSeProducts, err error) {
	sql := `SELECT id,content,score_price,money_price,is_show_in_exchange,is_vip_free,img_url FROM score_exchange_product WHERE is_show_in_exchange=0 AND is_vip_free=1`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&showSeProducts)
	return
}

//查询抽奖的产品并展示到前台
func GetLotteryProducts() (slp []ShowLotteryProducts, err error) {
	sql := `SELECT id,content,price,pro_type,score_pay,money_pay,is_show_in_lottery FROM score_lottery_product WHERE is_show_in_lottery=0`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&slp)
	return
}

//根据Uid查询用户头像
func GetUsersLogoImg(uid int) (user_img string, err error) {
	sql := `SELECT user_img FROM users WHERE id=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&user_img)
	return
}

//查询兑换记录并展示到前台
func GetScoreExchangeRecord(uid int) (showSeRecord []*ShowSeRecord, err error) {
	sql := `SELECT s.id,
			s.content,
            u.pay_type,
          u.score_amount,
           u.money_amount,
			u.create_time,
			s.img_url
		FROM
			score_exchange_product s
		RIGHT JOIN users_finance_record u
		ON u.deal_type=s.id
		WHERE
			u.uid =?
		AND u.pay_or_get=1
	    AND u.order_code IS NULL
        AND service_states = 0
		ORDER BY u.create_time DESC`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&showSeRecord)
	return
}

//根据id查询商品信息
func GetLotteryPriceById(id int) (showSeProducts ShowSeProducts, err error) {
	sql := `SELECT id,content,score_price,money_price,is_show_in_exchange,is_vip_free FROM score_exchange_product WHERE id=?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&showSeProducts)
	return
}

//根据用户ID查询余额
func GetUsersRongDouByUid(uid int) (srdb SeRongDouBalance, err error) {
	sql := `SELECT score,is_freeze FROM score WHERE uid=?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&srdb)
	return
}

//中奖快报轮播奖品查询
func GetLotteryReportProduct() (product []string, err error) {
	sql := `SELECT content FROM score_lottery_product WHERE content != '谢谢惠顾'`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&product)
	return
}

//查询抽奖价格
func GetLotteryPrice(id int) (moneyService *MoneyService, err error) {
	sql := `SELECT id,content,score_price,money_price,is_vip_free,describe_title1,describe_content1,describe_title2,describe_content2 FROM score_exchange_product WHERE id=?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&moneyService)
	if err == nil {
		if data, err2 := json.Marshal(moneyService); err2 == nil && utils.Re == nil {
			utils.Rc.Put(utils.WR_CACHE_KEY_SCORE_SERVICE+strconv.Itoa(moneyService.Id), data, 24*time.Hour)
		}
	}
	return
}

//积分兑换或积分抽奖融豆余额更新
func DeductBalanceByUid(uid, price int) error {
	sql := `UPDATE score SET score=? WHERE uid=?`
	_, err := orm.NewOrm().Raw(sql, price, uid).Exec()
	return err
}

//积分抽奖钱包余额更新
func LotteryAddBalanceByUid(uid int, price float64) error {
	sql := `UPDATE wallet SET account_balance=? WHERE uid=?`
	_, err := orm.NewOrm().Raw(sql, price, uid).Exec()
	return err
}

//存储用户中奖记录
func AddUsersLottery(uid, slProductId int) error {
	sql := `INSERT INTO score_lottery_record(uid,sl_product_id,create_time) VALUES(?,?,NOW())`
	_, err := orm.NewOrm().Raw(sql, uid, slProductId).Exec()
	return err
}

//获取用户vip信息
func GetUsersVIPInfoByUid(uid int) (vipInfo UsersVipInfo, err error) {
	sql := `SELECT begin_time,end_time FROM vip WHERE uid = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, uid).QueryRow(&vipInfo)
	return
}

//通过Uid获取用户钱包余额
func GetWalletBalanceByUid(uid int) (money float64, isFreeze int, err error) {
	sql := `SELECT account_balance,is_freeze FROM wallet WHERE uid = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, uid).QueryRow(&money, &isFreeze)
	return
}

//判断用户是否绑卡
func IsHadBankCard(uid int) (count int, err error) {
	o := orm.NewOrm()
	sql := `SELECT count(1) FROM users_bankcards WHERE uid = ? AND state='USING' `
	err = o.Raw(sql, uid).QueryRow(&count)
	return count, err
}

//通过Uid获取用户支付密码
func GetUsersPayPwdByUid(uid int) (password string, err error) {
	sql := `SELECT pay_pwd FROM users_metadata WHERE uid = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, uid).QueryRow(&password)
	return
}

//记录用户收支记录
func AddUsersFinanceRecord(ufr UsersFinanceRecord) error {
	sql := `INSERT INTO users_finance_record (uid,pay_type,deal_type,score_amount,money_amount,pay_or_get,create_time,befor_score_amount,after_score_amount,befor_money_amount,after_money_amount)VALUES(?,?,?,?,?,?,NOW(),?,?,?,?)`
	_, err := orm.NewOrm().Raw(sql, ufr.Uid, ufr.PayType, ufr.DealType, ufr.ScoreAmount, ufr.MoneyAmount, ufr.PayOrGet, ufr.BeforScoreAmount, ufr.AfterScoreAmount, ufr.BeforMoneyAmount, ufr.AfterMoneyAmount).Exec()
	return err
}

//查询用户中奖记录
func GetLotteryRecordByUid(uid int) (lotteryRecord []LotteryRecord, err error) {
	sql := `SELECT p.content,r.create_time FROM score_lottery_product p RIGHT JOIN score_lottery_record r ON p.id=r.sl_product_id WHERE r.uid=? ORDER BY r.create_time DESC`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&lotteryRecord)
	return
}

//活动开放校验
func GetLotteryValueByCode(code string) (cf *Config, err error) {
	sql := `SELECT id,config_key,config_value,config_desc,remark,config_url FROM config WHERE code = ?`
	err = orm.NewOrm().Raw(sql, code).QueryRow(&cf)
	return
}

//通过Uid查询用户信息:账号信息页面和个人中心页面
func SearchUserInfo(uid int) (usersMini *UsersMini, err error) {
	o := orm.NewOrm()
	sql := `		SELECT
		  u.id       AS uid,
		  um.head_portrait AS user_img,
		  u.account  AS account,
		  i.id_no    AS id_no,
		  i.id_name

		FROM users u
		  LEFT JOIN identification i
		    ON u.id = i.uid
		    LEFT JOIN users_metadata um ON um.uid=u.id
		WHERE u.id = ?`
	err = o.Raw(sql, uid).QueryRow(&usersMini)
	return
}

//查询用户未读铃铛消息条数
func GetMessageCountByUid(uid int) (int, error) {
	o := orm.NewOrm()
	var count int
	err := o.Raw(`select count(id) from users_message  where uid = ? and is_read=0 `, uid).QueryRow(&count)
	return count, err
}
