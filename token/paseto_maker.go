package token

import (
	"time"

	"github.com/o1egl/paseto"
)

type Paseto struct {
	paseto       *paseto.V2
	symmetrickey []byte
}

func NewPasetoMaker(symmetrickey string) (*Paseto, error) {

	if len(symmetrickey) < 4 {
		return nil, ErrInvalidKey
	}

	return &Paseto{
		paseto:       paseto.NewV2(),
		symmetrickey: []byte(symmetrickey),
	}, nil
}

func (p *Paseto) CreateToken(username, email string, duration time.Duration) (string, error) {

	payload := NewPayload(username, email, duration)

	return p.paseto.Encrypt(p.symmetrickey, payload, nil)
}

func (p *Paseto) VerifyToken(token string) error {
	payload := Payload{}

	err := p.paseto.Decrypt(token, p.symmetrickey, &payload, nil)
	if err != nil {
		return err
	}

	err = payload.Valid()
	if err != nil {
		return err
	}

	return nil
}

