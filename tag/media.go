package tag

import (
	"strings"

	"github.com/razdrom/m3u8/scanner"
)

type Media struct {
	Type            string
	GroupId         string
	Name            string
	Autoselect      bool
	Default         bool
	Uri             string
	Language        string
	AssocLanguage   string
	Forced          bool
	InstreamId      string
	Characteristics []string
	Channels        string
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
		case "FORCED":
			out.Forced = v == "YES"
		case "URI":
			out.Uri = strings.ReplaceAll(v, `"`, "")
		case "LANGUAGE":
			out.Language = strings.ReplaceAll(v, `"`, "")
		case "ASSOC-LANGUAGE":
			out.AssocLanguage = strings.ReplaceAll(v, `"`, "")
		case "INSTREAM-ID":
			out.InstreamId = strings.ReplaceAll(v, `"`, "")
		case "CHARACTERISTICS":
			v = strings.ReplaceAll(v, `"`, "")
			out.Characteristics = strings.Split(v, ",")
		case "CHANNELS":
			out.Channels = strings.ReplaceAll(v, `"`, "")
		}
	}

	return &out
}
