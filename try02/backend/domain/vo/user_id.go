package vo

func NewUserID(v string) *UserID {
	r := UserID(v)
	return &r
}

type UserID string

func ValidateUserID(v TodoText) {
	// FIXME:
}
