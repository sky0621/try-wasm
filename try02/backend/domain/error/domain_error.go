package domainerror

type DomainError struct {
	Err   error
	Field string
	Value string
}
