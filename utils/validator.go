package util

import (
	"github.com/go-playground/validator/v10"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

func GoValidator(s interface{}, config []gpc.ErrorMetaConfig) (interface{}, int) {
	var validate *validator.Validate
	validators := gpc.NewValidator(validate)
	bind := gpc.NewBindValidator(validators)

	errResponse, errCount := bind.BindValidator(s, config)
	return errResponse, errCount
}
