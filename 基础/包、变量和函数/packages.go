package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
}

/*
包
每个 Go 程序都是由包组成的。

程序运行的入口是包 `main`。

这个程序使用并导入了包 "fmt" 和 `"math/rand"`。

按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始。

注意： 这个程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （为了得到不同的数字，需要生成不同的种子数，参阅 rand.Seed。）*/
