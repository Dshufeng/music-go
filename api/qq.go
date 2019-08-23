package api

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
	}
	httpSend(&platform)
}