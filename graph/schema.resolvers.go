package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql_golang/graph/generated"
	"graphql_golang/graph/model"
)

func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpserCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
