package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// RandomString 生成随机字符串
func RandomString(length int) (string, error) {

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := fmt.Sprintf("%x", b)
	return s, nil
}

// Md5Salt  对密码进行“md5+盐”加密，并返回加密后的密文和盐
func Md5Salt(pwd string, SaltNum int) (string, string) {
	salt, _ := RandomString(SaltNum)

	pwdstring := pwd + salt
	data := []byte(pwdstring)

	h := md5.New()
	h.Write(data)
	output := hex.EncodeToString(h.Sum(nil))

	return output, salt
}

// Md5Stirng 对密码进行md5加密并返回密文
func Md5Stirng(pwd string) string {

	data := []byte(pwd)
	out := md5.Sum(data)
	return fmt.Sprintf("%x", out)

}
