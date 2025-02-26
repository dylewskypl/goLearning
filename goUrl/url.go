package main

type Url struct {
	OrginalUrl string
	ShortUrl   string
}

type JSONDb struct {
	Urls []Url
}
