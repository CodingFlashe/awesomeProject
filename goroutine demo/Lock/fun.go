package main

import (
	"fmt"
	"sync"
	"time"
)

func Add(a int, b int) (int, int) {
	return a + b, a - b
}
func Multiply(a, b int) int {
	return a * b
}

type student struct {
	Name string
	Age  int
}

//student结构体首字母小写，只能在funct包使用
//通过工厂模式解决

func NewStudent(n string, s int) *student {
	return &student{
		Name: n,
		Age:  s,
	}
}
func (stu *student) GetName() (string, int) {
	return stu.Name, stu.Age
}

type data struct {
	sync.Mutex
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()

}
