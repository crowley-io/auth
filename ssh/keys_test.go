package ssh

import (
	"golang.org/x/crypto/ssh"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPrivateKey(t *testing.T) {

	k, err := ReadPrivateKey("../files/host")

	assert.Nil(t, err)

	if !assert.NotNil(t, k) {
		t.Fatal(err)
	}

	assert.Implements(t, (*ssh.Signer)(nil), *k)
}

func TestReadPrivateKeyWithWrongPath(t *testing.T) {

	_, err := ReadPrivateKey("../files/id_rsa")

	assert.NotNil(t, err)
	t.Log(err)
}

func TestReadPrivateKeyWithCorruptedKey(t *testing.T) {

	_, err := ReadPrivateKey("../files/random")

	assert.NotNil(t, err)
	t.Log(err)
}

func TestAcceptPublicKey(t *testing.T) {

	k, err := ReadPrivateKey("../files/host")

	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	if !assert.NotNil(t, k) {
		t.FailNow()
	}

	pk := (*k).PublicKey()
	expected := string(pk.Marshal())

	p, err := AcceptPublicKey(nil, pk)

	assert.Nil(t, err)
	if assert.NotNil(t, p) {
		assert.NotEmpty(t, p.Extensions)
		s, ok := p.Extensions["pubkey"]
		assert.True(t, ok, "was expected a pubkey in extensions")
		assert.Equal(t, expected, s, "was expected another pubkey")
	}

}
