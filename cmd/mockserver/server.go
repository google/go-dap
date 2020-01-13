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
	debugSession := fakeDebugSession{
		rw: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)),
	}
	for {
		err := debugSession.handleRequest()
		// TODO(polina): check for connection vs decoding error?
		if err != nil {
			if err == io.EOF {
				log.Println("No more data to read:", err)
				return
			}
			log.Println("Server error:", err)
			continue // There may be more messages to process
		}
		debugSession.doContinue()
	}
}

func (ds *fakeDebugSession) handleRequest() error {
	log.Println("Reading request...")
	request, err := dap.ReadProtocolMessage(ds.rw.Reader)
	if err != nil {
		return err
	}
	log.Printf("Received request\n\t%#v\n", request)
	return ds.dispatchRequest(request)
}

func (ds *fakeDebugSession) dispatchRequest(request dap.Message) error {
	var response dap.Message
	switch request.(type) {
	case *dap.InitializeRequest:
		response = ds.onInitializeRequest(request.(*dap.InitializeRequest))
	case *dap.LaunchRequest:
		response = ds.onLaunchRequest(request.(*dap.LaunchRequest))
	case *dap.AttachRequest:
		response = ds.onAttachRequest(request.(*dap.AttachRequest))
	case *dap.DisconnectRequest:
		response = ds.onDisconnectRequest(request.(*dap.DisconnectRequest))
	case *dap.TerminateRequest:
		response = ds.onTerminateRequest(request.(*dap.TerminateRequest))
	case *dap.RestartRequest:
		response = ds.onRestartRequest(request.(*dap.RestartRequest))
	case *dap.SetBreakpointsRequest:
		response = ds.onSetBreakpointsRequest(request.(*dap.SetBreakpointsRequest))
	case *dap.SetFunctionBreakpointsRequest:
		response = ds.onSetFunctionBreakpointsRequest(request.(*dap.SetFunctionBreakpointsRequest))
	case *dap.SetExceptionBreakpointsRequest:
		response = ds.onSetExceptionBreakpointsRequest(request.(*dap.SetExceptionBreakpointsRequest))
	case *dap.ConfigurationDoneRequest:
		response = ds.onConfigurationDoneRequest(request.(*dap.ConfigurationDoneRequest))
	case *dap.ContinueRequest:
		response = ds.onContinueRequest(request.(*dap.ContinueRequest))
	case *dap.NextRequest:
		response = ds.onNextRequest(request.(*dap.NextRequest))
	case *dap.StepInRequest:
		response = ds.onStepInRequest(request.(*dap.StepInRequest))
	case *dap.StepOutRequest:
		response = ds.onStepOutRequest(request.(*dap.StepOutRequest))
	case *dap.StepBackRequest:
		response = ds.onStepBackRequest(request.(*dap.StepBackRequest))
	case *dap.ReverseContinueRequest:
		response = ds.onReverseContinueRequest(request.(*dap.ReverseContinueRequest))
	case *dap.RestartFrameRequest:
		response = ds.onRestartFrameRequest(request.(*dap.RestartFrameRequest))
	case *dap.GotoRequest:
		response = ds.onGotoRequest(request.(*dap.GotoRequest))
	case *dap.PauseRequest:
		response = ds.onPauseRequest(request.(*dap.PauseRequest))
	case *dap.StackTraceRequest:
		response = ds.onStackTraceRequest(request.(*dap.StackTraceRequest))
	case *dap.ScopesRequest:
		response = ds.onScopesRequest(request.(*dap.ScopesRequest))
	case *dap.VariablesRequest:
		response = ds.onVariablesRequest(request.(*dap.VariablesRequest))
	case *dap.SetVariableRequest:
		response = ds.onSetVariableRequest(request.(*dap.SetVariableRequest))
	case *dap.SetExpressionRequest:
		response = ds.onSetExpressionRequest(request.(*dap.SetExpressionRequest))
	case *dap.SourceRequest:
		response = ds.onSourceRequest(request.(*dap.SourceRequest))
	case *dap.ThreadsRequest:
		response = ds.onThreadsRequest(request.(*dap.ThreadsRequest))
	case *dap.TerminateThreadsRequest:
		response = ds.onTerminateThreadsRequest(request.(*dap.TerminateThreadsRequest))
	case *dap.EvaluateRequest:
		response = ds.onEvaluateRequest(request.(*dap.EvaluateRequest))
	case *dap.StepInTargetsRequest:
		response = ds.onStepInTargetsRequest(request.(*dap.StepInTargetsRequest))
	case *dap.GotoTargetsRequest:
		response = ds.onGotoTargetsRequest(request.(*dap.GotoTargetsRequest))
	case *dap.CompletionsRequest:
		response = ds.onCompletionsRequest(request.(*dap.CompletionsRequest))
	case *dap.ExceptionInfoRequest:
		response = ds.onExceptionInfoRequest(request.(*dap.ExceptionInfoRequest))
	case *dap.LoadedSourcesRequest:
		response = ds.onLoadedSourcesRequest(request.(*dap.LoadedSourcesRequest))
	case *dap.DataBreakpointInfoRequest:
		response = ds.onDataBreakpointInfoRequest(request.(*dap.DataBreakpointInfoRequest))
	case *dap.SetDataBreakpointsRequest:
		response = ds.onSetDataBreakpointsRequest(request.(*dap.SetDataBreakpointsRequest))
	case *dap.ReadMemoryRequest:
		response = ds.onReadMemoryRequest(request.(*dap.ReadMemoryRequest))
	case *dap.DisassembleRequest:
		response = ds.onDisassembleRequest(request.(*dap.DisassembleRequest))
	case *dap.CancelRequest:
		response = ds.onCancelRequest(request.(*dap.CancelRequest))
	case *dap.BreakpointLocationsRequest:
		response = ds.onBreakpointLocationsRequest(request.(*dap.BreakpointLocationsRequest))
	default:
		return errors.New("not a request")
	}
	ds.writeAndLogProtocolMessage(response)
	return nil
}

func (ds *fakeDebugSession) writeAndLogProtocolMessage(message dap.Message) {
	dap.WriteProtocolMessage(ds.rw.Writer, message)
	log.Printf("Message sent\n\t%#v\n", message)
	ds.rw.Flush()
}

// -----------------------------------------------------------------------
// Very Fake Debugger
//

// The debugging session will keep track of how many breakpoints
// have been set. Once start-up is done (i.e. configurationDone
// request is processed), it will "stop" at each breakpoint one
// by one, and once there are no more, it will trigger a terminate
// event.
type fakeDebugSession struct {
	rw *bufio.ReadWriter

	// breakpointSet is a counter of the remaining breakpoints
	// that the debug session is yet to stop at before the program
	// terminates. It must be 0 at start-up and termination.
	breakpointsSet int

	// canContinue is used to implement a small state machine between
	// multiple server functions. The debug session is paused
	// (canContinue is false) while multiple client requests are being
	// processed during the start-up sequence or when stopping at a
	// breakpoint. Once the client allows the session to continue,
	// the value is flipped to true. It is reset back to false
	// at termination and is ready for the next client session.
	canContinue bool
}

// doContinue is to be called between handling of each request/response
// to simulate events from the debug session that is in continue mode.
// If the program is "stopped", this will be a no-op. Otherwise, this
// will "stop" on a breakpoint or terminate if there are no more
// breakpoints.
func (ds *fakeDebugSession) doContinue() {
	if !ds.canContinue {
		return
	}
	ds.canContinue = false
	var e dap.Message
	if ds.breakpointsSet == 0 {
		e = &dap.TerminatedEvent{
			Event: *newEvent("terminated"),
		}
	} else {
		e = &dap.StoppedEvent{
			Event: *newEvent("stopped"),
			Body:  dap.StoppedEventBody{Reason: "breakpoint", ThreadId: 1, AllThreadsStopped: true},
		}
		ds.breakpointsSet--
	}
	ds.writeAndLogProtocolMessage(e)
}

// -----------------------------------------------------------------------
// Request Handlers
//
// Below is a dummy implementation of the request handlers.
// They take no action, but just return dummy responses.
// A real debug adaptor would call the debugger methods here
// and use their results to populate each response.

func (ds *fakeDebugSession) onInitializeRequest(request *dap.InitializeRequest) dap.Message {
	response := &dap.InitializeResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
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
	e := &dap.InitializedEvent{Event: *newEvent("initialized")}
	ds.writeAndLogProtocolMessage(e)
	return response
}

func (ds *fakeDebugSession) onLaunchRequest(request *dap.LaunchRequest) dap.Message {
	// This is where a real debug adaptor would check the soundness of the
	// arguments (e.g. program from launch.json) and then use them to launch the
	// debugger and attach to the program.
	response := &dap.LaunchResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	return response
}

func (ds *fakeDebugSession) onAttachRequest(request *dap.AttachRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "AttachRequest is not yet supported")
}

func (ds *fakeDebugSession) onDisconnectRequest(request *dap.DisconnectRequest) dap.Message {
	response := &dap.DisconnectResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	return response
}

func (ds *fakeDebugSession) onTerminateRequest(request *dap.TerminateRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "TerminateRequest is not yet supported")
}

func (ds *fakeDebugSession) onRestartRequest(request *dap.RestartRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "RestartRequest is not yet supported")
}

func (ds *fakeDebugSession) onSetBreakpointsRequest(request *dap.SetBreakpointsRequest) dap.Message {
	response := &dap.SetBreakpointsResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	response.Body.Breakpoints = make([]dap.Breakpoint, len(request.Arguments.Breakpoints))
	for i, b := range request.Arguments.Breakpoints {
		response.Body.Breakpoints[i].Line = b.Line
		response.Body.Breakpoints[i].Verified = true
		ds.breakpointsSet++
	}
	return response
}

func (ds *fakeDebugSession) onSetFunctionBreakpointsRequest(request *dap.SetFunctionBreakpointsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetFunctionBreakpointsRequest is not yet supported")
}

func (ds *fakeDebugSession) onSetExceptionBreakpointsRequest(request *dap.SetExceptionBreakpointsRequest) dap.Message {
	response := &dap.SetExceptionBreakpointsResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	return response
}

func (ds *fakeDebugSession) onConfigurationDoneRequest(request *dap.ConfigurationDoneRequest) dap.Message {
	// This would be the place to check if the session was configured to
	// stop on entry and if that is the case, to issue a
	// stopped-on-breakpoint event. This being a mock implementation,
	// we "let" the program continue.
	ds.onContinueRequest(&dap.ContinueRequest{Arguments: dap.ContinueArguments{ThreadId: 1}})
	e := &dap.ThreadEvent{Event: *newEvent("thread"), Body: dap.ThreadEventBody{Reason: "started", ThreadId: 1}}
	ds.writeAndLogProtocolMessage(e)
	response := &dap.ConfigurationDoneResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	ds.canContinue = true
	return response
}

func (ds *fakeDebugSession) onContinueRequest(request *dap.ContinueRequest) dap.Message {
	response := &dap.ContinueResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	ds.canContinue = true
	return response
}

func (ds *fakeDebugSession) onNextRequest(request *dap.NextRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "NextRequest is not yet supported")
}

func (ds *fakeDebugSession) onStepInRequest(request *dap.StepInRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepInRequest is not yet supported")
}

func (ds *fakeDebugSession) onStepOutRequest(request *dap.StepOutRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepOutRequest is not yet supported")
}

func (ds *fakeDebugSession) onStepBackRequest(request *dap.StepBackRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepBackRequest is not yet supported")
}

func (ds *fakeDebugSession) onReverseContinueRequest(request *dap.ReverseContinueRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ReverseContinueRequest is not yet supported")
}

func (ds *fakeDebugSession) onRestartFrameRequest(request *dap.RestartFrameRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "RestartFrameRequest is not yet supported")
}

func (ds *fakeDebugSession) onGotoRequest(request *dap.GotoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "GotoRequest is not yet supported")
}

func (ds *fakeDebugSession) onPauseRequest(request *dap.PauseRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "PauseRequest is not yet supported")
}

func (ds *fakeDebugSession) onStackTraceRequest(request *dap.StackTraceRequest) dap.Message {
	response := &dap.StackTraceResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
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

func (ds *fakeDebugSession) onScopesRequest(request *dap.ScopesRequest) dap.Message {
	response := &dap.ScopesResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	response.Body = dap.ScopesResponseBody{
		Scopes: []dap.Scope{
			dap.Scope{Name: "Local", VariablesReference: 1000, Expensive: false},
			dap.Scope{Name: "Global", VariablesReference: 1001, Expensive: true},
		},
	}
	return response
}

func (ds *fakeDebugSession) onVariablesRequest(request *dap.VariablesRequest) dap.Message {
	response := &dap.VariablesResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	response.Body = dap.VariablesResponseBody{
		Variables: []dap.Variable{dap.Variable{Name: "i", Value: "18434528", EvaluateName: "i", VariablesReference: 0}},
	}
	return response
}

func (ds *fakeDebugSession) onSetVariableRequest(request *dap.SetVariableRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "setVariableRequest is not yet supported")
}

func (ds *fakeDebugSession) onSetExpressionRequest(request *dap.SetExpressionRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetExpressionRequest is not yet supported")
}

func (ds *fakeDebugSession) onSourceRequest(request *dap.SourceRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SourceRequest is not yet supported")
}

func (ds *fakeDebugSession) onThreadsRequest(request *dap.ThreadsRequest) dap.Message {
	response := &dap.ThreadsResponse{}
	response.Response = *newResponse(request.Seq, request.Command)
	response.Body = dap.ThreadsResponseBody{Threads: []dap.Thread{dap.Thread{Id: 1, Name: "main"}}}
	return response
}

func (ds *fakeDebugSession) onTerminateThreadsRequest(request *dap.TerminateThreadsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "TerminateRequest is not yet supported")
}

func (ds *fakeDebugSession) onEvaluateRequest(request *dap.EvaluateRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "EvaluateRequest is not yet supported")
}

func (ds *fakeDebugSession) onStepInTargetsRequest(request *dap.StepInTargetsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "StepInTargetRequest is not yet supported")
}

func (ds *fakeDebugSession) onGotoTargetsRequest(request *dap.GotoTargetsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "GotoTargetRequest is not yet supported")
}

func (ds *fakeDebugSession) onCompletionsRequest(request *dap.CompletionsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "CompletionRequest is not yet supported")
}

func (ds *fakeDebugSession) onExceptionInfoRequest(request *dap.ExceptionInfoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ExceptionRequest is not yet supported")
}

func (ds *fakeDebugSession) onLoadedSourcesRequest(request *dap.LoadedSourcesRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "LoadedRequest is not yet supported")
}

func (ds *fakeDebugSession) onDataBreakpointInfoRequest(request *dap.DataBreakpointInfoRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "DataBreakpointInfoRequest is not yet supported")
}

func (ds *fakeDebugSession) onSetDataBreakpointsRequest(request *dap.SetDataBreakpointsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "SetDataBreakpointsRequest is not yet supported")
}

func (ds *fakeDebugSession) onReadMemoryRequest(request *dap.ReadMemoryRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "ReadMemoryRequest is not yet supported")
}

func (ds *fakeDebugSession) onDisassembleRequest(request *dap.DisassembleRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "DisassembleRequest is not yet supported")
}

func (ds *fakeDebugSession) onCancelRequest(request *dap.CancelRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "CancelRequest is not yet supported")
}

func (ds *fakeDebugSession) onBreakpointLocationsRequest(request *dap.BreakpointLocationsRequest) dap.Message {
	return newErrorResponse(request.Seq, request.Command, "BreakpointLocationsRequest is not yet supported")
}

func newEvent(event string) *dap.Event {
	return &dap.Event{
		ProtocolMessage: dap.ProtocolMessage{
			Seq:  0,
			Type: "event",
		},
		Event: event,
	}
}

func newResponse(requestSeq int, command string) *dap.Response {
	return &dap.Response{
		ProtocolMessage: dap.ProtocolMessage{
			Seq:  0,
			Type: "response",
		},
		Command:    command,
		RequestSeq: requestSeq,
		Success:    true,
	}
}

func newErrorResponse(requestSeq int, command string, message string) *dap.ErrorResponse {
	er := &dap.ErrorResponse{}
	er.Response = *newResponse(requestSeq, command)
	er.Success = false
	er.Message = "unsupported"
	er.Body.Error.Format = message
	er.Body.Error.Id = 12345
	return er
}
