package main

import (
	"fmt"
	"strings"
)

func func01() {
	arr2 := []int{2, 3, 5, 7, 11, 13} // 2, 3, 5, 7, 11, 13
	printSlice(arr2)

	// Slice the slice to give it zero length.
	arr2 = arr2[:0] // ""
	printSlice(arr2)

	// Extend its length.
	arr2 = arr2[:4] // 2, 3, 5, 7 （値、復活すんの？！）
	printSlice(arr2)

	// Drop its first two values.
	arr2 = arr2[2:] // 5, 7
	printSlice(arr2)
}

func printSlice(arr []int) {
	// 長さはlen()、スライスのために確保しているメモリ領域の取得にはcap()を使う
	/*
		↓ func01()の実行結果
		len=6 cap=6 [2 3 5 7 11 13]
		len=0 cap=6 []
		len=4 cap=6 [2 3 5 7]
		len=2 cap=4 [5 7]
	*/
	fmt.Printf("len=%d cap=%d %v\n", len(arr), cap(arr), arr)
}

func main() {
	arr1 := []int{2, 3, 5, 7, 11, 13}

	arr1 = arr1[1:4] // 3, 5, 7
	fmt.Println(arr1)

	arr1 = arr1[:2] // 3, 5
	fmt.Println(arr1)

	arr1 = arr1[1:] // 5
	fmt.Println(arr1)

	func01() // Slice length and capacity 詳細は関数の中身をチェック

	// メモ: スライスのゼロ値は 「nil」 です。 
	var arr2 []int
	fmt.Println(arr2, len(arr2), cap(arr2))
	if arr2 == nil {
		fmt.Println("nil!")
	}

	// メモ: スライスは、他のスライスを含む任意の型を含むことができます。 
	// Create a tic-tac-toe board.
	// ※ []stringはこの場合省略可能
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// メモ: スライスへ新しい要素を追加するには、Goの組み込みの append を使います
	// appendするとメモリ領域を新たに確保してスライスに値を追加します 
	var arr3 []int
	printSlice(arr3)

	// append works on nil slices.
	arr3 = append(arr3, 0)
	printSlice(arr3)

	// The slice grows as needed.
	arr3 = append(arr3, 1)
	printSlice(arr3)

	// We can add more than one element at a time.
	arr3 = append(arr3, 2, 3, 4)
	printSlice(arr3)
}
