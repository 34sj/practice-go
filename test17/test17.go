package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

// メモ（重要）: Goには、クラス( class )のしくみはありませんが、型にメソッド( method )を定義できます。

// メソッドは、特別なレシーバ( receiver )引数を関数に取ります。 ここでいうレシーバとは「(v Vertex)」の部分
// 返り値の型はメソッド名()の右に書く 例: func (v Vertex) Abs1() float64
func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 俺たちに馴染みのあるメソッド定義　機能はAbs1()と同様
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
例で挙げたstructの型だけではなく、任意の型(type)にもメソッドを宣言できます。
例は、 Abs メソッドを持つ、数値型の MyFloat 型です。
レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。 他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できません （組み込みの int などの型も同様です）。 
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// レシーバにポインタを持つメソッド
func (v *Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	// 計算結果はポインタのレシーバから参照できるため、値をreturnしなくても呼び出し元で計算結果を取り出して使える
}

// Scale1()メソッドのレシーバがポインタじゃなかったら、値をreturnしてあげないと計算結果を呼び出し元から参照できない（計算結果が無意味になる）
func (v Vertex) Scale1Wwwww(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	// この場合、レシーバを廃止して計算結果をreturnしてあげるか、引数にVertexのポインタを受け取れるような形にしないと、計算結果が無意味になる
}

// 俺たちに馴染みのあるメソッド定義　機能はScale1()と同様
func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs1()) // レシーバを持ったメソッドの呼び出し　こんな呼び出し方もできるよ　特殊！　実行結果: 5

	v = Vertex{3, 4}
	fmt.Println(Abs2(v)) // 俺たちに馴染みのあるメソッド呼び出し　実行結果: 5

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // レシーバを持ったメソッドの呼び出し 実行結果: 1.4142135623730951

	v = Vertex{3, 4}
	v.Scale1(10)  // レシーバにポインタを持ったメソッドの呼び出し　引数を使ってメソッド内で計算を行う　
	fmt.Println(v.Abs1())  // 実行結果: 50

	v = Vertex{3, 4}
	v.Scale1Wwwww(10)  // Scale1()メソッドのレシーバがポインタじゃなかったら、値をreturnしてあげないと計算結果を呼び出し元から参照できない（計算結果が無意味になる）　
	fmt.Println(v.Abs1())  // 実行結果: 5

	v = Vertex{3, 4}
	Scale2(&v, 10) // 俺たちに馴染みのあるメソッド呼び出し
	fmt.Println(v.Abs1())  // 実行結果: 5
}