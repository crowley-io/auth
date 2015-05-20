package mfa

import "errors"

var (
	ErrCreateRandSecret = errors.New("An error has occurred while creating random secret key")
	ErrSecretURIFormat  = errors.New("An error has occurred while creating secret URI")
	ErrQREncoding       = errors.New("An error has occurred while encoding QR code")
)
