package db

import "math/rand"

func ReadFullUrl(db *Database, shortUrl string) (string, error) {
	val, err := db.Client.Get(Ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func WriteFullUrl(db *Database, fullUrl string) (string, error) {
	shortUrl := generateRandomUrl()
	err := db.Client.Set(Ctx, shortUrl, fullUrl, 0).Err()
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func generateRandomUrl() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const size = 16
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(bytes)
}
