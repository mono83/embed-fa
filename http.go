package fa

import (
	"net/http"
	"strings"
)

// HTTPHandler constructs HTTP handler that will serve static data
func (a AssetsHolder) HTTPHandler(path string) http.Handler {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return &handler{
		AssetsHolder: a,
		RelativePath: path,
	}
}

// HTTPHandlerFunc constructs HTTP handler function that will serve static data
func (a AssetsHolder) HTTPHandlerFunc(path string) http.HandlerFunc {
	return a.HTTPHandler(path).ServeHTTP
}

type handler struct {
	AssetsHolder

	RelativePath string
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	if len(uri) > len(h.RelativePath) && strings.HasPrefix(uri, h.RelativePath) {
		uri = uri[len(h.RelativePath):]
		asset, present := h.Get(uri)
		if present {
			w.Header().Set("Content-Type", asset.ContentType)
			if asset.Gzipped {
				w.Header().Set("Content-Encoding", "gzip")
			}
			_, _ = w.Write(asset.Data)
			return
		}
	}

	w.WriteHeader(404)
}
