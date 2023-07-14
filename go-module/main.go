package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"go-module/c4f"
)

func main() {
    ac := accounting.Accounting{Symbol: "$", Precision: 2}
    fmt.Println(ac.FormatMoney(123456789.213123)) 

	c4f.Println("this is red color")


	
// import go-module
	// sudo go get github.com/leekchan/accounting
	// sudo go get github.com/fatih/color
	// sudo go mod init go-module
	// go build
	// go mod vendor
}