package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"federation-one/graph/generated"
	"federation-one/graph/model"
)

func (r *entityResolver) FindContentByID(ctx context.Context, id string) (*model.Content, error) {
	return &model.Content{
		ID:   id,
		Name: id + "_name",
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
