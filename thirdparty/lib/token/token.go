package token

import (
	"errors"
	"sync"
	"time"

	"github.com/micro/go-micro/config/source"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/micro/go-micro/config"
	log "github.com/micro/go-micro/util/log"
)

// CustomClaims
type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// Token jwt service
type Token struct {
	rwlock     sync.RWMutex
	privateKey []byte
	conf       config.Config
}

func (srv *Token) get() []byte {
	srv.rwlock.RLock()
	defer srv.rwlock.RUnlock()

	return srv.privateKey
}

func (srv *Token) put(newKey []byte) {
	srv.rwlock.Lock()
	defer srv.rwlock.Unlock()

	srv.privateKey = newKey
}

// InitConfig
func InitConfig(source source.Source, path ...string) (*Token, error) {
	token := &Token{}
	token.conf = config.NewConfig()
	err := token.conf.Load(source)
	if err != nil {
		return nil, err
	}
	value := token.conf.Get(path...).Bytes()
	if len(value) == 0 {
		return nil, errors.New("jwt key acquisition failed")
	}
	token.put(value)
	token.enableAutoUpdate(path...)

	return token, nil
}

func (srv *Token) enableAutoUpdate(path ...string) {
	go func() {
		for {
			w, err := srv.conf.Watch(path...)
			if err != nil {
				log.Error(err)
			}
			v, err := w.Next()
			if err != nil {
				log.Error(err)
			}
			value := v.Bytes()
			srv.put(value)
		}
	}()
}

//Decode
func (srv *Token) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.get(), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode
func (srv *Token) Encode(issuer, userName string, expireTime int64) (string, error) {
	claims := CustomClaims{
		userName,
		jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(srv.get())
}
