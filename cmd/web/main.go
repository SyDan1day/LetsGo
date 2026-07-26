package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// 实际上
	// fileServer := http.FileServer(http.Dir("./ui"))
	// mux.Handle("/static/", fileServer)也可以，但是路由注册得太宽，不太好
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//其实写成mux.Handle("/", http.HandlerFunc(home))也一样
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on port:4000")
	//ListenAndServe的第二个参数是Handler,mux也实现了ServeHTTP方法，因此也可以视作特殊的http.Handler
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
