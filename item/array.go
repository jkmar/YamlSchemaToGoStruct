package item

import (
	"fmt"
	"github.com/zimnx/YamlSchemaToGoStruct/set"
	"github.com/zimnx/YamlSchemaToGoStruct/util"
	"strings"
)

// Array is an implementation of Item interface
type Array struct {
	arrayItem Item
}

// IsNull implementation
func (array *Array) IsNull() bool {
	return false
}

// Type implementation
func (array *Array) Type(suffix string) string {
	return "[]" + array.arrayItem.Type(suffix)
}

// InterfaceType implementation
func (array *Array) InterfaceType(suffix string) string {
	return "[]" + array.arrayItem.InterfaceType(suffix)
}

// AddProperties implementation
func (array *Array) AddProperties(set set.Set, safe bool) error {
	return fmt.Errorf("cannot add properties to an array")
}

// Parse implementation
func (array *Array) Parse(
	prefix string,
	level int,
	required bool,
	data map[interface{}]interface{},
) (err error) {
	next, ok := data["items"].(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf(
			"array %s does not have items",
			prefix,
		)
	}
	objectType, ok := next["type"]
	if !ok {
		return fmt.Errorf(
			"items of array %s do not have a type",
			prefix,
		)
	}
	array.arrayItem, err = CreateItem(objectType)
	if err != nil {
		return fmt.Errorf("array %s: %v", prefix, err)
	}
	return array.arrayItem.Parse(prefix, level, required, next)
}

// CollectObjects implementation
func (array *Array) CollectObjects(limit, offset int) (set.Set, error) {
	return array.arrayItem.CollectObjects(limit, offset)
}

// CollectProperties implementation
func (array *Array) CollectProperties(limit, offset int) (set.Set, error) {
	return array.arrayItem.CollectProperties(limit, offset)
}

// GenerateSetter implementation
func (array *Array) GenerateSetter(
	variable string,
	argument string,
	depth int,
) string {
	if _, ok := array.arrayItem.(*PlainItem); ok {
		return fmt.Sprintf(
			"%s%s = %s",
			strings.Repeat("\t", depth),
			variable,
			argument,
		)
	}
	index := arrayIndex(depth)
	return fmt.Sprintf(
		"%s%s = make(%s, len(%s))\n%sfor %c := range %s {\n%s\n%s}",
		strings.Repeat("\t", depth),
		variable,
		array.Type(""),
		argument,
		strings.Repeat("\t", depth),
		util.IndexVariable(depth),
		argument,
		array.arrayItem.GenerateSetter(
			variable+index,
			argument+index,
			depth+1,
		),
		strings.Repeat("\t", depth),
	)
}

func arrayIndex(depth int) string {
	return fmt.Sprintf("[%c]", util.IndexVariable(depth))
}
