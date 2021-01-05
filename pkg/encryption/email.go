package encryption

import (
	"errors"
	"regexp"
)

var (
	// ErrEmailWrong  邮箱格式错误
	ErrEmailWrong = errors.New("the email does not conform to the format")
)

// VerifyEmailFormat 使用正则表达式对邮箱判断
func VerifyEmailFormat(email string) error {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(email) {
		return ErrEmailWrong
	}
	return nil
}
