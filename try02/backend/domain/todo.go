package domain

import (
	"errors"
	"unicode/utf8"
)

type Todo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

func CreateTodo(todo Todo) (*Todo, []*Error) {
	/*
	 * バリデーション
	 */
	var domainErrors []*Error

	// ToDoテキストのバリデーション
	vtt := ValidateTodoText(todo.Text)
	if vtt != nil {
		domainErrors = append(domainErrors, vtt...)
	}

	// MEMO: ユーザーIDのバリデーション
	// 今回はバリデーション１事例あれば十分なので省略。

	if domainErrors != nil {
		return nil, domainErrors
	}

	/*
	 * MEMO: 何かしらの登録処理を行った後、結果を返却！
	 * 今回の趣旨ではないので省略。
	 */
	return &Todo{
		Text:   todo.Text,
		UserID: todo.UserID,
	}, nil
}

func ValidateTodoText(text string) []*Error {
	// MEMO: 必須チェックを入れてみる。
	if text == "" {
		return []*Error{
			{Field: "text", Value: "nil", Err: errors.New("required")},
		}
	}

	// MEMO: 文字列長チェックを入れてみる。
	cnt := utf8.RuneCountInString(text)
	if cnt < 4 || cnt > 10 {
		return []*Error{
			{Field: "text", Value: text, Err: errors.New("length unmatch")},
		}
	}
	return nil
}
