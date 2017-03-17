//用于测试用go现有的web服务，自己写个框架
package main

import (
	"io"
	"net/http"

	"com.zyf/pureweb/simple"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ddddf")
}
func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "登陆界面")
}
func main() {
	engin := simple.New()
	engin.GET("/", hello, "欢迎")
	engin.POST("/login", login, "登入")
	engin.Run(":8888")
}
