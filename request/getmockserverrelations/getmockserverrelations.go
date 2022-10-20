package getmockserverrelations

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
}
type Response interface{}
