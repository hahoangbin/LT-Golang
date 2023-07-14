package main

import (
	"fmt"

	// "github.com/leekchan/accounting"

	// "go-module/c4f"
)

// func main() {
//     ac := accounting.Accounting{Symbol: "$", Precision: 2}
//     fmt.Println(ac.FormatMoney(123456789.213123))

// 	c4f.Println("this is red color")

// // import go-module
// 	// sudo go get github.com/leekchan/accounting
// 	// sudo go get github.com/fatih/color
// 	// sudo go mod init go-module
// 	// go build
// 	// go mod vendor
// }

// Channels
func main() {
	// 1. Unbuffered channel
	// 2. Buffered channel
	// 3. Select
	// 4. Close channel

	// 1. Unbuffered channel
	// ch := make(chan int)

	// go func() {
	// 	ch <- 100
	// }()

	// fmt.Println(<-ch)
	// fmt.Println("done")

	// 2. Buffered channel
	// ch := make(chan int, 2)

	// ch <- 10
	// ch <- 201
	// fmt.Println("find")

	// close(ch)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println("done")

	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
			if i > 5 {
				close(queue)
				// không thể gửi giá trị vào một channel đã close
			}
		}

		done <- true
	}()

	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			return
		}
	}
}