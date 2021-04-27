package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidatorReponse struct {
    Value interface{}
    Message string
}

func GoValidator(s interface{}) map[string]interface{} {

    validate = validator.New()

    err := validate.Struct(s)
    errObject := make(map[string]interface{})

    if err != nil {
        for _, errResult := range  err.(validator.ValidationErrors) {
            switch errResult.ActualTag() {
            case "email":
            errObject[errResult.StructField()] = ValidatorReponse{
                    Value: errResult.Value(),
                    Message: "email format is not valid",
                }
            case "required":
            errObject[errResult.StructField()] = ValidatorReponse{
                Value: errResult.Value(),
                Message: strings.ToLower(errResult.StructField()) + " is required",
            }
        }
    }
}
    return errObject
}
