package api

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
	httpSend(&platform)
}