package tag

import (
	"strconv"
)

// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1.2

type Version struct {
	Value int64
}

func ParseVersion(input string) *Version {
	out := Version{}

	version, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return nil
	}

	out.Value = version

	return &out
}
