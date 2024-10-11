package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseStreamInfo(t *testing.T) {
	src := `BANDWIDTH=6371345,AVERAGE-BANDWIDTH=341124,RESOLUTION=1920x1080,CODECS="avc1.640028,mp4a.40.2",VIDEO="chunked",AUDIO="chunked",FRAME-RATE=30.000,HDCP-LEVEL=TYPE-0,SUBTITLES="included",CLOSED-CAPTIONS=NONE`
	tagimpl := NewStreamInfo(src)
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
