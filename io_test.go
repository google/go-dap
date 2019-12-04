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

package dap

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func Test_WriteBaseMessage(t *testing.T) {
	tests := []struct {
		input       string
		wantWritten string
		wantErr     error
	}{
		{``, "Content-Length: 0\r\n\r\n", nil},
		{`a`, "Content-Length: 1\r\n\r\na", nil},
		{`{}`, "Content-Length: 2\r\n\r\n{}", nil},
		{`{"a":0 "b":"blah"}`, "Content-Length: 18\r\n\r\n{\"a\":0 \"b\":\"blah\"}", nil},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var buf bytes.Buffer
			gotErr := WriteBaseMessage(&buf, []byte(test.input))
			gotWritten := buf.String()
			if gotErr != test.wantErr {
				t.Errorf("got err=%#v, want %#v", gotErr, test.wantErr)
			}
			if gotErr == nil && gotWritten != test.wantWritten {
				t.Errorf("got written=%q, want %q", gotWritten, test.wantWritten)
			}
		})
	}
}

func Test_ReadBaseMessage(t *testing.T) {
	tests := []struct {
		input         string
		wantBytesRead []byte
		wantBytesLeft []byte
		wantErr       error
	}{
		{"", nil, []byte(""), io.EOF},
		{"random stuff\r\nabc", nil, []byte("c"), ErrHeaderDelimiterNotCrLfCrLf},
		{"Cache-Control: no-cache\r\n\r\n", nil, []byte(""), ErrHeaderNotContentLength},
		{"Content-Length 1\r\n\r\nabc", nil, []byte("abc"), ErrHeaderNotContentLength},
		{"Content-Length: 10\r\n\r\nabc", nil, []byte(""), io.ErrUnexpectedEOF},
		{"Content-Length: 3\r\n\r\nabc", []byte("abc"), []byte(""), nil},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(test.input))
			gotBytes, gotErr := ReadBaseMessage(reader)
			if gotErr != test.wantErr {
				t.Errorf("got err=%#v, want %#v", gotErr, test.wantErr)
			}
			if gotErr == nil && !reflect.DeepEqual(gotBytes, test.wantBytesRead) {
				t.Errorf("got bytes=%q, want %q", gotBytes, test.wantBytesRead)
			}
			bytesLeft, _ := ioutil.ReadAll(reader)
			if !reflect.DeepEqual(bytesLeft, test.wantBytesLeft) {
				t.Errorf("got bytesLeft=%q, want %q", bytesLeft, test.wantBytesLeft)
			}
		})
	}
}

func Test_readContentLengthHeader(t *testing.T) {
	tests := []struct {
		input         string
		wantBytesLeft string // Bytes left in the reader after header reading
		wantLen       int    // Extracted content length value
		wantErr       error
	}{
		{"", "", 0, io.EOF},
		{"Cache-Control: no-cache", "", 0, io.EOF},
		{"Cache-Control: no-cache\r", "", 0, io.EOF},
		{"Cache-Control: no-cache\rabc", "", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Cache-Control: no-cache\r\n", "", 0, io.ErrUnexpectedEOF},
		{"Cache-Control: no-cache\r\n\r", "", 0, io.ErrUnexpectedEOF},
		{"Cache-Control: no-cache\r\n\r\n", "", 0, ErrHeaderNotContentLength},
		{"Cache-Control: no-cache\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: 3 abc", "", 0, io.EOF},
		{"Content-Length: 3\nabc", "", 0, io.EOF},
		{"Content-Length: 3\rabc", "", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length: 3\r\nabc", "c", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length: 3\r\n\rabc", "bc", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length: 3\r \n\r\nabc", "\nabc", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length: 3\r\n \r\nabc", "\nabc", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length: 3\r\n\r \nabc", "\nabc", 0, ErrHeaderDelimiterNotCrLfCrLf},
		{"Content-Length 3\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"_Content-Length: 3\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: 3_\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: x\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: 3.0\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: -3\r\n\r\nabc", "abc", 0, ErrHeaderNotContentLength},
		{"Content-Length: 0\r\n\r\nabc", "abc", 0, nil},
		{"Content-Length: 3\r\n\r\nabc", "abc", 3, nil},
		{"Content-Length: 9223372036854775807\r\n\r\nabc", "abc", 9223372036854775807, nil},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(test.input))
			gotLen, gotErr := readContentLengthHeader(reader)
			if gotErr != test.wantErr {
				t.Errorf("got err=%#v, want %#v", gotErr, test.wantErr)
			}
			if gotErr == nil && gotLen != test.wantLen {
				t.Errorf("got len=%d, want %d", gotLen, test.wantLen)
			}
			bytesLeft, _ := ioutil.ReadAll(reader)
			if string(bytesLeft) != test.wantBytesLeft {
				t.Errorf("got bytesLeft=%q, want %q", bytesLeft, test.wantBytesLeft)
			}
		})
	}
}

func TestWriteRead(t *testing.T) {
	writeContent := [][]byte{
		[]byte("this is"),
		[]byte("a read write"),
		[]byte("test"),
	}
	var buf bytes.Buffer
	for _, wc := range writeContent {
		if err := WriteBaseMessage(&buf, wc); err != nil {
			t.Fatal(err)
		}
	}
	reader := bufio.NewReader(&buf)
	for _, wc := range writeContent {
		rc, err := ReadBaseMessage(reader)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(rc, wc) {
			t.Fatalf("got %q, want %q", rc, wc)
		}
	}
}

// readMessagesIntoChannel reads messages one by one until EOF.
// Die on error as we expect only well-formed messages.
func readMessagesIntoChannel(t *testing.T, r io.Reader, messages chan<- []byte) {
	reader := bufio.NewReader(r)
	for {
		msg, err := ReadBaseMessage(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			close(messages)
			// This goroutine might still be running after the test
			// completes, so we cannot use t.Fatal here without
			// additional synchronization
			panic(err)
		}
		messages <- msg
	}
}

func writeOrFail(t *testing.T, w io.Writer, data string) {
	if n, err := w.Write([]byte(data)); err != nil || n < len(data) {
		t.Fatal(err)
	}
}

func TestReadMessageInParts(t *testing.T) {
	// This test will use separate goroutines to write and read messages
	// and rely on blocking channel operations between them to ensure that
	// the expected number of messages is read for what is written.
	// Otherwise, the test will time out.
	// TODO(polina): use timeouts to catch such a failure mode
	// and fail gracefully?
	messages := make(chan []byte)
	r, w := io.Pipe()
	header := "Content-Length: 11"
	delim := "\r\n\r\n"
	content1 := "message one"
	content2 := "message two"

	// This will keep blocking to read a full message or EOF.
	go readMessagesIntoChannel(t, r, messages)

	// Write a message in full and verify via channel that it was read.
	writeOrFail(t, w, header+delim+content1)
	got := <-messages
	if !reflect.DeepEqual(got, []byte(content1)) {
		t.Fatalf("got %q, want %q", got, content1)
	}

	// Write a message in parts and verify via channel that it was read.
	writeOrFail(t, w, header)
	writeOrFail(t, w, delim)
	writeOrFail(t, w, content2)
	got = <-messages
	if !reflect.DeepEqual(got, []byte(content2)) {
		t.Fatalf("got %q, want %q", got, content2)
	}

	w.Close() // sends EOF
}
