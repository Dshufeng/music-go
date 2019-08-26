package api

import (
	"fmt"
	"github.com/tidwall/gjson"
	"encoding/json"
)

func (*NetsPlatform) Search(keyword string) {
	randNum := GenerateRangeNum(1884815360,1884890111)
	platform := Platform{
		Name:"nets",
		Method:"POST",
		Url:"http://music.163.com/api/cloudsearch/pc",
		Body:body{
			"format":"json",
			"s":keyword,
			"type":1,
			"limit":10,
			"total":"true",
			"offset":0,
		},
		Encode:"",
		Format:"",
		Headers:headers{
			"Referer"   : "https://music.163.com/",
			"Cookie"    : "appver=1.5.9; os=osx; __remember_me=true; osver=%E7%89%88%E6%9C%AC%2010.13.5%EF%BC%88%E7%89%88%E5%8F%B7%2017F77%EF%BC%89;",
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko)",
			"X-Real-IP" : Long2ip(randNum).String(),
			"Accept"    : "*/*",
			"Accept-Language" : "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
			"Connection"      : "keep-alive",
			"Content-Type"    : "application/x-www-form-urlencoded",
		},
	}
	resp := httpSend(&platform)
	if code := gjson.Get(resp,"code").Int(); code != 200{
		msg := gjson.Get(resp,"msg").String()
		fmt.Printf(`api request err, msg: "%v",code: %d`,msg,code)
		return
	}

	songs := gjson.Get(resp,"result.songs")
	var songsArr []Song
	for _,v := range songs.Array(){
		var artist []string
		for _,ar := range gjson.Get(v.String(),"ar.#.name").Array(){
			artist = append(artist,ar.String())
		}
		songsArr = append(songsArr,Song{
			Id: gjson.Get(v.String(),"id").Index,
			Name: gjson.Get(v.String(),"name").String(),
			Artist: artist,
			Album: gjson.Get(v.String(),"al.name").String(),
			PicId: gjson.Get(v.String(),"al.pic_str").String(),
			UrlId:gjson.Get(v.String(),"id").Index,
			LyricId:gjson.Get(v.String(),"id").Index,
			Source:"nets",
		})
	}
	j,_ := json.Marshal(songsArr)
	fmt.Println(string(j))
}
func (*NetsPlatform)searchFormat(resp string)  {
}