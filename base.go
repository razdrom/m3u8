package m3u8

import "github.com/razdrom/m3u8/tag"

type Playlist interface {
	ParseTag(key, value string)
	HandleUri(uri string)
}

type BasePlaylist struct {
	Version int64
}

func (pl *BasePlaylist) MatchBaseTags(key string, value string) {
	switch key {
	// Basic tags
	// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.1
	case "EXTM3U":
	case "EXT-X-VERSION":
		version := tag.ParseVersion(value)
		if version != nil && version.Value != 0 {
			pl.Version = version.Value
		}

	// Media or Master Playlist Tags
	// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.5
	case "EXT-X-INDEPENDENT-SEGMENTS":
	case "EXT-X-START":
	}
}
