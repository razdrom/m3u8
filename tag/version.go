package tag

import (
	"strconv"
)

// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.2

type Version struct {
	raw       string
	rawparsed bool
	value     int64
}

func NewVersion(raw string) *Version {
	return &Version{raw: raw}
}

func (t *Version) parse() {
	if version, err := strconv.ParseInt(t.raw, 10, 64); err == nil {
		t.value = version
	}
}

func (t *Version) GetValue() int64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.value
}
