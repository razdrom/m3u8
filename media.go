package m3u8

import (
	"github.com/razdrom/m3u8/tag"
)

type Segment struct {
	URI             string
	ProgramDateTime *tag.ProgramDateTime
	Info            *tag.Info
}

type MediaPlaylist struct {
	BasePlaylist
	tmpsegment     *Segment
	TargetDuration *tag.TargetDuration
	MediaSequence  *tag.MediaSequence
	DateRanges     []tag.DateRange
	Segments       []Segment
	EndList        bool
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
		pl.tmpsegment.Info = tag.NewInfo(value)
	case "EXT-X-BYTERANGE":
	case "EXT-X-DISCONTINUITY":
	case "EXT-X-KEY":
	case "EXT-X-MAP":
	case "EXT-X-PROGRAM-DATE-TIME":
		if pl.tmpsegment == nil {
			pl.tmpsegment = &Segment{}
		}
		pl.tmpsegment.ProgramDateTime = tag.NewProgramDateTime(value)
	case "EXT-X-DATERANGE":
		dateRange := tag.NewDateRange(value)
		pl.DateRanges = append(pl.DateRanges, *dateRange)
	case "EXT-X-TARGETDURATION":
		pl.TargetDuration = tag.NewTargetDuration(value)
	case "EXT-X-MEDIA-SEQUENCE":
		pl.MediaSequence = tag.NewMediaSequence(value)
	case "EXT-X-DISCONTINUITY-SEQUENCE":
	case "EXT-X-ENDLIST":
		pl.EndList = true
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
