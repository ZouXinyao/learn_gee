/*
使用Handler接口
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

/*
第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

/*
实现Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口。
可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。
*/

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
