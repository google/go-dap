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
	"sort"
	"strings"
)

// parseRef parses the value of a "$ref" key.
func parseRef(refValue interface{}) string {
	refContents := refValue.(string)
	if !strings.HasPrefix(refContents, "#/definitions/") {
		log.Fatal("want ref to start with '#/definitions/', got ", refValue)
	}

	return refContents[14:]
}

// goFieldName converts a property name from its JSON representation to a Go
// field name.
func goFieldName(jsonPropName string) string {
	//jsonPropName = strings.TrimLeft(jsonPropName, "_")
	clean := strings.ReplaceAll(jsonPropName, "_", " ")
	titled := strings.Title(clean)
	return strings.ReplaceAll(titled, " ", "")
}

// parsePropertyType takes the JSON value of a property field and extracts
// the Go type of the property. For example, given this map:
//
//  {
//   "type": "string",
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
		log.Fatal("property with no type:", propValue)
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
				log.Fatal("property with type array, no items:", propValue)
			}
			propItemsMap := propItems.(map[string]interface{})
			return "[]" + parsePropertyType(propItemsMap)
		case "object":
			return "map[string]string"

		default:
			log.Fatal("unknown property type", propType)
		}

	case []interface{}:
		return "interface{}"

	default:
		log.Fatal("unknown property type", propType)
	}

	panic("unreachable")
}

// emitType emits a single type into a string. It takes the type name and the
// json map representing the type. The type will have a "type" field,
// "properties" etc.
func emitType(name string, desc map[string]interface{}) string {
	var b strings.Builder
	var baseType string

	// A type description can can an "allOf" key, which means it inherits from
	// another type description. Process that, if exists, and then assign desc
	// to the "actual" description for this type.
	if allOfList, ok := desc["allOf"]; ok {
		allOfListSlice := allOfList.([]interface{})
		if len(allOfListSlice) != 2 {
			log.Fatal("want 2 elements in allOf list, got", allOfListSlice)
		}

		refInterface := allOfListSlice[0]
		ref := refInterface.(map[string]interface{})
		baseType = parseRef(ref["$ref"])
		desc = allOfListSlice[1].(map[string]interface{})
	}

	descType, ok := desc["type"]
	if !ok {
		log.Fatal("want description to have 'type', got ", desc)
	}

	descTypeString, ok := descType.(string)
	if !ok {
		log.Fatal("description type not string:", desc)
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
		log.Fatal("want description type to be object or string, got ", desc)
	}

	props, ok := desc["properties"]
	if !ok {
		b.WriteString("}\n")
		return b.String()
	}

	// Stores the properties that are required.
	requiredMap := make(map[string]bool)

	if required, ok := desc["required"]; ok {
		reqSlice := required.([]interface{})
		for _, r := range reqSlice {
			requiredMap[r.(string)] = true
		}
	}

	var bodyTypes []string

	propsMap := props.(map[string]interface{})

	// Sort property names to ensure stable emission order.
	var propNames []string
	for k := range propsMap {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	for _, propName := range propNames {
		propValue := propsMap[propName]
		// For top-level properties, don't emit unless we're in top-level types.
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
			bodyDesc := propValue.(map[string]interface{})

			var propType string
			if ref, ok := bodyDesc["$ref"]; ok {
				propType = parseRef(ref)
			} else {
				propType = name + "Body"
				bodyTypes = append(bodyTypes, emitType(propType, bodyDesc))
			}

			if requiredMap["body"] {
				fmt.Fprintf(&b, "\t%s %s `json:\"body\"`\n", "Body", propType)
			} else {
				fmt.Fprintf(&b, "\t%s %s `json:\"body,omitempty\"`\n", "Body", propType)
			}
		} else {
			propItems := propValue.(map[string]interface{})

			// Go type of this property.
			var goType string

			if propRef, ok := propItems["$ref"]; ok {
				goType = parseRef(propRef)
			} else {
				goType = parsePropertyType(propItems)
			}

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

	if len(bodyTypes) > 0 {
		b.WriteString("\n")
		for _, bt := range bodyTypes {
			b.WriteString(bt)
		}
	}

	return b.String()
}

// definitionKeys returns the keys in the "definitions" map in b, in their
// original order in the .json stream. Based on
// https://github.com/golang/go/issues/27179#issuecomment-415559968
func definitionKeys(b []byte) ([]string, error) {
	var top map[string]json.RawMessage
	if err := json.Unmarshal(b, &top); err != nil {
		log.Fatal(err)
	}

	d := json.NewDecoder(bytes.NewReader(top["definitions"]))
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

var end = errors.New("invalid end of array or object")

func skipValue(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch t {
	case json.Delim('['), json.Delim('{'):
		for {
			if err := skipValue(d); err != nil {
				if err == end {
					break
				}
				return err
			}
		}
	case json.Delim(']'), json.Delim('}'):
		return end
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
// See cmd/gentypes/README.md for details.

`

func main() {
	jp := os.Args[1]
	jpdata, err := ioutil.ReadFile(jp)
	if err != nil {
		log.Fatal(err)
	}

	keys, err := definitionKeys(jpdata)
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(jpdata, &m); err != nil {
		log.Fatal(err)
	}
	mdef := m["definitions"].(map[string]interface{})

	var b strings.Builder
	b.WriteString(preamble)
	for _, typeName := range keys {
		desc := mdef[typeName]
		b.WriteString(emitType(typeName, desc.(map[string]interface{})))
		b.WriteString("\n")
	}

	wholeFile := []byte(b.String())
	formatted, err := format.Source(wholeFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(formatted))
}
