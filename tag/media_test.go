package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseMedia(t *testing.T) {
	src := `TYPE=VIDEO,URI="https://example.com/media?abc=AAAAAQAAOpgCAAHFYAaVF==",GROUP-ID="720p30",LANGUAGE="en",ASSOC-LANGUAGE="ge",NAME="720p",AUTOSELECT=YES,DEFAULT=YES,FORCED=NO,INSTREAM-ID="CC1",CHARACTERISTICS="public.accessibility.transcribes-spoken-dialog,public.easy-to-read",CHANNELS="6"`
	tagimpl := NewMedia(src)
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
