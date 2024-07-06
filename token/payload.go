package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	errExpiredTokenMessage = errors.New("token has expired")
	errInvalidToken        = errors.New("invalid token")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errExpiredTokenMessage
	}
	return nil
}

// implementacoes da interface -> obrigatorias
func (p Payload) GetAudience() (jwt.ClaimStrings, error) {
	return []string{}, nil
}

func (p Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.ExpiredAt), nil
}

func (p Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.IssuedAt), nil
}

func (p Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (p Payload) GetSubject() (string, error) {
	return p.Username, nil
}
