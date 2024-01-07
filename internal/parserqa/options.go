package parserqa

type Options func(*ParserQAParams) error

type ParserQAParams struct {
	filename string
	buffered bool
}

func newParserQAParams(opts ...Options) (*ParserQAParams, error) {
	params := &ParserQAParams{}

	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithFileName(filename string) Options {
	return func(params *ParserQAParams) error {
		params.filename = filename
		return nil
	}
}

func WithBuffered() Options {
	return func(params *ParserQAParams) error {
		params.buffered = true
		return nil
	}
}

// getters -----

func (p *ParserQAParams) GetFileName() string {
	return p.filename
}

func (p *ParserQAParams) GetBuffered() bool {
	return p.buffered
}

// setters -----

func (p *ParserQAParams) SetFileName(filename string) {
	p.filename = filename
}

func (p *ParserQAParams) SetBuffered(buffered bool) {
	p.buffered = buffered
}
