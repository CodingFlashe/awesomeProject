package fileSend

import (
	"log"
	"net/http"
	"os"
)

// @File: fileWebDemo-Client.go
// @Author: Jason
// @Date: 2022/11/2

func main() {
	f, e := os.Open("D:\\Golang\\GoCode\\src\\awesomeProject\\goroutine demo\\channel\\main.go")
	failOnErr(e, "文件打开失败!")
	req, e := http.NewRequest(http.MethodPut, "http://10.151.12.45:8080/", f)
	failOnErr(e, "request构建失败!")
	client := http.Client{}
	_, e = client.Do(req)
	failOnErr(e, "请求失败！")
	log.Println()
}

func failOnErr(e error, msg string) {
	if e != nil {
		log.Panicln(e.Error(), msg)
	}
}
