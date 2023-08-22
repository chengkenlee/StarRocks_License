package app

import (
	"StarRocks_License/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type License struct {
	Code int `json:"code"`
	List []struct {
		Cores    int   `json:"cores"`
		ExpireAt int64 `json:"expire_at"`
		Hosts    int   `json:"hosts"`
	} `json:"list"`
	Total int `json:"total"`
}

func LicenseCrontab() {
	//go func() {
	//	crontab := cron.New()
	//	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	//	_, err := crontab.AddFunc(util.Config.GetString("starrocks.license.crontab"), license)
	//	if err != nil {
	//		util.Logger.Error(err.Error())
	//		return
	//	}
	//	// 启动定时器
	//	crontab.Start()
	//	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	//	// 根据实际情况进行控制
	//	select {}
	//}()
	license()
}

func license() {
	util.Logger.Info("license 检测")
	/*登录web, 拿到respone,cookie*/
	var wait sync.WaitGroup

	for cluster := range util.Config.GetStringMapStringSlice("starrocks") {
		s := util.Config.GetStringMapStringSlice("starrocks")[cluster]
		wait.Add(1)
		cluster := cluster
		go func() {
			defer wait.Done()
			manager := s[0]
			user := s[1]
			password := s[2]

			msg := fmt.Sprintf(`{"name":"%s","password":"%s"}`, user, password)

			/*新的合并模块---------获取Cookies*/
			request, err := http.NewRequest("POST", manager+"/api/user/login", strings.NewReader(msg))
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			request.Header.Set("Content-Type", "application/json;charset=utf-8")
			client := &http.Client{Timeout: time.Second * time.Duration(util.Config.GetInt("starrocks.web.timeout"))}
			respone, err := client.Do(request)
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			defer respone.Body.Close()
			b, err := ioutil.ReadAll(respone.Body)
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			var Cookies *http.Cookie
			for _, c := range respone.Cookies() {
				Cookies = &http.Cookie{
					Name:       c.Name,
					Value:      c.Value,
					Path:       c.Path,
					Domain:     c.Domain,
					Expires:    c.Expires,
					RawExpires: c.RawExpires,
					MaxAge:     c.MaxAge,
					Secure:     c.Secure,
					HttpOnly:   c.HttpOnly,
					SameSite:   c.SameSite,
					Raw:        c.Raw,
					Unparsed:   c.Unparsed,
				}
			}
			util.Logger.Info(string(b))
			/*新的合并模块---------重新利用Cookies请求license*/
			request2, err := http.NewRequest("GET", manager+"/api/license/list", strings.NewReader(msg))
			if err != nil {
				util.Logger.Error(err.Error())
				return
			}
			/*license已经过期*/
			var day int
			var m string
			if !strings.Contains(string(b), "add license") {
				request2.AddCookie(Cookies)
				request2.Header.Set("Content-Type", "application/json;charset=utf-8")
				client2 := &http.Client{Timeout: time.Second * time.Duration(util.Config.GetInt("starrocks.web.timeout"))}
				respone2, err := client2.Do(request2)
				if err != nil {
					util.Logger.Error(err.Error())
					return
				}
				defer respone2.Body.Close()
				b2, err := ioutil.ReadAll(respone2.Body)
				if err != nil {
					util.Logger.Error(err.Error())
					return
				}
				util.Logger.Info(string(b2))
				/*新的合并模块---------解析struct*/
				var license2 License
				err = json.Unmarshal(b2, &license2)
				if err != nil {
					util.Logger.Error(err.Error())
					return
				}
				l := license2.List[0]
				/*计算过期时间差距*/
				day = getday(fmt.Sprintf("%d", time.Now().UnixNano()/1e6), fmt.Sprintf("%d", l.ExpireAt))
				m = fmt.Sprintf("集群：%s,核数：%d,license过期时间：%s,节点数：%d, license %d天后将过期!", cluster, l.Cores, UnixToTime(strconv.FormatInt(l.ExpireAt, 10)).Format("2006-01-02 15:04:05"), l.Hosts, day)
			} else {
				m = fmt.Sprintf("集群：%s license目前已经属于过期状态！", cluster)
			}
			util.Logger.Info(m)
			util.Logger.Info(fmt.Sprintf("%s 现在day的值是：%d", cluster, day))
		}()
	}
	wait.Wait()
	util.Logger.Info("done!")
}

func UnixToTime(e string) (datatime time.Time) {
	data, _ := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}

func getday(date1Str, date2Str string) int {
	util.Logger.Info("现在日期:" + date1Str + ",到期日期:" + date2Str)
	// 将字符串转化为Time格式
	date1, err := time.ParseInLocation("2006-01-02", UnixToTime(date1Str).Format("2006-01-02"), time.Local)
	if err != nil {
		return 0
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation("2006-01-02", UnixToTime(date2Str).Format("2006-01-02"), time.Local)
	if err != nil {
		return 0
	}
	//计算相差天数
	return int(date2.Sub(date1).Hours() / 24)
}
