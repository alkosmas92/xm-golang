package context

type contextKey string

// contextKey is a custom type for defining context keys.

const (
	// UserIDKey is the key used to store and retrieve the user ID from context.
	UserIDKey contextKey = "userID"
	// UsernameKey is the key used to store and retrieve the username from context.
	UsernameKey contextKey = "username"
)
