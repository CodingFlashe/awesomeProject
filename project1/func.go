package main

import "fmt"

func sub(a, b int) int {
	return a - b
}

func TypeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) { //.(type)是固定写法，必须这样子写
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是%v\n", index, v)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是%v\n", index, v)
		case float64:
			fmt.Printf("第%v个参数是float64类型,值是%v\n", index, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型,值是%v\n", index, v)
		case string:
			fmt.Printf("第%v个参数是string类型,值是%v\n", index, v)
		default:
			fmt.Printf("第%v个参数类型不确定,值是%v\n", index, v)
		}
	}
}
