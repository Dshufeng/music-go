package api

type body map[string]interface{}
type headers map[string]string
type Platform struct {
	Name string
	Method string
	Url string
	Body body
	Encode string
	Format string
	Headers headers
}
type Song struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Artist []string `json:"artist"`
	Album string `json:"album"`
	PicId string `json:"pic_id"`
	UrlId string `json:"url_id"`
	LyricId string `json:"lyric_id"`
	Source string `json:"source"`
}

type platformApi interface {
	Search()
}

type (
	NetsPlatform Platform
	QqPlatform Platform
)

var (
	Nets NetsPlatform
	Qq QqPlatform
)




