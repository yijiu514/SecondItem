package encryption

import (
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Claims 创建claim结构体
type Claims struct {
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
	claims := &Claims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
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
	idtoken := strconv.Itoa(u.ID)
	w.Header().Set("id", idtoken)
	w.Header().Set("token", token)
	return nil
}

// ParseToken token解析
func ParseToken(tokenString string, id int) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		salt, err := models.QuerySessionSalt(id)
		if err != nil {
			return []byte(""), fmt.Errorf("parse wrong %w", err)
		}
		return []byte(salt), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse wrong %w", err)
	}
	return token, nil
}

// TokenVerify 令牌验证
func TokenVerify(tokenstring string, id int) (err error) {
	if tokenstring == "" {
		return ErrTokenEmpty
	}
	token, err := ParseToken(tokenstring, id)
	if err != nil || !token.Valid {
		return ErrTokenWrong
	}

	return nil
}
