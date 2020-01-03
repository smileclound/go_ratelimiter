
package test

import (
	"errors"
	"fmt"
	"github.com/juju/testlimiter/web_filter"
	"io"
	"log"
	"net/http"
	"testing"
)

type HttpServer struct {

	http.Server

}

func (server *HttpServer) StartServer()  {
	log.Println("web server start "+server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func testHandler(w http.ResponseWriter,r *http.Request)error{
	w.WriteHeader(200)
	io.WriteString(w,"业务处理逻辑....")
	return nil
}

func (server *HttpServer)ServeHTTP(wr http.ResponseWriter, r *http.Request)()  {
	web_filter.Handle(testHandler)(wr,r)

	fmt.Println("测试")
}

func TestWebFilter(t *testing.T) {
	web_filter.Register("/safe/**", func(rw http.ResponseWriter, r *http.Request)error {
		return errors.New("解密失败")
		//return nil
	})
	web_filter.Register("/safe/user/**", func(rw http.ResponseWriter, r *http.Request)error {
		return errors.New("请登录")
		//return nil
	})
	http.HandleFunc("/safe", web_filter.Handle(func(wr http.ResponseWriter,req *http.Request) error{
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user/test", web_filter.Handle(func(wr http.ResponseWriter,req *http.Request) error{
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user", web_filter.Handle(func(wr http.ResponseWriter,req *http.Request) error{
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	server := &HttpServer{}
	server.Addr=":8080"
	server.StartServer()
}