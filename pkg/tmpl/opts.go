package tmpl

import "text/template"

// TmplOpt ...
type TmplOpt func(*TmplOpts)

// Opts ...
type TmplOpts struct {
	Fields TmplFields
	Funcs  template.FuncMap
}

// TmplFields ...
type TmplFields map[string]interface{}

// NewOpts ...
func NewOpts() TmplOpts {
	return TmplOpts{
		Fields: make(TmplFields),
		Funcs:  make(template.FuncMap),
	}
}

// Configure os configuring the options.
func (o *TmplOpts) Configure(opts ...TmplOpt) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithExtraFields ...
func WithExtraFields(fields TmplFields) TmplOpt {
	return func(opts *TmplOpts) {
		for f, v := range fields {
			opts.Fields[f] = v
		}
	}
}

// WithExtraFuncs ...
func WithExtraFuncs(funcs template.FuncMap) TmplOpt {
	return func(opts *TmplOpts) {
		for f, v := range funcs {
			opts.Funcs[f] = v
		}
	}
}
