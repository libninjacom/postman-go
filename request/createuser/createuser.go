package createuser

func WithSchemas(schemas []string) RequestOption {
	return func(r *Request) {
		r.Schemas = &schemas
	}
}

func WithUserName(userName string) RequestOption {
	return func(r *Request) {
		r.UserName = &userName
	}
}

func WithActive(active bool) RequestOption {
	return func(r *Request) {
		r.Active = &active
	}
}

func WithExternalId(externalId string) RequestOption {
	return func(r *Request) {
		r.ExternalId = &externalId
	}
}

func WithGroups(groups []string) RequestOption {
	return func(r *Request) {
		r.Groups = &groups
	}
}

func WithLocale(locale string) RequestOption {
	return func(r *Request) {
		r.Locale = &locale
	}
}

func WithName(name interface{}) RequestOption {
	return func(r *Request) {
		r.Name = &name
	}
}

type RequestOption func(l *Request)
type Request struct {
	Schemas    *[]string
	UserName   *string
	Active     *bool
	ExternalId *string
	Groups     *[]string
	Locale     *string
	Name       *interface{}
}
type Response interface{}
