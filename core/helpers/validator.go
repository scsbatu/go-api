package helpers

import (
	"fmt"
	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/scsbatu/go-api/contracts"
	"github.com/scsbatu/go-api/core/middlewares"
)

var en ut.Translator

// CustomValidator is for validating the request parameters
type CustomValidator struct {
	Validator *validator.Validate
}

// Init is for initializing the Validator.
// Currently overriding the error message for required tag
func (cv *CustomValidator) Init() {
	english := en_US.New()
	uni := ut.New(english, english)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	en, _ = uni.GetTranslator("en")
	cv.Validator.RegisterTranslation("required", en, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

// Validate is used to validate the struct.
// It implements echo.Validator
func (cv *CustomValidator) Validate(i interface{}) error {
	// cv.Init()
	return cv.Validator.Struct(i)
}

// ExtractAndValidate is used to extract and validate the request to
// the given struct
func ExtractAndValidate(c echo.Context, req interface{}) *contracts.ErrorData {
	// Extracting the values
	if err := c.Bind(req); err != nil {
		httpError := middlewares.ErrTypeMismatch()
		e := &contracts.ErrorData{
			Code:        httpError.GetCode(),
			Description: httpError.Error(),
		}
		return e
	}
	// Validate the struct
	if err := c.Validate(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var msg string
		for _, e := range errs {
			msg = fmt.Sprintf("%s %s", msg, e.Translate(en))
		}
		httpError := middlewares.ErrParametersMissing(msg)
		e := &contracts.ErrorData{
			Code:        httpError.GetCode(),
			Description: httpError.Error(),
		}
		return e
	}
	return nil
}
