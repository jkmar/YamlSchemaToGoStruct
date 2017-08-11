package item

import (
	"fmt"
	"github.com/zimnx/YamlSchemaToGoStruct/set"
	"github.com/zimnx/YamlSchemaToGoStruct/util"
)

// PlainItem is an implementation of Item interface
type PlainItem struct {
	null     bool
	itemType string
}

// IsNull implementation
func (plainItem *PlainItem) IsNull() bool {
	return plainItem.null
}

// Type implementation
func (plainItem *PlainItem) Type(suffix string) string {
	return plainItem.itemType
}

// AddProperties implementation
func (plainItem *PlainItem) AddProperties(set set.Set, safe bool) error {
	return fmt.Errorf("cannot add properties to a plain item")
}

// Parse implementation
func (plainItem *PlainItem) Parse(
	prefix string,
	level int,
	required bool,
	data map[interface{}]interface{},
) (err error) {
	objectType, ok := data["type"]
	if !ok {
		return fmt.Errorf(
			"item %s does not have a type",
			prefix,
		)
	}
	plainItem.itemType, plainItem.null, err = util.ParseType(objectType)
	if err != nil {
		err = fmt.Errorf(
			"item %s: %v",
			prefix,
			err,
		)
	}

	if !required {
		if _, ok = data["default"]; !ok {
			plainItem.null = true
		}
	}

	return
}

// CollectObjects implementation
func (plainItem *PlainItem) CollectObjects(limit, offset int) (set.Set, error) {
	return nil, nil
}

// CollectProperties implementation
func (plainItem *PlainItem) CollectProperties(limit, offset int) (set.Set, error) {
	return nil, nil
}

func (plainItem *PlainItem) GenerateSetter(prefix, arg string) string {
	return fmt.Sprintf(
		"%s = %s",
		prefix,
		arg,
	)
}
