package tmpl

import "text/template"

// TmplOpt ...
type TmplOpt func(*TmplOpts)

// Opts ...
type TmplOpts struct {
	Fields                TmplFields
	Funcs                 template.FuncMap
	FailOnMissing         bool
	DisableReplaceNoValue bool
}

// TmplFields ...
type TmplFields map[string]interface{}

// NewOpts ...
func NewOpts() TmplOpts {
	return TmplOpts{
		Fields: make(TmplFields),
		Funcs:  tmplFuncs,
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

// WithDisableReplaceNoValue ...
func WithDisableReplaceNoValue() TmplOpt {
	return func(opts *TmplOpts) {
		opts.DisableReplaceNoValue = true
	}
}

// WithFailOnMissing ...
func WithFailOnMissing() TmplOpt {
	return func(opts *TmplOpts) {
		opts.FailOnMissing = true
	}
}
