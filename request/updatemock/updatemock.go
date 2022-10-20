package updatemock

func WithMock(mock interface{}) RequestOption {
	return func(r *Request) {
		r.Mock = &mock
	}
}

type RequestOption func(l *Request)
type Request struct {
	MockUid string
	Mock    *interface{}
}
type Response interface{}
