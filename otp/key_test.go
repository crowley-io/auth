package otp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {

	key, err := Generate("crowley.io", "user@gmail.com")

	if !assert.Nil(t, err) {
		t.Fatalf("%+v", err)
	}

	if !assert.NotEmpty(t, key) {
		t.Fatalf("%+v", err)
	}

	code := key.Code()

	if !assert.NotEmpty(t, code) {
		t.Fatalf("%+v", err)
	}

	success := key.Validate(code)

	if !assert.True(t, success) {
		t.Fatalf("%+v", success)
	}

}