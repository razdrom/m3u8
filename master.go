package m3u8

import "m3u8/tag"

type Variant struct {
	URI        string
	Media      *tag.Media
	StreamInfo *tag.StreamInfo
}

type MasterPlaylist struct {
	BasePlaylist
	tmpvariant *Variant
	Variants   []Variant
}

func (pl *MasterPlaylist) ParseTag(key string, value string) {
	pl.MatchBaseTags(key, value)
	pl.MatchMasterTags(key, value)
}

func (pl *MasterPlaylist) MatchMasterTags(key string, value string) {
	switch key {
	// Master Playlist Tags
	// https://datatracker.ietf.org/doc/html/rfc8216#section-4.3.4
	case "EXT-X-MEDIA":
		if pl.tmpvariant == nil {
			pl.tmpvariant = &Variant{}
		}
		pl.tmpvariant.Media = tag.ParseMedia(value)
	case "EXT-X-STREAM-INF":
		if pl.tmpvariant == nil {
			pl.tmpvariant = &Variant{}
		}
		pl.tmpvariant.StreamInfo = tag.ParseStreamInfo(value)

	case "EXT-X-I-FRAME-STREAM-INF":
	case "EXT-X-SESSION-DATA":
	case "EXT-X-SESSION-KEY":
	default:

	}
}

func (pl *MasterPlaylist) HandleUri(uri string) {
	if pl.tmpvariant == nil {
		return
	}

	pl.tmpvariant.URI = uri
	pl.Variants = append(pl.Variants, *pl.tmpvariant)
	pl.tmpvariant = nil
}
