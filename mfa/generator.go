package mfa

import (
	"github.com/pquerna/otp/totp"
)

func Generate(issuer string, user string) (*Key, error) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: user,
		Period:      DefaultPeriod,
		SecretSize:  DefaultSecretKeyLength,
		Algorithm:   DefaultAlgorithm,
		Digits:      DefaultDigits,
	})

	if err != nil {
		return nil, err
	}

	return &Key{key}, nil
}
