package tag

import (
	"strconv"
)

type MediaSequence struct {
	raw       string
	rawparsed bool
	value     int64
}

func NewMediaSequence(raw string) *MediaSequence {
	return &MediaSequence{raw: raw}
}

func (t *MediaSequence) parse() {
	value, err := strconv.ParseInt(t.raw, 10, 64)
	if err != nil {
		return
	}

	t.value = value
}

func (t *MediaSequence) GetValue() int64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.value
}
