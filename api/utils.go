package api

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func httpSend(p *Platform)  {
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
