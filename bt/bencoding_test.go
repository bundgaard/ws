package bt

import (
	"reflect"
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
		expected := "d4:spaml1:a1:bee"
		got := toBDict(map[string]interface{}{"spam": []interface{}{"a", "b"}})

		if got != expected {
			t.Errorf("expected %q. got %q", expected, got)
		}
	})
}

func TestFrom(t *testing.T) {

	t.Run("From B String", func(t *testing.T) {

		tests := []struct {
			Input          string
			ExpectedString string
			ExpectedSize   int
			wantPanic      bool
		}{
			{"4:spam", "spam", 6, false},
			{"8:dinosaur", "dinosaur", 10, false},
			{"9:spam spam", "spam spam", 11, false},
		}
		for _, test := range tests {
			got, size := fromBString(test.Input)
			if test.ExpectedString != got {
				t.Errorf("expected %q. got %q", test.ExpectedString, got)
			}
			if test.ExpectedSize != size {
				t.Errorf("expected %d. got %d", test.ExpectedSize, size)
			}
		}

	})

	t.Run("From B Integer", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("test recovered %v", err)
			}
		}()
		tests := []struct {
			Input        string
			ExpectedInt  int
			ExpectedSize int
		}{
			{"i3e", 3, 3},
			{"i0e", 0, 3},
			{"i100e", 100, 5},
		}

		for _, test := range tests {
			got, size := fromBInteger(test.Input)
			if got != test.ExpectedInt {
				t.Errorf("expected %q. got %q", test.ExpectedInt, got)
			}
			if test.ExpectedSize != size {
				t.Errorf("expected %d. got %d", test.ExpectedSize, size)
			}
		}
	})

	t.Run("From B List", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("test recovered %v", err)
			}
		}()
		tests := []struct {
			Input    string
			Expected []interface{}
		}{
			{
				"l4:spam4:eggse", []interface{}{"spam", "eggs"},
			},
			{"li3ei4ee", []interface{}{3, 4}},
		}

		for _, test := range tests {
			got := fromBList(test.Input)
			t.Log(got)
			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("exepcted %v got %q", test.Expected, got)
			}
		}
	})
}
