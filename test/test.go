package main

import "fmt"

func main() {
	a := []int{7, 5, 5, 89, 1}

	c := make(chan int)
	c1 := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c1)

	x, y := <-c, <-c1

	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	//将计算出的值发送到channel
	c <- total
}
