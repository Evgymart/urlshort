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
	fullUrl := "http://127.0.0.1:8000/"
	var expiresAfter int64 = 0

	data := url.Values{
		"FullUrl":      {fullUrl},
		"ExpiresAfter": {fmt.Sprint(expiresAfter)},
	}

	app := initTest(t)
	go app.Start()
	response, err := http.PostForm("http://127.0.0.1:8000/api/store", data)

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

	t.Log("Store test successful")
}
