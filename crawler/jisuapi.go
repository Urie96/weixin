package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

func AIQA(question string) string {
	url := "https://www.jisuapi.com/debug/iqa/?act=relay"
	body := fmt.Sprintf("url=%s&question=%s", "https://api.jisuapi.com/iqa/query", question)
	resp := post(url, body)
	b, _ := ioutil.ReadAll(resp.Body)
	if string(b[:9]) == "<!DOCTYPE" { //need login
		login()
		resp = post(url, body)
		b, _ = ioutil.ReadAll(resp.Body)
	}
	val, _ := jsonparser.GetString(b, "body")
	content, _ := jsonparser.GetString([]byte(val), "result", "content")
	return content
}

func login() {
	url := "https://www.jisuapi.com/my/login?act=login&rtype=json"
	body := fmt.Sprintf("email=&password=Youling1996&mobile=17828228827")
	resp := post(url, body)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func post(url, body string) *http.Response {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	checkError(err)
	req.Header.Add("Cookie", "PHPSESSID=6f95374435383219a1eb72ae76eb244f")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	checkError(err)
	return res
}
