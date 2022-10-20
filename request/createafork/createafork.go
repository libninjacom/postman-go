package createafork

func WithLabel(label string) RequestOption {
	return func(r *Request) {
		r.Label = &label
	}
}

type RequestOption func(l *Request)
type Request struct {
	Workspace     string
	CollectionUid string
	Label         *string
}
type Response interface{}
