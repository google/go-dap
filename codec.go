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
	case "cancel":
		var r CancelRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "runInTerminal":
		var r RunInTerminalRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "initialize":
		var r InitializeRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "configurationDone":
		var r ConfigurationDoneRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "launch":
		var r LaunchRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "attach":
		var r AttachRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "restart":
		var r RestartRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "disconnect":
		var r DisconnectRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "terminate":
		var r TerminateRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "breakpointLocations":
		var r BreakpointLocationsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setBreakpoints":
		var r SetBreakpointsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setFunctionBreakpoints":
		var r SetFunctionBreakpointsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setExceptionBreakpoints":
		var r SetExceptionBreakpointsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "dataBreakpointInfo":
		var r DataBreakpointInfoRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setDataBreakpoints":
		var r SetDataBreakpointsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "continue":
		var r ContinueRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "next":
		var r NextRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepIn":
		var r StepInRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepOut":
		var r StepOutRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepBack":
		var r StepBackRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "reverseContinue":
		var r ReverseContinueRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "restartFrame":
		var r RestartFrameRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "goto":
		var r GotoRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "pause":
		var r PauseRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "stackTrace":
		var r StackTraceRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "scopes":
		var r ScopesRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "variables":
		var r VariablesRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setVariable":
		var r SetVariableRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "source":
		var r SourceRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "threads":
		var r ThreadsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "terminateThreads":
		var r TerminateThreadsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "modules":
		var r ModulesRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "loadedSources":
		var r LoadedSourcesRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "evaluate":
		var r EvaluateRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "setExpression":
		var r SetExpressionRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepInTargets":
		var r StepInTargetsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "gotoTargets":
		var r GotoTargetsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "completions":
		var r CompletionsRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "exceptionInfo":
		var r ExceptionInfoRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "readMemory":
		var r ReadMemoryRequest
		err := json.Unmarshal(data, &r)
		return r, err
	case "disassemble":
		var r DisassembleRequest
		err := json.Unmarshal(data, &r)
		return r, err

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
	case "cancel":
		var r CancelResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "runInTerminal":
		var r RunInTerminalResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "initialize":
		var r InitializeResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "configurationDone":
		var r ConfigurationDoneResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "launch":
		var r LaunchResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "attach":
		var r AttachResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "restart":
		var r RestartResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "disconnect":
		var r DisconnectResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "terminate":
		var r TerminateResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "breakpointLocations":
		var r BreakpointLocationsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setBreakpoints":
		var r SetBreakpointsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setFunctionBreakpoints":
		var r SetFunctionBreakpointsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setExceptionBreakpoints":
		var r SetExceptionBreakpointsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "dataBreakpointInfo":
		var r DataBreakpointInfoResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setDataBreakpoints":
		var r SetDataBreakpointsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "continue":
		var r ContinueResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "next":
		var r NextResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepIn":
		var r StepInResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepOut":
		var r StepOutResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepBack":
		var r StepBackResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "reverseContinue":
		var r ReverseContinueResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "restartFrame":
		var r RestartFrameResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "goto":
		var r GotoResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "pause":
		var r PauseResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "stackTrace":
		var r StackTraceResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "scopes":
		var r ScopesResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "variables":
		var r VariablesResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setVariable":
		var r SetVariableResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "source":
		var r SourceResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "threads":
		var r ThreadsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "terminateThreads":
		var r TerminateThreadsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "modules":
		var r ModulesResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "loadedSources":
		var r LoadedSourcesResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "evaluate":
		var r EvaluateResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "setExpression":
		var r SetExpressionResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "stepInTargets":
		var r StepInTargetsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "gotoTargets":
		var r GotoTargetsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "completions":
		var r CompletionsResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "exceptionInfo":
		var r ExceptionInfoResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "readMemory":
		var r ReadMemoryResponse
		err := json.Unmarshal(data, &r)
		return r, err
	case "disassemble":
		var r DisassembleResponse
		err := json.Unmarshal(data, &r)
		return r, err
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
		var e InitializedEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "stopped":
		var e StoppedEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "continued":
		var e ContinuedEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "exited":
		var e ExitedEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "terminated":
		var e TerminatedEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "thread":
		var e ThreadEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "output":
		var e OutputEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "breakpoint":
		var e BreakpointEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "module":
		var e ModuleEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "loadedSource":
		var e LoadedSourceEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "process":
		var e ProcessEvent
		err := json.Unmarshal(data, &e)
		return e, err
	case "capabilities":
		var e CapabilitiesEvent
		err := json.Unmarshal(data, &e)
		return e, err
	default:
		return event, &DecodeProtocolMessageFieldError{"Event", "event", event.Event}
	}
}
