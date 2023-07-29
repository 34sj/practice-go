package main

import (
	"fmt"
)

// 構造体
type Vertex struct {
	X int
	Y int
}

/*
structリテラルは、フィールドの値を列挙することで新しいstructの初期値の割り当てを示しています。
Name: 構文を使って、フィールドの一部だけを列挙することができます(この方法でのフィールドの指定順序は関係ありません)。 例では X: 1 として X だけを初期化しています。
& を頭に付けると、新しく割り当てられたstructへのポインタを戻します。 
*/
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	// こんな感じにも定義できるよ
	v := Vertex{1, 2}
	// 構造体のフィールドにはこうやってアクセスするよ
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3)
}