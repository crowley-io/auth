package otp

import (
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"hash"
	"math"
	"time"
)

// Encode a TOTP secret to generate a one-time password.
func Encode(secret string) (string, error) {

	unix := time.Now().UTC().Unix()
	value := unix / DefaultPeriod
	hash := algorithm.Hash
	bytes, err := base32.StdEncoding.DecodeString(secret)

	if err != nil {
		return "", err
	}

	hs, err := hs(hash, bytes, value)

	if err != nil {
		return "", err
	}

	dcb1 := dcb1(hs)
	dcb2 := dcb2(dcb1)
	code := code(dcb2)

	return code, nil
}

// Compute message authentication code as HS.
func hs(hash func() hash.Hash, secret []byte, value int64) ([]byte, error) {

	h := hmac.New(hash, secret)
	err := binary.Write(h, binary.BigEndian, value)

	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Compute Dynamic Binary Code #1 from HS.
func dcb1(hs []byte) []byte {

	offset := hs[len(hs)-1] & 0x0f
	dcb1 := hs[offset : offset+4]

	return dcb1
}

// Compute Dynamic Binary Code #2.
func dcb2(dcb1 []byte) uint32 {

	dcb2 := binary.BigEndian.Uint32(dcb1)
	dcb2 &= 0x7fffffff

	return dcb2
}

// Compute HTOP final value.
func code(dcb2 uint32) string {

	l := digits.Length()
	p := uint32(math.Pow10(l))

	v := dcb2 % p

	return digits.Format(int32(v))
}
