package main

import (
	"io"
	"net/http"
)

// @desc go http server 简单示例

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":2333", mux)
}

// 路由处理器1
func helloHandler(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}

// 路由处理器2
func echoHandler(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, req.URL.Path+" "+req.RemoteAddr)
}
