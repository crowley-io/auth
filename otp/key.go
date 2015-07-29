package otp

import (
	"image"

	"github.com/pquerna/otp"
)

const (

	// Default secret key length
	DefaultSecretKeyLength = 100

	// Default algorithm used in HMAC operation
	DefaultAlgorithm = "SHA512"

	// Default digits representation of user's OTP passcode.
	DefaultDigits = 6

	// Number of seconds a TOTP hash is valid for.
	DefaultPeriod = 30

	// Extra time step as allowed for latency and clock skew.
	DefaultSkew = 1
)

type Key struct {
	base *otp.Key
}

// String returns the Key's URI
func (k *Key) String() string {
	return k.base.String()
}

// Secret returns the opaque secret for this Key.
func (k *Key) Secret() string {
	return k.base.Secret()
}

// Issuer returns the name of the issuing organization.
func (k *Key) Issuer() string {
	return k.base.Issuer()
}

// Type returns "hotp" or "totp".
func (k *Key) Type() string {
	return k.base.Type()
}

// AccountName returns the name of the user's account.
func (k *Key) AccountName() string {
	return k.base.AccountName()
}

// Image returns an QR-Code image of the specified width and height.
func (k *Key) Image(width uint32, height uint32) (image.Image, error) {
	return k.base.Image(int(width), int(height))
}

// Code returns a one-time password from this Key.
func (k *Key) Code() string {
	v, _ := Encode(k.base.Secret())
	return v
}

// Validate verify that the given one-time password is correct.
func (k *Key) Validate(code string) bool {
	return Validate(code, k.base.Secret())
}
