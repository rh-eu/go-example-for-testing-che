// +build dev

package sitedata

//go:generate go run -tags=dev ./assets_generate.go

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../../sitedata")