package main

import (
	"fmt"
	"math"
)

func pow_two(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	fmt.Println(
		pow_two(3, 2, 10),
		pow_two(3, 3, 20),
	)
}

/*if 和 else
在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。*/