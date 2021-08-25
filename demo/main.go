package main

import (
	_ "embed"
	fa "github.com/mono83/embed-fa"
	"log"
	"net/http"
	"strings"
)

func main() {
	panic(http.ListenAndServe(":8080", &handler{
		assets: fa.HTTPHandler(""),
	}))
}

type handler struct {
	assets http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	log.Println(uri)
	if uri == "/" || uri == "/index.html" || uri == "/index.htm" {
		_, _ = w.Write(index)
		return
	}
	if strings.HasPrefix(uri, "/css/") || strings.HasPrefix(uri, "/js/") || strings.HasPrefix(uri, "/webfonts/") {
		h.assets.ServeHTTP(w, req)
		return
	}
	w.WriteHeader(404)
}

//go:embed index.html
var index []byte
