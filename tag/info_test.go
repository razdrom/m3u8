package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseInfo(t *testing.T) {
	src := "13.682,live"
	impl := NewInfo(src)
	require.NotNil(t, impl)
	require.Equal(t, 13.682, impl.GetDuration())
	require.Equal(t, "live", impl.GetTitle())

	src = "10"
	impl = NewInfo(src)
	require.NotNil(t, impl)
	require.Equal(t, float64(10), impl.GetDuration())
	require.Equal(t, "", impl.GetTitle())
}
