package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2.0))
	fmt.Println(math.Sqrt(2.0))
}

/*作为练习函数和循环的简单途径，用牛顿法实现开方函数。

在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值：


为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。*/