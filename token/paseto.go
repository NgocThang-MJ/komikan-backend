package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/o1egl/paseto"
)

type TokenPayload struct {
	TokenId   uuid.UUID   `json:"token_id"`
	UserId    pgtype.UUID `json:"user_id"`
	IssuedAt  time.Time   `json:"issued_at"`
	ExpiredAt time.Time   `json:"expired_at"`
}

type Token interface {
	CreateToken(usreId pgtype.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenPayload, error)
}

type Paseto struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPaseto(key string) (Token, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(
			"Invalid key size: must be exactly %d characters",
			chacha20poly1305.KeySize,
		)
	}

	maker := &Paseto{paseto: paseto.NewV2(), symmetricKey: []byte(key)}
	return maker, nil
}

func (maker *Paseto) CreateToken(userId pgtype.UUID, duration time.Duration) (string, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &TokenPayload{
		TokenId:   tokenId,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, err
}

func (maker *Paseto) VerifyToken(token string) (payload *TokenPayload, err error) {
	err = maker.paseto.Decrypt(token, maker.symmetricKey, &payload, nil)
	if err != nil {
		return
	}

	err = payload.checkExpired()
	return
}

func (payload *TokenPayload) checkExpired() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("Token has expired")
	}
	return nil
}
