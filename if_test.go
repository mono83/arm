package arm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIf(t *testing.T) {
	require.Equal(t, "foo", If(true, "foo", "bar"))
	require.Equal(t, "bar", If(false, "foo", "bar"))
}
