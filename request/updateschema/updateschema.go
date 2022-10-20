package updateschema

func WithSchema(schema interface{}) RequestOption {
	return func(r *Request) {
		r.Schema = &schema
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
	SchemaId     string
	Schema       *interface{}
}
type Response interface{}
