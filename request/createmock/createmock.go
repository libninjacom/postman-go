package createmock

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithMock(mock interface{}) RequestOption {
	return func(r *Request) {
		r.Mock = &mock
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Mock        *interface{}
}
type Response interface{}
