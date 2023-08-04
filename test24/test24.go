package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 送り手は、これ以上の送信する値がないことを示すため、チャネルを close できます。
	// closeしたあとはchannelに入れれないので注意
	close(c)

	// 受信の式に2つ目のパラメータを割り当てることで、そのチャネルがcloseされているかどうかを確認できます
	// 例: v, ok := <-ch 
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
