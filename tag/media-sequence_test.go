package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseMediaSequence(t *testing.T) {
	src := "1"
	tagimpl := NewMediaSequence(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(1), tagimpl.GetValue())
}
