package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IssueAt   time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"Expired_at"`
}

func NewPayload(username, email string, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return fmt.Errorf("token has expired")
	}
	return nil
}
