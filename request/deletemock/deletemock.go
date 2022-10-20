package deletemock

type RequestOption func(l *Request)
type Request struct {
	MockUid string
}
type Response interface{}
