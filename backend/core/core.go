package core

import (
	"shorturl/backend/db"
)

type UrlData struct {
	FullUrl      string
	ShortUrl     string
	ExpiresAfter int64
}

var (
	coreRedis *db.Database = nil
)

func InitCore(redis *db.Database) {
	coreRedis = redis
}

func BuildUrlData(fullUrl string, expiresAfter int64) (*UrlData, error) {
	shortUrl, err := db.WriteFullUrl(coreRedis, fullUrl)
	if err != nil {
		return nil, err
	}

	urlData := UrlData{fullUrl, shortUrl, expiresAfter}
	return &urlData, nil
}
