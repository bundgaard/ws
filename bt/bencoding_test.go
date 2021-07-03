package bt

import (
	"testing"
)

func TestTo(t *testing.T) {
	t.Run("To B String", func(t *testing.T) {
		got := toBString("cow")
		if got != "3:cow" {
			t.Error("expected '3:cow'. got ", got)
		}
	})

	t.Run("To B Integer", func(t *testing.T) {
		tests := []struct {
			Input    int
			Expected string
		}{
			{3, "i3e"},
			{-3, "i-3e"},
			{0, "i0e"},
		}

		for _, test := range tests {
			got := toBInteger(test.Input)
			if got != test.Expected {
				t.Errorf("expected %s. got=%q", test.Expected, got)
			}
		}
	})

	t.Run("to B List", func(t *testing.T) {
		tests := []struct {
			Input    []interface{}
			Expected string
		}{
			{[]interface{}{3, 4}, "li3ei4ee"},
			{[]interface{}{"spam", "eggs"}, "l4:spam4:eggse"},
		}

		for _, test := range tests {
			got := toBList(test.Input...)

			if got != test.Expected {
				t.Errorf("expected %q. got %q", test.Expected, got)
			}
		}
	})

	t.Run("to B Dictionary", func(t *testing.T) {
		tests := []struct {
			Input         map[string]interface{}
			ExpectedValue string
		}{
			{map[string]interface{}{"spam": []interface{}{"a", "b"}}, "d4:spaml1:a1:bee"},
		}

		for _, test := range tests {
			got := toBDict(test.Input)
			if got != test.ExpectedValue {
				t.Errorf("expected %q. got %q", test.ExpectedValue, got)
			}
		}

	})
}
