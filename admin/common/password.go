package common

import (
	rsa2 "crypto/rsa"
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/SnakeHacker/deepkg/admin/common/rsa"
	"github.com/golang/glog"
)

func PasswordValidate(password string, privateKey *rsa2.PrivateKey) (decryptedPassword string, err error) {
	decryptedPassword, err = rsa.Decrypt(password, privateKey)
	if err != nil {
		glog.Error(err)
		return
	}

	if decryptedPassword == "" {
		err = errors.New("password is empty")
		glog.Error(err)
		return
	}

	pass, err := PasswordStrengthValidate(decryptedPassword,
		USER_PASSWORD_MAX_LENGTH, USER_PASSWORD_MIN_LENGTH)
	if err != nil || !pass {
		glog.Error(err)
		return
	}

	return
}

func PasswordStrengthValidate(password string, maxLen int, minLen int) (pass bool, err error) {
	if strings.Contains(password, SPACE) {
		err = errors.New("invalid password with space")
		glog.Error(err)
		pass = false
		return pass, err
	}

	r := IsChineseChar(password)
	if r {
		err = errors.New("password has Chinese char")
		glog.Error(err)
		pass = false
		return
	}

	passwordLen := len(password)
	if (passwordLen > maxLen) || (passwordLen < minLen) {
		err = errors.New("invalid password length")
		glog.Error(err)
		pass = false
		return
	}

	count := 0
	if strings.ContainsAny(password, LETTERS_UPPERCASE) {
		count++
	}

	if strings.ContainsAny(password, LETTERS_LOWERCASE) {
		count++
	}

	if strings.ContainsAny(password, NUMBER) {
		count++
	}

	if strings.ContainsAny(password, EN_SPECIAL_CHAR) {
		count++
	}

	if count >= 3 {
		pass = true
		return pass, err
	}

	err = errors.New("用户密码强度需包含：大小写字母、数字、特殊字符中的三种组合")
	glog.Error(err)

	pass = false
	return
}

func IsChineseChar(str string) bool {
	// \u3002 stands for 。
	// \uff1b stands for ；
	// \uff0c stands for ，
	// \uff1a stands for ：
	// \u201c stands for “
	// \u201d stands for ”
	// \uff08 stands for（
	// \uff09 stands for ）
	// \u3001 stands for 、
	// \uff1f stands for ？
	// \u300a stands for《
	// \u300b stands for 》
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) ||
			(regexp.
				MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").
				MatchString(string(r))) {
			return true
		}
	}

	return false
}
