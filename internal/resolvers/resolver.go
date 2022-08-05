package resolvers

import (
	"context"
	"time"

	"github.com/dollarkillerx/graphql_template/internal/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Resolver ...
type Resolver struct {
}

// Mutation is the root mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query is the root query resolver
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Healthcheck(ctx context.Context) (string, error) {
	return "ack", nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	return "ack", nil
}

func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	return &timestamppb.Timestamp{
		Seconds: time.Now().Unix(),
	}, nil
}

// field resolver
type fileInfoResolver struct{ *Resolver }
