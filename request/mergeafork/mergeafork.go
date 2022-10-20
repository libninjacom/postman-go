package mergeafork

func WithDestination(destination string) RequestOption {
	return func(r *Request) {
		r.Destination = &destination
	}
}

func WithSource(source string) RequestOption {
	return func(r *Request) {
		r.Source = &source
	}
}

func WithStrategy(strategy string) RequestOption {
	return func(r *Request) {
		r.Strategy = &strategy
	}
}

type RequestOption func(l *Request)
type Request struct {
	Destination *string
	Source      *string
	Strategy    *string
}
type Response interface{}
