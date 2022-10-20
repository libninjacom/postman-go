package getallapis

func WithWorkspace(workspace string) RequestOption {
	return func(r *Request) {
		r.Workspace = &workspace
	}
}

func WithSince(since string) RequestOption {
	return func(r *Request) {
		r.Since = &since
	}
}

func WithUntil(until string) RequestOption {
	return func(r *Request) {
		r.Until = &until
	}
}

func WithCreatedBy(createdBy string) RequestOption {
	return func(r *Request) {
		r.CreatedBy = &createdBy
	}
}

func WithUpdatedBy(updatedBy string) RequestOption {
	return func(r *Request) {
		r.UpdatedBy = &updatedBy
	}
}

func WithIsPublic(isPublic bool) RequestOption {
	return func(r *Request) {
		r.IsPublic = &isPublic
	}
}

func WithName(name string) RequestOption {
	return func(r *Request) {
		r.Name = &name
	}
}

func WithSummary(summary string) RequestOption {
	return func(r *Request) {
		r.Summary = &summary
	}
}

func WithDescription(description string) RequestOption {
	return func(r *Request) {
		r.Description = &description
	}
}

func WithSort(sort string) RequestOption {
	return func(r *Request) {
		r.Sort = &sort
	}
}

func WithDirection(direction string) RequestOption {
	return func(r *Request) {
		r.Direction = &direction
	}
}

type RequestOption func(l *Request)
type Request struct {
	Workspace   *string
	Since       *string
	Until       *string
	CreatedBy   *string
	UpdatedBy   *string
	IsPublic    *bool
	Name        *string
	Summary     *string
	Description *string
	Sort        *string
	Direction   *string
}
type Response interface{}
