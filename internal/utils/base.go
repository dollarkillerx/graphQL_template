package utils

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/dollarkillerx/graphql_template/internal/middlewares"
	"net/http"
	"time"
)

func GetDefaultHttpClient(timeout int) *http.Client {
	if timeout == 0 {
		timeout = 3
	}
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

// hasLoginFunc 如果是请求需要登录才能访问的接口，则需要判断是否带有 token ，并检测 token 的合法性，如果失败拒绝请求
func HasLoginFunc(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := middlewares.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return next(ctx)
}
