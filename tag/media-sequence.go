package tag

import (
	"strconv"
	"strings"
)

type MediaSequence struct {
	Value int64
}

func ParseMediaSequence(input string) *MediaSequence {
	out := MediaSequence{}
	args := strings.Split(input, ":")
	if len(args) < 1 {
		return nil
	}

	value, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return nil
	}

	out.Value = value
	return &out
}
