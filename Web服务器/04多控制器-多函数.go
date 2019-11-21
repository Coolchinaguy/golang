package main

import (
	"net/http"
)

func Handle1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("多函数返回的数据1"))
}

func Handle2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("多函数返回的数据2"))
}

type MyHander struct {
}

func main() {
	sever := http.Server{Addr: "localhost:8000"}
	http.HandleFunc("/first", Handle1)
	http.HandleFunc("/second", Handle2)
	sever.ListenAndServe()
}
