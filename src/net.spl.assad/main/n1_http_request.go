package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type User struct {
	Id, Name, City string
	Scope          int
}

// @desc http request 请求示例
func main() {
	//getSpl()
	//getWithParamsSpl()
	//postFormSpl()
	//postJsonSpl()
	httpClientSpl()
}

// GET 请求示例
func getSpl() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bodyString))
}

// GET 请求 - 携带 params
func getWithParamsSpl() {
	// 请求 Url
	myUrl, _ := url.Parse("https://ac05e368-dc52-4c80-a6ab-383522709917.mock.pstmn.io/user")
	// 请求参数
	params := url.Values{}
	params.Set("userId", "233")
	// 设置 URL 请求参数
	myUrl.RawQuery = params.Encode()
	// 发送 Get 请求
	resp, _ := http.Get(myUrl.String())
	defer resp.Body.Close()

	// 解析响应 json
	respBody, _ := ioutil.ReadAll(resp.Body)
	var user User
	json.Unmarshal(respBody, &user)
	fmt.Println(user)
}

// POST Form 请求
func postFormSpl() {
	myUrl := "https://ac05e368-dc52-4c80-a6ab-383522709917.mock.pstmn.io/userdetail"
	body := url.Values{"userid": {"233"}, "name": {"vancy"}, "scope": {"231"}}
	// 发送 post form 请求
	resp, _ := http.PostForm(myUrl, body)
	defer resp.Body.Close()
	// 读取响应
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

// Post Json 请求
func postJsonSpl() {
	// 发送 post 请求
	myUrl := "https://ac05e368-dc52-4c80-a6ab-383522709917.mock.pstmn.io/userdetail"
	user := User{"233", "assad", "guangzhou", 234521}
	requestBody, _ := json.Marshal(user)
	// 发送 post 请求
	resp, _ := http.Post(myUrl, "application/json", strings.NewReader(string(requestBody)))
	// 读取响应
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

// 直接使用 http.Client 进行 header 等参数的进一步调节
func httpClientSpl() {

	// 创建 Client 实例
	client := &http.Client{}
	// 请求 url
	myUrl, _ := url.Parse("https://api.bilibili.com/x/tag/detail?pn=0&ps=1&tag_id=17865158")
	myUrl.RawQuery = url.Values{"ps": {"1"}, "tag_id": {"17865158"}}.Encode()
	// 构建 Request
	req, _ := http.NewRequest(http.MethodGet, myUrl.String(), nil)
	// 设置 Header
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", "https://www.bilibili.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15")

	// 发送 http 请求
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	// 解析响应
	fmt.Println(resp.StatusCode, resp.Status)
	fmt.Println(resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

}
