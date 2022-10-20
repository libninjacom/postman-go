package createworkspace

func WithWorkspace(workspace interface{}) RequestOption {
	return func(r *Request) {
		r.Workspace = &workspace
	}
}

type RequestOption func(l *Request)
type Request struct {
	Workspace *interface{}
}
type Response interface{}
