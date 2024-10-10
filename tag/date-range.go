package tag

import (
	"strconv"
	"strings"
	"time"

	"github.com/razdrom/m3u8/scanner"
)

type DateRange struct {
	Id              string
	Class           string
	StartDate       *time.Time
	EndDate         *time.Time
	Duration        float64
	PlannedDuration float64
	EndOnNext       bool
	Scte35Cmd       string
	Scte35Out       string
	Scte35In        string
}

func ParseDateRange(input string) *DateRange {
	out := DateRange{}

	hm := scanner.ScanArgs(input)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "ID":
			out.Id = strings.ReplaceAll(v, `"`, "")
		case "CLASS":
			out.Class = strings.ReplaceAll(v, `"`, "")
		case "START-DATE":
			v = strings.ReplaceAll(v, `"`, "")
			date, err := time.Parse(time.RFC3339, v)
			if err != nil {
				continue
			}
			date = date.In(time.UTC)
			out.StartDate = &date
		case "END-DATE":
			v = strings.ReplaceAll(v, `"`, "")
			date, err := time.Parse(time.RFC3339, v)
			if err != nil {
				continue
			}
			date = date.In(time.UTC)
			out.EndDate = &date
		case "END-ON-NEXT":
			out.EndOnNext = v == "YES"
		case "DURATION":
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				out.Duration = 0
			}
			out.Duration = parsed
		case "PLANNED-DURATION":
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				out.PlannedDuration = 0
			}
			out.PlannedDuration = parsed

		case "SCTE35-CMD":
			out.Scte35Cmd = strings.ReplaceAll(v, `"`, "")
		case "SCTE35-OUT":
			out.Scte35Out = strings.ReplaceAll(v, `"`, "")
		case "SCTE35-IN":
			out.Scte35In = strings.ReplaceAll(v, `"`, "")
		}
	}

	return &out
}
