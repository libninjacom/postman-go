package createcollectionfromschema

type Required struct {
	ApiId        string
	ApiVersionId string
	SchemaId     string
	Name         string
	Relations    []interface{}
}

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId        string
	ApiVersionId string
	SchemaId     string
	WorkspaceId  *string
	Name         string
	Relations    []interface{}
}
type Response interface{}
