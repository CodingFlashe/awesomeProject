package main

func main() {
	////1、创建GET请求
	//resp, err := http.Get("https://www.baidu.com")
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//closer := resp.Body
	//bytes, err := ioutil.ReadAll(closer)
	//fmt.Println(string(bytes))
	//
	////2、创建POST请求
	//url := "https://www.shirdon.com/comment/add"
	//body := "{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}"
	//response, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//b, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(b))

}
