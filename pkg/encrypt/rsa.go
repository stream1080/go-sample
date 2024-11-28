package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func Decrypt(data string, privateKeyStr string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	decodeString, _ := base64.StdEncoding.DecodeString(data)
	bytes, err := decrypt(decodeString, privateKey)
	return string(bytes), err
}

func Encrypt(data string, publicKeyStr string) (string, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	bytes, err := encrypt([]byte(data), publicKey)
	return base64.StdEncoding.EncodeToString(bytes), err
}

// Sign 私钥对数据签名 & base64 Encode (SHA512)
func Sign(privateKeyStr string, data []byte) (string, error) {
	// 对原始数据进行哈希
	hashed := sha512.Sum512(data)

	block, _ := pem.Decode([]byte(privateKeyStr))
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 使用私钥对哈希数据进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify 校验签名 SHA512
func Verify(publicKeyStr string, sign string, data []byte) (bool, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	// 验证签名
	decodeString, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	// 对原始数据进行哈希
	hashed := sha512.Sum512(data)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashed[:], decodeString)
	if err != nil {
		return false, err
	}
	return err == nil, err
}

func ParsePrivate(privateKeyStr string) (bool, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return false, errors.New("无法解析密钥")
	}
	_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return err == nil, err
}

func ParsePublic(publicKeyStr string) (bool, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return false, errors.New("无法解析密钥")
	}
	_, err := x509.ParsePKCS1PublicKey(block.Bytes)
	return err == nil, err
}

func encrypt(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

func decrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}
