package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	name := req.FormValue("name")
	io.WriteString(res,
		"<doctype html><html><head><title>Hello!</title><style>.container{text-align:center; with: 960px;	margin-top: 200px}</style></head><body><div class=container><h1>Request parameter is <em>Name</em>=<span style=color:red>"+name+"</span></h1></div></body></html>")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
