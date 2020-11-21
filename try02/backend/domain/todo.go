package domain

type Todo struct {
	Text   TodoText `json:"text"`
	UserID UserID   `json:"userId"`
}

type TodoText string

type UserID string

func CreateTodo(todo Todo) (Todo, error) {

	return Todo{
		Text:   todo.Text,
		UserID: todo.UserID,
	}, nil
}
