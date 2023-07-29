package main

import (
	"fmt"
)

func main() {
	// メモ: Golangは配列の長さを変えることはできない（固定長）
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//メモ: 「スライス」は可変長　 実際には、スライスは配列よりもより一般的です。 スライスは配列のような挙動をする参照型です。
	var s []int = primes[1:4] // 1番地~4番地の前（この場合、3、5、7を取り出す）ここで取り出した値はスライス
	fmt.Println(s)

	// これは配列
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2] // 取り出される値は[John, Paul]（スライス）
	d := names[1:3] // 取り出される値は[Paul, George]（スライス）
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	// これはstructなslice
	strct := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(strct)
}