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

// This file contains high level utilities for reading and writing
// DAP messages.

// TODO(polina): move to 'dap' package? Define dap.ReadWriter?

package main

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/google/go-dap"
)

func writeProtocolMessage(w io.Writer, message dap.Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = dap.WriteBaseMessage(w, b)
	return err
}

func readProtocolMessage(r *bufio.Reader) (dap.Message, error) {
	content, err := dap.ReadBaseMessage(r)
	if err != nil {
		return nil, err
	}
	message, err := dap.DecodeProtocolMessage(content)
	if err != nil {
		return nil, err
	}
	return message, nil
}
