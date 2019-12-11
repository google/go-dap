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

// This file contains utilities for decoding JSON-encoded bytes into DAP message.
// TODO(polina): encoding utilities

package dap

import (
	"encoding/json"
	"fmt"
)

// DecodeProtocolMessageFieldError describes which JSON attribute
// has an unsupported value that the decoding cannot handle.
type DecodeProtocolMessageFieldError struct {
	SubType    string
	FieldName  string
	FieldValue string
}

func (e *DecodeProtocolMessageFieldError) Error() string {
	return fmt.Sprintf("%s %s '%s' is not supported", e.SubType, e.FieldName, e.FieldValue)
}

// DecodeProtocolMessage parses the JSON-encoded data and returns the result of
// the appropriate type within the ProtocolMessage hierarchy. If message type,
// command, etc cannot be cast, returns DecodeProtocolMessageFieldError.
// See also godoc for json.Unmarshal, which is used for underlying decoding.
func DecodeProtocolMessage(data []byte) (Message, error) {
	var protomsg ProtocolMessage
	if err := json.Unmarshal(data, &protomsg); err != nil {
		return nil, err
	}
	switch protomsg.Type {
	case "request":
		return decodeRequest(data)
	case "response":
		return decodeResponse(data)
	case "event":
		return decodeEvent(data)
	default:
		return protomsg, &DecodeProtocolMessageFieldError{"ProtocolMessage", "type", protomsg.Type}
	}
}

func decodeRequest(data []byte) (Message, error) {
	var request Request
	if err := json.Unmarshal(data, &request); err != nil {
		return request, err
	}
	switch request.Command {
	case "initialize":
		var ir InitializeRequest
		err := json.Unmarshal(data, &ir)
		return ir, err
	case "launch":
		panic("Not supported yet")
	case "attach":
		panic("Not supported yet")
	case "disconnect":
		panic("Not supported yet")
	case "terminate":
		panic("Not supported yet")
	case "restart":
		panic("Not supported yet")
	case "setBreakpoints":
		panic("Not supported yet")
	case "setFunctionBreakpoints":
		panic("Not supported yet")
	case "setExceptionBreakpoints":
		panic("Not supported yet")
	case "configurationDone":
		panic("Not supported yet")
	case "continue":
		panic("Not supported yet")
	case "next":
		panic("Not supported yet")
	case "stepIn":
		panic("Not supported yet")
	case "stepOut":
		panic("Not supported yet")
	case "stepBack":
		panic("Not supported yet")
	case "reverseContinue":
		panic("Not supported yet")
	case "restartFrame":
		panic("Not supported yet")
	case "goto":
		panic("Not supported yet")
	case "pause":
		panic("Not supported yet")
	case "stackTrace":
		panic("Not supported yet")
	case "scopes":
		panic("Not supported yet")
	case "variables":
		panic("Not supported yet")
	case "setVariable":
		panic("Not supported yet")
	case "setExpression":
		panic("Not supported yet")
	case "source":
		panic("Not supported yet")
	case "threads":
		panic("Not supported yet")
	case "terminateThreads":
		panic("Not supported yet")
	case "evaluate":
		panic("Not supported yet")
	case "stepInTargets":
		panic("Not supported yet")
	case "gotoTargets":
		panic("Not supported yet")
	case "completions":
		panic("Not supported yet")
	case "exceptionInfo":
		panic("Not supported yet")
	case "loadedSources":
		panic("Not supported yet")
	case "dataBreakpointInfo":
		panic("Not supported yet")
	case "setDataBreakpoints":
		panic("Not supported yet")
	case "readMemory":
		panic("Not supported yet")
	case "disassemble":
		panic("Not supported yet")
	case "cancel":
		panic("Not supported yet")
	case "breakpointLocations":
		panic("Not supported yet")
	default:
		return request, &DecodeProtocolMessageFieldError{"Request", "command", request.Command}
	}
}

func decodeResponse(data []byte) (Message, error) {
	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(data, &errorResponse); err != nil {
			return response, err
		}
		return errorResponse, nil
	}
	switch response.Command {
	case "initialize":
		var ir InitializeResponse
		err := json.Unmarshal(data, &ir)
		return ir, err
	case "launch":
		panic("Not supported yet")
	case "attach":
		panic("Not supported yet")
	case "disconnect":
		panic("Not supported yet")
	case "terminate":
		panic("Not supported yet")
	case "restart":
		panic("Not supported yet")
	case "setBreakpoints":
		panic("Not supported yet")
	case "setFunctionBreakpoints":
		panic("Not supported yet")
	case "setExceptionBreakpoints":
		panic("Not supported yet")
	case "configurationDone":
		panic("Not supported yet")
	case "continue":
		panic("Not supported yet")
	case "next":
		panic("Not supported yet")
	case "stepIn":
		panic("Not supported yet")
	case "stepOut":
		panic("Not supported yet")
	case "stepBack":
		panic("Not supported yet")
	case "reverseContinue":
		panic("Not supported yet")
	case "restartFrame":
		panic("Not supported yet")
	case "goto":
		panic("Not supported yet")
	case "pause":
		panic("Not supported yet")
	case "stackTrace":
		panic("Not supported yet")
	case "scopes":
		panic("Not supported yet")
	case "variables":
		panic("Not supported yet")
	case "setVariable":
		panic("Not supported yet")
	case "setExpression":
		panic("Not supported yet")
	case "source":
		panic("Not supported yet")
	case "threads":
		panic("Not supported yet")
	case "terminateThreads":
		panic("Not supported yet")
	case "evaluate":
		panic("Not supported yet")
	case "stepInTargets":
		panic("Not supported yet")
	case "gotoTargets":
		panic("Not supported yet")
	case "completions":
		panic("Not supported yet")
	case "exceptionInfo":
		panic("Not supported yet")
	case "loadedSources":
		panic("Not supported yet")
	case "dataBreakpointInfo":
		panic("Not supported yet")
	case "setDataBreakpoints":
		panic("Not supported yet")
	case "readMemory":
		panic("Not supported yet")
	case "disassemble":
		panic("Not supported yet")
	case "cancel":
		panic("Not supported yet")
	case "breakpointLocations":
		panic("Not supported yet")
	default:
		return response, &DecodeProtocolMessageFieldError{"Response", "command", response.Command}
	}
}

func decodeEvent(data []byte) (Message, error) {
	var event Event
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}
	switch event.Event {
	case "initialized":
		var ie InitializedEvent
		err := json.Unmarshal(data, &ie)
		return ie, err
	case "stopped":
		panic("Not supported yet")
	case "continued":
		panic("Not supported yet")
	case "exited":
		panic("Not supported yet")
	case "terminated":
		panic("Not supported yet")
	case "thread":
		panic("Not supported yet")
	case "output":
		panic("Not supported yet")
	case "breakpoint":
		panic("Not supported yet")
	case "module":
		panic("Not supported yet")
	case "loadedSource":
		panic("Not supported yet")
	case "process":
		panic("Not supported yet")
	case "capabilities":
		panic("Not supported yet")
	default:
		return event, &DecodeProtocolMessageFieldError{"Event", "event", event.Event}
	}
}
