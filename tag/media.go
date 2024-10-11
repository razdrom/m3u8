package tag

import (
	"strings"

	"github.com/razdrom/m3u8/scanner"
)

type Media struct {
	raw       string
	rawparsed bool

	mediaType       string
	groupId         string
	name            string
	isAutoselect    bool
	isDefault       bool
	uri             string
	language        string
	assocLanguage   string
	isForced        bool
	instreamId      string
	characteristics []string
	channels        string
}

func NewMedia(raw string) *Media {
	return &Media{raw: raw}
}

func (t *Media) parse() {
	if t.raw == "" {
		return
	}

	hm := scanner.ScanArgs(t.raw)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "TYPE":
			t.mediaType = v
		case "GROUP-ID":
			v = strings.ReplaceAll(v, `"`, "")
			t.groupId = v
		case "NAME":
			v = strings.ReplaceAll(v, `"`, "")
			t.name = v
		case "AUTOSELECT":

			t.isAutoselect = v == "YES"
		case "DEFAULT":
			t.isDefault = v == "YES"
		case "FORCED":
			t.isForced = v == "YES"
		case "URI":
			v = strings.ReplaceAll(v, `"`, "")
			t.uri = v
		case "LANGUAGE":
			v = strings.ReplaceAll(v, `"`, "")
			t.language = v
		case "ASSOC-LANGUAGE":
			v = strings.ReplaceAll(v, `"`, "")
			t.assocLanguage = v
		case "INSTREAM-ID":
			v = strings.ReplaceAll(v, `"`, "")
			t.instreamId = v
		case "CHARACTERISTICS":
			v = strings.ReplaceAll(v, `"`, "")
			t.characteristics = strings.Split(v, ",")
		case "CHANNELS":
			v = strings.ReplaceAll(v, `"`, "")
			t.channels = v
		}
	}
}

func (t *Media) GetType() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.mediaType
}

func (t *Media) GetGroupId() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.groupId
}

func (t *Media) GetName() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.name
}

func (t *Media) GetAutoselect() bool {
	if !t.rawparsed {
		t.parse()
	}
	return t.isAutoselect
}

func (t *Media) GetDefault() bool {
	if !t.rawparsed {
		t.parse()
	}
	return t.isDefault
}

func (t *Media) GetUri() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.uri
}

func (t *Media) GetLanguage() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.language
}

func (t *Media) GetAssocLanguage() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.assocLanguage
}

func (t *Media) GetForced() bool {
	if !t.rawparsed {
		t.parse()
	}
	return t.isForced
}

func (t *Media) GetInstreamId() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.instreamId
}

func (t *Media) GetCharacteristics() []string {
	if !t.rawparsed {
		t.parse()
	}
	return t.characteristics
}

func (t *Media) GetChannels() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.channels
}
