package main

import (
	"fmt"
	"math"
)

// インターフェイスメソッドの定義
// type インターフェイス名 interface {} の中にまとめてインターフェイスメソッドを定義
// インターフェイスを定義するときに"implements"キーワードは必要ありません -> Goでは登場しない
type Abser interface {
	Abs() float64
}

type MyFloat float64

// ①
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// ②
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	f1 := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f1 // インターフェイス「a」のレシーバに入れている値がfloat64なので、①のメソッドが呼ばれる
	fmt.Println(a.Abs())

	a = &v // インターフェイス「a」のレシーバに入れている値がstructへのポインタなので、②のメソッドが呼ばれる
	fmt.Println(a.Abs())

	// ゼロ個のメソッドを指定されたインターフェース型は、 空のインターフェース と呼ばれます
	// JavaでいうところのObjectみたいなもんらしい　万能
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)

	// 型アサーション
	s1 := i.(string)
	fmt.Println(s1)
	s1, ok := i.(string)
	fmt.Println(s1, ok)
	f2, ok := i.(float64)
	fmt.Println(f2, ok)
	// f2 = i.(float64) // panic なぜ？ → 「ok」にあたる変数が用意されていないため
	// fmt.Println(f2)

	// 型スイッチ
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}