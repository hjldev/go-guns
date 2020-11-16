package tools

import (
	"golang.org/x/crypto/bcrypt"
)

//加密
func Encrypt(pwd string) (encryptPwd string, err error) {
	if pwd == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		pwd = string(hash)
		return pwd, err
	}
}
