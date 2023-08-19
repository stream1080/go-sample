package jwt

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func NewClaims(uuid, username string, role int) *UserClaims {
	return &UserClaims{
		uuid,
		username,
		role,
		jwt.StandardClaims{},
	}
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
