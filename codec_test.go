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

package dap

import (
	"reflect"
	"testing"
)

var initializeRequestString = `{"command":"initialize","arguments":{"clientID":"vscode","clientName":"Visual Studio Code","adapterID":"go","pathFormat":"path","linesStartAt1":true,"columnsStartAt1":true,"supportsVariableType":true,"supportsVariablePaging":true,"supportsRunInTerminalRequest":true,"locale":"en-us"},"type":"request","seq":1}`
var initializeRequestStruct = InitializeRequest{
	Request: Request{
		ProtocolMessage: ProtocolMessage{
			Type: "request",
			Seq:  1,
		},
		Command: "initialize",
	},
	Arguments: InitializeRequestArguments{
		ClientID:                     "vscode",
		ClientName:                   "Visual Studio Code",
		AdapterID:                    "go",
		PathFormat:                   "path",
		LinesStartAt1:                true,
		ColumnsStartAt1:              true,
		SupportsVariableType:         true,
		SupportsVariablePaging:       true,
		SupportsRunInTerminalRequest: true,
		Locale:                       "en-us",
	},
}

var initializeResponseString = `{"seq":1,"type":"response","request_seq":2,"command":"initialize","success":true,"body":{"supportsConfigurationDoneRequest":true,"supportsSetVariable":true}}`
var initializeResponseStruct = InitializeResponse{
	Response: Response{
		ProtocolMessage: ProtocolMessage{
			Type: "response",
			Seq:  1,
		},
		Command:    "initialize",
		Success:    true,
		RequestSeq: 2,
	},
	Body: Capabilities{
		SupportsConfigurationDoneRequest: true,
		SupportsSetVariable:              true,
	},
}

var initializedEventString = `{"seq":1,"type":"event","event":"initialized"}`
var initializedEventStruct = InitializedEvent{
	Event: Event{
		ProtocolMessage: ProtocolMessage{
			Type: "event",
			Seq:  1,
		},
		Event: "initialized",
	},
}

func Test_DecodeProtocolMessage(t *testing.T) {
	tests := []struct {
		data    string
		wantMsg interface{}
		wantErr string
	}{
		// ProtocolMessage
		{``, nil, "unexpected end of JSON input"},
		{`,`, nil, "invalid character ',' looking for beginning of value"},
		{`{}`, ProtocolMessage{}, "ProtocolMessage type '' is not supported"},
		{`{"a": 1}`, ProtocolMessage{}, "ProtocolMessage type '' is not supported"},
		{`{"type":"foo", "seq": 2}`, ProtocolMessage{2, "foo"}, "ProtocolMessage type 'foo' is not supported"},
		// Request
		{`{"type":"request"}`, Request{ProtocolMessage{0, "request"}, ""}, "Request command '' is not supported"},
		{initializeRequestString, initializeRequestStruct, ""},
		// Response
		{`{"type":"response","success":true}`, Response{ProtocolMessage{0, "response"}, "", "", 0, true}, "Response command '' is not supported"},
		{initializeResponseString, initializeResponseStruct, ""},
		// TODO(polina): add ErrorResponse test case
		// Event
		{`{"type":"event"}`, Event{ProtocolMessage{0, "event"}, ""}, "Event event '' is not supported"},
		{initializedEventString, initializedEventStruct, ""},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			msg, err := DecodeProtocolMessage([]byte(test.data))
			// Partial structs maybe returned on error, so always check message
			if !reflect.DeepEqual(msg, test.wantMsg) {
				t.Errorf("got message=%#v, want %#v", msg, test.wantMsg)
			}
			if err != nil {
				if err.Error() != test.wantErr {
					t.Errorf("got error=%#v, want %q", err, test.wantErr)
				}
			} else {
				if test.wantErr != "" {
					t.Errorf("got error=nil, want %#q", test.wantErr)
				}
			}
		})
	}
}
