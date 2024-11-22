package validation

import (
	"errors"
	"reflect"
	"strings"

	"go-sample/pkg/response"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			var name string
			for _, t := range []string{"json", "form", "uri", "query"} {
				name = strings.SplitN(fld.Tag.Get(t), ",", 2)[0]
				if name != "" {
					if name == "-" {
						return ""
					}
					return name
				}
			}

			return name
		})

		trans, _ = ut.New(zh.New()).GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)
	}
}

func Translate(err error) string {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return response.InvalidArgs.Msg()
	} else {
		return errs[0].Translate(trans)
	}
}
