package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getParseParam(para string) string {
	return url.PathEscape(para)
}
func CopeHandle(method, urlVal, data string) {
	client := &http.Client{}
	var req *http.Request
	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}
	cookie := &http.Cookie{
		Name:     "X-Xsrftoken",
		Value:    "aaab6d695bbdbcdkhjiweh",
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	req.Header.Add("X-Xsrftoken", "aaab6d695bbdbcdkhjiweh")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(b))
}
func main() {
	CopeHandle("GET", "https://www.baidu.com", "")
}
