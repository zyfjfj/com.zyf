package simple

import (
	"io"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

//实现了ServeHTTP接口
type engine struct {
	muxTree map[string][]interface{}
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := e.muxTree[r.URL.String()]; ok {
		switch r.Method {
		case "GET":
			if h[0].(string) == "GET" {
				h[1].(HandlerFunc)(w, r)
				return
			}
		case "POST":
			if h[0].(string) == "POST" {
				h[1].(HandlerFunc)(w, r)
				return
			}
		}
	}
	io.WriteString(w, "no find: "+r.Method+","+r.URL.String())
}
func (e *engine) GET(url string, handle HandlerFunc, describe string) {
	e.addMuxTree(url, "GET", handle, describe)
}
func (e *engine) POST(url string, handle HandlerFunc, describe string) {
	e.addMuxTree(url, "POST", handle, describe)
}

//增加url和处理handle的字典
func (e *engine) addMuxTree(url string, method string, handle HandlerFunc, describe string) {
	e.muxTree[url] = append(e.muxTree[url], method, handle, describe)
}

type TcpListen struct {
}

func (t *TcpListen) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.Host)
}
func (t *TcpListen) Listen(port string) {
	http.ListenAndServe(port, t)
}
func New() *engine {
	e := new(engine)
	e.muxTree = make(map[string][]interface{})
	return e
}

//启动
//addr1 主监听，addr2 监听设备上传数据
func (e *engine) Run(addr1 string, addr2 string) {
	if addr2 != "" {
		t := new(TcpListen)
		go t.Listen(addr2)
	}

	server := http.Server{
		Addr:    addr1,
		Handler: e,
	}
	server.ListenAndServe()

}
