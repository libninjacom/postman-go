package updateenvironment

func WithEnvironment(environment interface{}) RequestOption {
	return func(r *Request) {
		r.Environment = &environment
	}
}

type RequestOption func(l *Request)
type Request struct {
	EnvironmentUid string
	Environment    *interface{}
}
type Response interface{}
