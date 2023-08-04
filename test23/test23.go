package main

import (
	"fmt"
)

/*
channel を使うと gorutine間でデータのやり取りができる
定義: [チャンネル名] := make(chan [型])
定義（引数）: [チャンネル名] chan [型]
チャンネルにデータを突っ込む: [チャンネル名]<-[データ]
チャンネルからデータを取り出す: [データ]<-[チャンネル名]
WaitGroupを使ってgorutineの制御をするよりもこっちの方が一般的らしい（値渡せるし簡単）
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// channelの挙動はバッファ（つまりキュー）　最初に入ったデータが最初に取り出される
	ch := make(chan int, 2) // 第二引数に整数を指定すると、バッファのサイズを指定できる
	ch <- 1
	ch <- 2
	// ch <- 3 // 3つ目を入れようとするとサイズが足りないのでエラーが出る
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
