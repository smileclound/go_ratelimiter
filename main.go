package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"io"
	"log"
	"net/http"
	"time"
)

type webHandler func(http.ResponseWriter, *http.Request)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(200)
	io.WriteString(wr, "hello world")

}

func defaultHandler(wr http.ResponseWriter,r *http.Request)  {
	wr.WriteHeader(200)
	io.WriteString(wr, "refused by limiter....")
}

func MyHandleFunc(handler webHandler) webHandler {

	return func(writer http.ResponseWriter, request *http.Request) {
		//流量控制
		if ok :=lim.AllowN(time.Now(),1);ok{
			fmt.Println("accept ")
			handler(writer,request)
		}else {
			log.Print("refused by limiter.....")
			fmt.Errorf("refused by limiter.....")
			 defaultHandler(writer,request)
		}

	}



}

var lim *rate.Limiter

func main() {
	lim = rate.NewLimiter(2, 1)
	http.HandleFunc("/",MyHandleFunc(sayhello))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
