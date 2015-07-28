package mfa

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"regexp"
)

// Default secret key length
const DefaultSecretKeyLength = 100

// Lambda which return a base32 secret an arbitrary key value encoded in Base32 according to RFC 3548
// (http://tools.ietf.org/html/rfc3548).
type SecretKeyGenerator func() (string, error)

// Lazy regexp that verify if secret key for 2FA seems correct.
var (
	keyRegExp = regexp.MustCompile("^otpauth:\\/\\/totp\\/[!-~]+:[!-~]+?secret=[\\w]+=*&issuer=[!-~]+$")
)

func AuthKey(issuer string, user string, generate SecretKeyGenerator) (string, error) {

	secret, err := generate()

	if err != nil {
		return "", err
	}

	auth := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, user, secret, issuer)

	if !keyRegExp.MatchString(auth) {
		return "", ErrSecretURIFormat
	}

	return auth, nil
}

// An implementation of SecretKeyGenerator.
func RandomBase32() (string, error) {

	// Generate random secret
	buffer := make([]byte, DefaultSecretKeyLength)
	_, err := rand.Read(buffer)

	if err != nil {
		return "", ErrCreateRandSecret
	}

	// Encode secret to base32 string
	secret := base32.StdEncoding.EncodeToString(buffer)

	return secret, nil
}
