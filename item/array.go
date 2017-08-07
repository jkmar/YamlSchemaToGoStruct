package item

import (
	"fmt"
	"github.com/zimnx/YamlSchemaToGoStruct/set"
)

// Array is an implementation of Item interface
type Array struct {
	item Item
}

// Type implementation
func (item *Array) Type(suffix string) string {
	return "[]" + item.item.Type(suffix)
}

// AddProperties implementation
func (item *Array) AddProperties(set set.Set, safe bool) error {
	return fmt.Errorf("cannot add properties to an array")
}

// Parse implementation
func (item *Array) Parse(prefix string, object map[interface{}]interface{}) (err error) {
	next, ok := object["items"].(map[interface{}]interface{})
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
	item.item, err = CreateItem(objectType)
	if err != nil {
		return fmt.Errorf("array %s: %v", prefix, err)
	}
	return item.item.Parse(prefix, next)
}

// CollectObjects implementation
func (item *Array) CollectObjects(limit, offset int) (set.Set, error) {
	return item.item.CollectObjects(limit, offset)
}

// CollectProperties implementation
func (item *Array) CollectProperties(limit, offset int) (set.Set, error) {
	return item.item.CollectProperties(limit, offset)
}
