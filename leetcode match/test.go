package main

import "fmt"

//定义数组表示每次移动的步数,2表示乘2
var pathX = []int{-1, 1, 2}

func FindFriend(N int, K int) int {
	used := make(map[int]bool) //标记走过的点
	ans := 0
	que := make([]int, 0)
	que = append(que, N)
	used[N] = true
	for len(que) > 0 {
		length := len(que)
		for i := 0; i < length; i++ {
			dx := que[0]
			if dx == K {
				return ans
			}
			que = que[1:]
			for j := 0; j < 3; j++ {
				px := 0
				if j == 2 {
					px = dx * 2
				} else {
					px = dx + pathX[j]
				}
				if px < 0 || px >= 2*K {
					continue
				}
				if !used[px] { //下一步要到达的点没有走过
					que = append(que, px)
					used[px] = true
				}
			}
		}
		ans++
	}
	return -1 //表示无法到达
}

func main() {
	var N int
	var K int
	fmt.Scanln(&N)
	fmt.Scanln(&K)
	ans := FindFriend(N, K)
	if ans == -1 {
		fmt.Println("小青无法到达小码家")
	} else {
		fmt.Println("小青到小码家时间为:", ans)
	}
}
