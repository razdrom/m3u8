package tag

import (
	"strconv"
	"strings"

	"github.com/razdrom/m3u8/scanner"
)

type StreamInfo struct {
	Bandwidth        int64
	AverageBandwidth int64
	Resolution       string
	Codecs           string
	Audio            string
	Video            string
	FrameRate        float64
	HdcpLevel        string
	Subtitles        string
	ClosedCaptions   string
}

func ParseStreamInfo(input string) *StreamInfo {
	out := StreamInfo{}

	hm := scanner.ScanArgs(input)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "AUDIO":
			out.Audio = strings.ReplaceAll(v, `"`, "")
		case "AVERAGE-BANDWIDTH":
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.AverageBandwidth = parsed
			}
		case "BANDWIDTH":
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.Bandwidth = parsed
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
		case "HDCP-LEVEL":
			out.HdcpLevel = v
		case "SUBTITLES":
			out.Subtitles = strings.ReplaceAll(v, `"`, "")
		case "CLOSED-CAPTIONS":
			out.ClosedCaptions = strings.ReplaceAll(v, `"`, "")
		}

	}

	return &out
}
