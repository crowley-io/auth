package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SSHTestHandler struct {
	called bool
	user   User
}

func (h *SSHTestHandler) OnConnect(channel Channel, user User) {
	h.called = true
	h.user = user
	fmt.Fprintf(channel, "Hello %s !\r\n", user.Name())
	channel.Close()
}

func config() (*SSHConfig, *SSHTestHandler) {

	h := &SSHTestHandler{}

	p, e := strconv.ParseUint(os.Getenv("SSH_PORT"), 10, 64)

	if e != nil {
		p = 2222
	}

	c := &SSHConfig{
		port:    p,
		path:    "../files/host",
		handler: h,
	}

	return c, h
}

func TestServerInit(t *testing.T) {

	c, _ := config()

	server, err := NewSSHServer(c)

	if !assert.Nil(t, err) {
		t.Logf("%+v", err)
	}

	if assert.NotNil(t, server) {
		t.Logf("%+v", server)
	}

}

func TestServerClose(t *testing.T) {

	c, _ := config()

	server, err := NewSSHServer(c)

	assert.Nil(t, err)

	if !assert.NotNil(t, server) {
		t.Fatalf("%+v", err)
	}

	server.Close()
	t.Logf("%+v", server)
}

// Lambda which return a new ssh.ClientConfig with Auth struct's member defined.
type createClientConfig func(c *SSHConfig) *ssh.ClientConfig

// Lambda which handle a new ssh client for testing.
// Return a boolean if we should use onNewSession callback.
type prepareNewSession func(t *testing.T, c *ssh.Client, e error) bool

// Lambda called on new ssh session for testing.
type onNewSession func(t *testing.T, s *ssh.Session, h *SSHTestHandler)

func handle(t *testing.T, conf createClientConfig, prepare prepareNewSession, session onNewSession) {

	c, h := config()

	s, err := NewSSHServer(c)

	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	if !assert.NotNil(t, s) {
		t.FailNow()
	}

	// Start ssh server.
	go func() {
		err := s.Listen()
		assert.Nil(t, err)
	}()
	defer s.Close()

	// Define ssh client configuration
	cf := conf(c)
	cf.User = "user"

	// Define remote host
	ht := "127.0.0.1:" + strconv.FormatUint(c.port, 10)

	// Wait ssh server to start
	time.Sleep(time.Millisecond * 200)

	// Start ssh handshake
	client, err := ssh.Dial("tcp", ht, cf)
	if client != nil {
		defer client.Close()
	}

	// Check if we should start a new ssh session
	if prepare(t, client, err) {

		s, err := client.NewSession()
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}
		defer s.Close()

		session(t, s, h)

	}

}

func TestServerPasswordConnect(t *testing.T) {

	c := func(c *SSHConfig) *ssh.ClientConfig {
		// Use password as auth method.
		e := &ssh.ClientConfig{}
		e.Auth = append(e.Auth, ssh.Password("1234"))
		return e
	}

	p := func(t *testing.T, c *ssh.Client, e error) bool {
		if !assert.NotNil(t, e) || !assert.Nil(t, c) {
			t.FailNow()
		}
		return false
	}

	f := func(t *testing.T, s *ssh.Session, h *SSHTestHandler) {
		t.Fatal("Unexpected behavior: ssh connexion should be closed.")
	}

	handle(t, c, p, f)
}

func TestServerRSAConnect(t *testing.T) {

	c := func(c *SSHConfig) *ssh.ClientConfig {
		// Use RSA key as auth method.
		e := &ssh.ClientConfig{}
		pk, err := ReadPrivateKey(c.path)
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}
		e.Auth = append(e.Auth, ssh.PublicKeys(*pk))
		return e
	}

	p := func(t *testing.T, c *ssh.Client, e error) bool {
		if !assert.Nil(t, e) {
			t.Fatal(e)
		}
		return true
	}

	f := func(t *testing.T, session *ssh.Session, h *SSHTestHandler) {

		// Get stdout io.Reader
		out, err := session.StdoutPipe()
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}

		// Set up terminal modes
		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     // disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}

		// Request a pty
		if err := session.RequestPty("xterm", 80, 40, modes); !assert.Nil(t, err) {
			t.Fatal(err)
		}

		// Start remote shell
		if err := session.Shell(); !assert.Nil(t, err) {
			t.Fatal(err)
		}

		assert.True(t, h.called)
		assert.NotNil(t, h.user)

		b, err := ioutil.ReadAll(out)

		if !assert.Nil(t, err) {
			t.Fatal(err)
		}

		s := string(b)

		assert.Equal(t, "user", h.user.Name())
		assert.NotEmpty(t, s)
		assert.Contains(t, s, "Hello user !")

		t.Log(s)
	}

	handle(t, c, p, f)
}

func TestServerRSAExec(t *testing.T) {

	c := func(c *SSHConfig) *ssh.ClientConfig {
		// Use RSA key as auth method.
		e := &ssh.ClientConfig{}
		pk, err := ReadPrivateKey(c.path)
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}
		e.Auth = append(e.Auth, ssh.PublicKeys(*pk))
		return e
	}

	p := func(t *testing.T, c *ssh.Client, e error) bool {
		if !assert.Nil(t, e) {
			t.Fatal(e)
		}
		return true
	}

	f := func(t *testing.T, session *ssh.Session, h *SSHTestHandler) {

		// Get stdout io.Reader
		out, err := session.StdoutPipe()
		if !assert.Nil(t, err) {
			t.Fatal(err)
		}

		// Execute remote command
		if _, err := session.Output("/bin/ls"); !assert.NotNil(t, err) {
			t.FailNow()
		}

		b, err := ioutil.ReadAll(out)

		if !assert.Nil(t, err) {
			t.Fatal(err)
		}

		s := string(b)
		assert.NotEmpty(t, s)

		t.Log(s)
	}

	handle(t, c, p, f)
}
