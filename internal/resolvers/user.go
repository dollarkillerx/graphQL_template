package resolvers

import (
	"context"
	"time"

	"github.com/dollarkillerx/graphql_template/internal/generated"
	"github.com/dollarkillerx/graphql_template/internal/pkg/enum"
	"github.com/dollarkillerx/graphql_template/internal/utils"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (r *mutationResolver) LoginByPassword(ctx context.Context, input *generated.LoginByPassword) (*generated.AuthPayload, error) {
	token, err := utils.JWT.CreateToken(enum.AuthJWT{}, time.Now().Add(time.Hour*24*7).Unix())
	if err != nil {
		return nil, err
	}
	return &generated.AuthPayload{
		AccessTokenString: token,
		UserID:            "user id",
	}, nil
}

func (r *mutationResolver) Registry(ctx context.Context, input *generated.Registry) (*wrapperspb.BoolValue, error) {
	return &wrapperspb.BoolValue{Value: true}, nil
}

func (r *queryResolver) User(ctx context.Context) (*generated.UserInformation, error) {
	return &generated.UserInformation{}, nil
}

func (r *queryResolver) Captcha(ctx context.Context) (*generated.Captcha, error) {
	return &generated.Captcha{}, nil
}
