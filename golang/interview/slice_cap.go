package main

import "fmt"

//猎豹移动的面试题

func main(){
	sl1 := make([]int, 3,10)

	for i:=0;i<3;i++{
		sl1[i] = i
	}


	fmt.Printf("%p\n", sl1)

	ff(sl1)

	fmt.Printf("%p\n", sl1)
}


func ff(x []int){
	fmt.Printf("%p\n", x)

	x[0] = 1
	fmt.Printf("%p\n", x)

	//x 容量不足，append之后重新分片，x地址变化
	x = append(x, 4)

	fmt.Printf("%p\n", x)

	fmt.Println(len(x), cap(x))
}
