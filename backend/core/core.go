package core

import (
	"shorturl/backend/db"
)

type UrlData struct {
	FullUrl      string
	ShortUrlCode string
	ExpiresAfter int64
}

var (
	coreRedis *db.Database = nil
)

func InitCore(redis *db.Database) {
	coreRedis = redis
}

func BuildUrlData(fullUrl string, expiresAfter int64) (*UrlData, error) {
	shortUrlCode, err := db.WriteFullUrl(coreRedis, fullUrl)
	if err != nil {
		return nil, err
	}

	urlData := UrlData{fullUrl, shortUrlCode, expiresAfter}
	return &urlData, nil
}

func GetRedirectUrl(code string) (string, error) {
	fullUrl, err := db.ReadFullUrl(coreRedis, code)
	if err != nil {
		return "", err
	}

	return fullUrl, nil
}
