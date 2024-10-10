package tag

import (
	"m3u8/scanner"
	"strings"
)

type Media struct {
	Type       string
	GroupId    string
	Name       string
	Autoselect bool
	Default    bool
}

func ParseMedia(input string) *Media {
	out := Media{}

	hm := scanner.ScanArgs(input)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "TYPE":
			out.Type = v
		case "GROUP-ID":
			out.GroupId = strings.ReplaceAll(v, `"`, "")
		case "NAME":
			out.Name = strings.ReplaceAll(v, `"`, "")
		case "AUTOSELECT":
			out.Autoselect = v == "YES"
		case "DEFAULT":
			out.Default = v == "YES"
		}
	}

	return &out
}
