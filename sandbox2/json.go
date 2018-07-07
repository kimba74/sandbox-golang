package main

import (
	"fmt"
	"strings"
)

// Element is the top-level interface for the JSON structure
type Element interface {
	GetName() string
	String() string
}

//******************************************************************************
//* JSON Object specifications
//******************************************************************************

// DocumentRoot creates a new json root object
func DocumentRoot() *Object {
	return &Object{name: "", elements: make(map[string]Element)}
}

// Object represents a JSON object
type Object struct {
	name     string
	elements map[string]Element
}

// GetChild looks up and returns a child object
func (o *Object) GetChild(name string) *Object {
	p := o.elements[name]
	if p == nil {
		return nil
	}
	c, ok := p.(*Object)
	if !ok {
		return nil
	}
	return c
}

// GetName returns the name of a json object
func (o *Object) GetName() string {
	return o.name
}

// GetValue looks up and returns a value
func (o *Object) GetValue(name string) *KeyValue {
	p := o.elements[name]
	if p == nil {
		return nil
	}
	v, ok := p.(*KeyValue)
	if !ok {
		return nil
	}
	return v
}

// NewChild creates a new child object
func (o *Object) NewChild(name string) *Object {
	c := &Object{name: name, elements: make(map[string]Element)}
	o.elements[name] = c
	return c
}

// NewValue creates a new value holder
func (o *Object) NewValue(name string) *KeyValue {
	k := &KeyValue{key: name}
	o.elements[name] = k
	return k
}

// String prints out a string representation
func (o *Object) String() string {
	var s string
	if o.name != "" {
		s = fmt.Sprintf("\"%s\" : ", o.name)
	}
	s = s + "{"
	m := ""
	for _, v := range o.elements {
		m = fmt.Sprintf("%s, %s", m, v)
	}
	m = strings.TrimPrefix(m, ", ")
	s = s + m + "}"
	return s
}

//******************************************************************************
//* JSON Value specifications
//******************************************************************************

// KeyValue represents a JSON value
type KeyValue struct {
	key, value string
}

// GetName returns the name of the JSON value
func (kv *KeyValue) GetName() string {
	return kv.key
}

// String prints out a string representation
func (kv *KeyValue) String() string {
	return fmt.Sprintf("\"%s\" : \"%s\"", kv.key, kv.value)
}

// Value sets the actual value
func (kv *KeyValue) Value(value string) {
	kv.value = value
}
