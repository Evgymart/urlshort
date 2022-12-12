package core

type UrlData struct {
	FullUrl      string
	ShortUrl     string
	ExpiresAfter int64
}

func BuildUrlData(fullUrl string, expiresAfter int64) *UrlData {
	urlData := UrlData{fullUrl, "blank", expiresAfter}
	return &urlData
}
