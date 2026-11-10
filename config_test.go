package main

import (
	"reflect"
	"testing"

	"github.com/tidwall/gjson"
)

func Test_parseApostas(t *testing.T) {
	casos := []struct {
		json     string
		esperado []string
	}{
		{`{"aposta": "05 12 29 32 38 42"}`, []string{"05 12 29 32 38 42"}},
		{`{"aposta": ["05 12 29 32 38 42", " 01 07 15 23 44 60 ", ""]}`, []string{"05 12 29 32 38 42", "01 07 15 23 44 60"}},
		{`{}`, nil},
	}
	for _, c := range casos {
		retornado := parseApostas(gjson.Get(c.json, "aposta"))
		if !reflect.DeepEqual(c.esperado, retornado) {
			t.Errorf("json: %s, expected: %q, returned: %q", c.json, c.esperado, retornado)
		}
	}
}
