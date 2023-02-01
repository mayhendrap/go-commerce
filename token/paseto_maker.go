package token

import (
	"aidanwoods.dev/go-paseto"
	"github.com/goccy/go-json"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance PasetoMaker
)

type PasetoMaker struct {
	secretKey paseto.V4AsymmetricSecretKey
}

func NewPasetoMaker() Maker {
	secretKey := paseto.NewV4AsymmetricSecretKey()
	once.Do(func() {
		instance = PasetoMaker{
			secretKey: secretKey,
		}
	})
	return &instance
}

func (maker *PasetoMaker) CreateToken(userID string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(payload.ExpiredAt)
	err = token.Set("data", payload)
	if err != nil {
		return "", err
	}

	signed := token.V4Sign(maker.secretKey, nil)
	return signed, nil
}

func (maker *PasetoMaker) VerifyToken(signedToken string) (*Payload, error) {
	parser := paseto.NewParser()
	token, errToken := parser.ParseV4Public(maker.secretKey.Public(), signedToken, nil)
	if errToken != nil {
		return nil, errToken
	}

	claimsJson := token.ClaimsJSON()
	payload := &Payload{}

	err := json.Unmarshal(claimsJson, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
