package updateuserinformation

func WithSchemas(schemas []string) RequestOption {
	return func(r *Request) {
		r.Schemas = &schemas
	}
}

func WithName(name interface{}) RequestOption {
	return func(r *Request) {
		r.Name = &name
	}
}

type RequestOption func(l *Request)
type Request struct {
	UserId  string
	Schemas *[]string
	Name    *interface{}
}
type Response interface{}
