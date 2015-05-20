package ssh

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// Read a ssh private key and return a Signer instance.
func ReadPrivateKey(path string) (*ssh.Signer, error) {

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(bytes)

	if err != nil {
		return nil, err
	}

	return &signer, nil
}

// Lambda used by ServerConfig for PublicKeyCallback.
// We recommend to use SSHHandler interface if you need to manage ssh connection...
func AcceptPublicKey(c ssh.ConnMetadata, k ssh.PublicKey) (*ssh.Permissions, error) {

	p := &ssh.Permissions{
		Extensions: map[string]string{"pubkey": string(k.Marshal())},
	}

	return p, nil
}
