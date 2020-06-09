package crawler

import (
	"io/ioutil"
	"net/http"
)

func GetLoveTalk() string {
	resp, _ := http.Get("https://api.lovelive.tools/api/SweetNothings")
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}
