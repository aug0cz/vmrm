package util

type ContextKey string

var (
	DBContextKey = ContextKey("db")
)

func (c ContextKey) String() string {
	return string(c)
}
