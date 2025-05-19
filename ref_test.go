package arm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRef(t *testing.T) {
	s := "Hello, world!"
	sr := &s
	require.Equal(t, sr, Ref("Hello, world!"))
	require.NotEqual(t, sr, Ref("Hello, world"))

	i := 12345
	ir := &i
	require.Equal(t, ir, Ref(i))
	require.NotEqual(t, ir, Ref(54321))
}
