package jwt

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func NewClaims(id uint64, username string, role int) *UserClaims {
	return &UserClaims{id, username, role, jwt.StandardClaims{}}
}

// NewToken 生成 token
func NewToken(userClaims *UserClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	return token.SignedString(secret)
}

// AnalyseToken 解析 token
func AnalyseToken(tokenString, secret string) (*UserClaims, error) {
	userClaims := &UserClaims{}
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !claims.Valid {
		return nil, err
	}

	return userClaims, nil
}
