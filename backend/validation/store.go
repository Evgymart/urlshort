package validation

import (
	"net/http"
	"strconv"
)

func ValidateStore(request *http.Request) (string, int64, error) {
	request.ParseForm()
	fullUrl := request.FormValue("FullUrl")
	expiresAfter := request.FormValue("ExpiresAfter")
	expiresAfterInt, err := strconv.ParseInt(expiresAfter, 0, 64)
	return fullUrl, expiresAfterInt, err
}
