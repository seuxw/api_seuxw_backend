package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	invalid  = errors.New("token invalid")
	mismatch = errors.New("token mismatch")
	ipdrift  = errors.New("token ipdrift")
	timeout  = errors.New("token timeout")
	tooshort = errors.New("token too short")
)

func EncryptPwd(pwd string) string {
	// hashing the password with the default cost of 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func ValidatePwd(pwd, hashed string) error {
	// comparing the password with the hash
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
}

func NewToken(id, ip string, t time.Duration, key []byte) string {
	token, err := encryptToken([]byte(fmt.Sprintf("%s %s %d", id, ip, time.Now().Add(t).Unix())), key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(token)

}

func NewExpiresToken(id, ip string, expires time.Time, key []byte) string {
	token, err := encryptToken([]byte(fmt.Sprintf("%s %s %d", id, ip, expires.Unix())), key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(token)

}

func ValidateToken(id, token, ip string, key []byte) error {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	if data, err = decryptToken(data, key); err != nil {
		return err
	}

	var _t int
	var _id, _ip string
	if _, err := fmt.Sscanf(string(data), "%s %s %d", &_id, &_ip, &_t); err != nil {
		return err
	}

	//println("verify:", data, id, _id, ip, _ip, _t)
	if id != _id {
		return mismatch
	}
	if ip != _ip {
		return ipdrift
	}
	if time.Now().After(time.Unix(int64(_t), 0)) {
		return timeout
	}

	return nil
}

func encryptToken(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decryptToken(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, tooshort
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
