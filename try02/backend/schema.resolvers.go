package app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/try-wasm/try02/backend/domain"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo, err := domain.CreateTodo(domain.Todo{
		Text:   domain.TodoText(input.Text),
		UserID: domain.UserID(input.UserID),
	})
	if err != nil {
		return nil, err
	}
	return &Todo{
		ID:   "newID",
		Text: string(todo.Text),
		Done: false,
		User: &User{
			ID:   string(todo.UserID),
			Name: "user1",
		},
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
