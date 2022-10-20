package updateanapi

func WithApi(api interface{}) RequestOption {
	return func(r *Request) {
		r.Api = &api
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId string
	Api   *interface{}
}
type Response interface{}
