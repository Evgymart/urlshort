package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shorturl/backend/config"
	"shorturl/backend/core"
	"testing"
)

func TestErrorStore(t *testing.T) {
	fullUrl := "notanurl"
	var expiresAfter int64 = 0
	data := url.Values{
		"FullUrl":      {fullUrl},
		"ExpiresAfter": {fmt.Sprint(expiresAfter)},
	}

	app := initTest(t)
	go app.Start()
	response, err := http.PostForm(buildUrlPath(config.GetServerAddr(), "api/store"), data)
	if err != nil {
		t.Fatalf("Error during store: %s", err.Error())
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error during response parsing: %s", err.Error())
	}

	var errorMessage core.ErrorMessage
	err = json.Unmarshal(responseBody, &errorMessage)
	if err != nil {
		t.Fatalf("Error during unmarshal: %s", err.Error())
	}

	if errorMessage.Error != fmt.Sprintf("Invalid url: %s", fullUrl) {
		t.Fatalf("Failed to recieve a proper error message, recieved: %s", errorMessage.Error)
	}
}
