package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

func main() {
	url := "http://openapi.tuling123.com/openapi/api/v2"
	postbody := fmt.Sprintf(`{
		"reqType":0,
		"perception": {
			"inputText": {
				"text": "%s"
			}
		},
		"userInfo": {
			"apiKey": "8de9915257594bf3a4b4ad1f2d6f1769",
			"userId": "%s"
		}
	}`, "你好", "111")
	resp, _ := http.Post(url, "application/json", strings.NewReader(postbody))
	b, _ := ioutil.ReadAll(resp.Body)
	var val string
	fmt.Println(string(b))
	jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		val, _ = jsonparser.GetString(value, "values", "text")
	}, "results")
	fmt.Println(val)
	return
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
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
