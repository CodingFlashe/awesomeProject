package main

import "fmt"

func twoValueSum[T int | float64 | string](a T, b T) T {
	return a + b
}
func main() {
	ret := twoValueSum[int](100, 200)
	fmt.Println(ret)

	ret1 := twoValueSum[string]("hello ", "world")
	fmt.Println(ret1)
	//stu := funct.student{
	//	Name: "tom",
	//	Age:  20,
	//}
	//当student结构体首字母小写，我们可以通过工厂模式来解决
	//var stu = funct.NewStudent("tom", 88) //返回的student实例的指针
	//fmt.Println(*stu)
	//fmt.Println(stu.GetName())
}
