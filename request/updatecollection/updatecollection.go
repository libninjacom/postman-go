package updatecollection

func WithCollection(collection interface{}) RequestOption {
	return func(r *Request) {
		r.Collection = &collection
	}
}

type RequestOption func(l *Request)
type Request struct {
	CollectionUid string
	Collection    *interface{}
}
type Response interface{}
