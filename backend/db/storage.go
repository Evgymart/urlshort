package db

import "math/rand"

func ReadFullUrl(db *Database, code string) (string, error) {
	val, err := db.Client.Get(Ctx, code).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func WriteFullUrl(db *Database, fullUrl string) (string, error) {
	code := generateRandomUrl()
	err := db.Client.Set(Ctx, code, fullUrl, 0).Err()
	if err != nil {
		return "", err
	}

	return code, nil
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
