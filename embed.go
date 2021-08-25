package fa

import (
	"bytes"
	"compress/gzip"
	"embed"
	"strings"
)

//go:embed assets/*
var fs embed.FS

func (a AssetsHolder) load(name string) error {
	bts, err := fs.ReadFile("assets/" + name)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(bts); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
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
		Data:        b.Bytes(),
		ContentType: ct,
		Gzipped:     true,
	}
	return nil
}

// BuildAssets constructs assets collection
func BuildAssets() (*AssetsHolder, error) {
	a := AssetsHolder(map[string]Asset{})
	for _, folder := range []string{"css", "js", "webfonts"} {
		files, err := fs.ReadDir("assets/" + folder)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if !file.IsDir() {
				if err := a.load(folder + "/" + file.Name()); err != nil {
					return nil, err
				}
			}
		}
	}
	return &a, nil
}
