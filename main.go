package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/wx", handleWx)
	http.ListenAndServe(":7001", nil)
}

func handleWx(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(1)
	req.ParseForm()
	for i, v := range req.Form {
		fmt.Println(i)
		fmt.Println(v)
	}
	resp.Write([]byte("get"))
}
