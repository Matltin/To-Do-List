package token

import (
	"time"
	"to_do_list/util"

	"github.com/o1egl/paseto"
)

type Paseto struct {
	paseto       *paseto.V2
	symmetrickey []byte
}

func NewPasetoMaker(symmetrickey string) (*Paseto, error) {

	if len(symmetrickey) < 4 {
		return nil, util.ErrInvalidKey
	}

	return &Paseto{
		paseto:       paseto.NewV2(),
		symmetrickey: []byte(symmetrickey),
	}, nil
}

func (p *Paseto) CreateToken(id uint, email string, duration time.Duration) (string, error) {

	payload := NewPayload(id, email, duration)

	token, err := p.paseto.Encrypt(p.symmetrickey, payload, nil)

	return token, err
}

func (p *Paseto) VerifyToken(token string) (*Payload, error) {
	payload := Payload{}

	err := p.paseto.Decrypt(token, p.symmetrickey, &payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
