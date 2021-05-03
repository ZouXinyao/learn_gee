package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
http.ListenAndServe用来启动 Web 服务:
第一个参数是地址，:9999表示在 9999 端口监听。
第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。
第二个参数，是基于net/http标准库实现Web框架的入口。
*/
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
