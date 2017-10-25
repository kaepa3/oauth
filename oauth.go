package main

import (
	"fmt"
	"net/http"

	"github.com/kaepa3/oauth/lib"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {

	goji.Get("/", indexPage)
	goji.Get("/callback", callbackPage)
	goji.Serve()
}

func indexPage(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "text/html: charset=utf-8")
	lib.GetConnect()
	fmt.Fprintf(w, "Hello, World")
}

func callbackPage(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "text/html: charset=utf-8")
	fmt.Fprintf(w, "call, back")
}

