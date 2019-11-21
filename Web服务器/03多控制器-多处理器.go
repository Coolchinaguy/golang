package main

import (
	"net/http"
)

type Firsttest struct {
}

type Secondtest struct {
}

func (m1 *Firsttest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("返回的数据1"))
}

func (m2 *Secondtest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("返回的数据2"))
}

func main() {
	h1 := Firsttest{}
	h2 := Secondtest{}
	sever := http.Server{Addr: "localhost:8000"}

	http.Handle("/first", &h1)
	http.Handle("/second", &h2)
	sever.ListenAndServe()
}
