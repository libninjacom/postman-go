package fetchuserresource

type RequestOption func(l *Request)
type Request struct {
	UserId string
}
type Response interface{}
