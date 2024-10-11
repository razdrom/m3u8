package m3u8

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func timeconv(datestring string) *time.Time {
	val, _ := time.Parse(time.RFC3339, datestring)
	val = val.In(time.UTC)
	return &val
}

func Test_Decode_Master_001(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	stream, err := os.Open(path.Join(dir, "assets", "simple-master.m3u8"))
	require.NoError(t, err)

	defer stream.Close()

	playlist := MasterPlaylist{}
	decoder := NewDecoder(&playlist)
	err = decoder.Decode(stream)
	require.NoError(t, err)

	variants := []map[string]map[string]any{
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/1920p30.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "chunked", "name": "1080p (source)", "isAutoselect": true, "isDefault": true},
			"stream": map[string]any{"bandwidth": int64(6371345), "resolution": "1920x1080", "codecs": "avc1.640028,mp4a.40.2", "video": "chunked", "frameRate": float64(30.000)},
		},
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/720p30.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "720p30", "name": "720p", "isAutoselect": true, "isDefault": true},
			"stream": map[string]any{"bandwidth": int64(2373000), "resolution": "1280x720", "codecs": "avc1.4D401F,mp4a.40.2", "video": "720p30", "frameRate": float64(30.000)},
		},
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/480p30.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "480p30", "name": "480p", "isAutoselect": true, "isDefault": true},
			"stream": map[string]any{"bandwidth": int64(1427999), "resolution": "852x480", "codecs": "avc1.4D401F,mp4a.40.2", "video": "480p30", "frameRate": float64(30.000)},
		},
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/360p30.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "360p30", "name": "360p", "isAutoselect": true, "isDefault": true},
			"stream": map[string]any{"bandwidth": int64(630000), "resolution": "640x360", "codecs": "avc1.4D401F,mp4a.40.2", "video": "360p30", "frameRate": float64(30.000)},
		},
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/160p30.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "160p30", "name": "160p", "isAutoselect": true, "isDefault": true},
			"stream": map[string]any{"bandwidth": int64(230000), "resolution": "284x160", "codecs": "avc1.4D401F,mp4a.40.2", "video": "160p30", "frameRate": float64(30.000)},
		},
		{
			"uri":    map[string]any{"value": "https://example.com/v1/playlist/audio_only.m3u8"},
			"media":  map[string]any{"mediaType": "VIDEO", "groupId": "audio_only", "name": "audio_only", "isAutoselect": false, "isDefault": false},
			"stream": map[string]any{"bandwidth": int64(160000), "resolution": "", "codecs": "mp4a.40.2", "video": "audio_only", "frameRate": float64(0.000)},
		},
	}

	require.Equal(t, len(variants), len(playlist.Variants))

	for i, variant := range playlist.Variants {
		uri := variants[i]["uri"]
		media := variants[i]["media"]
		stream := variants[i]["stream"]

		var resolution string = ""
		hw := variant.StreamInfo.GetResolution()
		if hw != nil {
			resolution = fmt.Sprintf("%dx%d", hw.Height, hw.Width)
		}

		require.Equal(t, uri["value"], variant.URI)
		require.Equal(t, media["mediaType"], variant.Media.GetType())
		require.Equal(t, media["groupId"], variant.Media.GetGroupId())
		require.Equal(t, media["name"], variant.Media.GetName())
		require.Equal(t, media["isAutoselect"], variant.Media.GetAutoselect())
		require.Equal(t, media["isDefault"], variant.Media.GetDefault())
		require.Equal(t, stream["bandwidth"], variant.StreamInfo.GetBandwidth())
		require.Equal(t, stream["resolution"], resolution)
		require.Equal(t, stream["codecs"], variant.StreamInfo.GetCodecs())
		require.Equal(t, stream["video"], variant.StreamInfo.GetVideo())
		require.Equal(t, stream["frameRate"], variant.StreamInfo.GetFrameRate())
	}
}

func Test_Decode_Media_001(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	stream, err := os.Open(path.Join(dir, "assets", "simple-media.m3u8"))
	require.NoError(t, err)

	defer stream.Close()

	playlist := MediaPlaylist{}
	decoder := NewDecoder(&playlist)
	err = decoder.Decode(stream)
	require.NoError(t, err)

	require.Equal(t, int64(3), playlist.Version.GetValue())
	require.Equal(t, int64(6), playlist.TargetDuration.GetValue())
	require.Equal(t, int64(60734), playlist.MediaSequence.GetValue())

	dateranges := []map[string]string{
		{"id": "playlist-creation-1728470443", "class": "timestamp", "startDate": "2024-10-09T03:40:43.795-07:00", "endOnNext": "YES"},
		{"id": "playlist-session-1728470443", "class": "twitch-session", "startDate": "2024-10-09T03:40:43.795-07:00", "endOnNext": "YES"},
		{"id": "source-1728470409", "class": "twitch-stream-source", "startDate": "2024-10-09T10:40:09.498Z", "endOnNext": "YES"},
		{"id": "trigger-1728470409", "class": "twitch-trigger", "startDate": "2024-10-09T10:40:09.498Z", "endOnNext": "YES"},
	}

	require.Equal(t, len(dateranges), len(playlist.DateRanges))
	for i, rng := range playlist.DateRanges {
		startDate := timeconv(dateranges[i]["startDate"])
		endOnNext := dateranges[i]["endOnNext"] == "YES"
		require.Equal(t, dateranges[i]["id"], rng.GetId())
		require.Equal(t, dateranges[i]["class"], rng.GetClass())
		require.Equal(t, startDate, rng.GetStartDate())
		require.Equal(t, endOnNext, rng.GetEndOnNext())
	}

	segments := []map[string]any{
		{"uri": "https://example.com/v1/segment/segment_001.ts", "programDateTime": "2024-10-09T10:40:09.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_002.ts", "programDateTime": "2024-10-09T10:40:11.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_003.ts", "programDateTime": "2024-10-09T10:40:13.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_004.ts", "programDateTime": "2024-10-09T10:40:15.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_005.ts", "programDateTime": "2024-10-09T10:40:17.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_006.ts", "programDateTime": "2024-10-09T10:40:19.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_007.ts", "programDateTime": "2024-10-09T10:40:21.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_008.ts", "programDateTime": "2024-10-09T10:40:23.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_009.ts", "programDateTime": "2024-10-09T10:40:25.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_010.ts", "programDateTime": "2024-10-09T10:40:27.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_011.ts", "programDateTime": "2024-10-09T10:40:29.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_012.ts", "programDateTime": "2024-10-09T10:40:31.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_013.ts", "programDateTime": "2024-10-09T10:40:33.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_014.ts", "programDateTime": "2024-10-09T10:40:35.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
		{"uri": "https://example.com/v1/segment/segment_015.ts", "programDateTime": "2024-10-09T10:40:37.498Z", "infoDuration": float64(2.000), "infoTitle": "live"},
	}

	require.Equal(t, len(segments), len(playlist.Segments))
	for i, seg := range playlist.Segments {
		programDateTime := timeconv(segments[i]["programDateTime"].(string))
		duration := segments[i]["infoDuration"].(float64)
		require.Equal(t, segments[i]["uri"], seg.URI)
		require.Equal(t, programDateTime, seg.ProgramDateTime.GetValue())
		require.Equal(t, duration, seg.Info.GetDuration())
		require.Equal(t, segments[i]["infoTitle"], seg.Info.GetTitle())
	}
}
