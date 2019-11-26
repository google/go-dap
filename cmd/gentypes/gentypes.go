// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// gentypes generates Go types from debugProtocol.json
//
// Usage:
//
// $ gentypes <path to debugProtocol.json>
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// parseRef parses the value of a "$ref" key.
// For example "#definitions/ProtocolMessage" => "ProtocolMessage".
func parseRef(refValue interface{}) string {
	refContents := refValue.(string)
	if !strings.HasPrefix(refContents, "#/definitions/") {
		log.Fatal("want ref to start with '#/definitions/', got ", refValue)
	}

	return refContents[14:]
}

// goFieldName converts a property name from its JSON representation to an
// exported Go field name.
// For example "__some_property_name" => "SomePropertyName".
func goFieldName(jsonPropName string) string {
	clean := strings.ReplaceAll(jsonPropName, "_", " ")
	titled := strings.Title(clean)
	return strings.ReplaceAll(titled, " ", "")
}

// parsePropertyType takes the JSON value of a property field and extracts
// the Go type of the property. For example, given this map:
//
//  {
//    "type": "string",
//    "description": "The command to execute."
//  },
//
// It will emit "string".
func parsePropertyType(propValue map[string]interface{}) string {
	if ref, ok := propValue["$ref"]; ok {
		return parseRef(ref)
	}

	propType, ok := propValue["type"]
	if !ok {
		log.Fatal("property with no type or ref:", propValue)
	}

	switch propType.(type) {
	case string:
		switch propType {
		case "string":
			return "string"
		case "integer":
			return "int"
		case "boolean":
			return "bool"
		case "array":
			propItems, ok := propValue["items"]
			if !ok {
				log.Fatal("missing items type for property of array type:", propValue)
			}
			propItemsMap := propItems.(map[string]interface{})
			return "[]" + parsePropertyType(propItemsMap)
		case "object":
			// When the type of a property is "object", we'll emit a map with a string
			// key and a value type that depends on the type of the
			// additionalProperties field.
			additionalProps, ok := propValue["additionalProperties"]
			if !ok {
				log.Fatal("missing additionalProperties field when type=object:", propValue)
			}
			valueType := parsePropertyType(additionalProps.(map[string]interface{}))
			return fmt.Sprintf("map[string]%v", valueType)
		default:
			log.Fatal("unknown property type value", propType)
		}

	case []interface{}:
		return "interface{}"

	default:
		log.Fatal("unknown property type", propType)
	}

	panic("unreachable")
}

// maybeParseInheritance helps parse types that inherit from other types.
// A type description can have an "allOf" key, which means it inherits from
// another type description. Returns the name of the base type specified in
// allOf, and the description of the inheriting type. If there is no "allOf",
// returns an empty baseTypeName and descMap itself.
func maybeParseInheritance(descMap map[string]json.RawMessage) (baseTypeName string, baseTypeJson map[string]json.RawMessage) {
	allOfListJson, ok := descMap["allOf"]
	if !ok {
		return "", descMap
	}

	var sliceAllOfJson []json.RawMessage
	if err := json.Unmarshal(allOfListJson, &sliceAllOfJson); err != nil {
		log.Fatal(err)
	}
	if len(sliceAllOfJson) != 2 {
		log.Fatal("want 2 elements in allOf list, got", sliceAllOfJson)
	}

	var baseTypeRef map[string]interface{}
	if err := json.Unmarshal(sliceAllOfJson[0], &baseTypeRef); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(sliceAllOfJson[1], &baseTypeJson); err != nil {
		log.Fatal(err)
	}
	return parseRef(baseTypeRef["$ref"]), baseTypeJson
}

// emitToplevelType emits a single type into a string. It takes the type name
// and a serialized json object representing the type. The json representation
// will have fields: "type", "properties" etc.
func emitToplevelType(name string, descJson json.RawMessage) string {
	var b strings.Builder
	var baseType string

	// We don't parse the description all the way to map[string]interface{}
	// because we have to retain the original JSON-order of properties (in this
	// type as well as any nested types like "body").
	var descMap map[string]json.RawMessage
	if err := json.Unmarshal(descJson, &descMap); err != nil {
		log.Fatal(err)
	}
	baseType, descMap = maybeParseInheritance(descMap)

	typeJson, ok := descMap["type"]
	if !ok {
		log.Fatal("want description to have 'type', got ", descMap)
	}

	var descTypeString string
	if err := json.Unmarshal(typeJson, &descTypeString); err != nil {
		log.Fatal(err)
	}

	if descTypeString == "string" {
		fmt.Fprintf(&b, "type %s string\n", name)
		return b.String()
	} else if descTypeString == "object" {
		fmt.Fprintf(&b, "type %s struct {\n", name)
		if len(baseType) > 0 {
			fmt.Fprintf(&b, "\t%s\n\n", baseType)
		}
	} else {
		log.Fatal("want description type to be object or string, got ", descTypeString)
	}

	var propsMapOfJson map[string]json.RawMessage
	if propsJson, ok := descMap["properties"]; ok {
		if err := json.Unmarshal(propsJson, &propsMapOfJson); err != nil {
			log.Fatal(err)
		}
	} else {
		b.WriteString("}\n")
		return b.String()
	}

	propsNamesInOrder, err := keysInOrder(descMap["properties"])
	if err != nil {
		log.Fatal(err)
	}

	// Stores the properties that are required.
	requiredMap := make(map[string]bool)

	if requiredJson, ok := descMap["required"]; ok {
		var required []interface{}
		if err := json.Unmarshal(requiredJson, &required); err != nil {
			log.Fatal(err)
		}
		for _, r := range required {
			requiredMap[r.(string)] = true
		}
	}

	// Some types will have a "body" which should be emitted as a separate type.
	// Since we can't emit a whole new Go type while in the middle of emitting
	// another type, we save it for later and emit it after the current type is
	// done.
	bodyType := ""

	for _, propName := range propsNamesInOrder {
		// The JSON schema is designed for the TypeScript type system, where a
		// subclass can redefine a field in a superclass with a refined type (such
		// as specific values for a field). To ensure we emit Go structs that can
		// be unmarshaled from JSON messages properly, we must limit each field
		// to appear only once in hierarchical types.
		if propName == "type" && (name == "Request" || name == "Response" || name == "Event") {
			continue
		}
		if propName == "command" && name != "Request" && name != "Response" {
			continue
		}
		if propName == "event" && name != "Event" {
			continue
		}
		if propName == "arguments" && name == "Request" {
			continue
		}

		if propName == "body" {
			if name == "Response" || name == "Event" {
				continue
			}
			var bodyDesc map[string]interface{}
			if err := json.Unmarshal(propsMapOfJson[propName], &bodyDesc); err != nil {
				log.Fatal(err)
			}

			var propType string
			if ref, ok := bodyDesc["$ref"]; ok {
				propType = parseRef(ref)
			} else {
				propType = name + "Body"

				if bodyType == "" {
					bodyType = emitToplevelType(propType, propsMapOfJson["body"])
				} else {
					log.Fatalf("have body type %s, see another body in %s\n", bodyType, propType)
				}
			}

			if requiredMap["body"] {
				fmt.Fprintf(&b, "\t%s %s `json:\"body\"`\n", "Body", propType)
			} else {
				fmt.Fprintf(&b, "\t%s %s `json:\"body,omitempty\"`\n", "Body", propType)
			}
		} else {
			var propDesc map[string]interface{}
			if err := json.Unmarshal(propsMapOfJson[propName], &propDesc); err != nil {
				log.Fatal(err)
			}

			// Go type of this property.
			goType := parsePropertyType(propDesc)

			jsonTag := fmt.Sprintf("`json:\"%s", propName)
			if requiredMap[propName] {
				jsonTag += "\"`"
			} else {
				jsonTag += ",omitempty\"`"
			}
			fmt.Fprintf(&b, "\t%s %s %s\n", goFieldName(propName), goType, jsonTag)
		}
	}

	b.WriteString("}\n")

	if len(bodyType) > 0 {
		b.WriteString("\n")
		b.WriteString(bodyType)
	}

	return b.String()
}

// keysInOrder returns the keys in json object in b, in their original order.
// Based on https://github.com/golang/go/issues/27179#issuecomment-415559968
func keysInOrder(b []byte) ([]string, error) {
	d := json.NewDecoder(bytes.NewReader(b))
	t, err := d.Token()
	if err != nil {
		return nil, err
	}
	if t != json.Delim('{') {
		return nil, errors.New("expected start of object")
	}
	var keys []string
	for {
		t, err := d.Token()
		if err != nil {
			return nil, err
		}
		if t == json.Delim('}') {
			return keys, nil
		}
		keys = append(keys, t.(string))
		if err := skipValue(d); err != nil {
			return nil, err
		}
	}
}

var errEnd = errors.New("invalid end of array or object")

func skipValue(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch t {
	case json.Delim('['), json.Delim('{'):
		for {
			if err := skipValue(d); err != nil {
				if err == errEnd {
					break
				}
				return err
			}
		}
	case json.Delim(']'), json.Delim('}'):
		return errEnd
	}
	return nil
}

const preamble = `// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// DO NOT EDIT: This file is auto-generated.
// DAP spec: https://microsoft.github.io/debug-adapter-protocol/specification
// See cmd/gentypes/README.md for additional details.

package dap

`

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	inputFilename := os.Args[1]
	inputData, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]json.RawMessage
	if err := json.Unmarshal(inputData, &m); err != nil {
		log.Fatal(err)
	}
	var typeMap map[string]json.RawMessage
	if err := json.Unmarshal(m["definitions"], &typeMap); err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	b.WriteString(preamble)

	typeNames, err := keysInOrder(m["definitions"])
	if err != nil {
		log.Fatal(err)
	}

	for _, typeName := range typeNames {
		b.WriteString(emitToplevelType(typeName, typeMap[typeName]))
		b.WriteString("\n")
	}

	wholeFile := []byte(b.String())
	formatted, err := format.Source(wholeFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(formatted))
}
