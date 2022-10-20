package createenvironment

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithEnvironment(environment interface{}) RequestOption {
	return func(r *Request) {
		r.Environment = &environment
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Environment *interface{}
}
type Response interface{}
