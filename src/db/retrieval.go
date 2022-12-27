package db

func WriteFullUrl(db *Database, fullUrl string) (string, error) {
	code := generateRandomUrl()
	err := db.Client.Set(Ctx, code, fullUrl, 0).Err()
	if err != nil {
		return "", err
	}

	return code, nil
}
