package updateanapiversion

func WithVersion(version interface{}) RequestOption {
	return func(r *Request) {
		r.Version = &version
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
	Version      *interface{}
}
type Response interface{}
