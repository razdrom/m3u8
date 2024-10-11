package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseProgramDateTime(t *testing.T) {
	src := "2021-02-09T10:40:11.498Z"
	tagimpl := NewProgramDateTime(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, "2021-02-09 10:40:11.498 +0000 UTC", tagimpl.GetValue().String())
}
