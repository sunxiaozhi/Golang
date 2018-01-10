package main

import (
	"fmt"
	"time"
)

func arraySum(array []int, resultChan chan int) {
	fmt.Println("I am in goroutine arraySum")
	sum := 0
	for _, value := range array {
		sum += value
	}
	resultChan <- sum;
}

func arrayQuadrature(array []int, resultChan chan int) {
	fmt.Println("I am in goroutine arrayQuadrature")
	sum := 1
	for _, value := range array {
		sum *= value
	}
	resultChan <- sum;
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	fmt.Println("ok let's go")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	go arraySum(array, resultChan)
	go arrayQuadrature(array, resultChan)
	fmt.Println("I am main i am running 1+1")
	a := 1 + 1;
	time.Sleep(1000) //假设这里在链接数据库
	fmt.Println("I am main i am running the result is", a)
	sum1 := <-resultChan
	sum2 := <-resultChan
	fmt.Println("I am main and go routine finished")
	fmt.Println("Result is :", sum1, sum2)
	for i := 0; i < 1; i = 0 {
	}
}

/*func longTimeToHandle()  {
	fmt.Println("goroutine finished")
}

func main()  {
	go longTimeToHandle();
	time.Sleep(1 * time.Second)
	fmt.Println("main finished")
}*/
