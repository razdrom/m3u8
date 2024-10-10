package test

import (
	"strings"
	"testing"

	"github.com/razdrom/m3u8/scanner"
	"github.com/stretchr/testify/require"
)

func Test_Args_ScanArgs(t *testing.T) {
	data := []string{
		`AUDIO="aac"`,
		`AVERAGE-BANDWIDTH=6000000`,
		`BANDWIDTH=1727000`,
		`CODECS="mp4a.40.2,avc1.100.40"`,
		`DEFAULT=YES`,
		`ID="splice-6FFFFFF0"`,
		`FRAME-RATE=24`,
		`HDCP-LEVEL=TYPE-0`,
		`KEYFORMATVERSIONS="1/2/5"`,
		`LANGUAGE="eng"`,
		`LOCAL=2022-04-26T13:11:31.333300Z`,
		`METHOD=AES-128`,
		`MPEGTS=7984482175`,
		`NAME="abcde"`,
		`PLANNED-DURATION=59.993`,
		`PROGRAM-ID=1`,
		`RESOLUTION=1680x750`,
		`START-DATE="2014-03-05T11:15:00Z"`,
		`SCTE35-IN="/DAlAAAAAAAAAP/wFAUAAAABf+/+ANgNkv4AFJlwAAEBAQAA5xULLA=="`,
		`SCTE35-OUT=0xFC002F0000000000FF000014056FFFFFF000E011622DCAFF000052636200000000000A0008029896F50000008700000000`,
		`SUBTITLES="subtitles0"`,
		`URI="https://example.com/key?ecm=AAAAAQAAOpgCAAHFYAaVFH6QrFv2wYU1lEaO2L3fGQB1%2FR3oaD9auWtXNAmcVLxgRTvRlHpqHgXX1YY00%2FpdUiOlgONVbViqou2%2FItyDOWc%3D"`,
		`VALUE="Este es un ejemplo"`,
		`VIDEO-RANGE=SDR`,
	}

	results := map[string]string{
		"AUDIO":             `"aac"`,
		"AVERAGE-BANDWIDTH": "6000000",
		"BANDWIDTH":         "1727000",
		"CODECS":            `"mp4a.40.2,avc1.100.40"`,
		"DEFAULT":           "YES",
		"ID":                `"splice-6FFFFFF0"`,
		"FRAME-RATE":        "24",
		"HDCP-LEVEL":        "TYPE-0",
		"KEYFORMATVERSIONS": `"1/2/5"`,
		"LANGUAGE":          `"eng"`,
		"LOCAL":             "2022-04-26T13:11:31.333300Z",
		"METHOD":            "AES-128",
		"MPEGTS":            "7984482175",
		"NAME":              `"abcde"`,
		"PLANNED-DURATION":  "59.993",
		"PROGRAM-ID":        "1",
		"RESOLUTION":        "1680x750",
		"START-DATE":        `"2014-03-05T11:15:00Z"`,
		"SCTE35-IN":         `"/DAlAAAAAAAAAP/wFAUAAAABf+/+ANgNkv4AFJlwAAEBAQAA5xULLA=="`,
		"SCTE35-OUT":        "0xFC002F0000000000FF000014056FFFFFF000E011622DCAFF000052636200000000000A0008029896F50000008700000000",
		"SUBTITLES":         `"subtitles0"`,
		"URI":               `"https://example.com/key?ecm=AAAAAQAAOpgCAAHFYAaVFH6QrFv2wYU1lEaO2L3fGQB1%2FR3oaD9auWtXNAmcVLxgRTvRlHpqHgXX1YY00%2FpdUiOlgONVbViqou2%2FItyDOWc%3D"`,
		"VALUE":             `"Este es un ejemplo"`,
		"VIDEO-RANGE":       "SDR",
	}

	joinedArgs := strings.Join(data, ",")
	scans := scanner.ScanArgs(joinedArgs)
	for key, value := range scans {
		require.Equal(t, results[key], value)
	}
}
