package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shorturl/backend/core"
	"testing"
)

func TestSuccessfulStore(t *testing.T) {
	fullUrl := buildUrl(HttpAddrTest)
	var expiresAfter int64 = 0
	data := url.Values{
		"FullUrl":      {fullUrl},
		"ExpiresAfter": {fmt.Sprint(expiresAfter)},
	}

	app := initTest(t)
	go app.Start()
	response, err := http.PostForm(buildUrlPath(HttpAddrTest, "api/store"), data)
	if err != nil {
		t.Fatalf("Error during store: %s", err.Error())
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error during response parsing: %s", err.Error())
	}

	var urlData core.UrlData
	err = json.Unmarshal(responseBody, &urlData)
	if err != nil {
		t.Fatalf("Error during unmarshal: %s", err.Error())
	}

	if urlData.FullUrl != fullUrl || urlData.ExpiresAfter != expiresAfter {
		t.Fatalf("Bad response data: %s", responseBody)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err = client.Get(buildUrlPath(HttpAddrTest, urlData.ShortUrlCode))
	if err != nil {
		t.Fatalf("Error during check redirect: %s", err.Error())
	}

	if response.StatusCode != http.StatusFound || response.Header["Location"][0] != fullUrl {
		t.Fatalf("Redirect check failed, code: %d", response.StatusCode)
	}
}
