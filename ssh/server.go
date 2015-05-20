package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	SSHSocketTimeout = 1
)

type SSHServer struct {

	// Server configuration
	port   uint64
	config *ssh.ServerConfig

	// Server behavior
	handler SSHHandler

	// Server shutdown handler
	shutdown chan bool
	lock     *sync.WaitGroup
}

func NewSSHServer(config *SSHConfig) (*SSHServer, error) {

	server := &SSHServer{
		port: config.port,
		config: &ssh.ServerConfig{
			NoClientAuth:      false,
			PublicKeyCallback: AcceptPublicKey,
		},
		handler:  config.handler,
		shutdown: make(chan bool),
		lock:     &sync.WaitGroup{},
	}

	key, err := ReadPrivateKey(config.path)

	if err != nil {
		return nil, err
	}

	server.config.AddHostKey(*key)

	return server, nil
}

func (s *SSHServer) Close() {
	close(s.shutdown)
	s.lock.Wait()
}

func (s *SSHServer) Listen() error {

	bind, err := net.ResolveTCPAddr("tcp", ":"+strconv.FormatUint(s.port, 10))

	if err != nil {
		return fmt.Errorf("binding ssh server socket: %s", err)
	}

	s.lock.Add(1)
	defer s.lock.Done()

	socket, err := net.ListenTCP("tcp", bind)

	if err != nil {
		return fmt.Errorf("opening ssh server socket: %s", err)
	}

	for {
		select {
		case <-s.shutdown:
			{
				socket.Close()
				return nil
			}
		default:
			{
				s.accept(socket)
			}
		}
	}
}

func (s *SSHServer) accept(socket *net.TCPListener) {

	// Define socket timeout
	socket.SetDeadline(time.Now().Add(SSHSocketTimeout * time.Second))

	// Wait for incoming client
	tcp, err := socket.AcceptTCP()

	// Check if we received an error or a timeout
	if err != nil {
		// Ignore timeout error
		if e, ok := err.(*net.OpError); !ok || !e.Timeout() {
			// TODO : Use logger
			fmt.Fprintln(os.Stderr, "accepting connection:", err)
		}
		return
	}

	// Handle tcp client
	go s.handle(tcp)
}

func (s *SSHServer) handle(tcp *net.TCPConn) {

	client, channels, requests, err := ssh.NewServerConn(tcp, s.config)

	if err != nil {
		// TODO : Use logger
		fmt.Fprintln(os.Stderr, "ssh handshake:", err)
		tcp.Close()
		return
	}

	defer client.Close()

	// Ignore requests
	go ssh.DiscardRequests(requests)

	fmt.Printf("new ssh connection from %+v\n", client)

	for channel := range channels {
		go s.dispatch(client, channel)
	}

}

func (s *SSHServer) dispatch(client *ssh.ServerConn, channel ssh.NewChannel) {

	if channel.ChannelType() != "session" {
		// TODO : Use logger
		fmt.Println("not a session")
		channel.Reject(ssh.Prohibited, "channel type is not a session")
		return
	}

	ch, reqs, err := channel.Accept()

	if err != nil {
		// TODO : Use logger
		fmt.Fprintln(os.Stderr, "accepting channel:", err)
		return
	}

	defer ch.Close()

	for r := range reqs {

		// Require shell request, ignore otherwise.
		if r.Type != "shell" && r.Type != "exec" {

			if r.WantReply {
				r.Reply((r.Type == "pty-req"), nil)
			}

			// Continue to next request
			continue

		} else if r.Type == "exec" && r.WantReply {

			// Ignore exec request with a "nice" message
			r.Reply(true, nil)
			fmt.Fprintf(ch, "Hello %s, ssh commands are disabled.\r\n", client.User())
			ch.Close()

			// Continue to next request
			continue
		}

		r.Reply(true, nil)

		u := User{
			name: client.User(),
			key:  []byte(client.Permissions.Extensions["pubkey"]),
		}

		s.handler.OnConnect(ch, u)
	}
}
