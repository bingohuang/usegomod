package main

import (
	"fmt"

	"github.com/bingohuang/testgomod/v2"
)

func main() {
	hi, err := testgomod.Hello("bingohuang", "cn")
	if err != nil {
		panic(err)
	}
	fmt.Println(hi)
}
