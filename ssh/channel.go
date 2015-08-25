package ssh

import "golang.org/x/crypto/ssh"

//
// Wrapper for ssh.Channel to avoid package name conflict with golang.org/x/crypto/ssh
// and github.com/crowley-io/auth/ssh
//
type Channel struct {
	ssh.Channel
}
