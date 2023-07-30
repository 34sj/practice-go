package main

import (
	"fmt"
	"time"
)

// Goのプログラムは、エラーの状態を error 値で表現します。

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// nil の error は成功したことを示し、 nilではない error は失敗したことを示します。 
	if err := run(); err != nil {
		fmt.Println(err) // error 表示
	}
}
