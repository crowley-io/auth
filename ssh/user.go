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

func (u *User) Name() string {
	return u.name
}

func (u *User) Key() string {
	return base64.StdEncoding.EncodeToString(u.key)
}

func (u *User) Fingerprint() string {
	h := md5.Sum(u.key)
	f := fmt.Sprintf("% x", h)
	return strings.Replace(f, " ", ":", -1)
}
