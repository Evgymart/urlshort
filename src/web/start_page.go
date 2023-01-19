package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"shorturl/backend/config"
	"shorturl/backend/core"
)

type StartPageTemplate struct {
	ShowMessage  bool
	ErrorMessage string
	ShortUrl     string
}

func startPage(writer http.ResponseWriter, request *http.Request) {
	root := config.GetPath().AppRoot
	template := template.Must(template.ParseFiles(root + "/web/templates/start_page.html"))

	isPost := request.Method == http.MethodPost
	if !isPost {
		template.Execute(writer, emptyTemplate())
		return
	}

	fullUrl := request.FormValue("FullUrl")
	if fullUrl == "" {
		template.Execute(writer, emptyTemplate())
		return
	}

	var tpl StartPageTemplate = storeUrlAttempt(fullUrl)
	template.Execute(writer, tpl)
}

func storeUrlAttempt(fullUrl string) StartPageTemplate {

	var expiresAfter int64 = 0
	data := url.Values{
		"FullUrl":      {fullUrl},
		"ExpiresAfter": {fmt.Sprint(expiresAfter)},
	}

	response, err := http.PostForm(config.GetServerUrl("api/store"), data)
	if err != nil {
		panic(fmt.Sprintf("Error: %s\n", err))
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("Error: %s\n", err))
	}

	var urlData core.UrlData
	err = json.Unmarshal(responseBody, &urlData)
	if err == nil && (core.UrlData{} != urlData) {
		return processSuccessfulStore(&urlData)
	}

	var errorMessage core.ErrorMessage
	err = json.Unmarshal(responseBody, &errorMessage)
	if err == nil && (core.ErrorMessage{} != errorMessage) {
		return processErrorStore(&errorMessage)
	}

	return StartPageTemplate{
		ShowMessage:  true,
		ErrorMessage: "UKNOWN ERROR",
		ShortUrl:     "",
	}
}

func processSuccessfulStore(urlData *core.UrlData) StartPageTemplate {
	shortUrl := urlData.ShortUrlCode
	return StartPageTemplate{
		ShowMessage:  true,
		ErrorMessage: "",
		ShortUrl:     shortUrl,
	}
}

func processErrorStore(errorMessage *core.ErrorMessage) StartPageTemplate {
	return StartPageTemplate{
		ShowMessage:  true,
		ErrorMessage: errorMessage.Error,
		ShortUrl:     "",
	}
}

func emptyTemplate() StartPageTemplate {
	return StartPageTemplate{
		ShowMessage:  false,
		ErrorMessage: "",
		ShortUrl:     "",
	}
}
