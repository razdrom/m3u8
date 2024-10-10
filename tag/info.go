package tag

import (
	"strconv"
	"strings"
)

// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.2.1

type Info struct {
	Druation float64
	Title    string
}

func ParseInfo(input string) *Info {
	out := Info{}

	args := strings.Split(input, ",")
	if len(args) < 1 {
		return nil
	}

	duration, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil
	}
	out.Druation = duration

	if len(args) == 2 {
		out.Title = args[1]
	}

	return &out
}
