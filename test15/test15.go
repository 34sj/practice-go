package main

import (
	"fmt"
)

func main() {
	// type（structの定義）はfunc内にも書けるみたい
	type Vertex struct {
		Lat, Long float64
	}

	// メモ: map はキーと値とを関連付けます(マップします)。 
	var m1 map[string]Vertex

	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])

	m1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m1)

	// 以下mapの操作とか
	m2 := make(map[string]int)

	// 追加
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	// 上書き
	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	// 削除
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	// 取り出し
	v, ok := m2["Answer"]
	// もし、 m2 に key があれば、変数 ok は true となり、存在しなければ、 ok は false となります。 
	fmt.Println("The value:", v, "Present?", ok)
}