package validation

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func ValidateStore(request *http.Request) (string, int64, error) {
	request.ParseForm()
	fullUrl := request.FormValue("FullUrl")
	_, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return "", 0, errors.New("Invalid url: " + fullUrl)
	}

	expiresAfter := request.FormValue("ExpiresAfter")
	expiresAfterInt, err := strconv.ParseInt(expiresAfter, 0, 64)
	if err != nil {
		return "", 0, err
	}

	if expiresAfterInt < 0 {
		return "", 0, errors.New("Expires after must be positive")
	}

	return fullUrl, expiresAfterInt, err
}
