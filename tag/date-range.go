package tag

import (
	"strconv"
	"strings"
	"time"

	"github.com/razdrom/m3u8/scanner"
)

type DateRange struct {
	raw       string
	rawparsed bool

	id              string
	class           string
	startDate       *time.Time
	endDate         *time.Time
	duration        float64
	plannedDuration float64
	endOnNext       bool
	scte35Cmd       string
	scte35Out       string
	scte35In        string
}

func NewDateRange(raw string) *DateRange {
	return &DateRange{raw: raw}
}

func (t *DateRange) parse() {
	t.rawparsed = true
	hm := scanner.ScanArgs(t.raw)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "ID":
			t.id = strings.ReplaceAll(v, `"`, "")
		case "CLASS":
			t.class = strings.ReplaceAll(v, `"`, "")
		case "START-DATE":
			v = strings.ReplaceAll(v, `"`, "")
			date, err := time.Parse(time.RFC3339, v)
			if err != nil {
				continue
			}
			date = date.In(time.UTC)
			t.startDate = &date
		case "END-DATE":
			v = strings.ReplaceAll(v, `"`, "")
			date, err := time.Parse(time.RFC3339, v)
			if err != nil {
				continue
			}
			date = date.In(time.UTC)
			t.endDate = &date
		case "END-ON-NEXT":
			t.endOnNext = v == "YES"
		case "DURATION":
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				t.duration = 0
			}
			t.duration = parsed
		case "PLANNED-DURATION":
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				t.plannedDuration = 0
			}
			t.plannedDuration = parsed

		case "SCTE35-CMD":
			t.scte35Cmd = strings.ReplaceAll(v, `"`, "")
		case "SCTE35-OUT":
			t.scte35Out = strings.ReplaceAll(v, `"`, "")
		case "SCTE35-IN":
			t.scte35In = strings.ReplaceAll(v, `"`, "")
		}
	}
}

func (t *DateRange) GetId() string {
	if !t.rawparsed {
		t.parse()
	}

	return t.id
}

func (t *DateRange) GetClass() string {
	if !t.rawparsed {
		t.parse()
	}

	return t.class
}

func (t *DateRange) GetStartDate() *time.Time {
	if !t.rawparsed {
		t.parse()
	}

	return t.startDate
}

func (t *DateRange) GetEndDate() *time.Time {
	if !t.rawparsed {
		t.parse()
	}

	return t.endDate
}

func (t *DateRange) GetDuration() float64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.duration
}

func (t *DateRange) GetPlannedDuration() float64 {
	if !t.rawparsed {
		t.parse()
	}

	return t.plannedDuration
}

func (t *DateRange) GetEndOnNext() bool {
	if !t.rawparsed {
		t.parse()
	}

	return t.endOnNext
}

func (t *DateRange) GetScte35Cmd() string {
	if !t.rawparsed {
		t.parse()
	}

	return t.scte35Cmd
}

func (t *DateRange) GetScte35Out() string {
	if !t.rawparsed {
		t.parse()
	}

	return t.scte35Out
}

func (t *DateRange) GetScte35In() string {
	if !t.rawparsed {
		t.parse()
	}

	return t.scte35In
}
