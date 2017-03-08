package main

import (
	_ "contract/config"
	"contract/config"
	"fmt"
)

func main() {
	keys, err := config.GetKeywords()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(keys)
}
