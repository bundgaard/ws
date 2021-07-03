package bt

import (
	"reflect"
	"testing"
)

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
			if !reflect.DeepEqual(got, test.ExpectedValue) {
				t.Errorf("exepcted %v got %q", test.ExpectedValue, got)
			}
			if size != test.ExpectedSize {
				t.Errorf("expected %d. got %d", test.ExpectedSize, size)
			}
		}
	})

	t.Run("From B Dictionary", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Error(err)
			}
		}()
		tests := []struct {
			Input         string
			ExpectedValue map[string]interface{}
		}{
			{"d4:spaml1:a1:bee", map[string]interface{}{"spam": []interface{}{"a", "b"}}},
			{"d1:a2:bee", map[string]interface{}{"a": "be"}},
			{"d1:a2:be3:cow4:spame", map[string]interface{}{"a": "be", "cow": "spam"}},
			{"d1:ai10ee", map[string]interface{}{"a": 10}},
			{"d17:dht_backup_enablei1ee", map[string]interface{}{"dht_backup_enable": 1}},
			{"d1:ad2:be0:ee", map[string]interface{}{"a": map[string]interface{}{"be": ""}}},
			{"d7:comment0:13:comment.utf-80:10:created by15:Azureus/2.5.0.013:creation datei1168469015e8:encoding5:UTF-88:announce54:http://files2.publicdomaintorrents.com/bt/announce.php18:azureus_propertiesd17:dht_backup_enablei1eee", map[string]interface{}{
				"announce": "http://files2.publicdomaintorrents.com/bt/announce.php",
				"azureus_properties": map[string]interface{}{
					"dht_backup_enable": 1,
				},
				"comment":       "",
				"comment.utf-8": "",
				"created by":    "Azureus/2.5.0.0",
				"creation date": 1168469015,
				"encoding":      "UTF-8",
			}},
		}

		for idx, test := range tests {
			got, _ := fromBDict(test.Input)
			if !reflect.DeepEqual(test.ExpectedValue, got) {
				t.Errorf("test[%04d] expected %+v. got %+v", idx, test.ExpectedValue, got)
			}
		}
	})
}
