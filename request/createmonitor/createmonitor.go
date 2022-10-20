package createmonitor

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithMonitor(monitor interface{}) RequestOption {
	return func(r *Request) {
		r.Monitor = &monitor
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Monitor     *interface{}
}
type Response interface{}
