package allworkspaces

func WithType(type_ string) RequestOption {
	return func(r *Request) {
		r.Type = &type_
	}
}

type RequestOption func(l *Request)
type Request struct {
	Type *string
}
type Response interface{}
