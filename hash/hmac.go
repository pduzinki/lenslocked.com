package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// NewHMAC creates and returns a new HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// HMAC is a wrapper around the crypto/hmac package
type HMAC struct {
	hmac hash.Hash
}

// Hash will hash the provided string using HMAC
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
