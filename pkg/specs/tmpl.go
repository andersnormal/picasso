package specs

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	// TmplFile ...
	TmplFile = ".template.yml"
)

// Tmpl ...
type Tmpl struct {
	// Spec ...
	Spec string `validate:"required"`
	// Version ...
	Version string `validate:"required"`
	// Includes ...
	Includes []string `validate:"required_with=Excludes"`
	// Excludes ...
	Excludes []string
	// Ignores ...
	Ignores []string
	// Inputs...
	Inputs TmplInputs
	// Plugins ...
	Plugins Plugins `yaml:"plugins"`
}

// Validate ..
func (t *Tmpl) Validate() error {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(t)
	if err != nil {
		return err
	}

	return v.Struct(t)
}

// TmplInputs ...
type TmplInputs []TmplInput

// TmplInput ...
type TmplInput struct {
	// Name ...
	Name string `validate:"required"`
	// Type ...
	Type string `validate:"required"`
	// Prompt ...
	Prompt string `validate:"required"`
	// Regex ...
	Regex string
}
