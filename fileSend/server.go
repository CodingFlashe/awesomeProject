package fileSend

//package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	http.HandleFunc("/", WithMiddleWare(GetFile, Logger))
	if e := http.ListenAndServe("0.0.0.0:8080", nil); e != nil {
		panic(e)
	}
}

func WithMiddleWare(handler http.HandlerFunc, mids ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, mid := range mids {
			mid(w, r)
		}
		handler(w, r)
	}
}

func Logger(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[JIN]")
	log.Printf("\t\t%s\t\t%s\n", r.UserAgent(), r.URL.EscapedPath())
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create("receive.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	_, e = io.Copy(f, r.Body)
	if e != nil {
		panic(e)
	}
	resp := R{
		Code: 200,
		Msg:  "成功接收文件",
	}
	buf, e := json.Marshal(resp)
	if e != nil {
		panic(e)
	}
	w.Write(buf)
}

func checkErr(e error, msg string) {
	if e != nil {
		log.Printf("error! %v\t[msg]%s", e, msg)
	}
}
