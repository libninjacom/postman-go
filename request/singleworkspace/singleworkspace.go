package singleworkspace

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId string
}
type Response interface{}
