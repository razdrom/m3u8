package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseDateRange(t *testing.T) {
	src := "ID=splice-6FFFFFF0,CLASS=sample.class,START-DATE=2014-03-05T11:15:00Z,END-DATE=2014-03-06T11:15:00Z,DURATION=59.993,PLANNED-DURATION=60.100,SCTE35-CMD=0xFC002F0000000000FF,SCTE35-OUT=0xFC002F0000000000FF000014056FFFFFF00,SCTE35-IN=0xFC002F0000000000FF,END-ON-NEXT=YES"
	tagimpl := NewDateRange(src)
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
