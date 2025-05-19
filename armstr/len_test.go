package armstr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lenDataProvider = []struct {
	Expected       int
	StandardLength int
	Given          string
}{
	{0, 0, ""},
	{1, 1, " "},
	{1, 1, "\n"},
	{1, 1, "Q"},
	{1, 2, "Ї"},
	{1, 4, "🔥"},
	{2, 6, "❤️"},    // Some emojis uses two runes
	{4, 14, "🏳️‍🌈"}, // Some emojis uses even more runes
	{12, 20, "Hello ❤Київ❤"},
}

func TestLen(t *testing.T) {
	for _, datum := range lenDataProvider {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			assert.Equal(t, datum.StandardLength, len(datum.Given))
			assert.Equal(t, datum.Expected, Len(datum.Given))
		})
	}
}
