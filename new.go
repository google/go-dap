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

// This file defines helpers for building base protocol messages.

package dap

// newRequest builds a Request struct with the specified fields.
func newRequest(seq int, command string) Request {
	return Request{
		ProtocolMessage: ProtocolMessage{
			Type: "request",
			Seq:  seq,
		},
		Command: command,
	}
}

// newEvent builds an Event struct with the specified fields.
func newEvent(seq int, event string) Event {
	return Event{
		ProtocolMessage: ProtocolMessage{
			Seq:  seq,
			Type: "event",
		},
		Event: event,
	}
}

// newResponse builds a Response struct with the specified fields.
func newResponse(seq int, requestSeq int, command string, success bool) Response {
	return Response{
		ProtocolMessage: ProtocolMessage{
			Seq:  seq,
			Type: "response",
		},
		Command:    command,
		RequestSeq: requestSeq,
		Success:    success,
	}
}

// newErrorResponse builds an ErrorResponse struct with the specified fields.
func newErrorResponse(seq int, requestSeq int, command string, message string) ErrorResponse {
	er := ErrorResponse{
		Response: newResponse(seq, requestSeq, command, false),
	}
	er.Message = message
	return er
}
