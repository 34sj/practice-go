package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") // defer ロードするときは上から順番なのでこのタイミングでロードされるが、returnするときに解決される

	fmt.Println("hello")
	
	// 出力結果: "hello world"
}