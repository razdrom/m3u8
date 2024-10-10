package tag

import (
	"m3u8/scanner"
	"strings"
	"time"
)

type DateRange struct {
	Id        string
	Class     string
	StartDate *time.Time
	EndOnNext bool
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
			date, err := time.Parse(time.RFC3339, v)
			if err != nil {
				continue
			}
			out.StartDate = &date
		case "END-ON-NEXT":
			out.EndOnNext = v == "YES"
		}
	}

	return &out
}
