// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	storage "github.com/ArnaudovSt/tx-parser/storage"
	types "github.com/ArnaudovSt/tx-parser/types"
	gomock "github.com/golang/mock/gomock"
)

// MockIReader is a mock of IReader interface.
type MockIReader struct {
	ctrl     *gomock.Controller
	recorder *MockIReaderMockRecorder
}

// MockIReaderMockRecorder is the mock recorder for MockIReader.
type MockIReaderMockRecorder struct {
	mock *MockIReader
}

// NewMockIReader creates a new mock instance.
func NewMockIReader(ctrl *gomock.Controller) *MockIReader {
	mock := &MockIReader{ctrl: ctrl}
	mock.recorder = &MockIReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIReader) EXPECT() *MockIReaderMockRecorder {
	return m.recorder
}

// GetLatestBlock mocks base method.
func (m *MockIReader) GetLatestBlock() (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestBlock")
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestBlock indicates an expected call of GetLatestBlock.
func (mr *MockIReaderMockRecorder) GetLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestBlock", reflect.TypeOf((*MockIReader)(nil).GetLatestBlock))
}

// GetTransactions mocks base method.
func (m *MockIReader) GetTransactions(address string) ([]*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", address)
	ret0, _ := ret[0].([]*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockIReaderMockRecorder) GetTransactions(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockIReader)(nil).GetTransactions), address)
}

// MockIWriter is a mock of IWriter interface.
type MockIWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIWriterMockRecorder
}

// MockIWriterMockRecorder is the mock recorder for MockIWriter.
type MockIWriterMockRecorder struct {
	mock *MockIWriter
}

// NewMockIWriter creates a new mock instance.
func NewMockIWriter(ctrl *gomock.Controller) *MockIWriter {
	mock := &MockIWriter{ctrl: ctrl}
	mock.recorder = &MockIWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWriter) EXPECT() *MockIWriterMockRecorder {
	return m.recorder
}

// AtomicWrite mocks base method.
func (m *MockIWriter) AtomicWrite(arg0 func(storage.IAtomicWriter) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtomicWrite", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AtomicWrite indicates an expected call of AtomicWrite.
func (mr *MockIWriterMockRecorder) AtomicWrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtomicWrite", reflect.TypeOf((*MockIWriter)(nil).AtomicWrite), arg0)
}

// Subscribe mocks base method.
func (m *MockIWriter) Subscribe(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockIWriterMockRecorder) Subscribe(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockIWriter)(nil).Subscribe), address)
}

// Unsubscribe mocks base method.
func (m *MockIWriter) Unsubscribe(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockIWriterMockRecorder) Unsubscribe(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockIWriter)(nil).Unsubscribe), address)
}

// MockIAtomicWriter is a mock of IAtomicWriter interface.
type MockIAtomicWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIAtomicWriterMockRecorder
}

// MockIAtomicWriterMockRecorder is the mock recorder for MockIAtomicWriter.
type MockIAtomicWriterMockRecorder struct {
	mock *MockIAtomicWriter
}

// NewMockIAtomicWriter creates a new mock instance.
func NewMockIAtomicWriter(ctrl *gomock.Controller) *MockIAtomicWriter {
	mock := &MockIAtomicWriter{ctrl: ctrl}
	mock.recorder = &MockIAtomicWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAtomicWriter) EXPECT() *MockIAtomicWriterMockRecorder {
	return m.recorder
}

// AppendBlock mocks base method.
func (m *MockIAtomicWriter) AppendBlock(block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendBlock", block)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendBlock indicates an expected call of AppendBlock.
func (mr *MockIAtomicWriterMockRecorder) AppendBlock(block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendBlock", reflect.TypeOf((*MockIAtomicWriter)(nil).AppendBlock), block)
}

// PopLatestBlock mocks base method.
func (m *MockIAtomicWriter) PopLatestBlock() (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopLatestBlock")
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PopLatestBlock indicates an expected call of PopLatestBlock.
func (mr *MockIAtomicWriterMockRecorder) PopLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopLatestBlock", reflect.TypeOf((*MockIAtomicWriter)(nil).PopLatestBlock))
}

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// AtomicWrite mocks base method.
func (m *MockIStorage) AtomicWrite(arg0 func(storage.IAtomicWriter) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtomicWrite", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AtomicWrite indicates an expected call of AtomicWrite.
func (mr *MockIStorageMockRecorder) AtomicWrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtomicWrite", reflect.TypeOf((*MockIStorage)(nil).AtomicWrite), arg0)
}

// GetLatestBlock mocks base method.
func (m *MockIStorage) GetLatestBlock() (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestBlock")
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestBlock indicates an expected call of GetLatestBlock.
func (mr *MockIStorageMockRecorder) GetLatestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestBlock", reflect.TypeOf((*MockIStorage)(nil).GetLatestBlock))
}

// GetTransactions mocks base method.
func (m *MockIStorage) GetTransactions(address string) ([]*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", address)
	ret0, _ := ret[0].([]*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockIStorageMockRecorder) GetTransactions(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockIStorage)(nil).GetTransactions), address)
}

// Subscribe mocks base method.
func (m *MockIStorage) Subscribe(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockIStorageMockRecorder) Subscribe(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockIStorage)(nil).Subscribe), address)
}

// Unsubscribe mocks base method.
func (m *MockIStorage) Unsubscribe(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockIStorageMockRecorder) Unsubscribe(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockIStorage)(nil).Unsubscribe), address)
}
