package createcollection

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithCollection(collection interface{}) RequestOption {
	return func(r *Request) {
		r.Collection = &collection
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Collection  *interface{}
}
type Response interface{}
