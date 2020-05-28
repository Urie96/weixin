package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func generateHtml(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(string(body))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
