package encryption

import (
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 创建claim结构体
type MyClaims struct {
	UserID             int
	jwt.StandardClaims //设置claim信息结构体
}

var (
	// ErrTokenEmpty 令牌为空
	ErrTokenEmpty = errors.New("the token is empty")
	// ErrTokenWrong 令牌失效或者错误
	ErrTokenWrong = errors.New("the token is wrong or expired")
)

// TokenCreate 生成token
func TokenCreate(u models.User) (tokenstring string, err error) {
	expirTime := time.Now().Add(2 * time.Hour) //设置有效时间为2小时
	claims := &MyClaims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
			Id:        "123",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err = token.SignedString([]byte(u.Passwordsalt))

	if err != nil {
		return "", fmt.Errorf("token encryption wrong %w", err)
	}
	return tokenstring, nil
}

// TokenIssue 令牌下发
func TokenIssue(u models.User, w http.ResponseWriter) error {

	token, err := TokenCreate(u)
	if err != nil {
		return fmt.Errorf("token create wrong %w", err)
	}
	w.Header().Set("token", token)
	return nil
}

// ParseToken token解析
func ParseToken(tokenString string) (*jwt.Token, int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		id := token.Claims.(*MyClaims).UserID
		salt, err := models.QuerySessionSalt(id)
		if err != nil {
			return []byte(""), fmt.Errorf("parse wrong %w", err)
		}
		return []byte(salt), nil
	})
	if err != nil {
		return nil, -1, fmt.Errorf("parse wrong %w", err)
	}
	id := token.Claims.(*MyClaims).UserID
	return token, id, nil
}

// TokenCheck 令牌验证
func TokenCheck(r *http.Request) (id int, err error) {
	tokenstring := r.Header.Get("token")
	if tokenstring == "" {
		return -1, ErrTokenEmpty
	}
	token, id, err := ParseToken(tokenstring)
	if err != nil || !token.Valid {
		return -1, ErrTokenWrong
	}
	return id, nil
}
