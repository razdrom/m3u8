package tag

import (
	"time"
)

type ProgramDateTime struct {
	Value *time.Time
}

func ParseProgramDateTime(input string) *ProgramDateTime {
	out := ProgramDateTime{}

	date, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return nil
	}
	out.Value = &date

	return &out
}
