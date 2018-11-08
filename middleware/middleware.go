package middleware

var (
	TokenAuth *tokenAuth
)

func InitMiddleware() {
	TokenAuth = &tokenAuth{TokenName: "auth-token"}
}
