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

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"

	"github.com/google/go-dap"
)

var initializeRequest = []byte(`{"seq":1,"type":"request","command":"initialize","arguments":{"clientID":"vscode","clientName":"Visual Studio Code","adapterID":"go","pathFormat":"path","linesStartAt1":true,"columnsStartAt1":true,"supportsVariableType":true,"supportsVariablePaging":true,"supportsRunInTerminalRequest":true,"locale":"en-us"}}`)
var initializedEvent = []byte(`{"seq":0,"type":"event","event":"initialized"}`)
var initializeResponse = []byte(`{"seq":0,"type":"response","request_seq":1,"success":true,"command":"initialize","body":{"supportsConfigurationDoneRequest":true}}`)

func expectMessage(t *testing.T, r *bufio.Reader, want []byte) {
	got, err := dap.ReadBaseMessage(r)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}

func TestServer(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	port := "54321"
	go server(port)

	var wg sync.WaitGroup
	wg.Add(2)
	go client(t, port, &wg)
	go client(t, port, &wg)
	wg.Wait()
}

func client(t *testing.T, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		t.Fatal("Could not connect to server:", err)
	}
	defer func() {
		t.Log("Closing connection to server at", conn.RemoteAddr())
		conn.Close()
	}()
	t.Log("Connected to server at", conn.RemoteAddr())

	r := bufio.NewReader(conn)
	dap.WriteBaseMessage(conn, initializeRequest)
	expectMessage(t, r, initializedEvent)
	expectMessage(t, r, initializeResponse)

	// TODO(polina): test more requests
}
