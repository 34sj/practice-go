package main

import "fmt"

func main() {

	// 変数定義は var hoge [型] = fuga のような形式でやる（型は省略可能）
	var i, j int = 1, 2
	// 「:=」でvarを省略できる
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
