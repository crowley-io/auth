package mfa

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomAuthKey(t *testing.T) {

	auth, err := AuthKey("crowley.io", "user@gmail.com", RandomBase32)

	assert.Nil(t, err)

	if !assert.NotEmpty(t, auth) {
		t.Fatalf("%+v", err)
	}

	u, err := url.Parse(auth)

	t.Logf("%+v", auth)
	t.Logf("%+v", u)

	if !assert.Empty(t, err) {
		t.Fatalf("%+v", err)
	}
}

func TestErrorAuthKey(t *testing.T) {

	auth, err := AuthKey("crowley.io", "user@gmail.com", func() (string, error) {
		return "foobar", ErrCreateRandSecret
	})

	assert.NotNil(t, err)
	assert.Equal(t, ErrCreateRandSecret, err)

	if !assert.Empty(t, auth) {
		t.Fatalf("%+v", auth)
	}

}

func TestSpecificAuthKey(t *testing.T) {

	auth, err := AuthKey("crowley.io", "user@gmail.com", func() (string, error) {
		return "JBSWY3DPEHPK3PXP", nil
	})

	assert.Nil(t, err)

	if !assert.NotEmpty(t, auth) {
		t.Fatalf("%+v", err)
	}

	assert.Contains(t, auth, "crowley.io")
	assert.Contains(t, auth, "user@gmail.com")
	assert.Contains(t, auth, "JBSWY3DPEHPK3PXP")

	t.Logf("%+v", auth)
}

func TestErrorAuthKeyIssuer(t *testing.T) {

	auth, err := AuthKey("crowley io", "user@gmail.com", RandomBase32)

	assert.NotNil(t, err)
	assert.Equal(t, ErrSecretURIFormat, err)

	if !assert.Empty(t, auth) {
		t.Fatalf("%+v", auth)
	}

}
