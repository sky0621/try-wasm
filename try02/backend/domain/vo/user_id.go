package vo

import de "github.com/sky0621/try-wasm/try02/backend/domain/error"

func NewUserID(v string) *UserID {
	r := UserID(v)
	return &r
}

type UserID string

func (v *UserID) Validate() []*de.DomainError {
	// FIXME:
	return nil
}

func ValidateUserID(v TodoText) []*de.DomainError {
	return v.Validate()
}
