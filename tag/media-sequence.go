package tag

import (
	"strconv"
)

type MediaSequence struct {
	raw       string
	rawparsed bool
	value     uint64
}

func NewMediaSequence(raw string) *MediaSequence {
	return &MediaSequence{raw: raw}
}

func (t *MediaSequence) parse() {
	t.rawparsed = true
	value, err := strconv.ParseUint(t.raw, 10, 64)
	if err != nil {
		return
	}

	t.value = value
}

func (t *MediaSequence) GetValue() uint64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.value
}
