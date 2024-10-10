package tag

import (
	"m3u8/scanner"
	"strconv"
	"strings"
)

type StreamInfo struct {
	Bandwith   int64
	Resolution string
	Codecs     string
	Video      string
	FrameRate  float64
}

func ParseStreamInfo(input string) *StreamInfo {
	out := StreamInfo{}

	hm := scanner.ScanArgs(input)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "BANDWIDTH":
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.Bandwith = parsed
			}
		case "RESOLUTION":
			out.Resolution = v
		case "CODECS":
			out.Codecs = strings.ReplaceAll(v, `"`, "")
		case "VIDEO":
			out.Video = strings.ReplaceAll(v, `"`, "")
		case "FRAME-RATE":
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				out.FrameRate = parsed
			}
		}
	}

	return &out
}
