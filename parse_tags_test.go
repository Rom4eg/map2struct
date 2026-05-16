package map2struct

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		s      interface{}
		expect []M2STag
	}{
		{
			name:   "FAIL: non struct",
			s:      "test",
			expect: nil,
		},

		{
			name:   "FAIL: empty struct",
			s:      struct{}{},
			expect: nil,
		},

		{
			name: "FAIL: incorrect tag",
			s: struct {
				Name string `test:"name"`
			}{},
			expect: nil,
		},
		{
			name: "OK: primitive tag",
			s: struct {
				Name string `m2s:"name"`
			}{},
			expect: []M2STag{
				{
					FName: "Name",
					Name:  "name",
					Type:  M2SString,
				},
			},
		},
		{
			name: "OK: slice tag",
			s: struct {
				Name []string `m2s:"name"`
			}{},
			expect: []M2STag{
				{
					FName: "Name",
					Name:  "name",
					Type:  M2SList,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tags := parseTags(tt.s)
			assert.Equal(t, tt.expect, tags)
		})
	}
}
