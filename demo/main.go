package main

import (
	_ "embed"
	fa "github.com/mono83/embed-fa"
	"log"
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":8080", &handler{
		assets: fa.MustBuildAssets().HTTPHandler("fa"),
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
	h.assets.ServeHTTP(w, req)
}

//go:embed index.html
var index []byte
