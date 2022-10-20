package deletecollection

type RequestOption func(l *Request)
type Request struct {
	CollectionUid string
}
type Response interface{}
