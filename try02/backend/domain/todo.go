package domain

import (
	de "github.com/sky0621/try-wasm/try02/backend/domain/error"
	"github.com/sky0621/try-wasm/try02/backend/domain/vo"
)

type Todo struct {
	Text   *vo.TodoText `json:"text"`
	UserID *vo.UserID   `json:"userId"`
}

func CreateTodo(todo Todo) (*Todo, []*de.DomainError) {
	/*
	 * バリデーション
	 */
	var domainErrors []*de.DomainError

	// ToDoテキストのバリデーション
	vtt := vo.ValidateTodoText(todo.Text)
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
