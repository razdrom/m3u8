package tag

import (
	"strconv"
	"strings"
)

// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.1

type Info struct {
	raw       string
	rawparsed bool
	duruation float64
	title     string
}

func NewInfo(raw string) *Info {
	return &Info{raw: raw}
}

func (t *Info) parse() {
	args := strings.Split(t.raw, ",")
	if len(args) < 1 {
		return
	}

	duration, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return
	}
	t.duruation = duration

	if len(args) == 2 {
		t.title = args[1]
	}
}

func (t *Info) GetDuration() float64 {
	if !t.rawparsed {
		t.parse()
	}
	return t.duruation
}
func (t *Info) GetTitle() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.title
}
