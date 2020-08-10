package guid

import (
	"crypto/rand"
	"encoding/base64"
)

func UrlGuid() (string, error) {
	g := new([4]byte)
	if _, err := rand.Read(g[:]); err != nil {
		return "", err
	}
	g[1] = (g[1] & 0x0f) | 0x40
	g[3] = (g[3] & 0x3f) | 0x80

	return base64.StdEncoding.EncodeToString(g[:]), nil
}
