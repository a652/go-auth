package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/a652/go-auth/utils"
)

type AuthToken struct {
	token               string
	createTime          int64
	expiredTimeInterval int64 // s
}

type TokenOption func(at *AuthToken)

func WithExpiredTimeInterval(interval int64) TokenOption {
	return func(at *AuthToken) {
		at.expiredTimeInterval = interval
	}
}

func WithDefaultExpiredTimeInterval() TokenOption {
	return func(at *AuthToken) {
		at.expiredTimeInterval = 1 * 60
	}
}

func GenToken(baseURL, appID, secret, timestamp string) string {
	return calcToken(baseURL, appID, secret, timestamp)
}

func calcToken(baseURL, appID, secret, timestamp string) string {
	var encryptText strings.Builder
	encryptText.WriteString(baseURL)
	encryptText.WriteString(appID)
	encryptText.WriteString(timestamp)
	tmp := utils.HmacSHA1Encrypt(secret, encryptText.String())
	sign := utils.Base64Encode(tmp)
	return sign
}

func GenerateAuthToken(baseUrl, appID, secret string, createTime int64, opts ...TokenOption) *AuthToken {
	token := calcToken(baseUrl, appID, secret, strconv.FormatInt(createTime, 10))
	return NewAuthToken(token, createTime, opts...)
}

func NewAuthToken(token string, createTime int64, opts ...TokenOption) *AuthToken {
	at := &AuthToken{
		token:      token,
		createTime: createTime,
	}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func (at *AuthToken) GetToken() string {
	return at.token
}

func (at *AuthToken) IsExpired() bool {
	now := time.Now().Unix()
	if now-at.createTime > at.expiredTimeInterval {
		return true
	}
	return false
}

func (at *AuthToken) Match(authToken *AuthToken) bool {
	fmt.Println("token1:", at.GetToken())
	fmt.Println("token2:", authToken.GetToken())
	return at.GetToken() == authToken.GetToken()
}
