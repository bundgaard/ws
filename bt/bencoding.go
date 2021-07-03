package bt

import (
	"fmt"
	"strconv"
	"strings"
)

// Strings are length-prefixed base ten followed by a colon and the string 3:cow
func toBString(t string) string {
	return fmt.Sprintf("%d:%s", len(t), t)
}

// wil take incoming bencoded string 4:spam and return spam and the total length
func fromBString(b string) (string, int) {

	// split on :
	idx := 0
	cc := b[idx]

	nCollector := ""
	for {
		if cc == ':' {
			break
		}
		if cc >= '0' && cc <= '9' {
			nCollector += string(cc)
		}
		idx++
		cc = b[idx]

	}
	idx++ // skip the :
	// convert n to int

	nn, err := strconv.Atoi(nCollector)
	if err != nil {
		panic(err)
	}

	sCollector := ""
	for i := 0; i < nn; i++ {
		sCollector += string(b[idx+i])
	}

	if len(sCollector) != nn {
		panic("wrong len")
	}
	// verify length

	// return
	size := len(fmt.Sprintf("%d:%s", nn, sCollector))
	return sCollector, size
}

// Integers are represented by an 'i' and followed by the number in base 10 followed by an 'e'
func toBInteger(n int) string {

	return fmt.Sprintf("i%de", n)
}

func fromBInteger(b string) (int, int) {
	idx := 0
	cc := b[idx]
	if cc != 'i' || b[len(b)-1] != 'e' {
		panic("not a bencoded number")
	}

	idx++
	cc = b[idx]
	nCollector := ""
	for cc != 'e' {

		// add validation according to site: http://www.dsc.ufcg.edu.br/~nazareno/bt/bt.htm
		if cc >= '0' && cc <= '9' {
			nCollector += string(cc)
		}
		idx++
		cc = b[idx]
	}

	n, err := strconv.Atoi(nCollector)
	if err != nil {
		panic(err)
	}

	return n, len(fmt.Sprintf("i%de", n))
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

func fromBList(b string) ([]interface{}, int) {
	result := make([]interface{}, 0)

	idx := 0
	cc := b[idx]
	if cc != 'l' || b[len(b)-1] != 'e' {
		panic("not a bencoded list")
	}

	// idx++ // consume 'l'
	advance := 0
	idx++
	for cc != 'e' {

		if cc == 'i' {

			bInteger, size := fromBInteger(b[idx:])

			result = append(result, bInteger)
			advance = size
			// idx += advance
		} else if cc >= '0' && cc <= '9' {
			t, size := fromBString(b[idx:])
			advance = size
			result = append(result, t)
		}
		idx += advance
		cc = b[idx]
	}

	return result, len(b)
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
		default:
			panic(fmt.Sprintf("not supported %T", vv))
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
func fromBDict(b string) map[string]interface{} {
	fmt.Printf("DICT %s %d\n", b, len(b))
	result := make(map[string]interface{})
	idx := 0
	cc := b[idx]

	if cc != 'd' || b[len(b)-1] != 'e' {
		panic("not a bencoded dictionary")
	}

	idx++ // consume 'd'

	// find first key

	for {
		if idx >= len(b)-1 {
			break
		}

		key, size := fromBString(b[idx:])
		idx += size
		cc = b[idx]
		// find first key
		if cc == 'i' {
			val, size := fromBInteger(b[idx:])
			idx += size
			result[key] = val
		} else if cc >= '0' && cc <= '9' {
			val, size := fromBString(b[idx:])
			idx += size
			result[key] = val
		} else if cc == 'l' {
			val, size := fromBList(b[idx:])
			idx += size // we need to stay on l..e
			result[key] = val
		}

	}

	return result
}
