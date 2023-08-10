package jwt

import "github.com/dgrijalva/jwt-go"

type UserInfo struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

var tokenKey = []byte("test-key")

// NewToken 生成 token
func NewToken(uuid, username string, role int) (string, error) {
	UserInfo := UserInfo{
		UUID:           uuid,
		UserName:       username,
		Role:           role,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserInfo)
	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken 解析 token
func AnalyseToken(tokenString string) (*UserInfo, error) {
	UserInfo := &UserInfo{}
	claims, err := jwt.ParseWithClaims(tokenString, UserInfo, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil || !claims.Valid {
		return nil, err
	}
	return UserInfo, nil
}
