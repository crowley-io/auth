package mfa

import (
	"github.com/pquerna/otp/totp"
	"github.com/pquerna/otp"
)

func Generate(issuer string, user string) (*Key, error) {

	invoker := func(issuer string, user string) (*otp.Key, error) {
		return totp.Generate(totp.GenerateOpts{
			Issuer:      issuer,
			AccountName: user,
			Period:      DefaultPeriod,
			SecretSize:  DefaultSecretKeyLength,
			Algorithm:   DefaultAlgorithm,
			Digits:      DefaultDigits,
		})
	}

	return generate(issuer, user, invoker)
}

func generate(issuer string, user string, invoke generator) (*Key, error) {

	key, err := invoke(issuer, user)

	if err != nil {
		return nil, err
	}

	return &Key{key}, nil
}

type generator func(issuer string, user string) (*otp.Key, error)
