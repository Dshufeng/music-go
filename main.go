package main

import (
	"music/api"
	"fmt"
	_"encoding/json"
	_"github.com/olekukonko/tablewriter"
	_"os"
)
var (
	selectPlatform string
	searchKeyword string
	platforms = []string{"nets","qq"}
)

func main()  {


	fmt.Printf("请选择平台:%v (default nets):\n",platforms)
	fmt.Scanln(&selectPlatform)
	if selectPlatform == "" {
		selectPlatform = "nets"
	}

	for  {
		fmt.Println("请输入关键字: ")
		fmt.Scanln(&searchKeyword)
		if searchKeyword != ""{
			break
		}
	}
	fmt.Printf("正在搜索: %v, 请耐心等待...\n",searchKeyword)
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