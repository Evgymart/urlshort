package core

import (
	"net/http"
)

func isUrlAlive(url string) bool {
	result, err := http.Get(url)
	if err != nil {
		return false
	}
	return result.StatusCode == http.StatusOK
}
