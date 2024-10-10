package m3u8

import (
	"time"

	"github.com/razdrom/m3u8/tag"
)

type Segment struct {
	URI             string
	ProgramDateTime *time.Time
	Info            *tag.Info
}

type MediaPlaylist struct {
	BasePlaylist
	tmpsegment     *Segment
	TargetDuration int64
	MediaSequence  int64
	DateRanges     []tag.DateRange
	Segments       []Segment
}

func (pl *MediaPlaylist) ParseTag(key string, value string) {
	pl.MatchBaseTags(key, value)
	pl.MatchCommonTags(key, value)
}

func (pl *MediaPlaylist) MatchCommonTags(key string, value string) {
	switch key {
	// Media Playlist Tags
	// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.3
	case "EXTINF":
		if pl.tmpsegment == nil {
			pl.tmpsegment = &Segment{}
		}
		pl.tmpsegment.Info = tag.ParseInfo(value)
	case "EXT-X-BYTERANGE":
	case "EXT-X-DISCONTINUITY":
	case "EXT-X-KEY":
	case "EXT-X-MAP":
	case "EXT-X-PROGRAM-DATE-TIME":
		programDateTime := tag.ParseProgramDateTime(value)
		if programDateTime != nil && programDateTime.Value != nil {
			if pl.tmpsegment == nil {
				pl.tmpsegment = &Segment{}
			}
			pl.tmpsegment.ProgramDateTime = programDateTime.Value
		}
	case "EXT-X-DATERANGE":
		dr := tag.ParseDateRange(value)
		if dr != nil {
			pl.DateRanges = append(pl.DateRanges, *dr)
		}
	case "EXT-X-TARGETDURATION":
		targetDuration := tag.ParseTargetDuration(value)
		if targetDuration != nil && targetDuration.Value != 0 {
			pl.TargetDuration = targetDuration.Value
		}
	case "EXT-X-MEDIA-SEQUENCE":
		mediaSequence := tag.ParseMediaSequence(value)
		if mediaSequence != nil && mediaSequence.Value != 0 {
			pl.MediaSequence = mediaSequence.Value
		}
	case "EXT-X-DISCONTINUITY-SEQUENCE":
	case "EXT-X-ENDLIST":
	case "EXT-X-PLAYLIST-TYPE":
	case "EXT-X-I-FRAMES-ONLY":
	}
}

func (pl *MediaPlaylist) HandleUri(uri string) {
	if pl.tmpsegment == nil {
		return
	}

	pl.tmpsegment.URI = uri
	pl.Segments = append(pl.Segments, *pl.tmpsegment)
	pl.tmpsegment = nil
}
