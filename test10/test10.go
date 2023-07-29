package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

/*
メモ

Goはポインタを扱います。 ポインタは値のメモリアドレスを指します。

変数 T のポインタは、 *T 型で、ゼロ値は nil です。

var p *int

& オペレータは、そのオペランド( operand )へのポインタを引き出します。

i := 42
p = &i

* オペレータは、ポインタの指す先の変数を示します。

fmt.Println(*p) // ポインタpを通してiから値を読みだす
*p = 21         // ポインタpを通してiへ値を代入する

これは "dereferencing" または "indirecting" としてよく知られています。

なお、C言語とは異なり、ポインタ演算はありません。
*/