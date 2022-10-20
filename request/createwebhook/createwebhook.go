package createwebhook

func WithWorkspaceId(workspaceId string) RequestOption {
	return func(r *Request) {
		r.WorkspaceId = &workspaceId
	}
}

func WithWebhook(webhook interface{}) RequestOption {
	return func(r *Request) {
		r.Webhook = &webhook
	}
}

type RequestOption func(l *Request)
type Request struct {
	WorkspaceId *string
	Webhook     *interface{}
}
type Response interface{}
