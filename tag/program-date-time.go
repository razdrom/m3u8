package tag

import (
	"time"
)

type ProgramDateTime struct {
	raw       string
	rawparsed bool
	value     *time.Time
}

func NewProgramDateTime(raw string) *ProgramDateTime {
	return &ProgramDateTime{raw: raw}
}

func (t *ProgramDateTime) parse() {
	t.rawparsed = true
	date, err := time.Parse(time.RFC3339, t.raw)
	if err != nil {
		return
	}

	date = date.In(time.UTC)
	t.value = &date
}

func (t *ProgramDateTime) GetValue() *time.Time {
	if !t.rawparsed {
		t.parse()
	}

	return t.value
}
