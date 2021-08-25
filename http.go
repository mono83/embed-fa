package fa

import (
	"net/http"
	"strings"
)

// HTTPHandler constructs HTTP handler that will serve static data
func HTTPHandler(path string) http.Handler {
	assets, err := BuildAssets()
	if err != nil {
		panic(err)
	}
	return &handler{
		AssetsHolder: *assets,
		RelativePath: path,
	}
}

type handler struct {
	AssetsHolder

	RelativePath string // TODO relative path
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	if len(uri) > 1 {
		if strings.HasPrefix(uri, "/") {
			uri = uri[1:]
		}
		if h.Has(uri) {
			asset, _ := h.Get(uri)
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
