package createapi

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithApi(api interface{}) RequestOption {
	return func(r *Request) {
		r.Api = &api
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Api         *interface{}
}
type Response interface{}
