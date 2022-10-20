package syncrelationswithschema

type Required struct {
	ApiId        string
	ApiVersionId string
	RelationType string
	EntityId     string
}

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
	RelationType string
	EntityId     string
}
type Response interface{}
