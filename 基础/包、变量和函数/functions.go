package main

import "fmt"

func add_one(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add_one(42, 13))
}

/*函数可以没有参数或接受多个参数。

在这个例子中，`add` 接受两个 int 类型的参数。

注意类型在变量名 _之后_。*/
