package domain

type Error struct {
	Err   error
	Field string
	Value string
}
