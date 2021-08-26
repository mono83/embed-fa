package fa

import (
	"embed"
)

//go:embed assets/*
var efs embed.FS

// BuildAssets constructs assets collection
func BuildAssets() (*AssetsHolder, error) {
	a := AssetsHolder(map[string]Asset{})
	for _, folder := range []string{"css", "js", "webfonts"} {
		files, err := efs.ReadDir("assets/" + folder)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if !file.IsDir() {
				if err := a.Import(efs, "assets/", folder+"/"+file.Name(), true, nil); err != nil {
					return nil, err
				}
			}
		}
	}
	return &a, nil
}

// MustBuildAssets constructs assets collection
func MustBuildAssets() AssetsHolder {
	a, err := BuildAssets()
	if err != nil {
		panic(err)
	}
	return *a
}
