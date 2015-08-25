package ssh

//
// Interface used to handle new user connection.
//
type SSHHandler interface {
	OnConnect(channel Channel, user User)
}

//
// Configuration for SSHServer.
//
type SSHConfig struct {

	// Port configuration.
	port uint64

	// Host's ssh private key path.
	path string

	// Server behavior on new user.
	handler SSHHandler
}

func NewConfig(port uint64, path string, handler SSHHandler) *SSHConfig {
	return &SSHConfig{
		port:    port,
		path:    path,
		handler: handler,
	}
}
