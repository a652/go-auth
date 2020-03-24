package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"strconv"
	"time"
)

//GetMD5Hash generates md5
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetNonce() string {
	result, _ := rand.Int(rand.Reader, big.NewInt(100000000))
	return result.String()
}

func GetTimeStamp() string {
	m := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(m, 10)
}

func HmacSHA1Encrypt(accessSecretKey, encryptText string) []byte {
	key := []byte(accessSecretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(encryptText))
	return mac.Sum(nil)
}

func Base64Encode(binput []byte) string {
	return base64.StdEncoding.EncodeToString(binput)
}
