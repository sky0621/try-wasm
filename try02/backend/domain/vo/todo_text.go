package vo

import (
	"errors"
	"unicode/utf8"

	de "github.com/sky0621/try-wasm/try02/backend/domain/error"
)

func NewTodoText(v string) *TodoText {
	r := TodoText(v)
	return &r
}

type TodoText string

func (v *TodoText) Validate() []*de.DomainError {
	// MEMO: 必須チェックを入れてみる。
	if v == nil {
		return []*de.DomainError{
			{Field: "text", Value: "nil", Err: errors.New("required")},
		}
	}

	// MEMO: 文字列長チェックを入れてみる。
	cnt := utf8.RuneCountInString(string(*v))
	if cnt < 4 || cnt > 10 {
		return []*de.DomainError{
			{Field: "text", Value: string(*v), Err: errors.New("length unmatch")},
		}
	}
	return nil
}

func ValidateTodoText(v *TodoText) []*de.DomainError {
	return v.Validate()
}
