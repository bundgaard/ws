package bt

import (
	"fmt"
	"strconv"
)

/*
Dictionaries are encoded as a 'd' followed by a list of alternating keys
and their corresponding values followed by an 'e'.

For example, d3:cow3:moo4:spam4:eggse corresponds to {'cow': 'moo', 'spam': 'eggs'}
and d4:spaml1:a1:bee corresponds to {'spam': ['a', 'b']} .
Keys must be strings and appear in sorted order (sorted as raw strings, not alphanumerics).
*/
func fromBDict(b string) (map[string]interface{}, int) {
	// fmt.Printf("DICT %s %d\n", b, len(b))
	result := make(map[string]interface{})
	idx := 0
	cc := b[idx]

	if cc != 'd' || b[len(b)-1] != 'e' {
		panic("not a bencoded dictionary")
	}

	idx++ // consume 'd'
	cc = b[idx]
	// find first key

	for {
		// fmt.Printf("fromBDict\t%c %d\n", cc, idx)
		if cc == 'e' {
			break
		}
		key, size := parseValue(b[idx:])
		// fmt.Printf("\t KEY %s %c %d\n", key, cc, idx)
		idx += size
		val, size := parseValue(b[idx:])
		// fmt.Printf("\t VAL %v %c %d\n", val, cc, idx)
		result[key.(string)] = val
		idx += size
		//	fmt.Printf("\t %s %c %d\n", b, b[idx], idx)
		cc = b[idx]
	}

	// fmt.Printf("fromBDict %d %d\n", idx, len(b))
	// fmt.Printf("END %c %d\n", b[idx], idx)
	return result, idx
}

func fromBList(b string) ([]interface{}, int) {
	// fmt.Printf("\tfromBList %s %c %d\n", b, b[0], 0)
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
	idx++
	return result, idx
}

func fromBInteger(b string) (int, int) {
	// 	log.Println("fromBInteger", b)
	idx := 0
	cc := b[idx]
	if cc != 'i' || b[len(b)-1] != 'e' {
		panic("not a bencoded number")
	}

	idx++
	cc = b[idx]
	// log.Printf("cc %c idx %d\n", cc, idx)
	nCollector := ""
	for cc != 'e' {

		// add validation according to site: http://www.dsc.ufcg.edu.br/~nazareno/bt/bt.htm
		if cc >= '0' && cc <= '9' {
			nCollector += string(cc)
		}
		idx++
		cc = b[idx]
	}

	// log.Printf("cc %c idx %d nColl %s\n", cc, idx, nCollector)
	n, err := strconv.Atoi(nCollector)
	if err != nil {
		panic(err)
	}

	return n, len(fmt.Sprintf("i%de", n))
}

// wil take incoming bencoded string 4:spam and return spam and the total length
func fromBString(b string) (string, int) {

	// split on :
	idx := 0
	cc := b[idx]
	// fmt.Printf("fromBString %s '%c' %d\n", b, cc, idx)
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
	// fmt.Printf("\t '%c' %d\n", cc, idx)
	nn, err := strconv.Atoi(nCollector)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("\t atoi %d\n", nn)

	sCollector := ""
	for i := 0; i < nn; i++ {
		sCollector += string(b[idx+i])
	}

	// fmt.Printf("\t %s\n", sCollector)

	// verify length
	if len(sCollector) != nn {
		panic("wrong len")
	}

	// return
	size := len(fmt.Sprintf("%d:%s", nn, sCollector))
	// fmt.Printf("\t RETURN %s %c %d\n", sCollector, b[idx], idx)
	return sCollector, size
}

func parseValue(b string) (interface{}, int) {

	idx := 0
	cc := b[idx]
	// fmt.Printf("parseValue %s %c %d\n", b, cc, idx)

	if cc == 'i' {
		return fromBInteger(b[idx:])
	} else if cc == 'd' {
		return fromBDict(b[idx:])
	} else if cc == 'l' {
		return fromBList(b[idx:])
	} else if cc >= '0' && cc <= '9' {
		// fmt.Printf("\t STRING '%c' %d %s\n", cc, idx, b)
		return fromBString(b[idx:])
	}

	panic("parsevalue unexpected token")
}
