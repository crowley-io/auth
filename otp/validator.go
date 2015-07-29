package otp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

// Validate a TOTP using the current time.
// A shortcut for ValidateCustom, Validate uses a configuration
// that is compatible with Google-Authenticator and most clients.
func Validate(code string, secret string) bool {
	v, _ := totp.ValidateCustom(
		code,
		secret,
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    DefaultPeriod,
			Skew:      DefaultSkew,
			Digits:    digits,
			Algorithm: algorithm,
		},
	)
	return v
}
