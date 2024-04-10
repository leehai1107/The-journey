package conv

import (
	"fmt"
	"reflect"
)

// MapStructs map struct by json tag
// Note: This func is not stable may cause panic
func mapStructsByJson(source interface{}, destination interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)

	if sourceValue.Kind() != reflect.Struct || destinationValue.Kind() != reflect.Ptr || destinationValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("source must be a struct, destination must be a pointer to a struct")
	}

	sourceType := sourceValue.Type()
	destinationType := destinationValue.Elem().Type()

	for i := 0; i < sourceType.NumField(); i++ {
		sourceField := sourceType.Field(i)
		sourceTag := sourceField.Tag.Get("json")
		if sourceTag == "" {
			continue
		}

		destinationField, ok := destinationType.FieldByNameFunc(func(fieldName string) bool {
			destinationField, found := destinationType.FieldByName(fieldName)
			destinationTag := destinationField.Tag.Get("json")
			return found && destinationTag == sourceTag
		})
		if !ok {
			continue
		}

		sourceFieldValue := sourceValue.Field(i)
		destinationFieldValue := destinationValue.Elem().FieldByName(destinationField.Name)

		if sourceFieldValue.Kind() == reflect.Ptr && destinationFieldValue.Kind() == reflect.Ptr {
			// If both fields are pointers, set the destination pointer to the source pointer
			if !sourceFieldValue.IsNil() {
				newPtrValue := reflect.New(destinationFieldValue.Type().Elem())
				if err := mapStructsByJson(sourceFieldValue.Elem().Interface(), newPtrValue.Interface()); err != nil {
					return err
				}
				destinationFieldValue.Set(newPtrValue)
			}
		} else if sourceFieldValue.Kind() == reflect.Ptr && destinationFieldValue.Kind() != reflect.Ptr {
			// If source field is a pointer and destination field is not,
			// create a new instance of destination field type and set its value
			if !sourceFieldValue.IsNil() {
				newValue := reflect.New(destinationFieldValue.Type().Elem()).Elem()
				if err := mapStructsByJson(sourceFieldValue.Elem().Interface(), newValue.Addr().Interface()); err != nil {
					return err
				}
				destinationFieldValue.Set(newValue)
			}
		} else if sourceFieldValue.Kind() == reflect.Struct && destinationFieldValue.Kind() == reflect.Struct {
			// If both fields are structs, recursively map them
			if err := mapStructsByJson(sourceFieldValue.Interface(), destinationFieldValue.Addr().Interface()); err != nil {
				return err
			}
		} else if sourceFieldValue.Kind() != reflect.Ptr && destinationFieldValue.Kind() == reflect.Ptr {
			// If source field is not a pointer and destination field is a pointer,
			// create a new instance of destination field type and set its value
			newValue := reflect.New(destinationFieldValue.Type().Elem()).Elem()
			newValue.Set(sourceFieldValue)
			destinationFieldValue.Set(newValue.Addr())
		} else {
			destinationFieldValue.Set(sourceFieldValue)
		}
	}

	return nil
}

func MapStructs(source interface{}, destination interface{}) error {
	return mapStructsByJson(source, destination)
}
