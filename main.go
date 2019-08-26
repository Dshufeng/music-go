package main

import (
	"music/api"
	"fmt"
	_"encoding/json"
)
var (
	selectPlatform string
	searchKeyword string
	platforms = []string{"nets","qq"}
)

func main()  {
	fmt.Printf("select platform (%v default nets):\n",platforms)
	fmt.Scanln(&selectPlatform)
	if selectPlatform == "" {
		selectPlatform = "nets"
	}
	fmt.Println("you select platform: ",selectPlatform)

	for  {
		fmt.Println("input keyword: ")
		fmt.Scanln(&searchKeyword)
		if searchKeyword != ""{
			break
		}
	}
	fmt.Println("you keyword: ",searchKeyword)

	search(selectPlatform,searchKeyword)
}

func search(selectPlatform string,searchKeyword string)  {
	switch selectPlatform {
	case "nets":
		api.Nets.Search(searchKeyword)
		break
	case "qq":
		api.Qq.Search(searchKeyword)
		break
	default:
		fmt.Println("err")
	}
}