package middlewares

import (
	"context"
)

func GetUserIDFromCtx(ctx context.Context) (string, error) {
	//userID, err := cu.GetUserIDFromContext(ctx)
	//if err != nil {
	//	return "", err
	//}
	//if userID != nil {
	//	return *userID, nil
	//}
	//
	//tokenString, _ := GetTokenFromCtx(ctx)
	//userIDFromToken, err := adam_jwt.GetUserIDFromAccessToken(tokenString, utils.CONFIG.JWTConfig.SecretKey)
	//if err != nil {
	//	return "", errors.WithStack(ec.ErrCodeNotLoggedIn.GetError())
	//}

	return "userIDFromToken", nil
}
