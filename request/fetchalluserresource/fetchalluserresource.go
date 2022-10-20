package fetchalluserresource

func WithStartIndex(startIndex float64) RequestOption {
	return func(r *Request) {
		r.StartIndex = &startIndex
	}
}

func WithCount(count float64) RequestOption {
	return func(r *Request) {
		r.Count = &count
	}
}

func WithFilter(filter string) RequestOption {
	return func(r *Request) {
		r.Filter = &filter
	}
}

type RequestOption func(l *Request)
type Request struct {
	StartIndex *float64
	Count      *float64
	Filter     *string
}
type Response interface{}
