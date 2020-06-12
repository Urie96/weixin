package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
GetSolar convert lunar to solar
*/
func GetSolar(year, month, day int) string {
	reqBody := fmt.Sprintf("type=lunar&year=%d&month=%d&day=%d", year, month, day)
	resp, err := http.Post("https://www.iamwawa.cn/home/nongli/ajax",
		"application/x-www-form-urlencoded", strings.NewReader(reqBody))
	checkError(err)
	b, _ := ioutil.ReadAll(resp.Body)
	model := &struct {
		Data struct {
			Solar string `json:"solar"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(b, model)
	checkError(err)
	solar := model.Data.Solar
	solar = strings.ReplaceAll(solar, "年", "/")
	solar = strings.ReplaceAll(solar, "月", "/")
	solar = strings.ReplaceAll(solar, "日", "/")
	return solar
}
