package storageid

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		prefix string
		length uint
	}{
		"10-character-prefix": {prefix: "production", length: 24},
		"3-character-prefix":  {prefix: "int", length: 24},
		"1-character-prefix":  {prefix: "i", length: 24},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := New(tc.prefix, tc.length)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.length != uint(len(got)) {
				t.Errorf("want length %v, got length %v", tc.length, uint(len(got)))
			}
			if !strings.HasPrefix(got, tc.prefix) {
				t.Errorf("id %s is missing prefix %s", got, tc.prefix)
			}
		})
	}
}

func TestNewErrors(t *testing.T) {
	tests := map[string]struct {
		prefix string
		length uint
	}{
		"empty-prefix":       {prefix: "", length: 24},
		"ampersand-prefix":   {prefix: "int@2", length: 24},
		"punctuation-prefix": {prefix: "int!2", length: 24},
		"mixed-case-prefix":  {prefix: "Int", length: 24},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := New(tc.prefix, tc.length)
			if err == nil {
				t.Fatal("expecting error")
			}
		})
	}
}
