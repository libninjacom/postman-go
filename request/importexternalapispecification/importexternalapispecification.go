package importexternalapispecification

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Body        interface{}
}
type Response interface{}
