package tools

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	pwd := "123456"
	newPwd, _ := Encrypt(pwd)
	fmt.Println(newPwd)
}
