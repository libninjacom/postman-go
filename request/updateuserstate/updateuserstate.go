package updateuserstate

func WithSchemas(schemas []string) RequestOption {
	return func(r *Request) {
		r.Schemas = &schemas
	}
}

func WithOperations(operations []interface{}) RequestOption {
	return func(r *Request) {
		r.Operations = &operations
	}
}

type RequestOption func(l *Request)
type Request struct {
	UserId     string
	Schemas    *[]string
	Operations *[]interface{}
}
type Response interface{}
