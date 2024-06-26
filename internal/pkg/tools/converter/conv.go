package conv

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

// ConvertStringToInt converts string to int
func ConvertStringToInt(numStr string) (int, error) {
	i64, err := convertStringToNumber(numStr, 0)
	return int(i64), err
}

// ConvertStringToInt8 converts string to int8
func ConvertStringToInt8(numStr string) (int8, error) {
	i64, err := convertStringToNumber(numStr, 8)
	return int8(i64), err
}

// ConvertStringToInt16 converts string to int16
func ConvertStringToInt16(numStr string) (int16, error) {
	i64, err := convertStringToNumber(numStr, 16)
	return int16(i64), err
}

// ConvertStringToInt32 converts string to int32
func ConvertStringToInt32(numStr string) (int32, error) {
	i64, err := convertStringToNumber(numStr, 32)
	return int32(i64), err
}

// ConvertStringToInt64 converts string to int64
func ConvertStringToInt64(numStr string) (int64, error) {
	return convertStringToNumber(numStr, 64)
}

// ConvertStringToUInt converts string to uint
func ConvertStringToUInt(numStr string) (uint, error) {
	ui64, err := convertStringToUNumber(numStr, 0)
	return uint(ui64), err
}

// ConvertStringToUInt8 converts string to uint8
func ConvertStringToUInt8(numStr string) (uint8, error) {
	ui64, err := convertStringToUNumber(numStr, 8)
	return uint8(ui64), err
}

// ConvertStringToUInt16 converts string to uint16
func ConvertStringToUInt16(numStr string) (uint16, error) {
	ui64, err := convertStringToUNumber(numStr, 16)
	return uint16(ui64), err
}

// ConvertStringToUInt32 converts string to uint32
func ConvertStringToUInt32(numStr string) (uint32, error) {
	ui64, err := convertStringToUNumber(numStr, 32)
	return uint32(ui64), err
}

// ConvertStringToUInt64 converts string to uint64
func ConvertStringToUInt64(numStr string) (uint64, error) {
	return convertStringToUNumber(numStr, 64)
}

// ConvertStringToBool converts string to bool.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ConvertStringToBool(boolStr string) (bool, error) {
	return strconv.ParseBool(boolStr)
}

// ConvertStringToFloat32 converts string to float32
func ConvertStringToFloat32(numStr string) (float32, error) {
	f64, err := convertStringToFloat(numStr, 32)
	return float32(f64), err
}

// ConvertStringToFloat64 converts string to float64
func ConvertStringToFloat64(numStr string) (float64, error) {
	return convertStringToFloat(numStr, 64)
}

// ConvertIntToString converts int to string
func ConvertIntToString(num int) string {
	return convertNumberToString(int64(num))
}

// ConvertInt8ToString converts int8 to string
func ConvertInt8ToString(num int8) string {
	return convertNumberToString(int64(num))
}

// ConvertInt16ToString converts int16 to string
func ConvertInt16ToString(num int16) string {
	return convertNumberToString(int64(num))
}

// ConvertInt32ToString converts int32 to string
func ConvertInt32ToString(num int32) string {
	return convertNumberToString(int64(num))
}

// ConvertInt64ToString converts int64 to string
func ConvertInt64ToString(num int64) string {
	return convertNumberToString(num)
}

// ConvertUIntToString converts uint to string
func ConvertUIntToString(num uint) string {
	return convertUNumberToString(uint64(num))
}

// ConvertUInt8ToString converts uint to string
func ConvertUInt8ToString(num uint8) string {
	return convertUNumberToString(uint64(num))
}

// ConvertUInt16ToString converts uint to string
func ConvertUInt16ToString(num uint16) string {
	return convertUNumberToString(uint64(num))
}

// ConvertUInt32ToString converts uint to string
func ConvertUInt32ToString(num uint32) string {
	return convertUNumberToString(uint64(num))
}

// ConvertUInt64ToString converts uint to string
func ConvertUInt64ToString(num uint64) string {
	return convertUNumberToString(num)
}

// ConvertBoolToString converts bool to string. It returns "true", "false"
func ConvertBoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// ConvertFloat32ToString convert float32 to string
func ConvertFloat32ToString(num float32) string {
	return convertFloatToString(float64(num), 32)
}

// ConvertFloat64ToString convert float32 to string
func ConvertFloat64ToString(num float64) string {
	return convertFloatToString(num, 64)
}

// ConvertToPointer convert any to pointer
func ConvertToPointer(value interface{}) *any {
	return &value
}

// ConvertToChar convert string to chars
func ConvertStringToChars(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ConvertToBytes convert any to bytes
func ConvertAnyToBytes(value any) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ConvertMapToStruct converts map to struct
func ConvertMapToStruct(m map[string]any, structObj any) error {
	for k, v := range m {
		err := setStructField(structObj, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func setStructField(structObj any, fieldName string, fieldValue any) error {
	structVal := reflect.ValueOf(structObj).Elem()

	fName := getFieldNameByJsonTag(structObj, fieldName)
	if fName == "" {
		return fmt.Errorf("struct field json tag don't match map key : %s in obj", fieldName)
	}

	fieldVal := structVal.FieldByName(fName)

	if !fieldVal.IsValid() {
		return fmt.Errorf("no such field: %s in obj", fieldName)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	val := reflect.ValueOf(fieldValue)

	if fieldVal.Type() != val.Type() {

		if val.CanConvert(fieldVal.Type()) {
			fieldVal.Set(val.Convert(fieldVal.Type()))
			return nil
		}

		if m, ok := fieldValue.(map[string]any); ok {

			if fieldVal.Kind() == reflect.Struct {
				return ConvertMapToStruct(m, fieldVal.Addr().Interface())
			}

			if fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
				if fieldVal.IsNil() {
					fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
				}

				return ConvertMapToStruct(m, fieldVal.Interface())
			}

		}

		return fmt.Errorf("map value type don't match struct field type")
	}

	fieldVal.Set(val)

	return nil
}

func getFieldNameByJsonTag(structObj any, jsonTag string) string {
	s := reflect.TypeOf(structObj).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		tag := field.Tag
		name := tag.Get("json")

		if name == jsonTag {
			return field.Name
		}
	}

	return ""
}

func AsString(src interface{}) string {
	switch value := src.(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case int:
		return strconv.Itoa(value)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case time.Time:
		return value.Format("2006/01/02 15:04:05")
	case bool:
		return strconv.FormatBool(value)
	default:
		bytes, _ := json.Marshal(src)
		return string(bytes)
	}
}

func convertStringToNumber(numStr string, bitSize int) (int64, error) {
	return strconv.ParseInt(numStr, 10, bitSize)
}

func convertStringToUNumber(numStr string, bitSize int) (uint64, error) {
	return strconv.ParseUint(numStr, 10, bitSize)
}

func convertStringToFloat(numStr string, bitSize int) (float64, error) {
	return strconv.ParseFloat(numStr, bitSize)
}

func convertNumberToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func convertUNumberToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func convertFloatToString(num float64, bitSize int) string {
	return strconv.FormatFloat(num, 'E', -1, bitSize)
}
