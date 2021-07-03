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

func fromBString(b string) string {

	// split on :
	idx := 0
	cc := b[idx]

	nCollector := ""
	for {
		if cc == ':' {
			println("finished collecting numbers")
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

	return sCollector
}

// Integers are represented by an 'i' and followed by the number in base 10 followed by an 'e'
func toBInteger(n int) string {

	return fmt.Sprintf("i%de", n)
}

func fromBInteger(b string) int {
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
	return n
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

func fromBList(b string) []interface{} {
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

			bInteger := fromBInteger(b[idx:])

			result = append(result, bInteger)
			advance = len(fmt.Sprintf("i%de", bInteger))
			// idx += advance
		} else if cc >= '0' && cc <= '9' {
			t := fromBString(b[idx:])
			advance = len(fmt.Sprintf("%d:%s", len(t), t)) // TODO: make the function output both or the size
			result = append(result, t)
		}
		idx += advance
		cc = b[idx]
	}

	return result
}

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