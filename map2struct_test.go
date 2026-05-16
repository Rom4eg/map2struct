package map2struct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap2Struct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		m      map[string]any
		s      interface{}
		expect interface{}
		err    error
	}{
		{
			name:   "FAIL: empty map",
			m:      map[string]any{},
			s:      struct{}{},
			expect: struct{}{},
			err:    ErrEmptyMap,
		},
		{
			name: "FAIL: struct not a pointer",
			m: map[string]any{
				"test": "test",
			},
			s:      struct{}{},
			expect: struct{}{},
			err:    ErrMustBePointerStruct,
		},
		{
			name: "PASS: OK",
			m: map[string]any{
				"str": "test",
				"num": 123,
				"lst": []string{"foo", "bar"},
			},
			s: &struct {
				Str string   `m2s:"str"`
				Num int      `m2s:"num"`
				Lst []string `m2s:"lst"`
			}{},
			expect: &struct {
				Str string   `m2s:"str"`
				Num int      `m2s:"num"`
				Lst []string `m2s:"lst"`
			}{
				Str: "test",
				Num: 123,
				Lst: []string{"foo", "bar"},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Map2Struct(tt.m, tt.s)
			if tt.err != nil {
				assert.ErrorIs(t, e, tt.err)
				return
			} else {
				assert.NoError(t, e)
			}
			assert.Equal(t, tt.s, tt.expect)
		})
	}
}
