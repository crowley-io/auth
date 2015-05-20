package mfa

import (
	"code.google.com/p/rsc/qr"
	// "github.com/boombuler/barcode/qr"
)

type QRCode []byte

func Encode(key string) (QRCode, error) {

	// Encode the QR image
	code, err := qr.Encode(key, qr.M)

	if err != nil {
		return nil, ErrQREncoding
	}

	return code.PNG(), nil
}
