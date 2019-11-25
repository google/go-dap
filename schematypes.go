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

// DO NOT EDIT: This file is auto-generated.
// DAP spec: https://microsoft.github.io/debug-adapter-protocol/specification
// See cmd/gentypes/README.md for additional details.

package dap

type ProtocolMessage struct {
	Seq  int    `json:"seq"`
	Type string `json:"type"`
}

type Request struct {
	ProtocolMessage

	Command string `json:"command"`
}

type Event struct {
	ProtocolMessage

	Event string `json:"event"`
}

type Response struct {
	ProtocolMessage

	RequestSeq int    `json:"request_seq"`
	Success    bool   `json:"success"`
	Command    string `json:"command"`
	Message    string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Response

	Body ErrorResponseBody `json:"body"`
}

type ErrorResponseBody struct {
	Error Message `json:"error,omitempty"`
}

type CancelRequest struct {
	Request

	Arguments CancelArguments `json:"arguments,omitempty"`
}

type CancelArguments struct {
	RequestId int `json:"requestId,omitempty"`
}

type CancelResponse struct {
	Response
}

type InitializedEvent struct {
	Event
}

type StoppedEvent struct {
	Event

	Body StoppedEventBody `json:"body"`
}

type StoppedEventBody struct {
	Reason            string `json:"reason"`
	Description       string `json:"description,omitempty"`
	ThreadId          int    `json:"threadId,omitempty"`
	PreserveFocusHint bool   `json:"preserveFocusHint,omitempty"`
	Text              string `json:"text,omitempty"`
	AllThreadsStopped bool   `json:"allThreadsStopped,omitempty"`
}

type ContinuedEvent struct {
	Event

	Body ContinuedEventBody `json:"body"`
}

type ContinuedEventBody struct {
	ThreadId            int  `json:"threadId"`
	AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`
}

type ExitedEvent struct {
	Event

	Body ExitedEventBody `json:"body"`
}

type ExitedEventBody struct {
	ExitCode int `json:"exitCode"`
}

type TerminatedEvent struct {
	Event

	Body TerminatedEventBody `json:"body,omitempty"`
}

type TerminatedEventBody struct {
	Restart interface{} `json:"restart,omitempty"`
}

type ThreadEvent struct {
	Event

	Body ThreadEventBody `json:"body"`
}

type ThreadEventBody struct {
	Reason   string `json:"reason"`
	ThreadId int    `json:"threadId"`
}

type OutputEvent struct {
	Event

	Body OutputEventBody `json:"body"`
}

type OutputEventBody struct {
	Category           string      `json:"category,omitempty"`
	Output             string      `json:"output"`
	VariablesReference int         `json:"variablesReference,omitempty"`
	Source             Source      `json:"source,omitempty"`
	Line               int         `json:"line,omitempty"`
	Column             int         `json:"column,omitempty"`
	Data               interface{} `json:"data,omitempty"`
}

type BreakpointEvent struct {
	Event

	Body BreakpointEventBody `json:"body"`
}

type BreakpointEventBody struct {
	Reason     string     `json:"reason"`
	Breakpoint Breakpoint `json:"breakpoint"`
}

type ModuleEvent struct {
	Event

	Body ModuleEventBody `json:"body"`
}

type ModuleEventBody struct {
	Reason string `json:"reason"`
	Module Module `json:"module"`
}

type LoadedSourceEvent struct {
	Event

	Body LoadedSourceEventBody `json:"body"`
}

type LoadedSourceEventBody struct {
	Reason string `json:"reason"`
	Source Source `json:"source"`
}

type ProcessEvent struct {
	Event

	Body ProcessEventBody `json:"body"`
}

type ProcessEventBody struct {
	Name            string `json:"name"`
	SystemProcessId int    `json:"systemProcessId,omitempty"`
	IsLocalProcess  bool   `json:"isLocalProcess,omitempty"`
	StartMethod     string `json:"startMethod,omitempty"`
	PointerSize     int    `json:"pointerSize,omitempty"`
}

type CapabilitiesEvent struct {
	Event

	Body CapabilitiesEventBody `json:"body"`
}

type CapabilitiesEventBody struct {
	Capabilities Capabilities `json:"capabilities"`
}

type RunInTerminalRequest struct {
	Request

	Arguments RunInTerminalRequestArguments `json:"arguments"`
}

type RunInTerminalRequestArguments struct {
	Kind  string            `json:"kind,omitempty"`
	Title string            `json:"title,omitempty"`
	Cwd   string            `json:"cwd"`
	Args  []string          `json:"args"`
	Env   map[string]string `json:"env,omitempty"`
}

type RunInTerminalResponse struct {
	Response

	Body RunInTerminalResponseBody `json:"body"`
}

type RunInTerminalResponseBody struct {
	ProcessId      int `json:"processId,omitempty"`
	ShellProcessId int `json:"shellProcessId,omitempty"`
}

type InitializeRequest struct {
	Request

	Arguments InitializeRequestArguments `json:"arguments"`
}

type InitializeRequestArguments struct {
	ClientID                     string `json:"clientID,omitempty"`
	ClientName                   string `json:"clientName,omitempty"`
	AdapterID                    string `json:"adapterID"`
	Locale                       string `json:"locale,omitempty"`
	LinesStartAt1                bool   `json:"linesStartAt1,omitempty"`
	ColumnsStartAt1              bool   `json:"columnsStartAt1,omitempty"`
	PathFormat                   string `json:"pathFormat,omitempty"`
	SupportsVariableType         bool   `json:"supportsVariableType,omitempty"`
	SupportsVariablePaging       bool   `json:"supportsVariablePaging,omitempty"`
	SupportsRunInTerminalRequest bool   `json:"supportsRunInTerminalRequest,omitempty"`
	SupportsMemoryReferences     bool   `json:"supportsMemoryReferences,omitempty"`
}

type InitializeResponse struct {
	Response

	Body Capabilities `json:"body,omitempty"`
}

type ConfigurationDoneRequest struct {
	Request

	Arguments ConfigurationDoneArguments `json:"arguments,omitempty"`
}

type ConfigurationDoneArguments struct {
}

type ConfigurationDoneResponse struct {
	Response
}

type LaunchRequest struct {
	Request

	Arguments LaunchRequestArguments `json:"arguments"`
}

type LaunchRequestArguments struct {
	NoDebug bool        `json:"noDebug,omitempty"`
	Restart interface{} `json:"__restart,omitempty"`
}

type LaunchResponse struct {
	Response
}

type AttachRequest struct {
	Request

	Arguments AttachRequestArguments `json:"arguments"`
}

type AttachRequestArguments struct {
	Restart interface{} `json:"__restart,omitempty"`
}

type AttachResponse struct {
	Response
}

type RestartRequest struct {
	Request

	Arguments RestartArguments `json:"arguments,omitempty"`
}

type RestartArguments struct {
}

type RestartResponse struct {
	Response
}

type DisconnectRequest struct {
	Request

	Arguments DisconnectArguments `json:"arguments,omitempty"`
}

type DisconnectArguments struct {
	Restart           bool `json:"restart,omitempty"`
	TerminateDebuggee bool `json:"terminateDebuggee,omitempty"`
}

type DisconnectResponse struct {
	Response
}

type TerminateRequest struct {
	Request

	Arguments TerminateArguments `json:"arguments,omitempty"`
}

type TerminateArguments struct {
	Restart bool `json:"restart,omitempty"`
}

type TerminateResponse struct {
	Response
}

type BreakpointLocationsRequest struct {
	Request

	Arguments BreakpointLocationsArguments `json:"arguments,omitempty"`
}

type BreakpointLocationsArguments struct {
	Source    Source `json:"source"`
	Line      int    `json:"line"`
	Column    int    `json:"column,omitempty"`
	EndLine   int    `json:"endLine,omitempty"`
	EndColumn int    `json:"endColumn,omitempty"`
}

type BreakpointLocationsResponse struct {
	Response

	Body BreakpointLocationsResponseBody `json:"body"`
}

type BreakpointLocationsResponseBody struct {
	Breakpoints []BreakpointLocation `json:"breakpoints"`
}

type SetBreakpointsRequest struct {
	Request

	Arguments SetBreakpointsArguments `json:"arguments"`
}

type SetBreakpointsArguments struct {
	Source         Source             `json:"source"`
	Breakpoints    []SourceBreakpoint `json:"breakpoints,omitempty"`
	Lines          []int              `json:"lines,omitempty"`
	SourceModified bool               `json:"sourceModified,omitempty"`
}

type SetBreakpointsResponse struct {
	Response

	Body SetBreakpointsResponseBody `json:"body"`
}

type SetBreakpointsResponseBody struct {
	Breakpoints []Breakpoint `json:"breakpoints"`
}

type SetFunctionBreakpointsRequest struct {
	Request

	Arguments SetFunctionBreakpointsArguments `json:"arguments"`
}

type SetFunctionBreakpointsArguments struct {
	Breakpoints []FunctionBreakpoint `json:"breakpoints"`
}

type SetFunctionBreakpointsResponse struct {
	Response

	Body SetFunctionBreakpointsResponseBody `json:"body"`
}

type SetFunctionBreakpointsResponseBody struct {
	Breakpoints []Breakpoint `json:"breakpoints"`
}

type SetExceptionBreakpointsRequest struct {
	Request

	Arguments SetExceptionBreakpointsArguments `json:"arguments"`
}

type SetExceptionBreakpointsArguments struct {
	Filters          []string           `json:"filters"`
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`
}

type SetExceptionBreakpointsResponse struct {
	Response
}

type DataBreakpointInfoRequest struct {
	Request

	Arguments DataBreakpointInfoArguments `json:"arguments"`
}

type DataBreakpointInfoArguments struct {
	VariablesReference int    `json:"variablesReference,omitempty"`
	Name               string `json:"name"`
}

type DataBreakpointInfoResponse struct {
	Response

	Body DataBreakpointInfoResponseBody `json:"body"`
}

type DataBreakpointInfoResponseBody struct {
	DataId      interface{}                `json:"dataId"`
	Description string                     `json:"description"`
	AccessTypes []DataBreakpointAccessType `json:"accessTypes,omitempty"`
	CanPersist  bool                       `json:"canPersist,omitempty"`
}

type SetDataBreakpointsRequest struct {
	Request

	Arguments SetDataBreakpointsArguments `json:"arguments"`
}

type SetDataBreakpointsArguments struct {
	Breakpoints []DataBreakpoint `json:"breakpoints"`
}

type SetDataBreakpointsResponse struct {
	Response

	Body SetDataBreakpointsResponseBody `json:"body"`
}

type SetDataBreakpointsResponseBody struct {
	Breakpoints []Breakpoint `json:"breakpoints"`
}

type ContinueRequest struct {
	Request

	Arguments ContinueArguments `json:"arguments"`
}

type ContinueArguments struct {
	ThreadId int `json:"threadId"`
}

type ContinueResponse struct {
	Response

	Body ContinueResponseBody `json:"body"`
}

type ContinueResponseBody struct {
	AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`
}

type NextRequest struct {
	Request

	Arguments NextArguments `json:"arguments"`
}

type NextArguments struct {
	ThreadId int `json:"threadId"`
}

type NextResponse struct {
	Response
}

type StepInRequest struct {
	Request

	Arguments StepInArguments `json:"arguments"`
}

type StepInArguments struct {
	ThreadId int `json:"threadId"`
	TargetId int `json:"targetId,omitempty"`
}

type StepInResponse struct {
	Response
}

type StepOutRequest struct {
	Request

	Arguments StepOutArguments `json:"arguments"`
}

type StepOutArguments struct {
	ThreadId int `json:"threadId"`
}

type StepOutResponse struct {
	Response
}

type StepBackRequest struct {
	Request

	Arguments StepBackArguments `json:"arguments"`
}

type StepBackArguments struct {
	ThreadId int `json:"threadId"`
}

type StepBackResponse struct {
	Response
}

type ReverseContinueRequest struct {
	Request

	Arguments ReverseContinueArguments `json:"arguments"`
}

type ReverseContinueArguments struct {
	ThreadId int `json:"threadId"`
}

type ReverseContinueResponse struct {
	Response
}

type RestartFrameRequest struct {
	Request

	Arguments RestartFrameArguments `json:"arguments"`
}

type RestartFrameArguments struct {
	FrameId int `json:"frameId"`
}

type RestartFrameResponse struct {
	Response
}

type GotoRequest struct {
	Request

	Arguments GotoArguments `json:"arguments"`
}

type GotoArguments struct {
	ThreadId int `json:"threadId"`
	TargetId int `json:"targetId"`
}

type GotoResponse struct {
	Response
}

type PauseRequest struct {
	Request

	Arguments PauseArguments `json:"arguments"`
}

type PauseArguments struct {
	ThreadId int `json:"threadId"`
}

type PauseResponse struct {
	Response
}

type StackTraceRequest struct {
	Request

	Arguments StackTraceArguments `json:"arguments"`
}

type StackTraceArguments struct {
	ThreadId   int              `json:"threadId"`
	StartFrame int              `json:"startFrame,omitempty"`
	Levels     int              `json:"levels,omitempty"`
	Format     StackFrameFormat `json:"format,omitempty"`
}

type StackTraceResponse struct {
	Response

	Body StackTraceResponseBody `json:"body"`
}

type StackTraceResponseBody struct {
	StackFrames []StackFrame `json:"stackFrames"`
	TotalFrames int          `json:"totalFrames,omitempty"`
}

type ScopesRequest struct {
	Request

	Arguments ScopesArguments `json:"arguments"`
}

type ScopesArguments struct {
	FrameId int `json:"frameId"`
}

type ScopesResponse struct {
	Response

	Body ScopesResponseBody `json:"body"`
}

type ScopesResponseBody struct {
	Scopes []Scope `json:"scopes"`
}

type VariablesRequest struct {
	Request

	Arguments VariablesArguments `json:"arguments"`
}

type VariablesArguments struct {
	VariablesReference int         `json:"variablesReference"`
	Filter             string      `json:"filter,omitempty"`
	Start              int         `json:"start,omitempty"`
	Count              int         `json:"count,omitempty"`
	Format             ValueFormat `json:"format,omitempty"`
}

type VariablesResponse struct {
	Response

	Body VariablesResponseBody `json:"body"`
}

type VariablesResponseBody struct {
	Variables []Variable `json:"variables"`
}

type SetVariableRequest struct {
	Request

	Arguments SetVariableArguments `json:"arguments"`
}

type SetVariableArguments struct {
	VariablesReference int         `json:"variablesReference"`
	Name               string      `json:"name"`
	Value              string      `json:"value"`
	Format             ValueFormat `json:"format,omitempty"`
}

type SetVariableResponse struct {
	Response

	Body SetVariableResponseBody `json:"body"`
}

type SetVariableResponseBody struct {
	Value              string `json:"value"`
	Type               string `json:"type,omitempty"`
	VariablesReference int    `json:"variablesReference,omitempty"`
	NamedVariables     int    `json:"namedVariables,omitempty"`
	IndexedVariables   int    `json:"indexedVariables,omitempty"`
}

type SourceRequest struct {
	Request

	Arguments SourceArguments `json:"arguments"`
}

type SourceArguments struct {
	Source          Source `json:"source,omitempty"`
	SourceReference int    `json:"sourceReference"`
}

type SourceResponse struct {
	Response

	Body SourceResponseBody `json:"body"`
}

type SourceResponseBody struct {
	Content  string `json:"content"`
	MimeType string `json:"mimeType,omitempty"`
}

type ThreadsRequest struct {
	Request
}

type ThreadsResponse struct {
	Response

	Body ThreadsResponseBody `json:"body"`
}

type ThreadsResponseBody struct {
	Threads []Thread `json:"threads"`
}

type TerminateThreadsRequest struct {
	Request

	Arguments TerminateThreadsArguments `json:"arguments"`
}

type TerminateThreadsArguments struct {
	ThreadIds []int `json:"threadIds,omitempty"`
}

type TerminateThreadsResponse struct {
	Response
}

type ModulesRequest struct {
	Request

	Arguments ModulesArguments `json:"arguments"`
}

type ModulesArguments struct {
	StartModule int `json:"startModule,omitempty"`
	ModuleCount int `json:"moduleCount,omitempty"`
}

type ModulesResponse struct {
	Response

	Body ModulesResponseBody `json:"body"`
}

type ModulesResponseBody struct {
	Modules      []Module `json:"modules"`
	TotalModules int      `json:"totalModules,omitempty"`
}

type LoadedSourcesRequest struct {
	Request

	Arguments LoadedSourcesArguments `json:"arguments,omitempty"`
}

type LoadedSourcesArguments struct {
}

type LoadedSourcesResponse struct {
	Response

	Body LoadedSourcesResponseBody `json:"body"`
}

type LoadedSourcesResponseBody struct {
	Sources []Source `json:"sources"`
}

type EvaluateRequest struct {
	Request

	Arguments EvaluateArguments `json:"arguments"`
}

type EvaluateArguments struct {
	Expression string      `json:"expression"`
	FrameId    int         `json:"frameId,omitempty"`
	Context    string      `json:"context,omitempty"`
	Format     ValueFormat `json:"format,omitempty"`
}

type EvaluateResponse struct {
	Response

	Body EvaluateResponseBody `json:"body"`
}

type EvaluateResponseBody struct {
	Result             string                   `json:"result"`
	Type               string                   `json:"type,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	VariablesReference int                      `json:"variablesReference"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
	MemoryReference    string                   `json:"memoryReference,omitempty"`
}

type SetExpressionRequest struct {
	Request

	Arguments SetExpressionArguments `json:"arguments"`
}

type SetExpressionArguments struct {
	Expression string      `json:"expression"`
	Value      string      `json:"value"`
	FrameId    int         `json:"frameId,omitempty"`
	Format     ValueFormat `json:"format,omitempty"`
}

type SetExpressionResponse struct {
	Response

	Body SetExpressionResponseBody `json:"body"`
}

type SetExpressionResponseBody struct {
	Value              string                   `json:"value"`
	Type               string                   `json:"type,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	VariablesReference int                      `json:"variablesReference,omitempty"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
}

type StepInTargetsRequest struct {
	Request

	Arguments StepInTargetsArguments `json:"arguments"`
}

type StepInTargetsArguments struct {
	FrameId int `json:"frameId"`
}

type StepInTargetsResponse struct {
	Response

	Body StepInTargetsResponseBody `json:"body"`
}

type StepInTargetsResponseBody struct {
	Targets []StepInTarget `json:"targets"`
}

type GotoTargetsRequest struct {
	Request

	Arguments GotoTargetsArguments `json:"arguments"`
}

type GotoTargetsArguments struct {
	Source Source `json:"source"`
	Line   int    `json:"line"`
	Column int    `json:"column,omitempty"`
}

type GotoTargetsResponse struct {
	Response

	Body GotoTargetsResponseBody `json:"body"`
}

type GotoTargetsResponseBody struct {
	Targets []GotoTarget `json:"targets"`
}

type CompletionsRequest struct {
	Request

	Arguments CompletionsArguments `json:"arguments"`
}

type CompletionsArguments struct {
	FrameId int    `json:"frameId,omitempty"`
	Text    string `json:"text"`
	Column  int    `json:"column"`
	Line    int    `json:"line,omitempty"`
}

type CompletionsResponse struct {
	Response

	Body CompletionsResponseBody `json:"body"`
}

type CompletionsResponseBody struct {
	Targets []CompletionItem `json:"targets"`
}

type ExceptionInfoRequest struct {
	Request

	Arguments ExceptionInfoArguments `json:"arguments"`
}

type ExceptionInfoArguments struct {
	ThreadId int `json:"threadId"`
}

type ExceptionInfoResponse struct {
	Response

	Body ExceptionInfoResponseBody `json:"body"`
}

type ExceptionInfoResponseBody struct {
	ExceptionId string             `json:"exceptionId"`
	Description string             `json:"description,omitempty"`
	BreakMode   ExceptionBreakMode `json:"breakMode"`
	Details     ExceptionDetails   `json:"details,omitempty"`
}

type ReadMemoryRequest struct {
	Request

	Arguments ReadMemoryArguments `json:"arguments"`
}

type ReadMemoryArguments struct {
	MemoryReference string `json:"memoryReference"`
	Offset          int    `json:"offset,omitempty"`
	Count           int    `json:"count"`
}

type ReadMemoryResponse struct {
	Response

	Body ReadMemoryResponseBody `json:"body,omitempty"`
}

type ReadMemoryResponseBody struct {
	Address         string `json:"address"`
	UnreadableBytes int    `json:"unreadableBytes,omitempty"`
	Data            string `json:"data,omitempty"`
}

type DisassembleRequest struct {
	Request

	Arguments DisassembleArguments `json:"arguments"`
}

type DisassembleArguments struct {
	MemoryReference   string `json:"memoryReference"`
	Offset            int    `json:"offset,omitempty"`
	InstructionOffset int    `json:"instructionOffset,omitempty"`
	InstructionCount  int    `json:"instructionCount"`
	ResolveSymbols    bool   `json:"resolveSymbols,omitempty"`
}

type DisassembleResponse struct {
	Response

	Body DisassembleResponseBody `json:"body,omitempty"`
}

type DisassembleResponseBody struct {
	Instructions []DisassembledInstruction `json:"instructions"`
}

type Capabilities struct {
	SupportsConfigurationDoneRequest   bool                         `json:"supportsConfigurationDoneRequest,omitempty"`
	SupportsFunctionBreakpoints        bool                         `json:"supportsFunctionBreakpoints,omitempty"`
	SupportsConditionalBreakpoints     bool                         `json:"supportsConditionalBreakpoints,omitempty"`
	SupportsHitConditionalBreakpoints  bool                         `json:"supportsHitConditionalBreakpoints,omitempty"`
	SupportsEvaluateForHovers          bool                         `json:"supportsEvaluateForHovers,omitempty"`
	ExceptionBreakpointFilters         []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`
	SupportsStepBack                   bool                         `json:"supportsStepBack,omitempty"`
	SupportsSetVariable                bool                         `json:"supportsSetVariable,omitempty"`
	SupportsRestartFrame               bool                         `json:"supportsRestartFrame,omitempty"`
	SupportsGotoTargetsRequest         bool                         `json:"supportsGotoTargetsRequest,omitempty"`
	SupportsStepInTargetsRequest       bool                         `json:"supportsStepInTargetsRequest,omitempty"`
	SupportsCompletionsRequest         bool                         `json:"supportsCompletionsRequest,omitempty"`
	CompletionTriggerCharacters        []string                     `json:"completionTriggerCharacters,omitempty"`
	SupportsModulesRequest             bool                         `json:"supportsModulesRequest,omitempty"`
	AdditionalModuleColumns            []ColumnDescriptor           `json:"additionalModuleColumns,omitempty"`
	SupportedChecksumAlgorithms        []ChecksumAlgorithm          `json:"supportedChecksumAlgorithms,omitempty"`
	SupportsRestartRequest             bool                         `json:"supportsRestartRequest,omitempty"`
	SupportsExceptionOptions           bool                         `json:"supportsExceptionOptions,omitempty"`
	SupportsValueFormattingOptions     bool                         `json:"supportsValueFormattingOptions,omitempty"`
	SupportsExceptionInfoRequest       bool                         `json:"supportsExceptionInfoRequest,omitempty"`
	SupportTerminateDebuggee           bool                         `json:"supportTerminateDebuggee,omitempty"`
	SupportsDelayedStackTraceLoading   bool                         `json:"supportsDelayedStackTraceLoading,omitempty"`
	SupportsLoadedSourcesRequest       bool                         `json:"supportsLoadedSourcesRequest,omitempty"`
	SupportsLogPoints                  bool                         `json:"supportsLogPoints,omitempty"`
	SupportsTerminateThreadsRequest    bool                         `json:"supportsTerminateThreadsRequest,omitempty"`
	SupportsSetExpression              bool                         `json:"supportsSetExpression,omitempty"`
	SupportsTerminateRequest           bool                         `json:"supportsTerminateRequest,omitempty"`
	SupportsDataBreakpoints            bool                         `json:"supportsDataBreakpoints,omitempty"`
	SupportsReadMemoryRequest          bool                         `json:"supportsReadMemoryRequest,omitempty"`
	SupportsDisassembleRequest         bool                         `json:"supportsDisassembleRequest,omitempty"`
	SupportsCancelRequest              bool                         `json:"supportsCancelRequest,omitempty"`
	SupportsBreakpointLocationsRequest bool                         `json:"supportsBreakpointLocationsRequest,omitempty"`
}

type ExceptionBreakpointsFilter struct {
	Filter  string `json:"filter"`
	Label   string `json:"label"`
	Default bool   `json:"default,omitempty"`
}

type Message struct {
	Id            int               `json:"id"`
	Format        string            `json:"format"`
	Variables     map[string]string `json:"variables,omitempty"`
	SendTelemetry bool              `json:"sendTelemetry,omitempty"`
	ShowUser      bool              `json:"showUser,omitempty"`
	Url           string            `json:"url,omitempty"`
	UrlLabel      string            `json:"urlLabel,omitempty"`
}

type Module struct {
	Id             interface{} `json:"id"`
	Name           string      `json:"name"`
	Path           string      `json:"path,omitempty"`
	IsOptimized    bool        `json:"isOptimized,omitempty"`
	IsUserCode     bool        `json:"isUserCode,omitempty"`
	Version        string      `json:"version,omitempty"`
	SymbolStatus   string      `json:"symbolStatus,omitempty"`
	SymbolFilePath string      `json:"symbolFilePath,omitempty"`
	DateTimeStamp  string      `json:"dateTimeStamp,omitempty"`
	AddressRange   string      `json:"addressRange,omitempty"`
}

type ColumnDescriptor struct {
	AttributeName string `json:"attributeName"`
	Label         string `json:"label"`
	Format        string `json:"format,omitempty"`
	Type          string `json:"type,omitempty"`
	Width         int    `json:"width,omitempty"`
}

type ModulesViewDescriptor struct {
	Columns []ColumnDescriptor `json:"columns"`
}

type Thread struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Source struct {
	Name             string      `json:"name,omitempty"`
	Path             string      `json:"path,omitempty"`
	SourceReference  int         `json:"sourceReference,omitempty"`
	PresentationHint string      `json:"presentationHint,omitempty"`
	Origin           string      `json:"origin,omitempty"`
	Sources          []Source    `json:"sources,omitempty"`
	AdapterData      interface{} `json:"adapterData,omitempty"`
	Checksums        []Checksum  `json:"checksums,omitempty"`
}

type StackFrame struct {
	Id                          int         `json:"id"`
	Name                        string      `json:"name"`
	Source                      Source      `json:"source,omitempty"`
	Line                        int         `json:"line"`
	Column                      int         `json:"column"`
	EndLine                     int         `json:"endLine,omitempty"`
	EndColumn                   int         `json:"endColumn,omitempty"`
	InstructionPointerReference string      `json:"instructionPointerReference,omitempty"`
	ModuleId                    interface{} `json:"moduleId,omitempty"`
	PresentationHint            string      `json:"presentationHint,omitempty"`
}

type Scope struct {
	Name               string `json:"name"`
	PresentationHint   string `json:"presentationHint,omitempty"`
	VariablesReference int    `json:"variablesReference"`
	NamedVariables     int    `json:"namedVariables,omitempty"`
	IndexedVariables   int    `json:"indexedVariables,omitempty"`
	Expensive          bool   `json:"expensive"`
	Source             Source `json:"source,omitempty"`
	Line               int    `json:"line,omitempty"`
	Column             int    `json:"column,omitempty"`
	EndLine            int    `json:"endLine,omitempty"`
	EndColumn          int    `json:"endColumn,omitempty"`
}

type Variable struct {
	Name               string                   `json:"name"`
	Value              string                   `json:"value"`
	Type               string                   `json:"type,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	EvaluateName       string                   `json:"evaluateName,omitempty"`
	VariablesReference int                      `json:"variablesReference"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
	MemoryReference    string                   `json:"memoryReference,omitempty"`
}

type VariablePresentationHint struct {
	Kind       string   `json:"kind,omitempty"`
	Attributes []string `json:"attributes,omitempty"`
	Visibility string   `json:"visibility,omitempty"`
}

type BreakpointLocation struct {
	Line      int `json:"line"`
	Column    int `json:"column,omitempty"`
	EndLine   int `json:"endLine,omitempty"`
	EndColumn int `json:"endColumn,omitempty"`
}

type SourceBreakpoint struct {
	Line         int    `json:"line"`
	Column       int    `json:"column,omitempty"`
	Condition    string `json:"condition,omitempty"`
	HitCondition string `json:"hitCondition,omitempty"`
	LogMessage   string `json:"logMessage,omitempty"`
}

type FunctionBreakpoint struct {
	Name         string `json:"name"`
	Condition    string `json:"condition,omitempty"`
	HitCondition string `json:"hitCondition,omitempty"`
}

type DataBreakpointAccessType string

type DataBreakpoint struct {
	DataId       string                   `json:"dataId"`
	AccessType   DataBreakpointAccessType `json:"accessType,omitempty"`
	Condition    string                   `json:"condition,omitempty"`
	HitCondition string                   `json:"hitCondition,omitempty"`
}

type Breakpoint struct {
	Id        int    `json:"id,omitempty"`
	Verified  bool   `json:"verified"`
	Message   string `json:"message,omitempty"`
	Source    Source `json:"source,omitempty"`
	Line      int    `json:"line,omitempty"`
	Column    int    `json:"column,omitempty"`
	EndLine   int    `json:"endLine,omitempty"`
	EndColumn int    `json:"endColumn,omitempty"`
}

type StepInTarget struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type GotoTarget struct {
	Id                          int    `json:"id"`
	Label                       string `json:"label"`
	Line                        int    `json:"line"`
	Column                      int    `json:"column,omitempty"`
	EndLine                     int    `json:"endLine,omitempty"`
	EndColumn                   int    `json:"endColumn,omitempty"`
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`
}

type CompletionItem struct {
	Label    string             `json:"label"`
	Text     string             `json:"text,omitempty"`
	SortText string             `json:"sortText,omitempty"`
	Type     CompletionItemType `json:"type,omitempty"`
	Start    int                `json:"start,omitempty"`
	Length   int                `json:"length,omitempty"`
}

type CompletionItemType string

type ChecksumAlgorithm string

type Checksum struct {
	Algorithm ChecksumAlgorithm `json:"algorithm"`
	Checksum  string            `json:"checksum"`
}

type ValueFormat struct {
	Hex bool `json:"hex,omitempty"`
}

type StackFrameFormat struct {
	ValueFormat

	Parameters      bool `json:"parameters,omitempty"`
	ParameterTypes  bool `json:"parameterTypes,omitempty"`
	ParameterNames  bool `json:"parameterNames,omitempty"`
	ParameterValues bool `json:"parameterValues,omitempty"`
	Line            bool `json:"line,omitempty"`
	Module          bool `json:"module,omitempty"`
	IncludeAll      bool `json:"includeAll,omitempty"`
}

type ExceptionOptions struct {
	Path      []ExceptionPathSegment `json:"path,omitempty"`
	BreakMode ExceptionBreakMode     `json:"breakMode"`
}

type ExceptionBreakMode string

type ExceptionPathSegment struct {
	Negate bool     `json:"negate,omitempty"`
	Names  []string `json:"names"`
}

type ExceptionDetails struct {
	Message        string             `json:"message,omitempty"`
	TypeName       string             `json:"typeName,omitempty"`
	FullTypeName   string             `json:"fullTypeName,omitempty"`
	EvaluateName   string             `json:"evaluateName,omitempty"`
	StackTrace     string             `json:"stackTrace,omitempty"`
	InnerException []ExceptionDetails `json:"innerException,omitempty"`
}

type DisassembledInstruction struct {
	Address          string `json:"address"`
	InstructionBytes string `json:"instructionBytes,omitempty"`
	Instruction      string `json:"instruction"`
	Symbol           string `json:"symbol,omitempty"`
	Location         Source `json:"location,omitempty"`
	Line             int    `json:"line,omitempty"`
	Column           int    `json:"column,omitempty"`
	EndLine          int    `json:"endLine,omitempty"`
	EndColumn        int    `json:"endColumn,omitempty"`
}
