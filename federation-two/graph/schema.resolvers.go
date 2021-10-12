package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"federation-two/graph/generated"
	"federation-two/graph/model"
	"fmt"
	"math/rand"
)

func (r *queryResolver) GetReview(ctx context.Context, input *model.ReviewInput) (*model.Review, error) {
	res := &model.Review{
		ID:     input.ID,
		Rating: rand.Int() % 1000,
	}
	for i := 0; i < 3; i++ {
		res.Content = append(res.Content, &model.Content{ID: fmt.Sprintf("%d", rand.Int()%1000)})
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
