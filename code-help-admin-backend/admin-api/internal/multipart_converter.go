package multipart_converter

import (
	"bytes"
	"fmt"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"io"
	"mime/multipart"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type item struct {
	key         string
	value       reflect.Value
	isOmitEmpty bool
}

func EncodeMultipartFormData(data any) (io.Reader, string) {
	value := reflect.ValueOf(data)

	var body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	defer func() { _ = writer.Close() }()

	values := extractRootValues(value)

	for i := 0; i < len(values); i++ {
		it := values[i]

		if !it.value.IsValid() {
			continue
		}

		if it.value.Kind() == reflect.Pointer {
			if it.value.IsNil() {
				continue
			}

			it.value = it.value.Elem()
		}

		if it.value.Kind() == reflect.Struct {
			switch it.value.Interface().(type) {
			case openapi_types.File:
				field, err := writer.CreateFormField(it.key)
				file := it.value.Interface().(openapi_types.File)
				if err != nil {
					continue
				}

				if bts, err := file.Bytes(); err == nil {
					_, _ = field.Write(bts)
				}

				continue
			}

			values = append(values, extractValues(it)...)
			continue
		}

		if it.value.Kind() == reflect.Slice || it.value.Kind() == reflect.Array {
			values = append(values, toValuesArray(it)...)
			continue
		}

		if it.value.Kind() == reflect.String {
			_ = writer.WriteField(it.key, it.value.String())
			continue
		}

		if it.value.CanInt() {
			_ = writer.WriteField(it.key, strconv.FormatInt(it.value.Int(), 10))
			continue
		}

		if it.value.CanUint() {

			_ = writer.WriteField(it.key, strconv.FormatUint(it.value.Uint(), 10))
			continue
		}

		if it.value.CanFloat() {
			_ = writer.WriteField(it.key, fmt.Sprintf("%f", it.value.Float()))
			continue
		}

	}

	fmt.Println(body.String())

	return body, writer.FormDataContentType()
}

func extractRootValues(value reflect.Value) []item {
	values := make([]item, value.NumField())

	for i := 0; i < value.NumField(); i++ {
		childValue := value.Field(i)
		childField := value.Type().Field(i)

		key, isOmitEmpty := getKey("", childField)

		values[i] = item{key, childValue, isOmitEmpty}
	}

	return values
}

func extractValues(it item) []item {
	values := make([]item, it.value.NumField())

	for i := 0; i < it.value.NumField(); i++ {
		value := it.value.Field(i)
		field := it.value.Type().Field(i)

		key, isOmitEmpty := getKey(it.key, field)

		values[i] = item{key, value, isOmitEmpty}
	}

	return values
}

func getKey(parent string, field reflect.StructField) (string, bool) {
	fieldName, isOmitEmpty := parseJsonTag(field)
	if len(parent) == 0 {
		return fieldName, isOmitEmpty
	}

	return fmt.Sprintf("%s['%s']", parent, fieldName), isOmitEmpty
}

func parseJsonTag(field reflect.StructField) (string, bool) {
	jsonTag := field.Tag.Get("json")
	sections := strings.Split(jsonTag, ",")

	return sections[0], slices.Contains(sections[1:], "omitempty")
}

func toValuesArray(it item) []item {
	values := make([]item, it.value.Len())
	for i := 0; i < it.value.Len(); i++ {
		childValue := it.value.Index(i)

		values = append(values, item{
			fmt.Sprintf("%s[%d]", it.key, i),
			childValue,
			true,
		})
	}

	return values
}
