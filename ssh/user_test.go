package ssh

import (
	"testing"

	"encoding/base64"
	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {

	name := "crowley"

	user := User{
		name: name,
	}

	assert.Equal(t, name, user.Name())
}

func TestKey(t *testing.T) {

	s := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"

	b, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		t.Fatal(err)
	}

	user := User{
		key: b,
	}

	assert.Equal(t, s, user.Key())
}

func TestFingerprint(t *testing.T) {

	s := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	f := "bd:84:fe:a3:c3:aa:ba:1e:75:7c:38:e1:ea:db:cd:2b"

	b, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		t.Fatal(err)
	}

	user := User{
		key: b,
	}

	assert.Equal(t, f, user.Fingerprint())
}
