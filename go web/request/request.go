package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func request(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request解析")
	fmt.Println("method", r.Method)
	fmt.Println("RequestURI:", r.RequestURI)
	fmt.Println("URL_Path", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	fmt.Println("proto", r.Proto)
	fmt.Println("protoMajor", r.ProtoMajor)
	fmt.Println("protoMinor", r.ProtoMinor)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("header key:" + k + " value:" + vv)
		}
	}
	isMultipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}
	if isMultipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("解析方式：ParseMultipartForm")
	} else {
		r.ParseForm()
		fmt.Println("解析方式：ParseForm")
	}
	fmt.Println("ContentLength", r.ContentLength)
	fmt.Println("Close", r.Close)
	fmt.Println("host", r.Host)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Fprintf(w, "hello,let's go!")
}
func main() {
	http.HandleFunc("/hello", request)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
