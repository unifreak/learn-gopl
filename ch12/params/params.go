package params

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Unpack populates the fields of the struct pointed to by ptr from the HTTP request
// parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm; err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			// the same parameter name may appear more than once. If this happens, and
			// the field is a slice, then all the values of that parameter are accumulated
			// into the slice.
			if f.Kind() == reflect.Slice {
				elemen := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Error("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
				// otherwise, the field is repeatedly overwritten so that only the last value
				// has any effect.
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
