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

func TestMessageInterface(t *testing.T) {
	var errorResponseStruct = ErrorResponse{
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

	f := func(m Message) int {
		return m.GetSeq()
	}
	seq := f(&errorResponseStruct)

	if seq != 199 {
		t.Errorf("got seq=%d, want 199", seq)
	}
}
