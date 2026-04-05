package armhash

import (
	"hash/crc32"
	"io"
)

// CRC32 calculates IEEE CRC32 checksum
// IEEE is by far and away the most common CRC-32 polynomial.
// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, ...
// Produces same result as MySQL and PHP.
func CRC32(r io.Reader) (uint32, error) {
	h := crc32.NewIEEE()
	if _, err := io.Copy(h, r); err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}
