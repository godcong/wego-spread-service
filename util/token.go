package util

import (
	"github.com/godcong/wego-spread-service/model"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"time"
)

// WebToken ...
type WebToken struct {
	UID      string `json:"oid"`
	Nickname string `json:"nickname"`
	//TODO:will add
}

// ExpireTime ...
var ExpireTime = time.Hour * 24 * 7

// NewWebToken ...
func NewWebToken(user *model.User) *WebToken {
	return &WebToken{
		UID:      user.ID,
		Nickname: user.Nickname,
	}
}

// ToToken ...
func ToToken(key string, token *WebToken) (string, error) {
	sub, err := jsoniter.Marshal(token)
	if err != nil {
		return "", err
	}
	jwt, err := EncryptJWT([]byte(key), sub, ExpireTime)
	return jwt, err
}

// FromToken ...
func FromToken(key, token string) (*WebToken, error) {
	t := WebToken{}
	sub, err := DecryptJWT([]byte(key), token)
	log.Info("sub", sub)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal([]byte(sub), &t)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}

	return &t, nil
}
