package enum

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

// contextKey ...
const (
	TokenCtxKey             ContextKey = "Token"
	UserAgentCtxKey         ContextKey = "UserAgent"
	RequestId               ContextKey = "requestID"
	RequestReceivedAtCtxKey ContextKey = "ReqReceivedAt"
)
