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
// See cmd/gentypes/README.md for details.

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

	Command    string `json:"command"`
	Message    string `json:"message,omitempty"`
	RequestSeq int    `json:"request_seq"`
	Success    bool   `json:"success"`
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
	AllThreadsStopped bool   `json:"allThreadsStopped,omitempty"`
	Description       string `json:"description,omitempty"`
	PreserveFocusHint bool   `json:"preserveFocusHint,omitempty"`
	Reason            string `json:"reason"`
	Text              string `json:"text,omitempty"`
	ThreadId          int    `json:"threadId,omitempty"`
}

type ContinuedEvent struct {
	Event

	Body ContinuedEventBody `json:"body"`
}

type ContinuedEventBody struct {
	AllThreadsContinued bool `json:"allThreadsContinued,omitempty"`
	ThreadId            int  `json:"threadId"`
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

	Body TerminatedEventBody `json:"body"`
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
	Column             int         `json:"column,omitempty"`
	Data               interface{} `json:"data,omitempty"`
	Line               int         `json:"line,omitempty"`
	Output             string      `json:"output"`
	Source             Source      `json:"source,omitempty"`
	VariablesReference int         `json:"variablesReference,omitempty"`
}

type BreakpointEvent struct {
	Event

	Body BreakpointEventBody `json:"body"`
}

type BreakpointEventBody struct {
	Breakpoint Breakpoint `json:"breakpoint"`
	Reason     string     `json:"reason"`
}

type ModuleEvent struct {
	Event

	Body ModuleEventBody `json:"body"`
}

type ModuleEventBody struct {
	Module Module `json:"module"`
	Reason string `json:"reason"`
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
	IsLocalProcess  bool   `json:"isLocalProcess,omitempty"`
	Name            string `json:"name"`
	PointerSize     int    `json:"pointerSize,omitempty"`
	StartMethod     string `json:"startMethod,omitempty"`
	SystemProcessId int    `json:"systemProcessId,omitempty"`
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
	Args  []string          `json:"args"`
	Cwd   string            `json:"cwd"`
	Env   map[string]string `json:"env,omitempty"`
	Kind  string            `json:"kind,omitempty"`
	Title string            `json:"title,omitempty"`
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
	AdapterID                    string `json:"adapterID"`
	ClientID                     string `json:"clientID,omitempty"`
	ClientName                   string `json:"clientName,omitempty"`
	ColumnsStartAt1              bool   `json:"columnsStartAt1,omitempty"`
	LinesStartAt1                bool   `json:"linesStartAt1,omitempty"`
	Locale                       string `json:"locale,omitempty"`
	PathFormat                   string `json:"pathFormat,omitempty"`
	SupportsMemoryReferences     bool   `json:"supportsMemoryReferences,omitempty"`
	SupportsRunInTerminalRequest bool   `json:"supportsRunInTerminalRequest,omitempty"`
	SupportsVariablePaging       bool   `json:"supportsVariablePaging,omitempty"`
	SupportsVariableType         bool   `json:"supportsVariableType,omitempty"`
}

type InitializeResponse struct {
	Response

	Body Capabilities `json:"body"`
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
	Restart interface{} `json:"__restart,omitempty"`
	NoDebug bool        `json:"noDebug,omitempty"`
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
	Column    int    `json:"column,omitempty"`
	EndColumn int    `json:"endColumn,omitempty"`
	EndLine   int    `json:"endLine,omitempty"`
	Line      int    `json:"line"`
	Source    Source `json:"source"`
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
	Breakpoints    []SourceBreakpoint `json:"breakpoints,omitempty"`
	Lines          []int              `json:"lines,omitempty"`
	Source         Source             `json:"source"`
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
	ExceptionOptions []ExceptionOptions `json:"exceptionOptions,omitempty"`
	Filters          []string           `json:"filters"`
}

type SetExceptionBreakpointsResponse struct {
	Response
}

type DataBreakpointInfoRequest struct {
	Request

	Arguments DataBreakpointInfoArguments `json:"arguments"`
}

type DataBreakpointInfoArguments struct {
	Name               string `json:"name"`
	VariablesReference int    `json:"variablesReference,omitempty"`
}

type DataBreakpointInfoResponse struct {
	Response

	Body DataBreakpointInfoResponseBody `json:"body"`
}

type DataBreakpointInfoResponseBody struct {
	AccessTypes []DataBreakpointAccessType `json:"accessTypes,omitempty"`
	CanPersist  bool                       `json:"canPersist,omitempty"`
	DataId      interface{}                `json:"dataId"`
	Description string                     `json:"description"`
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
	TargetId int `json:"targetId,omitempty"`
	ThreadId int `json:"threadId"`
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
	TargetId int `json:"targetId"`
	ThreadId int `json:"threadId"`
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
	Format     StackFrameFormat `json:"format,omitempty"`
	Levels     int              `json:"levels,omitempty"`
	StartFrame int              `json:"startFrame,omitempty"`
	ThreadId   int              `json:"threadId"`
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
	Count              int         `json:"count,omitempty"`
	Filter             string      `json:"filter,omitempty"`
	Format             ValueFormat `json:"format,omitempty"`
	Start              int         `json:"start,omitempty"`
	VariablesReference int         `json:"variablesReference"`
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
	Format             ValueFormat `json:"format,omitempty"`
	Name               string      `json:"name"`
	Value              string      `json:"value"`
	VariablesReference int         `json:"variablesReference"`
}

type SetVariableResponse struct {
	Response

	Body SetVariableResponseBody `json:"body"`
}

type SetVariableResponseBody struct {
	IndexedVariables   int    `json:"indexedVariables,omitempty"`
	NamedVariables     int    `json:"namedVariables,omitempty"`
	Type               string `json:"type,omitempty"`
	Value              string `json:"value"`
	VariablesReference int    `json:"variablesReference,omitempty"`
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
	ModuleCount int `json:"moduleCount,omitempty"`
	StartModule int `json:"startModule,omitempty"`
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
	Context    string      `json:"context,omitempty"`
	Expression string      `json:"expression"`
	Format     ValueFormat `json:"format,omitempty"`
	FrameId    int         `json:"frameId,omitempty"`
}

type EvaluateResponse struct {
	Response

	Body EvaluateResponseBody `json:"body"`
}

type EvaluateResponseBody struct {
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
	MemoryReference    string                   `json:"memoryReference,omitempty"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	Result             string                   `json:"result"`
	Type               string                   `json:"type,omitempty"`
	VariablesReference int                      `json:"variablesReference"`
}

type SetExpressionRequest struct {
	Request

	Arguments SetExpressionArguments `json:"arguments"`
}

type SetExpressionArguments struct {
	Expression string      `json:"expression"`
	Format     ValueFormat `json:"format,omitempty"`
	FrameId    int         `json:"frameId,omitempty"`
	Value      string      `json:"value"`
}

type SetExpressionResponse struct {
	Response

	Body SetExpressionResponseBody `json:"body"`
}

type SetExpressionResponseBody struct {
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	Type               string                   `json:"type,omitempty"`
	Value              string                   `json:"value"`
	VariablesReference int                      `json:"variablesReference,omitempty"`
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
	Column int    `json:"column,omitempty"`
	Line   int    `json:"line"`
	Source Source `json:"source"`
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
	Column  int    `json:"column"`
	FrameId int    `json:"frameId,omitempty"`
	Line    int    `json:"line,omitempty"`
	Text    string `json:"text"`
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
	BreakMode   ExceptionBreakMode `json:"breakMode"`
	Description string             `json:"description,omitempty"`
	Details     ExceptionDetails   `json:"details,omitempty"`
	ExceptionId string             `json:"exceptionId"`
}

type ReadMemoryRequest struct {
	Request

	Arguments ReadMemoryArguments `json:"arguments"`
}

type ReadMemoryArguments struct {
	Count           int    `json:"count"`
	MemoryReference string `json:"memoryReference"`
	Offset          int    `json:"offset,omitempty"`
}

type ReadMemoryResponse struct {
	Response

	Body ReadMemoryResponseBody `json:"body"`
}

type ReadMemoryResponseBody struct {
	Address         string `json:"address"`
	Data            string `json:"data,omitempty"`
	UnreadableBytes int    `json:"unreadableBytes,omitempty"`
}

type DisassembleRequest struct {
	Request

	Arguments DisassembleArguments `json:"arguments"`
}

type DisassembleArguments struct {
	InstructionCount  int    `json:"instructionCount"`
	InstructionOffset int    `json:"instructionOffset,omitempty"`
	MemoryReference   string `json:"memoryReference"`
	Offset            int    `json:"offset,omitempty"`
	ResolveSymbols    bool   `json:"resolveSymbols,omitempty"`
}

type DisassembleResponse struct {
	Response

	Body DisassembleResponseBody `json:"body"`
}

type DisassembleResponseBody struct {
	Instructions []DisassembledInstruction `json:"instructions"`
}

type Capabilities struct {
	AdditionalModuleColumns            []ColumnDescriptor           `json:"additionalModuleColumns,omitempty"`
	CompletionTriggerCharacters        []string                     `json:"completionTriggerCharacters,omitempty"`
	ExceptionBreakpointFilters         []ExceptionBreakpointsFilter `json:"exceptionBreakpointFilters,omitempty"`
	SupportTerminateDebuggee           bool                         `json:"supportTerminateDebuggee,omitempty"`
	SupportedChecksumAlgorithms        []ChecksumAlgorithm          `json:"supportedChecksumAlgorithms,omitempty"`
	SupportsBreakpointLocationsRequest bool                         `json:"supportsBreakpointLocationsRequest,omitempty"`
	SupportsCancelRequest              bool                         `json:"supportsCancelRequest,omitempty"`
	SupportsCompletionsRequest         bool                         `json:"supportsCompletionsRequest,omitempty"`
	SupportsConditionalBreakpoints     bool                         `json:"supportsConditionalBreakpoints,omitempty"`
	SupportsConfigurationDoneRequest   bool                         `json:"supportsConfigurationDoneRequest,omitempty"`
	SupportsDataBreakpoints            bool                         `json:"supportsDataBreakpoints,omitempty"`
	SupportsDelayedStackTraceLoading   bool                         `json:"supportsDelayedStackTraceLoading,omitempty"`
	SupportsDisassembleRequest         bool                         `json:"supportsDisassembleRequest,omitempty"`
	SupportsEvaluateForHovers          bool                         `json:"supportsEvaluateForHovers,omitempty"`
	SupportsExceptionInfoRequest       bool                         `json:"supportsExceptionInfoRequest,omitempty"`
	SupportsExceptionOptions           bool                         `json:"supportsExceptionOptions,omitempty"`
	SupportsFunctionBreakpoints        bool                         `json:"supportsFunctionBreakpoints,omitempty"`
	SupportsGotoTargetsRequest         bool                         `json:"supportsGotoTargetsRequest,omitempty"`
	SupportsHitConditionalBreakpoints  bool                         `json:"supportsHitConditionalBreakpoints,omitempty"`
	SupportsLoadedSourcesRequest       bool                         `json:"supportsLoadedSourcesRequest,omitempty"`
	SupportsLogPoints                  bool                         `json:"supportsLogPoints,omitempty"`
	SupportsModulesRequest             bool                         `json:"supportsModulesRequest,omitempty"`
	SupportsReadMemoryRequest          bool                         `json:"supportsReadMemoryRequest,omitempty"`
	SupportsRestartFrame               bool                         `json:"supportsRestartFrame,omitempty"`
	SupportsRestartRequest             bool                         `json:"supportsRestartRequest,omitempty"`
	SupportsSetExpression              bool                         `json:"supportsSetExpression,omitempty"`
	SupportsSetVariable                bool                         `json:"supportsSetVariable,omitempty"`
	SupportsStepBack                   bool                         `json:"supportsStepBack,omitempty"`
	SupportsStepInTargetsRequest       bool                         `json:"supportsStepInTargetsRequest,omitempty"`
	SupportsTerminateRequest           bool                         `json:"supportsTerminateRequest,omitempty"`
	SupportsTerminateThreadsRequest    bool                         `json:"supportsTerminateThreadsRequest,omitempty"`
	SupportsValueFormattingOptions     bool                         `json:"supportsValueFormattingOptions,omitempty"`
}

type ExceptionBreakpointsFilter struct {
	Default bool   `json:"default,omitempty"`
	Filter  string `json:"filter"`
	Label   string `json:"label"`
}

type Message struct {
	Format        string            `json:"format"`
	Id            int               `json:"id"`
	SendTelemetry bool              `json:"sendTelemetry,omitempty"`
	ShowUser      bool              `json:"showUser,omitempty"`
	Url           string            `json:"url,omitempty"`
	UrlLabel      string            `json:"urlLabel,omitempty"`
	Variables     map[string]string `json:"variables,omitempty"`
}

type Module struct {
	AddressRange   string      `json:"addressRange,omitempty"`
	DateTimeStamp  string      `json:"dateTimeStamp,omitempty"`
	Id             interface{} `json:"id"`
	IsOptimized    bool        `json:"isOptimized,omitempty"`
	IsUserCode     bool        `json:"isUserCode,omitempty"`
	Name           string      `json:"name"`
	Path           string      `json:"path,omitempty"`
	SymbolFilePath string      `json:"symbolFilePath,omitempty"`
	SymbolStatus   string      `json:"symbolStatus,omitempty"`
	Version        string      `json:"version,omitempty"`
}

type ColumnDescriptor struct {
	AttributeName string `json:"attributeName"`
	Format        string `json:"format,omitempty"`
	Label         string `json:"label"`
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
	AdapterData      interface{} `json:"adapterData,omitempty"`
	Checksums        []Checksum  `json:"checksums,omitempty"`
	Name             string      `json:"name,omitempty"`
	Origin           string      `json:"origin,omitempty"`
	Path             string      `json:"path,omitempty"`
	PresentationHint string      `json:"presentationHint,omitempty"`
	SourceReference  int         `json:"sourceReference,omitempty"`
	Sources          []Source    `json:"sources,omitempty"`
}

type StackFrame struct {
	Column                      int         `json:"column"`
	EndColumn                   int         `json:"endColumn,omitempty"`
	EndLine                     int         `json:"endLine,omitempty"`
	Id                          int         `json:"id"`
	InstructionPointerReference string      `json:"instructionPointerReference,omitempty"`
	Line                        int         `json:"line"`
	ModuleId                    interface{} `json:"moduleId,omitempty"`
	Name                        string      `json:"name"`
	PresentationHint            string      `json:"presentationHint,omitempty"`
	Source                      Source      `json:"source,omitempty"`
}

type Scope struct {
	Column             int    `json:"column,omitempty"`
	EndColumn          int    `json:"endColumn,omitempty"`
	EndLine            int    `json:"endLine,omitempty"`
	Expensive          bool   `json:"expensive"`
	IndexedVariables   int    `json:"indexedVariables,omitempty"`
	Line               int    `json:"line,omitempty"`
	Name               string `json:"name"`
	NamedVariables     int    `json:"namedVariables,omitempty"`
	PresentationHint   string `json:"presentationHint,omitempty"`
	Source             Source `json:"source,omitempty"`
	VariablesReference int    `json:"variablesReference"`
}

type Variable struct {
	EvaluateName       string                   `json:"evaluateName,omitempty"`
	IndexedVariables   int                      `json:"indexedVariables,omitempty"`
	MemoryReference    string                   `json:"memoryReference,omitempty"`
	Name               string                   `json:"name"`
	NamedVariables     int                      `json:"namedVariables,omitempty"`
	PresentationHint   VariablePresentationHint `json:"presentationHint,omitempty"`
	Type               string                   `json:"type,omitempty"`
	Value              string                   `json:"value"`
	VariablesReference int                      `json:"variablesReference"`
}

type VariablePresentationHint struct {
	Attributes []string `json:"attributes,omitempty"`
	Kind       string   `json:"kind,omitempty"`
	Visibility string   `json:"visibility,omitempty"`
}

type BreakpointLocation struct {
	Column    int `json:"column,omitempty"`
	EndColumn int `json:"endColumn,omitempty"`
	EndLine   int `json:"endLine,omitempty"`
	Line      int `json:"line"`
}

type SourceBreakpoint struct {
	Column       int    `json:"column,omitempty"`
	Condition    string `json:"condition,omitempty"`
	HitCondition string `json:"hitCondition,omitempty"`
	Line         int    `json:"line"`
	LogMessage   string `json:"logMessage,omitempty"`
}

type FunctionBreakpoint struct {
	Condition    string `json:"condition,omitempty"`
	HitCondition string `json:"hitCondition,omitempty"`
	Name         string `json:"name"`
}

type DataBreakpointAccessType string

type DataBreakpoint struct {
	AccessType   DataBreakpointAccessType `json:"accessType,omitempty"`
	Condition    string                   `json:"condition,omitempty"`
	DataId       string                   `json:"dataId"`
	HitCondition string                   `json:"hitCondition,omitempty"`
}

type Breakpoint struct {
	Column    int    `json:"column,omitempty"`
	EndColumn int    `json:"endColumn,omitempty"`
	EndLine   int    `json:"endLine,omitempty"`
	Id        int    `json:"id,omitempty"`
	Line      int    `json:"line,omitempty"`
	Message   string `json:"message,omitempty"`
	Source    Source `json:"source,omitempty"`
	Verified  bool   `json:"verified"`
}

type StepInTarget struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type GotoTarget struct {
	Column                      int    `json:"column,omitempty"`
	EndColumn                   int    `json:"endColumn,omitempty"`
	EndLine                     int    `json:"endLine,omitempty"`
	Id                          int    `json:"id"`
	InstructionPointerReference string `json:"instructionPointerReference,omitempty"`
	Label                       string `json:"label"`
	Line                        int    `json:"line"`
}

type CompletionItem struct {
	Label    string             `json:"label"`
	Length   int                `json:"length,omitempty"`
	SortText string             `json:"sortText,omitempty"`
	Start    int                `json:"start,omitempty"`
	Text     string             `json:"text,omitempty"`
	Type     CompletionItemType `json:"type,omitempty"`
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

	IncludeAll      bool `json:"includeAll,omitempty"`
	Line            bool `json:"line,omitempty"`
	Module          bool `json:"module,omitempty"`
	ParameterNames  bool `json:"parameterNames,omitempty"`
	ParameterTypes  bool `json:"parameterTypes,omitempty"`
	ParameterValues bool `json:"parameterValues,omitempty"`
	Parameters      bool `json:"parameters,omitempty"`
}

type ExceptionOptions struct {
	BreakMode ExceptionBreakMode     `json:"breakMode"`
	Path      []ExceptionPathSegment `json:"path,omitempty"`
}

type ExceptionBreakMode string

type ExceptionPathSegment struct {
	Names  []string `json:"names"`
	Negate bool     `json:"negate,omitempty"`
}

type ExceptionDetails struct {
	EvaluateName   string             `json:"evaluateName,omitempty"`
	FullTypeName   string             `json:"fullTypeName,omitempty"`
	InnerException []ExceptionDetails `json:"innerException,omitempty"`
	Message        string             `json:"message,omitempty"`
	StackTrace     string             `json:"stackTrace,omitempty"`
	TypeName       string             `json:"typeName,omitempty"`
}

type DisassembledInstruction struct {
	Address          string `json:"address"`
	Column           int    `json:"column,omitempty"`
	EndColumn        int    `json:"endColumn,omitempty"`
	EndLine          int    `json:"endLine,omitempty"`
	Instruction      string `json:"instruction"`
	InstructionBytes string `json:"instructionBytes,omitempty"`
	Line             int    `json:"line,omitempty"`
	Location         Source `json:"location,omitempty"`
	Symbol           string `json:"symbol,omitempty"`
}

