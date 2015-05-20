package ssh

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

type User struct {
	name string
	key  []byte
}

// Username from incoming ssh connection
func (u *User) Name() string {
	return u.name
}

// SSH Key in base64 used by incoming ssh connection
func (u *User) Key() string {
	return base64.StdEncoding.EncodeToString(u.key)
}

// SSH Key's fingerprint from incoming ssh connection
func (u *User) Fingerprint() string {
	h := md5.Sum(u.key)
	f := fmt.Sprintf("% x", h)
	return strings.Replace(f, " ", ":", -1)
}
