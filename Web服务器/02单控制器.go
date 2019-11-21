package main

import (
	"net/http"
)

type MyHander struct {
}

func (m *MyHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("返回的数据"))
}

func main() {

	h := MyHander{}
	server := http.Server{Addr: "127.0.0.1:8000", Handler: &h}
	server.ListenAndServe()

}
