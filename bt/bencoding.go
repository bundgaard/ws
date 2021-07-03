package bt

import (
	"fmt"
	"strings"
)

// Strings are length-prefixed base ten followed by a colon and the string 3:cow
func toBString(t string) string {
	return fmt.Sprintf("%d:%s", len(t), t)
}

// Integers are represented by an 'i' and followed by the number in base 10 followed by an 'e'
func toBInteger(n int) string {

	return fmt.Sprintf("i%de", n)
}

func toBList(args ...interface{}) string {
	var out strings.Builder
	out.WriteRune('l')
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			out.WriteString(toBString(t))
		case int:
			out.WriteString(toBInteger(t))
		default:
			panic(fmt.Sprintf("not supported %T", t))
		}
	}
	out.WriteRune('e')
	return out.String()
}

/*
Dictionaries are encoded as a 'd' followed by a list of alternating keys
and their corresponding values followed by an 'e'.

For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow': 'moo', 'spam': 'eggs'}
and d4:spaml1:a1:bee corresponds to {'spam': ['a', 'b']} .
Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics).
*/
func toBDict(m map[string]interface{}) string {
	var out strings.Builder

	out.WriteRune('d')
	for k, v := range m {
		kk := toBString(k)
		out.WriteString(kk)
		switch vv := v.(type) {
		case string:
			out.WriteString(toBString(vv))
		case int:
			out.WriteString(toBInteger(vv))
		case []interface{}:
			out.WriteString(toBList(vv...))
		case map[string]interface{}:
			out.WriteString(toBDict(vv))
		default:
			panic(fmt.Sprintf("not supported %T", vv))
		}
	}

	out.WriteRune('e')
	return out.String()
}
