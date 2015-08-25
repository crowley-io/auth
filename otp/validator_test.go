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

	success := Validate(code, key.Secret())

	if !assert.True(t, success) {
		t.Fatalf("%+v", success)
	}

}
