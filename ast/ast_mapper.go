package ast




























//type (
//	Builder interface {
//		AddField(name string, _type interface{}) Builder
//
//		Build() DynamicStruct
//	}
//
//	DynamicStruct interface {
//		New() interface{}
//	}
//
//	builderImplementation struct {
//		fields map[string]*fieldImplementation
//	}
//
//	fieldImplementation struct {
//		_type interface{}
//		tag   string
//	}
//
//	dynamicStructImpl struct {
//		definition reflect.Type
//	}
//)
//
//func NewStruct() Builder {
//	return &builderImplementation{
//		fields: map[string]*fieldImplementation{},
//	}
//}
//
//func ExtendStruct(value interface{}) Builder {
//	return MergeStructs(value)
//}
//
//// MergeStructs merges a list of existing instances of structs and
//// returns new instance of Builder interface.
////
//// builder := dynamicstruct.MergeStructs(MyStructOne{}, MyStructTwo{}, MyStructThree{})
////
//func MergeStructs(values ...interface{}) Builder {
//	builder := NewStruct()
//
//	for _, value := range values {
//		valueOf := reflect.Indirect(reflect.ValueOf(value))
//		typeOf := valueOf.Type()
//
//		for i := 0; i < valueOf.NumField(); i++ {
//			fval := valueOf.Field(i)
//			ftyp := typeOf.Field(i)
//			builder.AddField(ftyp.Name, fval.Interface())
//		}
//	}
//
//	return builder
//}
//
//
//
//
//func (b *builderImplementation) AddField(name string, _type interface{}) Builder {
//	b.fields[name] = &fieldImplementation{
//		_type: _type,
//		tag:   "",
//	}
//
//	return b
//}
//
//func (b *builderImplementation) Build() DynamicStruct {
//	var fields []reflect.StructField
//
//	for name, field := range b.fields {
//		fields = append(fields, reflect.StructField{
//			Name:      name,
//			Type:      reflect.TypeOf(field),
//			Tag:       reflect.StructTag(field.tag),
//		})
//	}
//
//	return &dynamicStructImpl{ definition:reflect.StructOf(fields) }
//}
//
//
//
//func (d *dynamicStructImpl) New() interface{} {
//	return reflect.New(d.definition).Interface()
//}


