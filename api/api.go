package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

type body map[string]interface{}

type Platform struct {
	Name string
	Method string
	Url string
	Body body
	Encode string
	Format string
}

type platformApi interface {
	search()
}

func main()  {
	selectPlatform("qq").search()
}

func (p *Platform)search() {
	fmt.Println("serach by keyword...")
	curl(p)
}

func selectPlatform(p string)  *Platform{
	var platform Platform
	switch p {
	case "nets":
		platform = Platform{
			Name:"nets",
			Method:"POST",
			Url:"http://music.163.com/api/cloudsearch/pc",
			Body:body{
				"format":"json",
				"s":"中国话",
				"type":1,
				"limit":10,
				"total":"true",
				"offset":0,
			},
			Encode:"",
			Format:"",
		}
		break
	case "qq":
		platform = Platform{
			Name:"qq",
			Method:"GET",
			Url:"https://c.y.qq.com/soso/fcgi-bin/client_search_cp",
			Body:body{
				"format":"json",
				"p":1,
				"n":10,
				"w":"中国话",
				"aggr":1,
				"lossless":1,
				"cr":1,
				"new_json":1,
			},
			Encode:"",
			Format:"",
		}
		break
	}
	return &platform
}
func curl(p *Platform)  {
		urlParam := buildUrl(p.Body)
		var err error
		var resp *http.Response
		if p.Method == "POST"{
			resp,err = http.Post(p.Url,"application/x-www-form-urlencoded",strings.NewReader(urlParam))
			defer resp.Body.Close()
		}else {
			resp,err = http.Get(p.Url + "?" + urlParam)
		}

		if err != nil{
			fmt.Println("http err: ",err)
		}
		body,err:= ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Println("read err: ",err)
		}
		fmt.Println(string(body))
}

func buildUrl(requestBody body)  string{
	i := 0
	var str string
	for k,v := range requestBody{
		if i == 0 {
			str = fmt.Sprintf("%v=%v",k,v)
		}else{
			str += fmt.Sprintf("&%v=%v",k,v)
		}
		i++
	}
	return str
}


