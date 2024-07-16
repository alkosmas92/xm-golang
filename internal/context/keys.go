package context

type contextKey string

const (
	UserIDKey   contextKey = "userID"
	UsernameKey contextKey = "username"
)
