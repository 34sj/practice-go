package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Scale2(), Abs2()についてのメモ

ポインタレシーバを使う2つの理由があります。
ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。 （要約するとパフォーマンスの問題）
*/
func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale1(2)
	ScaleFunc(&v, 10)
	// ScaleFunc(v, 10) // ① メソッド側の引数がポインタなのに渡している変数に&なし（変数） -> コンパイルエラーになるので注意

	p := &Vertex{4, 3}
	p.Scale1(3) // レシーバを定義していなくてもpをVertexへのポインタにしていれば使える
	ScaleFunc(p, 8)
	// ScaleFunc(&p, 8) // ② メソッド側の引数がポインタなのに渡しているポインタに&がついている（ポインタのポインタ？） -> コンパイルエラーになるので注意

	fmt.Println(v, p)

	v = Vertex{3, 4}
	fmt.Println(v.Abs1())
	fmt.Println(AbsFunc(v))
	// fmt.Println(AbsFunc(&v)) // ③ メソッド側の引数が変数なのに渡している変数に&がついている（ポインタ） -> コンパイルエラーになるので注意

	p = &Vertex{4, 3}
	fmt.Println(p.Abs1())
	fmt.Println(AbsFunc(*p)) // ポインタが参照している値を参照するには名前の先頭に「*」をつける
	// fmt.Println(AbsFunc(p)) // ④ メソッド側の引数が変数なのにポインタを渡している -> コンパイルエラーになるので注意

	/*
	ポインタレシーバを使う2つの理由があります。
	ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
	ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。 （要約するとパフォーマンスの問題）
	*/
	vp := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vp, vp.Abs2())
	v.Scale1(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vp, vp.Abs2())
}