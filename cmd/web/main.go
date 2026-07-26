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
	//依赖注入
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	infoLog.Printf("Starting server on port%s", *addr)
	//ListenAndServe的第二个参数是Handler,mux也实现了ServeHTTP方法，因此也可以视作特殊的http.Handler
	//http.ListenAndServe实际上是使用默认参数创建了一个http.Server，并调用它的ListenAndServe方法。
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on port%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
