package app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/try-wasm/try02/backend/apperror"
	"github.com/sky0621/try-wasm/try02/backend/domain"
	"github.com/sky0621/try-wasm/try02/backend/domain/vo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo, errs := domain.CreateTodo(domain.Todo{
		Text:   vo.NewTodoText(input.Text),
		UserID: vo.NewUserID(input.UserID),
	})
	if errs != nil {
		for _, err := range errs {
			apperror.NewValidationError(err.Field, err.Value).AddGraphQLError(ctx)
		}
		return nil, nil
	}

	return &Todo{
		ID:   "newID",
		Text: string(*todo.Text),
		Done: false,
		User: &User{
			ID:   string(*todo.UserID),
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
