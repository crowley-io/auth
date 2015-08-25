package otp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomValidator(t *testing.T) {

	key, err := Generate("crowley.io", "user@gmail.com")

	if !assert.Nil(t, err) {
		t.Fatalf("%+v", err)
	}

	if !assert.NotEmpty(t, key) {
		t.Fatalf("%+v", err)
	}

	code, err := Encode(key.Secret())

	if !assert.Nil(t, err) {
		t.Fatalf("%+v", err)
	}

	if !assert.NotEmpty(t, code) {
		t.Fatalf("%+v", err)
	}

	sucess := Validate(code, key.Secret())

	if !assert.True(t, sucess) {
		t.Fatalf("%+v", sucess)
	}

}