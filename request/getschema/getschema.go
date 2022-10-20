package getschema

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
	SchemaId     string
}
type Response interface{}
