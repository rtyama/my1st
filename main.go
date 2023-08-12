package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		ret := rand.Intn(100)
		if hoge(ret) {
			fmt.Println("My favorite number is", ret)
			break
		}
		fmt.Println("失敗またやります")
	}
}

func hoge(r int) bool {
	if r > 50 {
		return true
	}
	return false
}
