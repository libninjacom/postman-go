package schemasecurityvalidation

func WithSchema(schema interface{}) RequestOption {
	return func(r *Request) {
		r.Schema = &schema
	}
}

type RequestOption func(l *Request)
type Request struct {
	Schema *interface{}
}
type Response interface{}
