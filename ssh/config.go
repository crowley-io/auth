package ssh

import "golang.org/x/crypto/ssh"

/**
 * Interface used to handle new user connection.
 */
type SSHHandler interface {
	OnConnect(channel ssh.Channel, user User)
}

/**
 * Configuration for SSHServer.
 */
type SSHConfig struct {

	/**
	 * Port configuration.
	 */
	port uint64

	/**
	 * Host's ssh private key path.
	 */
	path string

	/**
	 * Server behavior on new user.
	 */
	handler SSHHandler
}
