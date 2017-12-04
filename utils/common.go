package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//随机数种子
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

//序列化
func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

//md5加密
func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}

// 获取数字随机字符
func GetRandDigit(n int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(n)+"d", rnd.Intn(int(math.Pow10(n))))
}

//处理5次操作(恶意操作处理)
func Check5Time(key string) int {
	count := 5
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, GetTodayLastSecond())
	return count
}

//处理5次操作(恶意操作处理)
func Check20Time(key string) int {
	count := 100
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	Rc.Put(key, count, GetTodayLastSecond())
	return count
}

// 获取随机数
func GetRandNumber(n int) int {
	return rnd.Intn(n)
}

//处理验证码获取5次处理Or 登录错误5次处理
func CheckPwd5Time(key string) int {
	count := 5
	if Rc.IsExist(key) {
		count, _ = Rc.RedisInt(key)
	}
	count--
	if count < 0 {
		return count
	}
	Rc.Put(key, count, GetTodayLastSecond())
	return count
}

// //获取相差时间-秒
// func GetSecondDiffer(start_time, end_time string) (int64, error) {

// 	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
// 	if err != nil {
// 		return 0, err
// 	}
// 	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)

// 	if err != nil {
// 		return 0, err
// 	}
// 	return GetSecondDifferByTime(t1, t2), nil
// }

//获取相差时间-秒
func GetSecondDifferByTime(start_time, end_time time.Time) int64 {
	diff := end_time.Unix() - start_time.Unix()
	return diff
}

func FixFloat(f float64, m int) float64 {
	newn := SubFloatToString(f+0.00000001, m)
	newf, _ := strconv.ParseFloat(newn, 64)
	return newf
}

var whoareyou = make(map[string]string)

func init() {
	// Rc, Re = cache.NewCache("redis", BEEGO_CACHE)
	var yidong []string = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188"}
	var liantong []string = []string{"130", "131", "132", "145", "155", "156", "176", "185", "186"}
	var dianxin []string = []string{"133", "153", "177", "180", "181", "189", "173"}
	for i := 0; i < len(yidong); i++ {
		whoareyou[yidong[i]] = "P100014"
	}
	for i := 0; i < len(liantong); i++ {
		whoareyou[liantong[i]] = "P100015"
	}
	for i := 0; i < len(dianxin); i++ {
		whoareyou[dianxin[i]] = "P100016"
	}
}

func WhoAreYou(account string) string {
	key := string(account[:3])
	return whoareyou[key]
}

//验证是否是手机号

func Validate(mobileNum string) bool {
	reg := regexp.MustCompile(Regular)
	return reg.MatchString(mobileNum)
}

func PageCount(count, pagesize int) int {
	if count%pagesize > 0 {
		return count/pagesize + 1
	} else {
		return count / pagesize
	}
}

func GetToday(format string) string {
	today := time.Now().Format(format)
	return today
}

//获取今天剩余秒数
func GetTodayLastSecond() time.Duration {
	today := GetToday(FormatDate) + " 23:59:59"
	end, _ := time.ParseInLocation(FormatDateTime, today, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}

// 处理出生日期函数
func GetBrithDate(idcard string) string {
	l := len(idcard)
	var s string
	if l == 15 {
		s = "19" + idcard[6:8] + "-" + idcard[8:10] + "-" + idcard[10:12]
		return s
	}
	if l == 18 {
		s = idcard[6:10] + "-" + idcard[10:12] + "-" + idcard[12:14]
		return s
	}
	return GetToday(FormatDate)
}

//处理性别
func WhichSexByIdcard(idcard string) string {
	var sexs = [2]string{"女", "男"}
	length := len(idcard)
	if length == 18 {
		sex, _ := strconv.Atoi(string(idcard[16]))
		return sexs[sex%2]
	} else if length == 15 {
		sex, _ := strconv.Atoi(string(idcard[14]))
		return sexs[sex%2]
	}
	return "男"
}

//截取小数点后几位
func SubFloatToString(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if m == 0 {
		return newn[0]
	}
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//截取小数点后几位
func SubFloatToFloat(f float64, m int) float64 {
	newn := SubFloatToString(f, m)
	newf, _ := strconv.ParseFloat(newn, 64)
	return newf
}

// func init() {
// 	fmt.Println(GetMinuteDiffer("2017-05-04 13:50", "2017-05-04 13:55"))
// }

//获取相差时间-年
func GetYearDiffer(start_time, end_time string) int64 {
	var Age int64

	t1, err := time.ParseInLocation("2006-01-02", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02", end_time, time.Local)

	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		Age = diff / (3600 * 365 * 24)
		return Age
	} else {
		return Age
	}
}

func init() {
	// Rc, Re = cache.NewCache("redis", BEEGO_CACHE)
	var yidong []string = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188"}
	var liantong []string = []string{"130", "131", "132", "145", "155", "156", "176", "185", "186"}
	var dianxin []string = []string{"133", "153", "177", "180", "181", "189", "173"}
	for i := 0; i < len(yidong); i++ {
		whoareyou[yidong[i]] = "P100014"
	}
	for i := 0; i < len(liantong); i++ {
		whoareyou[liantong[i]] = "P100015"
	}
	for i := 0; i < len(dianxin); i++ {
		whoareyou[dianxin[i]] = "P100016"
	}
}

func StartIndex(page, pagesize int) int {
	if page > 1 {
		return (page - 1) * pagesize
	}
	return 0
}

//发送邮件
func SendEmail(title, content, touser string) {
	host := "smtp.exmail.qq.com:25"
	to := strings.Split(touser, ";") //收件人  ;号隔开
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + touser + "\r\nFrom: jgl@zcmlc.com>\r\nSubject:" + title + "\r\n" + content_type + "\r\n\r\n" + content)
	err := smtp.SendMail(host, smtp.PlainAuth("", "jgl@zcmlc.com", "8050107Hyc", "smtp.exmail.qq.com"), "jgl@zcmlc.com", to, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}

func ErrNoRow() string {
	return "<QuerySeter> no row found"
}

//截取字符串
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 将图片下载并保存到本地
func SaveImage(imgUrl, saveDir, saveName string) {
	res, err := http.Get(imgUrl)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("%d HTTP ERROR:%s", imgUrl, err)
		return
	}
	//按分辨率目录保存图片
	if !isDirExist(saveDir) {
		os.MkdirAll(saveDir, os.ModePerm)
	}
	//根据URL文件名创建文件
	dst, err := os.Create(saveDir + "/" + saveName)
	if err != nil {
		fmt.Println("%d HTTP ERROR:%s", "A", err)
		return
	}
	// 写入文件
	io.Copy(dst, res.Body)
}

func isDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

//判断版本号
func GetAppVersionEquery(newVer, oldVer string) bool {
	if oldVer != "" && newVer != "" {
		oldArr := strings.Split(oldVer, ".")
		newArr := strings.Split(newVer, ".")
		if newArr[0] > oldArr[0] {
			return true
		}
		if newArr[1] > oldArr[1] {
			return true
		}
		if newArr[2] > oldArr[2] {
			return true
		}
		return false
	} else {
		return false
	}
}

//期数换算成天
func GetLoanTermDays(param string) int {
	if param == "" {
		return 0
	}
	plen := len(param)
	unit := Substr(param, plen-3, 3)
	days := 0
	switch unit {
	case "天":
		dayStr := Substr(param, 0, plen-3)
		dayInt, err := strconv.Atoi(dayStr)
		if err != nil {
			dayInt = 0
		}
		days = dayInt
	case "月":
		month := Substr(param, 0, plen-3)
		monthInt, err := strconv.Atoi(month)
		if err != nil {
			monthInt = 0
		}
		days = monthInt * 30
	case "年":
		year := Substr(param, 0, plen-3)
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			yearInt = 0
		}
		days = yearInt * 360
	}
	return days
}

//期数换算成年或月
func GetLoanTermYOrM(param int) string {
	result := 0
	unit := ""
	if param >= 360 {
		result = param / 360
		unit = "年"
	} else if param >= 30 && param < 360 {
		result = param / 30
		unit = "月"
	} else {
		result = param
		unit = "天"
	}
	return strconv.Itoa(result) + unit
}

func HttpValueDecode(res []byte) []byte {
	strP := strings.Split(string(res), "&")
	m := map[string]string{}
	for i := 0; i < len(strP)-1; i++ {
		matr := strings.Split(strP[i], "=")
		matr[1], _ = url.QueryUnescape(matr[1])
		m[matr[0]] = matr[1]
	}
	str, _ := json.Marshal(m)
	return str
}

// 判断用户是否付款
func CheckPay(uid, serviceId int, token string) bool {
	key := CACHE_KEY_SERVICE_PAY_TOKEN + strconv.Itoa(uid) + strconv.Itoa(serviceId)
	if Rc.IsExist(key) {
		str, _ := Rc.RedisString(key)
		if str == token {
			Rc.Delete(key)
			return true
		}
	}
	return false
}
