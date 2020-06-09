package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

func TuLingAPI(text, openid string) string {
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
	}`, text, openid)
	resp, _ := http.Post(url, "application/json", strings.NewReader(postbody))
	b, _ := ioutil.ReadAll(resp.Body)
	var val string
	jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		val, _ = jsonparser.GetString(value, "values", "text")
	}, "results")
	return val
}
