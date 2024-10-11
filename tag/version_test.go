package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseVersion(t *testing.T) {
	src := "3"
	tagimpl := NewVersion(src)
	require.NotNil(t, tagimpl)
	require.Equal(t, int64(3), tagimpl.GetValue())
}
