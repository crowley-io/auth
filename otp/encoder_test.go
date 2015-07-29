package otp

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeWithZeroPadding(t *testing.T) {

	e := "012345"
	v := uint32(12345)
	s := code(v)

	if !assert.Equal(t, e, s) {
		t.Fatalf("%+v", s)
	}

}

func TestCodeWithModuloL6(t *testing.T) {

	e := "456789"
	v := uint32(123456789)
	s := code(v)

	if !assert.Equal(t, e, s) {
		t.Fatalf("%+v", s)
	}

}

func TestEncoderWithNoBase32Secret(t *testing.T) {

	s := "ef8:FwV$KDY_"

	v, err := Encode(s)

	if !assert.Empty(t, v) {
		t.Fatalf("%+v", v)
	}

	assert.NotNil(t, err)
}

func TestEncoderCodeMatch(t *testing.T) {

	s := "62DAVEG4ABFZ7WHR"
	e := regexp.MustCompile("^[0-9]{6}$")

	v, err := Encode(s)

	if !assert.NotEmpty(t, v) {
		t.Fatalf("%+v", err)
	}

	assert.Nil(t, err)
	assert.True(t, e.MatchString(v))

}
