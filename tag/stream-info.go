package tag

import (
	"strconv"
	"strings"

	"github.com/razdrom/m3u8/scanner"
)

type Resolution struct {
	Height int64
	Width  int64
}

type StreamInfo struct {
	raw       string
	rawparsed bool

	bandwidth        int64
	averageBandwidth int64
	resolution       *Resolution
	codecs           string
	audio            string
	video            string
	frameRate        float64
	hdcpLevel        string
	subtitles        string
	closedCaptions   string
}

func NewStreamInfo(raw string) *StreamInfo {
	return &StreamInfo{raw: raw}
}

func (t *StreamInfo) parse() {
	t.rawparsed = true
	hm := scanner.ScanArgs(t.raw)
	for k, v := range hm {
		if k == "" || v == "" {
			continue
		}

		switch k {
		case "AUDIO":
			t.audio = strings.ReplaceAll(v, `"`, "")
		case "AVERAGE-BANDWIDTH":
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				t.averageBandwidth = parsed
			}
		case "BANDWIDTH":
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				t.bandwidth = parsed
			}
		case "RESOLUTION":
			hw := strings.Split(v, "x")
			if len(hw) == 2 {
				height, err := strconv.ParseInt(hw[0], 10, 64)
				if err != nil {
					continue
				}
				width, err := strconv.ParseInt(hw[1], 10, 64)
				if err != nil {
					continue
				}

				t.resolution = &Resolution{Height: height, Width: width}
			}
		case "CODECS":
			t.codecs = strings.ReplaceAll(v, `"`, "")
		case "VIDEO":
			t.video = strings.ReplaceAll(v, `"`, "")
		case "FRAME-RATE":
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				t.frameRate = parsed
			}
		case "HDCP-LEVEL":
			t.hdcpLevel = v
		case "SUBTITLES":
			t.subtitles = strings.ReplaceAll(v, `"`, "")
		case "CLOSED-CAPTIONS":
			t.closedCaptions = strings.ReplaceAll(v, `"`, "")
		}
	}
}

func (t *StreamInfo) GetBandwidth() int64 {
	if !t.rawparsed {
		t.parse()
	}
	return t.bandwidth
}

func (t *StreamInfo) GetAverageBandwidth() int64 {
	if !t.rawparsed {
		t.parse()
	}
	return t.averageBandwidth
}

func (t *StreamInfo) GetResolution() *Resolution {
	if !t.rawparsed {
		t.parse()
	}
	return t.resolution
}

func (t *StreamInfo) GetCodecs() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.codecs
}

func (t *StreamInfo) GetAudio() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.audio
}

func (t *StreamInfo) GetVideo() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.video
}

func (t *StreamInfo) GetFrameRate() float64 {
	if !t.rawparsed {
		t.parse()
	}
	return t.frameRate
}

func (t *StreamInfo) GetHdcpLevel() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.hdcpLevel
}

func (t *StreamInfo) GetSubtitles() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.subtitles
}

func (t *StreamInfo) GetClosedCaptions() string {
	if !t.rawparsed {
		t.parse()
	}
	return t.closedCaptions
}
