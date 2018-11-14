package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w,r,url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w,r)
		}
	}
	return h
}
// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	p, err := parseYAML(yml)
	if err !=nil {
		return nil, err
	}
	m := buildMap(p)
	return MapHandler(m, fallback), nil
}

// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
func parseYAML(yml []byte) ([]map[string]string, error){
	var parsed []map[string]string
	if err := yaml.Unmarshal(yml, &parsed); err != nil {
		return nil, err
	}
	return parsed, nil
	}

func buildMap(p []map[string]string) map[string]string{
	m := make(map[string]string)
	for _, v := range p {
		m[v["path"]] = v["url"]
	}
	return m
}