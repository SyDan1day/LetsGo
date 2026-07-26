package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()
	// 实际上
	// fileServer := http.FileServer(http.Dir("./ui"))
	// mux.Handle("/static/", fileServer)也可以，但是路由注册得太宽，不太好
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//其实写成mux.Handle("/", http.HandlerFunc(home))也一样
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	infoLog.Printf("Starting server on port%s", *addr)
	//ListenAndServe的第二个参数是Handler,mux也实现了ServeHTTP方法，因此也可以视作特殊的http.Handler
	//http.ListenAndServe实际上是使用默认参数创建了一个http.Server，并调用它的ListenAndServe方法。
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on port%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
