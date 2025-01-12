package utils

import (
	"errors"
	"fmt"
	"reflect"

	validator "github.com/go-playground/validator/v10"
)

func ValidateStruct(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}

func Validate(s interface{}) error {
	// Obtém o tipo e o valor da estrutura
	v := reflect.ValueOf(s)

	// Garante que o valor seja uma struct
	if v.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	// Itera sobre os campos da struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		// Verifica se o campo é vazio (zero valor)
		if isEmpty(field) {
			return fmt.Errorf("field %s is empty", fieldName)
		}
	}
	return nil
}

func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return value.IsNil()
	default:
		// Para tipos numéricos, verifica se é zero
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
