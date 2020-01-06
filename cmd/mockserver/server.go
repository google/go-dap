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

// This file defines helpers and request handlers for a dummy server
// that accepts DAP requests and responds with dummy or error responses.
// Fake-supported requests:
// - initialize
// - launch
// - setBreakpoints
// - setExceptionBreakpoints
// - configurationDone
// - threads
// - stackTrace
// - scopes
// - variables
// - continue
// - disconnect
// All other requests result in ErrorResponse's.

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"

	"github.com/google/go-dap"
)

// server starts a server that listens on a specified port
// and blocks indefinitely. This server cannot accept multiple
// client connections at the same time.
func server(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Println("Started server at", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection failed:", err)
			continue
		}
		log.Println("Accepted connection from", conn.RemoteAddr())
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println("Closing connection from", conn.RemoteAddr())
		conn.Close()
	}()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		err := handleRequest(rw)
		// TODO(polina): check for connection vs decoding error?
		if err != nil {
			if err == io.EOF {
				log.Println("No more data to read:", err)
				return
			}
			log.Println("Server error:", err)
			continue // There may be more messages to process
		}
	}
}

func handleRequest(rw *bufio.ReadWriter) error {
	log.Println("Reading request...")
	request, err := dap.ReadProtocolMessage(rw.Reader)
	if err != nil {
		return err
	}
	log.Printf("Received request\n\t%#v\n", request)
	return dispatchRequest(rw, request)
}

func dispatchRequest(rw *bufio.ReadWriter, request dap.Message) error {
	var response dap.Message
	switch request.(type) {
	case dap.InitializeRequest:
		response = onInitializeRequest(rw.Writer, request.(dap.InitializeRequest))
	case dap.LaunchRequest:
		response = onLaunchRequest(rw.Writer, request.(dap.LaunchRequest))
	case dap.AttachRequest:
		response = onAttachRequest(rw.Writer, request.(dap.AttachRequest))
	case dap.DisconnectRequest:
		response = onDisconnectRequest(rw.Writer, request.(dap.DisconnectRequest))
	case dap.TerminateRequest:
		response = onTerminateRequest(rw.Writer, request.(dap.TerminateRequest))
	case dap.RestartRequest:
		response = onRestartRequest(rw.Writer, request.(dap.RestartRequest))
	case dap.SetBreakpointsRequest:
		response = onSetBreakpointsRequest(rw.Writer, request.(dap.SetBreakpointsRequest))
	case dap.SetFunctionBreakpointsRequest:
		response = onSetFunctionBreakpointsRequest(rw.Writer, request.(dap.SetFunctionBreakpointsRequest))
	case dap.SetExceptionBreakpointsRequest:
		response = onSetExceptionBreakpointsRequest(rw.Writer, request.(dap.SetExceptionBreakpointsRequest))
	case dap.ConfigurationDoneRequest:
		response = onConfigurationDoneRequest(rw.Writer, request.(dap.ConfigurationDoneRequest))
	case dap.ContinueRequest:
		response = onContinueRequest(rw.Writer, request.(dap.ContinueRequest))
	case dap.NextRequest:
		response = onNextRequest(rw.Writer, request.(dap.NextRequest))
	case dap.StepInRequest:
		response = onStepInRequest(rw.Writer, request.(dap.StepInRequest))
	case dap.StepOutRequest:
		response = onStepOutRequest(rw.Writer, request.(dap.StepOutRequest))
	case dap.StepBackRequest:
		response = onStepBackRequest(rw.Writer, request.(dap.StepBackRequest))
	case dap.ReverseContinueRequest:
		response = onReverseContinueRequest(rw.Writer, request.(dap.ReverseContinueRequest))
	case dap.RestartFrameRequest:
		response = onRestartFrameRequest(rw.Writer, request.(dap.RestartFrameRequest))
	case dap.GotoRequest:
		response = onGotoRequest(rw.Writer, request.(dap.GotoRequest))
	case dap.PauseRequest:
		response = onPauseRequest(rw.Writer, request.(dap.PauseRequest))
	case dap.StackTraceRequest:
		response = onStackTraceRequest(rw.Writer, request.(dap.StackTraceRequest))
	case dap.ScopesRequest:
		response = onScopesRequest(rw.Writer, request.(dap.ScopesRequest))
	case dap.VariablesRequest:
		response = onVariablesRequest(rw.Writer, request.(dap.VariablesRequest))
	case dap.SetVariableRequest:
		response = onSetVariableRequest(rw.Writer, request.(dap.SetVariableRequest))
	case dap.SetExpressionRequest:
		response = onSetExpressionRequest(rw.Writer, request.(dap.SetExpressionRequest))
	case dap.SourceRequest:
		response = onSourceRequest(rw.Writer, request.(dap.SourceRequest))
	case dap.ThreadsRequest:
		response = onThreadsRequest(rw.Writer, request.(dap.ThreadsRequest))
	case dap.TerminateThreadsRequest:
		response = onTerminateThreadsRequest(rw.Writer, request.(dap.TerminateThreadsRequest))
	case dap.EvaluateRequest:
		response = onEvaluateRequest(rw.Writer, request.(dap.EvaluateRequest))
	case dap.StepInTargetsRequest:
		response = onStepInTargetsRequest(rw.Writer, request.(dap.StepInTargetsRequest))
	case dap.GotoTargetsRequest:
		response = onGotoTargetsRequest(rw.Writer, request.(dap.GotoTargetsRequest))
	case dap.CompletionsRequest:
		response = onCompletionsRequest(rw.Writer, request.(dap.CompletionsRequest))
	case dap.ExceptionInfoRequest:
		response = onExceptionInfoRequest(rw.Writer, request.(dap.ExceptionInfoRequest))
	case dap.LoadedSourcesRequest:
		response = onLoadedSourcesRequest(rw.Writer, request.(dap.LoadedSourcesRequest))
	case dap.DataBreakpointInfoRequest:
		response = onDataBreakpointInfoRequest(rw.Writer, request.(dap.DataBreakpointInfoRequest))
	case dap.SetDataBreakpointsRequest:
		response = onSetDataBreakpointsRequest(rw.Writer, request.(dap.SetDataBreakpointsRequest))
	case dap.ReadMemoryRequest:
		response = onReadMemoryRequest(rw.Writer, request.(dap.ReadMemoryRequest))
	case dap.DisassembleRequest:
		response = onDisassembleRequest(rw.Writer, request.(dap.DisassembleRequest))
	case dap.CancelRequest:
		response = onCancelRequest(rw.Writer, request.(dap.CancelRequest))
	case dap.BreakpointLocationsRequest:
		response = onBreakpointLocationsRequest(rw.Writer, request.(dap.BreakpointLocationsRequest))
	default:
		return errors.New("not a request")
	}
	dap.WriteProtocolMessage(rw, response)
	log.Printf("Response sent\n\t%#v\n", response)
	debugSession.run(rw)
	rw.Flush()
	return nil
}

func writeAndLogProtocolMessage(w io.Writer, message dap.Message) {
	dap.WriteProtocolMessage(w, message)
	log.Printf("Message sent\n\t%#v\n", message)
}

// -----------------------------------------------------------------------
// Very Fake Debugger
//

var debugSession fakeDebugSession

type fakeDebugSession struct {
	breakpointsSet int
	canRun         bool
}

func (ds *fakeDebugSession) init() {
	ds.breakpointsSet = 0
	ds.canRun = false
}

func (ds *fakeDebugSession) run(w io.Writer) {
	if !ds.canRun {
		return
	}
	ds.canRun = false
	var e dap.Message
	if ds.breakpointsSet == 0 {
		e = dap.TerminatedEvent{
			Event: newEvent("terminated"),
		}
	} else {
		e = dap.StoppedEvent{
			Event: newEvent("stopped"),
			Body:  dap.StoppedEventBody{Reason: "breakpoint", ThreadId: 1, AllThreadsStopped: true},
		}
		ds.breakpointsSet--
	}
	writeAndLogProtocolMessage(w, e)
}

// -----------------------------------------------------------------------
// Request Handlers
//
// Below is a dummy implementation of the request handlers.
// They take no action, but just return dummy responses.
// A real debug adaptor would call the debugger methods here
// and use their results to populate each response.

func onInitializeRequest(w io.Writer, request dap.InitializeRequest) dap.Message {
	response := dap.InitializeResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body.SupportsConfigurationDoneRequest = true
	response.Body.SupportsFunctionBreakpoints = false
	response.Body.SupportsConditionalBreakpoints = false
	response.Body.SupportsHitConditionalBreakpoints = false
	response.Body.SupportsEvaluateForHovers = false
	response.Body.ExceptionBreakpointFilters = []dap.ExceptionBreakpointsFilter{}
	response.Body.SupportsStepBack = false
	response.Body.SupportsSetVariable = false
	response.Body.SupportsRestartFrame = false
	response.Body.SupportsGotoTargetsRequest = false
	response.Body.SupportsStepInTargetsRequest = false
	response.Body.SupportsCompletionsRequest = false
	response.Body.CompletionTriggerCharacters = []string{}
	response.Body.SupportsModulesRequest = false
	response.Body.AdditionalModuleColumns = []dap.ColumnDescriptor{}
	response.Body.SupportedChecksumAlgorithms = []dap.ChecksumAlgorithm{}
	response.Body.SupportsRestartRequest = false
	response.Body.SupportsExceptionOptions = false
	response.Body.SupportsValueFormattingOptions = false
	response.Body.SupportsExceptionInfoRequest = false
	response.Body.SupportTerminateDebuggee = false
	response.Body.SupportsDelayedStackTraceLoading = false
	response.Body.SupportsLoadedSourcesRequest = false
	response.Body.SupportsLogPoints = false
	response.Body.SupportsTerminateThreadsRequest = false
	response.Body.SupportsSetExpression = false
	response.Body.SupportsTerminateRequest = false
	response.Body.SupportsDataBreakpoints = false
	response.Body.SupportsReadMemoryRequest = false
	response.Body.SupportsDisassembleRequest = false
	response.Body.SupportsCancelRequest = false
	response.Body.SupportsBreakpointLocationsRequest = false
	// This is a fake set up, so we can start "accepting" configuration
	// requests for setting breakpoints, etc from the client at any time.
	// Notify the client with an 'initialized' event. The client will end
	// the configuration sequence with 'configurationDone' request.
	e := dap.InitializedEvent{Event: newEvent("initialized")}
	writeAndLogProtocolMessage(w, e)
	debugSession.init()
	return response
}

func onLaunchRequest(w io.Writer, request dap.LaunchRequest) dap.Message {
	// This is where a real debug adaptor would check the soundness of the
	// arguments (e.g. program from launch.json) and then use them to launch the
	// debugger and attach to the program.
	response := dap.LaunchResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	return response
}

func onAttachRequest(w io.Writer, request dap.AttachRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "AttachRequest is not yet supported")
}

func onDisconnectRequest(w io.Writer, request dap.DisconnectRequest) dap.Message {
	response := dap.DisconnectResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	return response
}

func onTerminateRequest(w io.Writer, request dap.TerminateRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "TerminateRequest is not yet supported")
}

func onRestartRequest(w io.Writer, request dap.RestartRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "RestartRequest is not yet supported")
}

func onSetBreakpointsRequest(w io.Writer, request dap.SetBreakpointsRequest) dap.Message {
	response := dap.SetBreakpointsResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body.Breakpoints = make([]dap.Breakpoint, len(request.Arguments.Breakpoints))
	for i, b := range request.Arguments.Breakpoints {
		response.Body.Breakpoints[i].Line = b.Line
		response.Body.Breakpoints[i].Verified = true
		debugSession.breakpointsSet++
	}
	return response
}

func onSetFunctionBreakpointsRequest(w io.Writer, request dap.SetFunctionBreakpointsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetFunctionBreakpointsRequest is not yet supported")
}

func onSetExceptionBreakpointsRequest(w io.Writer, request dap.SetExceptionBreakpointsRequest) dap.Message {
	response := dap.SetExceptionBreakpointsResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	return response
}

func onConfigurationDoneRequest(w io.Writer, request dap.ConfigurationDoneRequest) dap.Message {
	// This would be the place to check if the session was configured to stop on entry
	// and if that is the case, to issue a stopped-on-breakpoint event.
	// This being a mock implementation, we "let" the program continue.
	onContinueRequest(w, dap.ContinueRequest{Arguments: dap.ContinueArguments{ThreadId: 1}})
	e := dap.ThreadEvent{Event: newEvent("thread"), Body: dap.ThreadEventBody{Reason: "started", ThreadId: 1}}
	writeAndLogProtocolMessage(w, e)
	response := dap.ConfigurationDoneResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	debugSession.canRun = true
	return response
}

func onContinueRequest(w io.Writer, request dap.ContinueRequest) dap.Message {
	response := dap.ContinueResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	debugSession.canRun = true
	return response
}

func onNextRequest(w io.Writer, request dap.NextRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "NextRequest is not yet supported")
}

func onStepInRequest(w io.Writer, request dap.StepInRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepInRequest is not yet supported")
}

func onStepOutRequest(w io.Writer, request dap.StepOutRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepOutRequest is not yet supported")
}

func onStepBackRequest(w io.Writer, request dap.StepBackRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepBackRequest is not yet supported")
}

func onReverseContinueRequest(w io.Writer, request dap.ReverseContinueRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ReverseContinueRequest is not yet supported")
}

func onRestartFrameRequest(w io.Writer, request dap.RestartFrameRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "RestartFrameRequest is not yet supported")
}

func onGotoRequest(w io.Writer, request dap.GotoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "GotoRequest is not yet supported")
}

func onPauseRequest(w io.Writer, request dap.PauseRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "PauseRequest is not yet supported")
}

func onStackTraceRequest(w io.Writer, request dap.StackTraceRequest) dap.Message {
	response := dap.StackTraceResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body = dap.StackTraceResponseBody{
		StackFrames: []dap.StackFrame{
			dap.StackFrame{
				Id:     1000,
				Source: dap.Source{Name: "hello.go", Path: "/Users/foo/go/src/hello/hello.go", SourceReference: 0},
				Line:   5,
				Column: 0,
				Name:   "main.main",
			},
		},
		TotalFrames: 1,
	}
	return response
}

func onScopesRequest(w io.Writer, request dap.ScopesRequest) dap.Message {
	response := dap.ScopesResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body = dap.ScopesResponseBody{
		Scopes: []dap.Scope{
			dap.Scope{Name: "Local", VariablesReference: 1000, Expensive: false},
			dap.Scope{Name: "Global", VariablesReference: 1001, Expensive: true},
		},
	}
	return response
}

func onVariablesRequest(w io.Writer, request dap.VariablesRequest) dap.Message {
	response := dap.VariablesResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body = dap.VariablesResponseBody{
		Variables: []dap.Variable{dap.Variable{Name: "i", Value: "18434528", EvaluateName: "i", VariablesReference: 0}},
	}
	return response
}

func onSetVariableRequest(w io.Writer, request dap.SetVariableRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "setVariableRequest is not yet supported")
}

func onSetExpressionRequest(w io.Writer, request dap.SetExpressionRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetExpressionRequest is not yet supported")
}

func onSourceRequest(w io.Writer, request dap.SourceRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SourceRequest is not yet supported")
}

func onThreadsRequest(w io.Writer, request dap.ThreadsRequest) dap.Message {
	response := dap.ThreadsResponse{}
	response.Response = newResponse(request.Seq, request.Command)
	response.Body = dap.ThreadsResponseBody{Threads: []dap.Thread{dap.Thread{Id: 1, Name: "main"}}}
	return response
}

func onTerminateThreadsRequest(w io.Writer, request dap.TerminateThreadsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "TerminateRequest is not yet supported")
}

func onEvaluateRequest(w io.Writer, request dap.EvaluateRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "EvaluateRequest is not yet supported")
}

func onStepInTargetsRequest(w io.Writer, request dap.StepInTargetsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepInTargetRequest is not yet supported")
}

func onGotoTargetsRequest(w io.Writer, request dap.GotoTargetsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "GotoTargetRequest is not yet supported")
}

func onCompletionsRequest(w io.Writer, request dap.CompletionsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "CompletionRequest is not yet supported")
}

func onExceptionInfoRequest(w io.Writer, request dap.ExceptionInfoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ExceptionRequest is not yet supported")
}

func onLoadedSourcesRequest(w io.Writer, request dap.LoadedSourcesRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "LoadedRequest is not yet supported")
}

func onDataBreakpointInfoRequest(w io.Writer, request dap.DataBreakpointInfoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "DataBreakpointInfoRequest is not yet supported")
}

func onSetDataBreakpointsRequest(w io.Writer, request dap.SetDataBreakpointsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetDataBreakpointsRequest is not yet supported")
}

func onReadMemoryRequest(w io.Writer, request dap.ReadMemoryRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ReadMemoryRequest is not yet supported")
}

func onDisassembleRequest(w io.Writer, request dap.DisassembleRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "DisassembleRequest is not yet supported")
}

func onCancelRequest(w io.Writer, request dap.CancelRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "CancelRequest is not yet supported")
}

func onBreakpointLocationsRequest(w io.Writer, request dap.BreakpointLocationsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "BreakpointLocationsRequest is not yet supported")
}

func newEvent(event string) dap.Event {
	return dap.Event{
		ProtocolMessage: dap.ProtocolMessage{
			Seq:  0,
			Type: "event",
		},
		Event: event,
	}
}

func newResponse(requestSeq int, command string) dap.Response {
	return dap.Response{
		ProtocolMessage: dap.ProtocolMessage{
			Seq:  0,
			Type: "response",
		},
		Command:    command,
		RequestSeq: requestSeq,
		Success:    true,
	}
}

func newErrorResponse(requestSeq int, command string, message string) dap.ErrorResponse {
	er := dap.ErrorResponse{}
	er.Response = newResponse(requestSeq, command)
	er.Success = false
	er.Message = "unsupported"
	er.Body.Error.Format = message
	er.Body.Error.Id = 12345
	return er
}
