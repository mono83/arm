package armhash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var crc32dataProvider = []struct {
	Expected uint32
	Given    string
}{
	{Expected: 0, Given: ""},
	{Expected: 3916222277, Given: " "},
	{Expected: 4157704578, Given: "Hello"},
	{Expected: 710749765, Given: "Привіт"},
}

func TestCRC32(t *testing.T) {
	for _, datum := range crc32dataProvider {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			if hash, err := String[uint32](CRC32, datum.Given); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, hash)
			}
			if hash, err := Bytes[uint32](CRC32, []byte(datum.Given)); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, hash)
			}
		})
	}
}
