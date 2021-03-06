package schema

import (
	"github.com/zimnx/YamlSchemaToGoStruct/item"
	"github.com/zimnx/YamlSchemaToGoStruct/set"
)

// Convert converts given maps describing schemas to go structs
// args:
//   other []map[interface{}]interface{} - maps describing schemas than
//                                         should not be converted to go structs
//   toConvert []map[interface{}]interface{} - maps describing schemas that
//                                             should be converted to go structs
//   annotationDB string - annotation added to each field in schemas
//   annotationObject string - annotation added to each field in objects
//   suffix string - suffix added to each type name
// return:
//   1. list of go interfaces as strings
//   2. list of go structs as strings
//   3. list of implementations of interfaces as strings
//   4. error during execution
func Convert(
	other,
	toConvert []map[interface{}]interface{},
	rawSuffix,
	interfaceSuffix string,
) (
	generated,
	interfaces,
	structs,
	implementations []string,
	err error,
) {
	var otherSet set.Set
	otherSet, err = parseAll(other)
	if err != nil {
		return
	}
	var toConvertSet set.Set
	toConvertSet, err = parseAll(toConvert)
	if err != nil {
		return
	}
	if err = collectSchemas(toConvertSet, otherSet); err != nil {
		return
	}
	dbObjects := set.New()
	jsonObjects := set.New()
	for _, toConvertSchema := range toConvertSet {
		objectFromSchema, _ := toConvertSchema.(*Schema).collectObjects(1, 0)
		dbObjects.InsertAll(objectFromSchema)
		var object set.Set
		object, err = toConvertSchema.(*Schema).collectObjects(-1, 1)
		if err != nil {
			return
		}
		jsonObjects.InsertAll(object)
	}
	dbObjects.InsertAll(jsonObjects)
	for _, object := range dbObjects.ToArray() {
		item := object.(*item.Object)
		generated = append(generated, item.GenerateInterface(interfaceSuffix))
		interfaces = append(interfaces, item.GenerateMutableInterface(interfaceSuffix, rawSuffix))
		structs = append(structs, item.GenerateStruct(rawSuffix))
		implementations = append(implementations, item.GenerateImplementation(interfaceSuffix, rawSuffix))
	}
	return
}
