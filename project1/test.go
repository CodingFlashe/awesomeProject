package main

import (
	"encoding/json"
	"fmt"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
func test() {
	//使用defer+recover 来捕获处理异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  //nil是err的零值，不等于nil说明捕获到错误
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}
type Stu struct {
}

func (s Stu) test01() {
	panic("implement me")
}

func (s Stu) test02() {
	panic("implement me")
}

func (s Stu) test03() {
	panic("implement me")
}

type Student struct {
	Name  string //`json:"n"`
	Grade int    //`json:"s"`
}

func StructJson() {
	stu := Student{
		Name:  "jack",
		Grade: 95,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func UnMarshal() {
	str := "{\"Name\":\"jack\",\"Grade\":95}"
	var stu Student
	err := json.Unmarshal([]byte(str), &stu) //因为要在函数内部改变函数外部的变量，因此要用引用传递
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}

func UnMarshalMap() {
	str := "{\"Name\":\"jack\",\"Grade\":95}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a) //因为要在函数内部改变函数外部的变量，因此要用引用传递
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func main() {

}
