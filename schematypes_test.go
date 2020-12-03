// Copyright 2020 Google LLC
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
	"testing"
)

// makeErrorResponse creates a pre-populated ErrorResponse for testing.
func makeErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			ProtocolMessage: ProtocolMessage{
				Seq:  199,
				Type: "response",
			},
			Command:    "stackTrace",
			RequestSeq: 9,
			Success:    false,
			Message:    "Unable to produce stack trace: \"{e}\"",
		},
		Body: ErrorResponseBody{
			Error: ErrorMessage{
				Id:        2004,
				Format:    "Unable to produce stack trace: \"{e}\"",
				Variables: map[string]string{"e": "Unknown goroutine 1"},
				ShowUser:  true,
			},
		},
	}
}

func TestMessageInterface(t *testing.T) {
	resp := makeErrorResponse()
	f := func(m Message) int {
		return m.GetSeq()
	}
	// Test adherence to the Message interface.
	seq := f(resp)

	if seq != 199 {
		t.Errorf("got seq=%d, want 199", seq)
	}
}

func TestReponseMessageInterface(t *testing.T) {
	resp := makeErrorResponse()
	f := func(rm ResponseMessage) (int, int) {
		return rm.GetSeq(), rm.GetResponse().RequestSeq
	}
	// Test adherence to the ResponseMessage interface.
	seq, rseq := f(resp)

	if seq != 199 {
		t.Errorf("got seq=%d, want 199", seq)
	}
	if rseq != 9 {
		t.Errorf("got ResponseSeq=%d, want 9", rseq)
	}
}

func TestLaunchAttachRequestInterface(t *testing.T) {
	lr := &LaunchRequest{
		Request: Request{
			ProtocolMessage: ProtocolMessage{
				Seq:  19,
				Type: "request",
			},
			Command: "launch",
		},
		Arguments: map[string]interface{}{"foo": "bar"},
	}
	ar := &AttachRequest{
		Request: Request{
			ProtocolMessage: ProtocolMessage{
				Seq:  19,
				Type: "request",
			},
			Command: "attach",
		},
		Arguments: map[string]interface{}{"foo": "bar"},
	}

	f := func(r LaunchAttachRequest) (int, string, interface{}) {
		return r.GetSeq(), r.GetRequest().Command, r.GetArguments()["foo"]
	}
	// Test adherence to the LaunchAttachRequest interface.
	lseq, lcmd, lfoo := f(lr)
	aseq, acmd, afoo := f(ar)

	if lseq != 19 || aseq != 19 {
		t.Errorf("got lseq=%d aseq=%d, want 19", lseq, aseq)
	}
	if lcmd != "launch" || acmd != "attach" {
		t.Errorf("got lcmd=%s acmd=%s, want (\"launch\", \"attach\")", lcmd, acmd)
	}
	if lfoo != "bar" || afoo != "bar" {
		t.Errorf("got lfoo=%s afoo=%s, want \"bar\"", lfoo, afoo)
	}
}
