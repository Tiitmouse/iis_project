package structs

import "github.com/google/uuid"

type ContactType string

const (
	EmailType  ContactType = "email"
	PhoneType  ContactType = "phone"
	SocialType ContactType = "social"
)

type Contact struct {
	ID      string      `json:"id"`
	Type    ContactType `json:"type" binding:"required"`
	Value   string      `json:"value" binding:"required"`
	Name    string      `json:"name,omitempty"`
	Sources []string    `json:"sources,omitempty"`
}

func NewContact() Contact {
	return Contact{
		ID:      uuid.NewString(),
		Sources: []string{},
	}
}
