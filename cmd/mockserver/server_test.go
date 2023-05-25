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

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/google/go-dap"
)

var initializeRequest = []byte(`{"seq":1,"type":"request","command":"initialize","arguments":{"clientID":"vscode","clientName":"Visual Studio Code","adapterID":"go","pathFormat":"path","linesStartAt1":true,"columnsStartAt1":true,"supportsVariableType":true,"supportsVariablePaging":true,"supportsRunInTerminalRequest":true,"locale":"en-us"}}`)
var initializedEvent = []byte(`{"seq":0,"type":"event","event":"initialized"}`)
var initializeResponse = []byte(`{"seq":0,"type":"response","request_seq":1,"success":true,"command":"initialize","body":{"supportsConfigurationDoneRequest":true}}`)

var launchRequest = []byte(`{"seq":2,"type":"request","command":"launch","arguments":{"noDebug": true,"name":"Launch","type":"go","request":"launch","mode":"debug","program":"/Users/foo/go/src/hello","__sessionId":"4c88179f-1202-4f75-9e67-5bf535cde30a","args":["somearg"],"env":{"GOPATH":"/Users/foo/go","HOME":"/Users/foo","SHELL":"/bin/bash"}}}`)
var launchResponse = []byte(`{"seq":0,"type":"response","request_seq":2,"success":true,"command":"launch"}`)

var setBreakpointsRequest = []byte(`{"seq":3,"type":"request","command":"setBreakpoints","arguments":{"source":{"name":"hello.go","path":"/Users/foo/go/src/hello/hello.go"},"lines":[7],"breakpoints":[{"line":7}],"sourceModified":false}}`)
var setBreakpointsResponse = []byte(`{"seq":0,"type":"response","request_seq":3,"success":true,"command":"setBreakpoints","body":{"breakpoints":[{"verified":true,"line":7}]}}`)

var setExceptionBreakpointsRequest = []byte(`{"seq":4,"type":"request","command":"setExceptionBreakpoints","arguments":{"filters":[]}}`)
var setExceptionBreakpointsResponse = []byte(`{"seq":0,"type":"response","request_seq":4,"success":true,"command":"setExceptionBreakpoints","body":{}}`)

var configurationDoneRequest = []byte(`{"seq":5,"type":"request","command":"configurationDone"}`)
var threadEvent = []byte(`{"seq":0,"type":"event","event":"thread","body":{"reason":"started","threadId":1}}`)
var configurationDoneResponse = []byte(`{"seq":0,"type":"response","request_seq":5,"success":true,"command":"configurationDone"}`)

var stoppedEvent = []byte(`{"seq":0,"type":"event","event":"stopped","body":{"reason":"breakpoint","threadId":1,"allThreadsStopped":true}}`)

var threadsRequest = []byte(`{"seq":6,"type":"request","command":"threads"}`)
var threadsResponse = []byte(`{"seq":0,"type":"response","request_seq":6,"success":true,"command":"threads","body":{"threads":[{"id":1,"name":"main"}]}}`)

var stackTraceRequest = []byte(`{"seq":7,"type":"request","command":"stackTrace","arguments":{"threadId":1,"startFrame":0,"levels":20}}`)
var stackTraceResponse = []byte(`{"seq":0,"type":"response","request_seq":7,"success":true,"command":"stackTrace","body":{"stackFrames":[{"id":1000,"name":"main.main","source":{"name":"hello.go","path":"/Users/foo/go/src/hello/hello.go"},"line":5,"column":0}],"totalFrames":1}}`)

var scopesRequest = []byte(`{"seq":8,"type":"request","command":"scopes","arguments":{"frameId":1000}}`)
var scopesResponse = []byte(`{"seq":0,"type":"response","request_seq":8,"success":true,"command":"scopes","body":{"scopes":[{"name":"Local","variablesReference":1000,"expensive":false},{"name":"Global","variablesReference":1001,"expensive":true}]}}`)

var variablesRequest = []byte(`{"seq":9,"type":"request","command":"variables","arguments":{"variablesReference":1000}}`)
var variablesResponse = []byte(`{"seq":0,"type":"response","request_seq":9,"success":true,"command":"variables","body":{"variables":[{"name":"i","value":"18434528","evaluateName":"i","variablesReference":0}]}}`)

var continueRequest = []byte(`{"seq":10,"type":"request","command":"continue","arguments":{"threadId":1}}`)
var continueResponse = []byte(`{"seq":0,"type":"response","request_seq":10,"success":true,"command":"continue","body":{"allThreadsContinued":false}}`)

var terminatedEvent = []byte(`{"seq":0,"type":"event","event":"terminated","body":{}}`)
var disconnectRequest = []byte(`{"seq":11,"type":"request","command":"disconnect","arguments":{"restart":false}}`)
var disconnectResponse = []byte(`{"seq":0,"type":"response","request_seq":11,"success":true,"command":"disconnect"}`)

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
	go func() {
		err := server(port)
		if err != nil {
			log.Fatal("Could not start server:", err)
		}
	}()
	// Give server time to start listening before clients connect
	time.Sleep(100 * time.Millisecond)

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
		log.Fatal("Could not connect to server:", err)
	}
	defer func() {
		t.Log("Closing connection to server at", conn.RemoteAddr())
		conn.Close()
	}()
	t.Log("Connected to server at", conn.RemoteAddr())

	r := bufio.NewReader(conn)

	// Start up

	dap.WriteBaseMessage(conn, initializeRequest)
	expectMessage(t, r, initializedEvent)
	expectMessage(t, r, initializeResponse)

	dap.WriteBaseMessage(conn, launchRequest)
	expectMessage(t, r, launchResponse)

	dap.WriteBaseMessage(conn, setBreakpointsRequest)
	expectMessage(t, r, setBreakpointsResponse)
	dap.WriteBaseMessage(conn, setExceptionBreakpointsRequest)
	expectMessage(t, r, setExceptionBreakpointsResponse)

	dap.WriteBaseMessage(conn, configurationDoneRequest)
	expectMessage(t, r, threadEvent)
	expectMessage(t, r, configurationDoneResponse)

	// Stop on preconfigured breakpoint & Continue

	expectMessage(t, r, stoppedEvent)

	dap.WriteBaseMessage(conn, threadsRequest)
	expectMessage(t, r, threadsResponse)

	dap.WriteBaseMessage(conn, stackTraceRequest)
	expectMessage(t, r, stackTraceResponse)

	dap.WriteBaseMessage(conn, scopesRequest)
	expectMessage(t, r, scopesResponse)

	// Processing of this request will be slow due to a fake delay.
	// Send the next request right away and confirm that processing
	// happens concurrently and the two responses are received
	// out of order.
	dap.WriteBaseMessage(conn, variablesRequest)
	dap.WriteBaseMessage(conn, continueRequest)
	expectMessage(t, r, continueResponse)
	expectMessage(t, r, variablesResponse)

	// Shut down

	expectMessage(t, r, terminatedEvent)
	dap.WriteBaseMessage(conn, disconnectRequest)
	expectMessage(t, r, disconnectResponse)
}
