package test

import (
	"testing"

	"github.com/razdrom/m3u8/tag"
	"github.com/stretchr/testify/require"
)

func Test_ParseDateRange(t *testing.T) {
	src := "ID=splice-6FFFFFF0,CLASS=sample.class,START-DATE=2014-03-05T11:15:00Z,END-DATE=2014-03-06T11:15:00Z,DURATION=59.993,PLANNED-DURATION=60.100,SCTE35-CMD=0xFC002F0000000000FF,SCTE35-OUT=0xFC002F0000000000FF000014056FFFFFF00,SCTE35-IN=0xFC002F0000000000FF,END-ON-NEXT=YES"
	parsed := tag.ParseDateRange(src)
	require.NotNil(t, parsed)
	require.Equal(t, "splice-6FFFFFF0", parsed.Id)
	require.Equal(t, "sample.class", parsed.Class)
	require.Equal(t, "2014-03-05 11:15:00 +0000 UTC", parsed.StartDate.String())
	require.Equal(t, "2014-03-06 11:15:00 +0000 UTC", parsed.EndDate.String())
	require.Equal(t, 59.993, parsed.Duration)
	require.Equal(t, 60.100, parsed.PlannedDuration)
	require.Equal(t, true, parsed.EndOnNext)
	require.Equal(t, "0xFC002F0000000000FF", parsed.Scte35Cmd)
	require.Equal(t, "0xFC002F0000000000FF000014056FFFFFF00", parsed.Scte35Out)
	require.Equal(t, "0xFC002F0000000000FF", parsed.Scte35In)
}

func Test_ParseInfo(t *testing.T) {
	src := "13.682,live"
	parsed := tag.ParseInfo(src)
	require.NotNil(t, parsed)
	require.Equal(t, 13.682, parsed.Druation)
	require.Equal(t, "live", parsed.Title)

	src = "10"
	parsed = tag.ParseInfo(src)
	require.NotNil(t, parsed)
	require.Equal(t, float64(10), parsed.Druation)
	require.Equal(t, "", parsed.Title)
}

func Test_ParseMediaSequence(t *testing.T) {
	src := "1"
	parsed := tag.ParseMediaSequence(src)
	require.NotNil(t, parsed)
	require.Equal(t, int64(1), parsed.Value)
}

func Test_ParseMedia(t *testing.T) {
	src := `TYPE=VIDEO,URI="https://example.com/media?abc=AAAAAQAAOpgCAAHFYAaVF==",GROUP-ID="720p30",LANGUAGE="en",ASSOC-LANGUAGE="ge",NAME="720p",AUTOSELECT=YES,DEFAULT=YES,FORCED=NO,INSTREAM-ID="CC1",CHARACTERISTICS="public.accessibility.transcribes-spoken-dialog,public.easy-to-read",CHANNELS="6"`
	parsed := tag.ParseMedia(src)
	require.NotNil(t, parsed)
	require.Equal(t, "VIDEO", parsed.Type)
	require.Equal(t, "https://example.com/media?abc=AAAAAQAAOpgCAAHFYAaVF==", parsed.Uri)
	require.Equal(t, "720p30", parsed.GroupId)
	require.Equal(t, "en", parsed.Language)
	require.Equal(t, "ge", parsed.AssocLanguage)
	require.Equal(t, "720p", parsed.Name)
	require.Equal(t, true, parsed.Autoselect)
	require.Equal(t, true, parsed.Default)
	require.Equal(t, false, parsed.Forced)
	require.Equal(t, "CC1", parsed.InstreamId)
	require.ElementsMatch(t, []string{"public.accessibility.transcribes-spoken-dialog", "public.easy-to-read"}, parsed.Characteristics)
	require.Equal(t, "6", parsed.Channels)
}

func Test_ParseProgramDateTime(t *testing.T) {
	src := "2021-02-09T10:40:11.498Z"
	parsed := tag.ParseProgramDateTime(src)
	require.NotNil(t, parsed)
	require.Equal(t, "2021-02-09 10:40:11.498 +0000 UTC", parsed.Value.String())
}

func Test_ParseStreamInfo(t *testing.T) {
	src := `BANDWIDTH=6371345,AVERAGE-BANDWIDTH=341124,RESOLUTION=1920x1080,CODECS="avc1.640028,mp4a.40.2",VIDEO="chunked",AUDIO="chunked",FRAME-RATE=30.000,HDCP-LEVEL=TYPE-0,SUBTITLES="included",CLOSED-CAPTIONS=NONE`
	parsed := tag.ParseStreamInfo(src)
	require.NotNil(t, parsed)
	require.Equal(t, int64(341124), parsed.AverageBandwidth)
	require.Equal(t, int64(6371345), parsed.Bandwidth)
	require.Equal(t, "1920x1080", parsed.Resolution)
	require.Equal(t, "avc1.640028,mp4a.40.2", parsed.Codecs)
	require.Equal(t, "chunked", parsed.Video)
	require.Equal(t, "chunked", parsed.Audio)
	require.Equal(t, float64(30.000), parsed.FrameRate)
	require.Equal(t, "TYPE-0", parsed.HdcpLevel)
	require.Equal(t, "included", parsed.Subtitles)
	require.Equal(t, "NONE", parsed.ClosedCaptions)
}

func Test_ParseTargetDuration(t *testing.T) {
	src := "12"
	parsed := tag.ParseTargetDuration(src)
	require.NotNil(t, parsed)
	require.Equal(t, int64(12), parsed.Value)
}

func Test_ParseVersion(t *testing.T) {
	src := "3"
	parsed := tag.ParseVersion(src)
	require.NotNil(t, parsed)
	require.Equal(t, int64(3), parsed.Value)
}
