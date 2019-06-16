package main

import (
	"github.com/wonderivan/logger"
	"github.com/devedge/imagehash"
	"github.com/anthonynsimon/bild/transform"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	"encoding/hex"
	"encoding/json"
	"encoding/base64"
	"image"
	"net/http"
	"net/url"
	"os"
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"flag"
	"fmt"
)

var (
	u string
	p string
	t bool
)

func init(){
	flag.StringVar(&u, "u", "", "设置校园网登录学号")
	flag.StringVar(&p, "p", "", "设置校园网登录密码(身份证后8位)")
	flag.BoolVar(&t, "t", false, "循环运行(按Ctrl+C结束程序)")
	flag.Usage = usage
}

func usage(){
	fmt.Println("Usage: zquAutoLogin-go -u [studentId] -p [password]\nOptions:")
	flag.PrintDefaults()
}

func main(){
	flag.Parse()
	
	if u == "" || p == ""{
		flag.Usage()
		return
	}
	
	logger.Info("-v1.2- 学号:", u, "密码:", p)
	
	networkTest()
	
	for t {
		time.Sleep(time.Duration(1)*time.Minute)
		networkTest()
	}
}

func networkTest() {
	url := "http://quan.suning.com/getSysTime.do"
	resp, err := http.Head(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	test_status := resp.Request.URL.Host
	
	if test_status == "quan.suning.com" {
		logger.Info("可以上网")
	} else if test_status == "10.0.1.51" {
		resp.Request.ParseForm();
		formStr := "wlanuserip="
		formStr += resp.Request.Form.Get("wlanuserip")
		formStr += "&wlanacname="
		formStr += resp.Request.Form.Get("wlanacname")
		formStr += "&ssid="
		formStr += resp.Request.Form.Get("ssid")
		formStr += "&nasip="
		formStr += resp.Request.Form.Get("nasip")
		formStr += "&mac="
		formStr += resp.Request.Form.Get("mac")
		formStr += "&t="
		formStr += resp.Request.Form.Get("t")
		formStr += "&url="
		formStr += resp.Request.Form.Get("url")	
		result, message := autoLogin_1(formStr)
		if result == "success" {
			logger.Info("已经通过局域网验证")
			time.Sleep(time.Duration(3)*time.Second)
			networkTest()
		} else {
			logger.Error(message)
		}
	} else if test_status == "enet.10000.gd.cn:10001"{
		if autoLogin_2(u, p) == "success" {
			time.Sleep(time.Duration(3)*time.Second)
			networkTest()
		}
	}
}


func autoLogin_1(str string) (string, string) {
	userid := u
	password := string([]rune(p)[2:])
	
	login_url := "http://10.0.1.51/eportal/InterFace.do?method=login"
	form := url.Values{
		"userId" : {userid},
		"password" : {password},
		"queryString" : {str},
	}

	referer := "http://10.0.1.51/eportal/index.jsp?"
	referer += str

	client := &http.Client{}
	req, _ := http.NewRequest("POST", login_url, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-Hans-CN, zh-Hans; q=0.5")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Host", "10.0.1.51")
	req.Header.Set("Origin", "http://10.0.1.51")
	req.Header.Set("Referer", referer)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var m map[string]interface{}
	json.Unmarshal(body, &m)

	return m["result"].(string), m["message"].(string)

}


func autoLogin_2 (userid, password string) string {

	test_url := "http://quan.suning.com/getSysTime.do"
	login_url := "http://enet.10000.gd.cn:10001/login.do"
	captcha_url := "http://enet.10000.gd.cn:10001/common/image.jsp"

	resp, _ := http.Head(test_url)	
	defer resp.Body.Close()
	resp.Request.ParseForm();
	formStr := "wlanacip="
	formStr += resp.Request.Form.Get("wlanacip")
	formStr += "&wlanuserip="
	formStr += resp.Request.Form.Get("wlanuserip")
	wlanacip := resp.Request.Form.Get("wlanacip")
	wlanuserip := resp.Request.Form.Get("wlanuserip")
	//logger.Debug("wlanacip:", wlanacip, "wlanuserip:", wlanuserip)
	
	client := &http.Client{}
	req, _ := http.NewRequest("GET", captcha_url, nil)
	resp, _ = client.Do(req)
	cookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]
	
	req, _ = http.NewRequest("GET", captcha_url, nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cookie", cookie)

	resp, _ = client.Do(req)
	img, _ := os.Create("image.jpg")
	io.Copy(img, resp.Body)
	captcha := getCaptcha()
	encode := base64.StdEncoding.EncodeToString([]byte(password))
	
	form := url.Values{
		"edubas" : {wlanacip},
		"eduuser" : {wlanuserip},
		"userName1" : {userid},
		"password1" : {encode},
		"patch" : {"wifi"},
		"rand" : {captcha},
	}
	
	referer := "http://enet.10000.gd.cn:10001/zq/zq251/index.jsp?"
	referer += formStr
	
	req, _ = http.NewRequest("POST", login_url, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "110")
	req.Header.Set("Host", "enet.10000.gd.cn:10001")
	req.Header.Set("Origin", "http://enet.10000.gd.cn:10001")
	req.Header.Set("Referer", referer)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Cookie", cookie)
	
	logger.Info("正在登录电信网络")
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := string(body)
	if strings.Index(result, "success") > 0 {
		logger.Info("登录成功")
		return "success"
	} else {
		r := result[402 : len(result)-600]
		val := strings.Index(r, "\"")
		r = r[:val]
		logger.Error(r)
		return "fail"
	}

}

func getCaptcha() string {
	src,_ := imagehash.OpenImg("./image.jpg")
	
	x := 13
	h1,_ := imagehash.DhashHorizontal(segment.Threshold(effect.Grayscale(transform.Crop(src, image.Rect(5, 3, 5+x, 3+x))), 128), 8)
	h2,_ := imagehash.DhashHorizontal(segment.Threshold(effect.Grayscale(transform.Crop(src, image.Rect(5+x, 3, 5+2*x, 3+x))), 128), 8)
	h3,_ := imagehash.DhashHorizontal(segment.Threshold(effect.Grayscale(transform.Crop(src, image.Rect(5+2*x, 3, 5+3*x, 3+x))), 128), 8)
	h4,_ := imagehash.DhashHorizontal(segment.Threshold(effect.Grayscale(transform.Crop(src, image.Rect(5+3*x, 3, 5+4*x, 3+x))), 128), 8)
		
	var numMap map[int]string
	numMap = make(map[int]string)
	numMap [0] = "0e33232b2b23330e"
	numMap [1] = "0d0d0d2d2d0c0f03"
	numMap [2] = "0f6202264d983606"
	numMap [3] = "0f72060d0602660f"
	numMap [4] = "26461696370b2706"
	numMap [5] = "0626203e0602660f"
	numMap [6] = "06b3312e33233306"
	numMap [7] = "030b02062d4a0a9a"
	numMap [8] = "8632329637233386"
	numMap [9] = "8e3323338b43328f"
	
	min := 9999
	min_p1 := 0
	for i := 0; i < 10; i++ {
		numVal, _ := hex.DecodeString(numMap[i])
		if imagehash.GetDistance(h1 , numVal) < min {
			min = imagehash.GetDistance(h1, numVal)
			min_p1 = i
		}
	}
	min = 9999
	min_p2 := 0
	for i := 0; i < 10; i++ {
                numVal, _ := hex.DecodeString(numMap[i])
                if imagehash.GetDistance(h2 , numVal) < min {
			min = imagehash.GetDistance(h2, numVal)
			min_p2 = i 
		}
	}
	min = 9999
	min_p3 := 0
        for i := 0; i < 10; i++ {
                numVal, _ := hex.DecodeString(numMap[i])
                if imagehash.GetDistance(h3 , numVal) < min {
                        min = imagehash.GetDistance(h3, numVal)
                        min_p3 = i
                }
	}
	min = 9999
        min_p4 := 0
        for i := 0; i < 10; i++ {
                numVal, _ := hex.DecodeString(numMap[i])
                if imagehash.GetDistance(h4 , numVal) < min {
                        min = imagehash.GetDistance(h4, numVal)
                        min_p4 = i
                }
	}
	
	code := strconv.Itoa(min_p1)
	code += strconv.Itoa(min_p2)
	code += strconv.Itoa(min_p3)
	code += strconv.Itoa(min_p4)	
	
	return code
}