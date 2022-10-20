package updatemonitor

func WithMonitor(monitor interface{}) RequestOption {
	return func(r *Request) {
		r.Monitor = &monitor
	}
}

type RequestOption func(l *Request)
type Request struct {
	MonitorUid string
	Monitor    *interface{}
}
type Response interface{}
