package test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestSuccessfulStore(t *testing.T) {
	data := url.Values{
		"FullUrl":      {"https://ya.ru"},
		"ExpiresAfter": {"0"},
	}

	app := initTest(t)
	go app.Start()
	response, err := http.PostForm("http://127.0.0.1:8000/api/store", data)

	if err != nil {
		t.Errorf("Error during store: %s", err.Error())
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error during response parsing: %s", err.Error())
		return
	}

	t.Errorf("Desu: %s", string(body))
}
