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
			Input         string
			ExpectedValue []interface{}
			ExpectedSize  int
		}{
			{"l4:spam4:eggse", []interface{}{"spam", "eggs"}, 14},
			{"li3ei4ee", []interface{}{3, 4}, 8},
		}

		for _, test := range tests {
			got, size := fromBList(test.Input)
			t.Log(got)
			if !reflect.DeepEqual(got, test.ExpectedValue) {
				t.Errorf("exepcted %v got %q", test.ExpectedValue, got)
			}
			if size != test.ExpectedSize {
				t.Errorf("expected %d. got %d", test.ExpectedSize, size)
			}
		}
	})

	t.Run("From B Dictionary", func(t *testing.T) {
		tests := []struct {
			Input         string
			ExpectedValue map[string]interface{}
		}{
			{"d4:spaml1:a1:bee", map[string]interface{}{"spam": []interface{}{"a", "b"}}},
			{"d1:a2:be", map[string]interface{}{"a": "be"}},
			{"d1:a2:be3:cow4:spame", map[string]interface{}{"a": "be", "cow": "spam"}},
		}

		for _, test := range tests {
			got := fromBDict(test.Input)
			t.Logf("got %v", got)
			if !reflect.DeepEqual(got, test.ExpectedValue) {
				t.Errorf("expected %q. got %q", test.ExpectedValue, got)
			}
		}
	})
}
