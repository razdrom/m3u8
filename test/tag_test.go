package test

import (
	"testing"

	"github.com/razdrom/m3u8/tag"
	"github.com/stretchr/testify/require"
)

func Test_ParseDateRange(t *testing.T) {
	src := "ID=splice-6FFFFFF0,CLASS=sample.class,START-DATE=2014-03-05T11:15:00Z,END-DATE=2014-03-06T11:15:00Z,DURATION=59.993,PLANNED-DURATION=60.100,SCTE35-CMD=0xFC002F0000000000FF,SCTE35-OUT=0xFC002F0000000000FF000014056FFFFFF00,SCTE35-IN=0xFC002F0000000000FF,END-ON-NEXT=YES"
	tagimpl := tag.NewDateRange(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, "splice-6FFFFFF0", tagimpl.GetId())
	require.Equal(t, "sample.class", tagimpl.GetClass())
	require.Equal(t, "2014-03-05 11:15:00 +0000 UTC", tagimpl.GetStartDate().String())
	require.Equal(t, "2014-03-06 11:15:00 +0000 UTC", tagimpl.GetEndDate().String())
	require.Equal(t, 59.993, tagimpl.GetDuration())
	require.Equal(t, 60.100, tagimpl.GetPlannedDuration())
	require.Equal(t, true, tagimpl.GetEndOnNext())
	require.Equal(t, "0xFC002F0000000000FF", tagimpl.GetScte35Cmd())
	require.Equal(t, "0xFC002F0000000000FF000014056FFFFFF00", tagimpl.GetScte35Out())
	require.Equal(t, "0xFC002F0000000000FF", tagimpl.GetScte35In())
}

func Test_ParseInfo(t *testing.T) {
	src := "13.682,live"
	impl := tag.NewInfo(src)
	require.NotNil(t, impl)
	require.Equal(t, 13.682, impl.GetDuration())
	require.Equal(t, "live", impl.GetTitle())

	src = "10"
	impl = tag.NewInfo(src)
	require.NotNil(t, impl)
	require.Equal(t, float64(10), impl.GetDuration())
	require.Equal(t, "", impl.GetTitle())
}

func Test_ParseMediaSequence(t *testing.T) {
	src := "1"
	tagimpl := tag.NewMediaSequence(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(1), tagimpl.GetValue())
}

func Test_ParseMedia(t *testing.T) {
	src := `TYPE=VIDEO,URI="https://example.com/media?abc=AAAAAQAAOpgCAAHFYAaVF==",GROUP-ID="720p30",LANGUAGE="en",ASSOC-LANGUAGE="ge",NAME="720p",AUTOSELECT=YES,DEFAULT=YES,FORCED=NO,INSTREAM-ID="CC1",CHARACTERISTICS="public.accessibility.transcribes-spoken-dialog,public.easy-to-read",CHANNELS="6"`
	tagimpl := tag.NewMedia(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, "VIDEO", tagimpl.GetType())
	require.Equal(t, "https://example.com/media?abc=AAAAAQAAOpgCAAHFYAaVF==", tagimpl.GetUri())
	require.Equal(t, "720p30", tagimpl.GetGroupId())
	require.Equal(t, "en", tagimpl.GetLanguage())
	require.Equal(t, "ge", tagimpl.GetAssocLanguage())
	require.Equal(t, "720p", tagimpl.GetName())
	require.Equal(t, true, tagimpl.GetAutoselect())
	require.Equal(t, true, tagimpl.GetDefault())
	require.Equal(t, false, tagimpl.GetForced())
	require.Equal(t, "CC1", tagimpl.GetInstreamId())
	require.ElementsMatch(t, []string{"public.accessibility.transcribes-spoken-dialog", "public.easy-to-read"}, tagimpl.GetCharacteristics())
	require.Equal(t, "6", tagimpl.GetChannels())
}

func Test_ParseProgramDateTime(t *testing.T) {
	src := "2021-02-09T10:40:11.498Z"
	tagimpl := tag.NewProgramDateTime(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, "2021-02-09 10:40:11.498 +0000 UTC", tagimpl.GetValue().String())
}

func Test_ParseStreamInfo(t *testing.T) {
	src := `BANDWIDTH=6371345,AVERAGE-BANDWIDTH=341124,RESOLUTION=1920x1080,CODECS="avc1.640028,mp4a.40.2",VIDEO="chunked",AUDIO="chunked",FRAME-RATE=30.000,HDCP-LEVEL=TYPE-0,SUBTITLES="included",CLOSED-CAPTIONS=NONE`
	tagimpl := tag.NewStreamInfo(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(341124), tagimpl.GetAverageBandwidth())
	require.Equal(t, int64(6371345), tagimpl.GetBandwidth())
	require.Equal(t, int64(1920), tagimpl.GetResolution().Height)
	require.Equal(t, int64(1080), tagimpl.GetResolution().Width)
	require.Equal(t, "avc1.640028,mp4a.40.2", tagimpl.GetCodecs())
	require.Equal(t, "chunked", tagimpl.GetVideo())
	require.Equal(t, "chunked", tagimpl.GetAudio())
	require.Equal(t, float64(30.000), tagimpl.GetFrameRate())
	require.Equal(t, "TYPE-0", tagimpl.GetHdcpLevel())
	require.Equal(t, "included", tagimpl.GetSubtitles())
	require.Equal(t, "NONE", tagimpl.GetClosedCaptions())
}

func Test_ParseTargetDuration(t *testing.T) {
	src := "12"
	tagimpl := tag.NewTargetDuration(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(12), tagimpl.GetValue())
}

func Test_ParseVersion(t *testing.T) {
	src := "3"
	tagimpl := tag.NewVersion(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(3), tagimpl.GetValue())
}
