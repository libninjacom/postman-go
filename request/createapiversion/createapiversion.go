package createapiversion

func WithVersion(version interface{}) RequestOption {
	return func(r *Request) {
		r.Version = &version
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId   string
	Version *interface{}
}
type Response interface{}
