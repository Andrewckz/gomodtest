package gomodtest

import (
	"crypto/md5"
	"fmt"
)

func JiaMi(pwd string) string {
	pwdAndSalt := []byte("salt1" + pwd)
	apwd := fmt.Sprintf("%x", md5.Sum(pwdAndSalt))
	fmt.Printf("加密后密码为: %x\n", apwd)
	return apwd
}
