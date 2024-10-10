package scanner

import (
	"strings"
	"text/scanner"
	"unicode"
)

func ScanArgs(input string) map[string]string {
	var scan scanner.Scanner
	scan.IsIdentRune = func(r rune, i int) bool {
		return r == '-' || unicode.IsLetter(r)
	}

	scan.Init(strings.NewReader(input))

	var k, v strings.Builder
	hm := map[string]string{}
	iskey := true

	for {
		if scanner.EOF == scan.Scan() {
			hm[k.String()] = v.String()
			k.Reset()
			v.Reset()
			break
		}

		tt := scan.TokenText()
		if tt == "=" {
			iskey = false
			continue
		}

		if tt == "," {
			iskey = true

			hm[k.String()] = v.String()
			k.Reset()
			v.Reset()
			continue
		}

		if iskey {
			k.WriteString(tt)
		} else {
			v.WriteString(tt)
		}
	}

	return hm
}
