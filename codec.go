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

// This file contains utilities for decoding JSON-encoded bytes into DAP message.

package dap

import (
	"encoding/json"
	"fmt"
)

// DecodeProtocolMessageFieldError describes which JSON attribute
// has an unsupported value that the decoding cannot handle.
type DecodeProtocolMessageFieldError struct {
	Seq        int
	SubType    string
	FieldName  string
	FieldValue string
}

func (e *DecodeProtocolMessageFieldError) Error() string {
	return fmt.Sprintf("%s %s '%s' is not supported (seq: %d)", e.SubType, e.FieldName, e.FieldValue, e.Seq)
}

// defaultCodec is used to decode vanilla DAP messages.
var defaultCodec = NewCodec()

// Codec is responsible for turning byte blobs into DAP messages.
type Codec struct {
	eventCtor    map[string]messageCtor
	requestCtor  map[string]messageCtor
	responseCtor map[string]messageCtor
}

// NewCodec constructs a new codec that extends the vanilla DAP protocol.
// Unless you need to register custom DAP messages, use
// DecodeProtocolMessage instead.
func NewCodec() *Codec {
	ret := &Codec{
		eventCtor:    make(map[string]messageCtor),
		requestCtor:  make(map[string]messageCtor),
		responseCtor: make(map[string]messageCtor),
	}
	for k, v := range eventCtor {
		ret.eventCtor[k] = v
	}
	for k, v := range requestCtor {
		ret.requestCtor[k] = v
	}
	for k, v := range responseCtor {
		ret.responseCtor[k] = v
	}
	return ret
}

// RegisterRequest registers a new custom DAP command, so that it can be
// unmarshalled by DecodeMessage. Returns an error when the command already
// exists.
//
// The ctor functions need to return a new instance of the underlying DAP
// message type. A typical implementation looks like this:
//
//	func() Message { return &StoppedEvent{} }
func (c *Codec) RegisterRequest(command string, requestCtor, responseCtor func() Message) error {
	_, hasReqCtor := c.requestCtor[command]
	_, hasRespCtor := c.responseCtor[command]
	if hasReqCtor || hasRespCtor {
		return fmt.Errorf("command %q is already registered", command)
	}
	c.requestCtor[command] = requestCtor
	c.responseCtor[command] = responseCtor
	return nil
}

// RegisterEvent registers a new custom DAP event, so that it can be
// unmarshalled by DecodeMessage. Returns an error when the event already
// exists.
//
// The ctor function needs to return a new instance of the underlying DAP
// message type. A typical implementation looks like this:
//
//	func() Message { return &StoppedEvent{} }
func (c *Codec) RegisterEvent(event string, ctor func() Message) error {
	if _, hasEventCtor := c.eventCtor[event]; hasEventCtor {
		return fmt.Errorf("event %q is already registered", event)
	}
	c.eventCtor[event] = ctor
	return nil
}

// DecodeMessage parses the JSON-encoded data and returns the result of
// the appropriate type within the ProtocolMessage hierarchy. If message type,
// command, etc cannot be cast, returns DecodeProtocolMessageFieldError.
// See also godoc for json.Unmarshal, which is used for underlying decoding.
func (c *Codec) DecodeMessage(data []byte) (Message, error) {
	var protomsg ProtocolMessage
	if err := json.Unmarshal(data, &protomsg); err != nil {
		return nil, err
	}
	switch protomsg.Type {
	case "request":
		return c.decodeRequest(data)
	case "response":
		return c.decodeResponse(data)
	case "event":
		return c.decodeEvent(data)
	default:
		return nil, &DecodeProtocolMessageFieldError{protomsg.GetSeq(), "ProtocolMessage", "type", protomsg.Type}
	}
}

// decodeRequest determines what request type in the ProtocolMessage hierarchy
// data corresponds to and uses json.Unmarshal to populate the corresponding
// struct to be returned.
func (c *Codec) decodeRequest(data []byte) (Message, error) {
	var r Request
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	if ctor, ok := c.requestCtor[r.Command]; ok {
		requestPtr := ctor()
		err := json.Unmarshal(data, requestPtr)
		return requestPtr, err
	}
	return nil, &DecodeProtocolMessageFieldError{r.GetSeq(), "Request", "command", r.Command}
}

// decodeResponse determines what response type in the ProtocolMessage hierarchy
// data corresponds to and uses json.Unmarshal to populate the corresponding
// struct to be returned.
func (c *Codec) decodeResponse(data []byte) (Message, error) {
	var r Response
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	if !r.Success {
		var er ErrorResponse
		err := json.Unmarshal(data, &er)
		return &er, err
	}
	if ctor, ok := c.responseCtor[r.Command]; ok {
		responsePtr := ctor()
		err := json.Unmarshal(data, responsePtr)
		return responsePtr, err
	}
	return nil, &DecodeProtocolMessageFieldError{r.GetSeq(), "Response", "command", r.Command}
}

// decodeEvent determines what event type in the ProtocolMessage hierarchy
// data corresponds to and uses json.Unmarshal to populate the corresponding
// struct to be returned.
func (c *Codec) decodeEvent(data []byte) (Message, error) {
	var e Event
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	if ctor, ok := c.eventCtor[e.Event]; ok {
		eventPtr := ctor()
		err := json.Unmarshal(data, eventPtr)
		return eventPtr, err
	}
	return nil, &DecodeProtocolMessageFieldError{e.GetSeq(), "Event", "event", e.Event}
}

// DecodeProtocolMessage parses the JSON-encoded ProtocolMessage and returns
// the message embedded in it. If message type, command, etc cannot be cast,
// returns DecodeProtocolMessageFieldError. See also godoc for json.Unmarshal,
// which is used for underlying decoding.
func DecodeProtocolMessage(data []byte) (Message, error) {
	return defaultCodec.DecodeMessage(data)
}

type messageCtor func() Message
