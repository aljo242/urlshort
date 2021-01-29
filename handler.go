package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc
// (note HandlerFunc implements http.Handler interface)
// will attempt to map any paths (keys in a map)
// to corresponding URL (map values)
// If path is not provided in the map, the fallback
// http.Handler is used
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// TODO
	return nil
}

// YAMLHandler will parse the provided YAML and then
// return an http.HandlerFunc (note HandlerFunc implements
// http.Handler interface)
// will attempt to map any paths to their corresponding
// URL.  If path is not provided in the YAML, the
// fallback http.Handler will be used instead
//
// YAML is expected to be in format:
// 		- path: /somepath
//		- url: https://www.some-url.com/demo
//
// Only errors that can be returned all related to having
// invalid YAML dataa
//
// See MapHandler to create a similar http.Handler via
// a mapping of paths to urls
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO
	return nil, nil
}
