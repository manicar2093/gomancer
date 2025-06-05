package tags

const (
	EchoQuery  = "query"
	EchoParam  = "param"
	EchoHeader = "header"
	EchoForm   = "form"
)

type (
	EchoOptions struct {
		Name string
		Tag  string
	}

	Echo struct {
		EchoOptions
	}
)

func NewEcho(opts EchoOptions) Generable {
	return &Echo{EchoOptions: opts}
}

func (e Echo) Generate() (string, string) {
	return e.Tag, e.Name
}
