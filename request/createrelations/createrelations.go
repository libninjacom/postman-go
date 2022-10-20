package createrelations

func WithDocumentation(documentation []string) RequestOption {
	return func(r *Request) {
		r.Documentation = &documentation
	}
}

func WithEnvironment(environment []string) RequestOption {
	return func(r *Request) {
		r.Environment = &environment
	}
}

func WithMock(mock []string) RequestOption {
	return func(r *Request) {
		r.Mock = &mock
	}
}

func WithMonitor(monitor []string) RequestOption {
	return func(r *Request) {
		r.Monitor = &monitor
	}
}

func WithTest(test []string) RequestOption {
	return func(r *Request) {
		r.Test = &test
	}
}

func WithContracttest(contracttest []string) RequestOption {
	return func(r *Request) {
		r.Contracttest = &contracttest
	}
}

func WithTestsuite(testsuite []string) RequestOption {
	return func(r *Request) {
		r.Testsuite = &testsuite
	}
}

type RequestOption func(l *Request)
type Request struct {
	ApiId         string
	ApiVersionId  string
	Documentation *[]string
	Environment   *[]string
	Mock          *[]string
	Monitor       *[]string
	Test          *[]string
	Contracttest  *[]string
	Testsuite     *[]string
}
type Response interface{}
