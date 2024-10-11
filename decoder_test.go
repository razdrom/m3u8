package m3u8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Decoder_splitTag(t *testing.T) {
	data := []string{
		`#EXTM3U`,
		`#EXT-X-VERSION:3`,
		`#EXT-X-MEDIA-SEQUENCE:60734`,
		`#EXT-X-DATERANGE:ID="trigger-1728470409",CLASS="twitch-trigger",START-DATE="2024-10-09T10:40:09.498Z",END-ON-NEXT=YES,X-TV-TWITCH-TRIGGER-URL="https://video-weaver.muc01.hls.ttvnw.net/trigger/CqkEoO_AtI-IEPP3-2ftMNVgLPqfawuz5wthE_s18Kc4WyuZgRf3a0NdgsCFC3LDHlm_DGHCan-eJH8RWnxV37Tx9hkDpS_HqOBAG89G-eO7iJt6NYIT_g8scdYn32OwseskSNxnel4z7PnX-_XUsuWN2hvy1WQAN4U5rZqWJRfsbJyxqVH-03_DIJi6xrNUrXwqxzHTTGqljmCJptf1zoekJfeGqyHqucHniCUf8NjWBwdPGE7vR1ddW-nieXthAlcwtP90nXuzZJqjCvD41BcspBSrMdIm5Cw-vjQzgNjsAPZS4PLxVV5h0uxtWXqne_r0Wb_HuYM4JyHkhHY6W-DMbFbxysUxnoWyHW6PkDvnskmw-1bNPKZiow8ybeQyC8Ak59lLLkhBBwzhT44BKANdoEBl5rv4WmeFR8NgwtLWNieRBG8NzQQPSOTKvWqNoivdY77B46mZT0IIyhHT3mdzUgcrSfrK2eu5iZcFzKDiOiguXnG6U8o0aLTb18lihxVfkLMaMllG3SWDiyhPhD1wrKcczsiOk6uK3COMGb84tcqqbR_1_B_Fd9do0ZCf6ILHbo-PS5xR22OWuO6VydRZEfC2OOsuAcknyf9LZtYSECdhfIWG6qULZ8SzhtQol3uchG6JkjXh33hs75JWKpLcgzOWFyCxZXilA_rOSZawohaw78aT_ItLmCZoIw90A67PBoMBr32YCAQxPlkOGnIEwwXuFqauNv6EchoMPbXmjwz0RYJuoCOTIAEqCWV1LXdlc3QtMjDKCg"`,
		`#EXT-X-PROGRAM-DATE-TIME:2024-10-09T10:40:09.498Z`,
		`#EXTINF:2.000,live`,
		`#EXT-X-MEDIA:TYPE=VIDEO,GROUP-ID="chunked",NAME="1080p (source)",AUTOSELECT=YES,DEFAULT=YES`,
		`#EXT-X-STREAM-INF:BANDWIDTH=6371345,RESOLUTION=1920x1080,CODECS="avc1.640028,mp4a.40.2",VIDEO="chunked",FRAME-RATE=30.000`,
	}

	results := map[string]string{
		"EXTM3U":                  "",
		"EXT-X-VERSION":           "3",
		"EXT-X-MEDIA-SEQUENCE":    "60734",
		"EXT-X-DATERANGE":         `ID="trigger-1728470409",CLASS="twitch-trigger",START-DATE="2024-10-09T10:40:09.498Z",END-ON-NEXT=YES,X-TV-TWITCH-TRIGGER-URL="https://video-weaver.muc01.hls.ttvnw.net/trigger/CqkEoO_AtI-IEPP3-2ftMNVgLPqfawuz5wthE_s18Kc4WyuZgRf3a0NdgsCFC3LDHlm_DGHCan-eJH8RWnxV37Tx9hkDpS_HqOBAG89G-eO7iJt6NYIT_g8scdYn32OwseskSNxnel4z7PnX-_XUsuWN2hvy1WQAN4U5rZqWJRfsbJyxqVH-03_DIJi6xrNUrXwqxzHTTGqljmCJptf1zoekJfeGqyHqucHniCUf8NjWBwdPGE7vR1ddW-nieXthAlcwtP90nXuzZJqjCvD41BcspBSrMdIm5Cw-vjQzgNjsAPZS4PLxVV5h0uxtWXqne_r0Wb_HuYM4JyHkhHY6W-DMbFbxysUxnoWyHW6PkDvnskmw-1bNPKZiow8ybeQyC8Ak59lLLkhBBwzhT44BKANdoEBl5rv4WmeFR8NgwtLWNieRBG8NzQQPSOTKvWqNoivdY77B46mZT0IIyhHT3mdzUgcrSfrK2eu5iZcFzKDiOiguXnG6U8o0aLTb18lihxVfkLMaMllG3SWDiyhPhD1wrKcczsiOk6uK3COMGb84tcqqbR_1_B_Fd9do0ZCf6ILHbo-PS5xR22OWuO6VydRZEfC2OOsuAcknyf9LZtYSECdhfIWG6qULZ8SzhtQol3uchG6JkjXh33hs75JWKpLcgzOWFyCxZXilA_rOSZawohaw78aT_ItLmCZoIw90A67PBoMBr32YCAQxPlkOGnIEwwXuFqauNv6EchoMPbXmjwz0RYJuoCOTIAEqCWV1LXdlc3QtMjDKCg"`,
		"EXT-X-PROGRAM-DATE-TIME": "2024-10-09T10:40:09.498Z",
		"EXTINF":                  "2.000,live",
		"EXT-X-MEDIA":             `TYPE=VIDEO,GROUP-ID="chunked",NAME="1080p (source)",AUTOSELECT=YES,DEFAULT=YES`,
		"EXT-X-STREAM-INF":        `BANDWIDTH=6371345,RESOLUTION=1920x1080,CODECS="avc1.640028,mp4a.40.2",VIDEO="chunked",FRAME-RATE=30.000`,
	}

	decoder := NewDecoder(nil)
	for _, item := range data {
		key, value := decoder.SplitTag([]byte(item))
		require.Equal(t, results[key], value)
	}
}

func Test_Decoder_IsBlank(t *testing.T) {
	decoder := NewDecoder(nil)
	require.Equal(t, false, decoder.IsBlank([]byte("#EXTM3U")))
	require.Equal(t, false, decoder.IsBlank([]byte("https://example.com.hls.net/playlist/Cv4EtbciG9p.m3u8")))
	require.Equal(t, true, decoder.IsBlank([]byte("")))
}

func Test_Decoder_IsTag(t *testing.T) {
	decoder := NewDecoder(nil)
	require.Equal(t, true, decoder.IsTag([]byte("#EXTM3U")))
	require.Equal(t, false, decoder.IsTag([]byte("https://example.com.hls.net/playlist/Cv4EtbciG9p.m3u8")))
	require.Equal(t, false, decoder.IsTag([]byte("")))
}

func Test_Decoder_IsURI(t *testing.T) {
	decoder := NewDecoder(nil)
	require.Equal(t, false, decoder.IsURI([]byte("#EXTM3U")))
	require.Equal(t, true, decoder.IsURI([]byte("https://example.com.hls.net/playlist/Cv4EtbciG9p.m3u8")))
	require.Equal(t, false, decoder.IsURI([]byte("")))
}
