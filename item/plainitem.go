package item

import (
	"fmt"
	"github.com/zimnx/YamlSchemaToGoStruct/hash"
	"github.com/zimnx/YamlSchemaToGoStruct/name"
	"github.com/zimnx/YamlSchemaToGoStruct/set"
	"github.com/zimnx/YamlSchemaToGoStruct/util"
)

// PlainItem is an implementation of Item interface
type PlainItem struct {
	required bool
	null     bool
	itemType string
}

// Copy implementation
func (plainItem *PlainItem) Copy() Item {
	newItem := *plainItem
	return &newItem
}

// ToString implementation
func (plainItem *PlainItem) ToString() string {
	return fmt.Sprintf("#%s,%v", plainItem.itemType, plainItem.IsNull())
}

// Compress implementation
func (plainItem *PlainItem) Compress(hash.IHashable, hash.IHashable) {
}

// GetChildren implementation
func (plainItem *PlainItem) GetChildren() []hash.IHashable {
	return nil
}

// ChangeName implementation
func (plainItem *PlainItem) ChangeName(mark name.Mark) {
}

// ContainsObject implementation
func (plainItem *PlainItem) ContainsObject() bool {
	return false
}

// IsNull implementation
func (plainItem *PlainItem) IsNull() bool {
	return plainItem.null || !plainItem.required
}

// MakeRequired implementation
func (plainItem *PlainItem) MakeRequired() {
	plainItem.required = true
}

// Type implementation
func (plainItem *PlainItem) Type(suffix string) string {
	return plainItem.itemType
}

// InterfaceType implementation
func (plainItem *PlainItem) InterfaceType(suffix string) string {
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

	if _, ok = data["default"]; ok || required {
		plainItem.required = true
	}
	//if !required {
	//	if _, ok = data["default"]; !ok {
	//		plainItem.required = true
	//	}
	//}

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

// GenerateGetter implementation
func (plainItem *PlainItem) GenerateGetter(
	variable,
	argument,
	interfaceSuffix string,
	depth int,
) string {
	return fmt.Sprintf(
		"%s%s %s",
		util.Indent(depth),
		util.ResultPrefix(argument, depth, false),
		variable,
	)
}

// GenerateSetter implementation
func (plainItem *PlainItem) GenerateSetter(
	variable,
	argument,
	typeSuffix string,
	depth int,
) string {
	return fmt.Sprintf(
		"%s%s = %s",
		util.Indent(depth),
		variable,
		argument,
	)
}
