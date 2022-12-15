package db

import "math/rand"

func ReadFullUrl(db *Database, code string) (string, error) {
	val, err := db.Client.Get(Ctx, code).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func generateRandomUrl() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const size = 16
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(bytes)
}
