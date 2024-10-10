package tag

import (
	"strconv"
	"strings"
)

type TargetDuration struct {
	Value int64
}

func ParseTargetDuration(input string) *TargetDuration {
	out := TargetDuration{}
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
