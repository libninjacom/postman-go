package singleenvironment

type RequestOption func(l *Request)
type Request struct {
	EnvironmentUid string
}
type Response interface{}
