package ssh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ConfigHandler struct {
}

func (h *ConfigHandler) OnConnect(channel Channel, user User) {
	channel.Close()
}

func TestNewConfig(t *testing.T) {

	p := uint64(2222)
	k := "/etc/ssh/ssh_host_rsa_key"
	h := &ConfigHandler{}

	c := NewConfig(p, k, h)

	if !assert.NotNil(t, k) {
		t.FailNow()
	}

	assert.Equal(t, p, c.port, "was expected another port")
	assert.Equal(t, k, c.path, "was expected another path")
	assert.Equal(t, h, c.handler, "was expected another handler")

}
