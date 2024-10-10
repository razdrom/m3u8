package test

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/razdrom/m3u8"
	"github.com/razdrom/m3u8/tag"
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

	playlist := m3u8.MasterPlaylist{}
	decoder := m3u8.NewDecoder(&playlist)
	err = decoder.Decode(stream)
	require.NoError(t, err)

	variants := []m3u8.Variant{
		{URI: "https://example.com/v1/playlist/1920p30.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "chunked", Name: "1080p (source)", Autoselect: true, Default: true}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(6371345), Resolution: "1920x1080", Codecs: "avc1.640028,mp4a.40.2", Video: "chunked", FrameRate: float64(30.000)}},
		{URI: "https://example.com/v1/playlist/720p30.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "720p30", Name: "720p", Autoselect: true, Default: true}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(2373000), Resolution: "1280x720", Codecs: "avc1.4D401F,mp4a.40.2", Video: "720p30", FrameRate: float64(30.000)}},
		{URI: "https://example.com/v1/playlist/480p30.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "480p30", Name: "480p", Autoselect: true, Default: true}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(1427999), Resolution: "852x480", Codecs: "avc1.4D401F,mp4a.40.2", Video: "480p30", FrameRate: float64(30.000)}},
		{URI: "https://example.com/v1/playlist/360p30.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "360p30", Name: "360p", Autoselect: true, Default: true}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(630000), Resolution: "640x360", Codecs: "avc1.4D401F,mp4a.40.2", Video: "360p30", FrameRate: float64(30.000)}},
		{URI: "https://example.com/v1/playlist/160p30.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "160p30", Name: "160p", Autoselect: true, Default: true}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(230000), Resolution: "284x160", Codecs: "avc1.4D401F,mp4a.40.2", Video: "160p30", FrameRate: float64(30.000)}},
		{URI: "https://example.com/v1/playlist/audio_only.m3u8", Media: &tag.Media{Type: "VIDEO", GroupId: "audio_only", Name: "audio_only", Autoselect: false, Default: false}, StreamInfo: &tag.StreamInfo{Bandwidth: int64(160000), Resolution: "", Codecs: "mp4a.40.2", Video: "audio_only", FrameRate: float64(0.000)}},
	}

	require.Equal(t, len(variants), len(playlist.Variants))
	for i, variant := range playlist.Variants {
		require.Equal(t, variants[i].URI, variant.URI)
		require.Equal(t, variants[i].Media.Type, variant.Media.Type)
		require.Equal(t, variants[i].Media.GroupId, variant.Media.GroupId)
		require.Equal(t, variants[i].Media.Name, variant.Media.Name)
		require.Equal(t, variants[i].Media.Autoselect, variant.Media.Autoselect)
		require.Equal(t, variants[i].Media.Default, variant.Media.Default)
		require.Equal(t, variants[i].StreamInfo.Bandwidth, variant.StreamInfo.Bandwidth)
		require.Equal(t, variants[i].StreamInfo.Resolution, variant.StreamInfo.Resolution)
		require.Equal(t, variants[i].StreamInfo.Codecs, variant.StreamInfo.Codecs)
		require.Equal(t, variants[i].StreamInfo.Video, variant.StreamInfo.Video)
		require.Equal(t, variants[i].StreamInfo.FrameRate, variant.StreamInfo.FrameRate)
	}
}

func Test_Decode_Media_001(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	stream, err := os.Open(path.Join(dir, "assets", "simple-media.m3u8"))
	require.NoError(t, err)

	defer stream.Close()

	playlist := m3u8.MediaPlaylist{}
	decoder := m3u8.NewDecoder(&playlist)
	err = decoder.Decode(stream)
	require.NoError(t, err)

	dateranges := []tag.DateRange{
		{Id: "playlist-creation-1728470443", Class: "timestamp", StartDate: timeconv("2024-10-09T03:40:43.795-07:00"), EndOnNext: true},
		{Id: "playlist-session-1728470443", Class: "twitch-session", StartDate: timeconv("2024-10-09T03:40:43.795-07:00"), EndOnNext: true},
		{Id: "source-1728470409", Class: "twitch-stream-source", StartDate: timeconv("2024-10-09T10:40:09.498Z"), EndOnNext: true},
		{Id: "trigger-1728470409", Class: "twitch-trigger", StartDate: timeconv("2024-10-09T10:40:09.498Z"), EndOnNext: true},
	}

	segments := []m3u8.Segment{
		{URI: "https://example.com/v1/segment/segment_001.ts", ProgramDateTime: timeconv("2024-10-09T10:40:09.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_002.ts", ProgramDateTime: timeconv("2024-10-09T10:40:11.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_003.ts", ProgramDateTime: timeconv("2024-10-09T10:40:13.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_004.ts", ProgramDateTime: timeconv("2024-10-09T10:40:15.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_005.ts", ProgramDateTime: timeconv("2024-10-09T10:40:17.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_006.ts", ProgramDateTime: timeconv("2024-10-09T10:40:19.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_007.ts", ProgramDateTime: timeconv("2024-10-09T10:40:21.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_008.ts", ProgramDateTime: timeconv("2024-10-09T10:40:23.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_009.ts", ProgramDateTime: timeconv("2024-10-09T10:40:25.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_010.ts", ProgramDateTime: timeconv("2024-10-09T10:40:27.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_011.ts", ProgramDateTime: timeconv("2024-10-09T10:40:29.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_012.ts", ProgramDateTime: timeconv("2024-10-09T10:40:31.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_013.ts", ProgramDateTime: timeconv("2024-10-09T10:40:33.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_014.ts", ProgramDateTime: timeconv("2024-10-09T10:40:35.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
		{URI: "https://example.com/v1/segment/segment_015.ts", ProgramDateTime: timeconv("2024-10-09T10:40:37.498Z"), Info: &tag.Info{Druation: float64(2.000), Title: "live"}},
	}

	require.Equal(t, int64(3), playlist.Version)
	require.Equal(t, int64(6), playlist.TargetDuration)
	require.Equal(t, int64(60734), playlist.MediaSequence)

	require.Equal(t, len(dateranges), len(playlist.DateRanges))
	for i, rng := range playlist.DateRanges {
		require.Equal(t, dateranges[i].Id, rng.Id)
		require.Equal(t, dateranges[i].Class, rng.Class)
		require.Equal(t, dateranges[i].StartDate, rng.StartDate)
		require.Equal(t, dateranges[i].EndOnNext, rng.EndOnNext)
	}

	require.Equal(t, len(segments), len(playlist.Segments))
	for i, seg := range playlist.Segments {
		require.Equal(t, segments[i].URI, seg.URI)
		require.Equal(t, segments[i].ProgramDateTime, seg.ProgramDateTime)
		require.Equal(t, segments[i].Info.Druation, seg.Info.Druation)
		require.Equal(t, segments[i].Info.Title, seg.Info.Title)
	}
}
