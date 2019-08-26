package api

import (
	"fmt"
	"github.com/tidwall/gjson"
	"encoding/json"
)

type song struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Artist []string `json:"artist"`
	Album string `json:"album"`
	PicId string `json:"pic_id"`
	UrlId int `json:"url_id"`
	LyricId int `json:"lyric_id"`
	Source string `json:"source"`
}

func (*NetsPlatform) Search(keyword string) {
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
	}
	resp := httpSend(&platform)
	if code := gjson.Get(resp,"code").Int(); code != 200{
		msg := gjson.Get(resp,"msg").String()
		fmt.Printf(`api request err, msg: "%v",code: %d`,msg,code)
		return
	}

	songs := gjson.Get(resp,"result.songs")

	var songsArr []song
	for _,v := range songs.Array(){
		var artist []string
		for _,ar := range gjson.Get(v.String(),"ar.#.name").Array(){
			artist = append(artist,ar.String())
		}
		songsArr = append(songsArr,song{
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
func (*NetsPlatform)searchFormat()  {

}