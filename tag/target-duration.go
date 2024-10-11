package tag

import (
	"strconv"
	"strings"
)

type TargetDuration struct {
	raw       string
	rawparsed bool
	value     int64
}

func NewTargetDuration(raw string) *TargetDuration {
	return &TargetDuration{raw: raw}
}

func (t *TargetDuration) parse() {
	t.rawparsed = true
	args := strings.Split(t.raw, ":")
	if len(args) < 1 {
		return
	}

	value, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}

	t.value = value
}

func (t *TargetDuration) GetValue() int64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.value
}
