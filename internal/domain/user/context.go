package user

import "context"

type contextKey string

const userIDKey = contextKey("user_id")

func contextWithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}
func userIDFromContext(ctx context.Context) int64 {
	val := ctx.Value(userIDKey)
	if id, ok := val.(int64); ok {
		return id
	}
	return 0
}
