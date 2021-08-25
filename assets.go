package fa

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
