package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kaepa3/oauth/lib"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	v2 "google.golang.org/api/oauth2/v2"
)

func main() {
	goji.Get("/", indexPage)
	goji.Get("/callback", callbackPage)
	goji.Serve()
}

func indexPage(c web.C, w http.ResponseWriter, r *http.Request) {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func callbackPage(c web.C, w http.ResponseWriter, r *http.Request) {
	config := google.GetConnect()

	context := context.Background()

	token, err := config.Exchange(context, createCode(r))
	if err != nil {
		log.Fatal(err)
	}

	if token.Valid() == false {
		log.Fatal("vaild token")
	}

	service, _ := v2.New(config.Client(context, token))
	tokenInfo, _ := service.Tokeninfo().AccessToken(token.AccessToken).Context(context).Do()

	w.Header().Set("Context-Type", "text/html: charset=utf-8")
	buf, _ := tokenInfo.MarshalJSON()
	fmt.Fprintf(w, string(buf))
}

func createCode(r *http.Request) string {
	for key, values := range r.URL.Query() {
		if key == "code" {
			for _, v := range values {
				return v
			}
		}
	}
	return ""
}
