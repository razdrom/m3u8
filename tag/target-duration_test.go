package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseTargetDuration(t *testing.T) {
	src := "12"
	tagimpl := NewTargetDuration(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(12), tagimpl.GetValue())
}
