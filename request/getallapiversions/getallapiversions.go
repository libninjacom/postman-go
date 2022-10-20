package getallapiversions

type RequestOption func(l *Request)
type Request struct {
	ApiId string
}
type Response interface{}
