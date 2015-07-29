package mfa

import (
	"testing"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/pquerna/otp"
)

func TestRandomGenerate(t *testing.T) {

	key, err := Generate("crowley.io", "user@gmail.com")

	if !assert.Nil(t, err) {
		t.Fatalf("%+v", err)
	}

	if !assert.NotEmpty(t, key) {
		t.Fatalf("%+v", err)
	}

}

func TestErrorGenerate(t *testing.T) {

	e := errors.New("An error has occurred...")

	invoker := func(issuer string, user string) (*otp.Key, error) {
		return nil, e
	}

	key, err := generate("crowley.io", "user@gmail.com", invoker)

	assert.NotNil(t, err)
	assert.Equal(t, e, err)

	if !assert.Nil(t, key) {
		t.Fatalf("%+v", key)
	}

}

func TestGenerateKeyInfo(t *testing.T) {

	key, err := Generate("crowley.io", "user@gmail.com")
	
	if !assert.NotNil(t, key) {
		t.Fatalf("%+v", err)
	}

	assert.Equal(t, "crowley.io", key.Issuer())
	assert.Equal(t, "user@gmail.com", key.AccountName())
	assert.Equal(t, "totp", key.Type())
	assert.NotEmpty(t, key.Secret())

	t.Logf("%+v", key)
}