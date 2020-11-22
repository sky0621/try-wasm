package app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/sky0621/try-wasm/try02/backend/domain"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo, errs := domain.CreateTodo(domain.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	})
	if errs != nil {
		for _, e := range errs {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: "VALIDATION_FAILURE",
				Extensions: map[string]interface{}{
					"status_code":   400,
					"error_message": e.Err.Error(),
					"field":         e.Field,
					"value":         e.Value,
				},
			})
		}
		return nil, nil
	}

	return &Todo{
		ID:   "newID",
		Text: todo.Text,
		Done: false,
		User: &User{
			ID:   todo.UserID,
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
