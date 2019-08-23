package api

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




