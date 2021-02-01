package urlshort

import (
	"net/http"
)

type pathToUrl struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc
// (note HandlerFunc implements http.Handler interface)
// will attempt to map any paths (keys in a map)
// to corresponding URL (map values)
// If path is not provided in the map, the fallback
// http.Handler is used
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// TODO

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return handlerFunc
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
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO
	parsedYAML, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYAML)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yaml []byte) ([]string, error) {
	return nil, nil
}

func buildMap(parsedYAML []string) map[string]string {
	parsedMap := make(map[string]string)
	return parsedMap
}
