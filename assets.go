package fa

import (
	"bytes"
	"compress/gzip"
	"io/fs"
	"strings"
)

// Asset defines asset data
type Asset struct {
	Data        []byte
	ContentType string
	Gzipped     bool
}

// AssetsHolder is holder for all assets
type AssetsHolder map[string]Asset

// Has returns true if asset with given name exists
func (a AssetsHolder) Has(name string) bool {
	_, ok := a[name]
	return ok
}

// Get returns asset by name
func (a AssetsHolder) Get(name string) (Asset, bool) {
	bts, ok := a[name]
	return bts, ok
}

// Import imports filesystem contents
func (a AssetsHolder) Import(f fs.ReadFileFS, prefix, name string, encode bool, predicate func(string) bool) error {
	if predicate == nil {
		predicate = func(string) bool { return true }
	}

	var bts []byte
	raw, err := f.ReadFile(prefix + name)
	if err != nil {
		return err
	}
	if encode {
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		if _, err := gz.Write(raw); err != nil {
			return err
		}
		if err := gz.Close(); err != nil {
			return err
		}
		bts = b.Bytes()
	} else {
		bts = raw
	}

	ct := "text/plain"
	if strings.HasSuffix(name, "css") {
		ct = "text/css"
	} else if strings.HasSuffix(name, "js") {
		ct = "text/javascript"
	} else if strings.HasSuffix(name, "woff2") {
		ct = "font/woff2"
	} else if strings.HasSuffix(name, "woff") {
		ct = "font/woff"
	} else if strings.HasSuffix(name, "eot") {
		ct = "font/eot"
	} else if strings.HasSuffix(name, "ttf") {
		ct = "font/ttf"
	} else if strings.HasSuffix(name, "svg") {
		ct = "image/svg+xml"
	}
	a[name] = Asset{
		Data:        bts,
		ContentType: ct,
		Gzipped:     true,
	}
	return nil
}
