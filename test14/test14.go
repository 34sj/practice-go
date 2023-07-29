package main

import (
	"fmt"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// メモ: for ループに利用する range は、スライスや、マップ( map )をひとつずつ反復処理するために使います。 
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// メモ: インデックスや値は、 " _ "(アンダーバー) へ代入することで捨てることができます。 
	// メモ: もしインデックスだけが必要なのであれば、2つ目の値を省略します。 
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}