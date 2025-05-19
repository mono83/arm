package armfx

import (
	"go.uber.org/fx"
	"reflect"
)

// ProvideStruct constructs provider for given struct.
func ProvideStruct(x any, anno ...fx.Annotation) fx.Option {
	t := reflect.TypeOf(x)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	funcOut := []reflect.Type{t}
	var funcIn []reflect.Type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		funcIn = append(funcIn, fieldType)
	}

	fn := reflect.MakeFunc(
		reflect.FuncOf(funcIn, funcOut, false),
		func(args []reflect.Value) []reflect.Value {
			instance := reflect.New(t)
			value := instance.Elem()
			for i := 0; i < t.NumField(); i++ {
				value.Field(i).Set(args[i])
			}

			return []reflect.Value{value}
		},
	)

	if len(anno) == 0 {
		return fx.Provide(fn.Interface())
	}
	return fx.Provide(
		fx.Annotate(fn.Interface(), anno...),
	)
}
