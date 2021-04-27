package util

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func GoValidator(s interface{}) map[string]interface{} {

    validate = validator.New()

    err := validate.Struct(s)
    errObject := make(map[string]interface{})

    if err != nil {
        for _, errResult := range  err.(validator.ValidationErrors) {
            errObject[errResult.StructField()] = errResult.Value()
        }
    }

    return errObject
}
