// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/msg/consumer (interfaces: Message,MessageProcessor)

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package consumer is a generated GoMock package.
package consumer

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockMessage) Ack() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ack")
}

// Ack indicates an expected call of Ack.
func (mr *MockMessageMockRecorder) Ack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockMessage)(nil).Ack))
}

// Bytes mocks base method.
func (m *MockMessage) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockMessageMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockMessage)(nil).Bytes))
}

// MockMessageProcessor is a mock of MessageProcessor interface.
type MockMessageProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockMessageProcessorMockRecorder
}

// MockMessageProcessorMockRecorder is the mock recorder for MockMessageProcessor.
type MockMessageProcessorMockRecorder struct {
	mock *MockMessageProcessor
}

// NewMockMessageProcessor creates a new mock instance.
func NewMockMessageProcessor(ctrl *gomock.Controller) *MockMessageProcessor {
	mock := &MockMessageProcessor{ctrl: ctrl}
	mock.recorder = &MockMessageProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageProcessor) EXPECT() *MockMessageProcessorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessageProcessor) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessageProcessorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageProcessor)(nil).Close))
}

// Process mocks base method.
func (m *MockMessageProcessor) Process(arg0 Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process", arg0)
}

// Process indicates an expected call of Process.
func (mr *MockMessageProcessorMockRecorder) Process(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockMessageProcessor)(nil).Process), arg0)
}
