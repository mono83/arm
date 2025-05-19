package armhash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sha256dataProvider = []struct {
	Expected string
	Given    string
}{
	{Expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", Given: ""},
	{Expected: "36a9e7f1c95b82ffb99743e0c5c4ce95d83c9a430aac59f84ef3cbfab6145068", Given: " "},
	{Expected: "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969", Given: "Hello"},
	{Expected: "968cd037c3ec06e3cda58e60a0c015ece18887a42b1cf37bf5159d7f70f309fb", Given: "Привіт"},
}

func TestSHA256(t *testing.T) {
	for _, datum := range sha256dataProvider {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			if hash, err := String[[]byte](SHA256, datum.Given); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, fmt.Sprintf("%x", hash))
			}
			if hash, err := Bytes[[]byte](SHA256, []byte(datum.Given)); assert.NoError(t, err) {
				assert.Equal(t, datum.Expected, fmt.Sprintf("%x", hash))
			}
		})
	}
}
