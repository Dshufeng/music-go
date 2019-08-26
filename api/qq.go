package api

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func (*QqPlatform) Search(keyword string) {
	platform := Platform{
		Name:"qq",
		Method:"GET",
		Url:"https://c.y.qq.com/soso/fcgi-bin/client_search_cp",
		Body:body{
			"format":"json",
			"p":1,
			"n":10,
			"w":keyword,
			"aggr":1,
			"lossless":1,
			"cr":1,
			"new_json":1,
		},
		Encode:"",
		Format:"",
		Headers:headers{
			"Referer"         : "http://y.qq.com",
			"Cookie"          : "pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1",
			"User-Agent"      : "QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)",
			"Accept"          : "*/*",
			"Accept-Language" : "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
			"Connection"      : "keep-alive",
			"Content-Type"    : "application/x-www-form-urlencoded",
		},
	}
	resp := httpSend(&platform)
	if code := gjson.Get(resp,"code").Int();code != 0{
		msg := gjson.Get(resp,"message").String()
		fmt.Printf(`api request err, msg: "%v",code: %d`,msg,code)
		return
	}

	songs := gjson.Get(resp,"data.song.list")
	var songsArr []Song
	for _,v := range songs.Array(){
		var artist []string
		for _,ar := range gjson.Get(v.String(),"singer.#.name").Array(){
			artist = append(artist,ar.String())
		}
		songsArr = append(songsArr,Song{
			Id: gjson.Get(v.String(),"mid").String(),
			Name: gjson.Get(v.String(),"name").String(),
			Artist: artist,
			Album: gjson.Get(v.String(),"album.name").String(),
			PicId: gjson.Get(v.String(),"album.album.mid").String(),
			UrlId:gjson.Get(v.String(),"mid").String(),
			LyricId:gjson.Get(v.String(),"mid").String(),
			Source:"qq",
		})
	}
	RenderTable(songsArr)
	return
}