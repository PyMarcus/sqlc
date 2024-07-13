package token

import (
	"errors"
	"log"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

var (
	errInvalidKeySize = errors.New("invalid key size")
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	log.Println(len(symmetricKey))
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, errInvalidKeySize
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (pm PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return pm.paseto.Encrypt(pm.symmetricKey, payload, nil)
}

func (pm PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := pm.paseto.Decrypt(token, pm.symmetricKey, payload, nil)
	if err != nil {
		return nil, errInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, errExpiredTokenMessage
	}

	return payload, nil
}
