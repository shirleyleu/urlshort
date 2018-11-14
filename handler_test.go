package urlshort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseYAML(t *testing.T) {
	tests := []struct {
		name    string
		yml    []byte
		want    []map[string]string
		wantErr bool
	}{
		{
			name: "Normal",
			yml: []byte(`- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution`),
			want:[]map[string]string{{"path":"/urlshort", "url":"https://github.com/gophercises/urlshort"},{"path":"/urlshort-final", "url":"https://github.com/gophercises/urlshort/tree/solution"}},
			wantErr: false,
		},
		{
			name: "Incorrect YAML format",
			yml: []byte(` path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution`),
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseYAML(tt.yml)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("parseYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
