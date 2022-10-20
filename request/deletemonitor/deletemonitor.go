package deletemonitor

type RequestOption func(l *Request)
type Request struct {
	MonitorUid string
}
type Response interface{}
