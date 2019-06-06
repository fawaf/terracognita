// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cycloidio/terracognita/provider (interfaces: Resource)

// Package mock is a generated GoMock package.
package mock

import (
	filter "github.com/cycloidio/terracognita/filter"
	provider "github.com/cycloidio/terracognita/provider"
	writer "github.com/cycloidio/terracognita/writer"
	gomock "github.com/golang/mock/gomock"
	schema "github.com/hashicorp/terraform/helper/schema"
	reflect "reflect"
)

// Resource is a mock of Resource interface
type Resource struct {
	ctrl     *gomock.Controller
	recorder *ResourceMockRecorder
}

// ResourceMockRecorder is the mock recorder for Resource
type ResourceMockRecorder struct {
	mock *Resource
}

// NewResource creates a new mock instance
func NewResource(ctrl *gomock.Controller) *Resource {
	mock := &Resource{ctrl: ctrl}
	mock.recorder = &ResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Resource) EXPECT() *ResourceMockRecorder {
	return m.recorder
}

// Data mocks base method
func (m *Resource) Data() *schema.ResourceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(*schema.ResourceData)
	return ret0
}

// Data indicates an expected call of Data
func (mr *ResourceMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*Resource)(nil).Data))
}

// HCL mocks base method
func (m *Resource) HCL(arg0 writer.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HCL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HCL indicates an expected call of HCL
func (mr *ResourceMockRecorder) HCL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HCL", reflect.TypeOf((*Resource)(nil).HCL), arg0)
}

// ID mocks base method
func (m *Resource) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *ResourceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*Resource)(nil).ID))
}

// Provider mocks base method
func (m *Resource) Provider() provider.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(provider.Provider)
	return ret0
}

// Provider indicates an expected call of Provider
func (mr *ResourceMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*Resource)(nil).Provider))
}

// Read mocks base method
func (m *Resource) Read(arg0 filter.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read
func (mr *ResourceMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Resource)(nil).Read), arg0)
}

// State mocks base method
func (m *Resource) State(arg0 writer.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *ResourceMockRecorder) State(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*Resource)(nil).State), arg0)
}

// TFResource mocks base method
func (m *Resource) TFResource() *schema.Resource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TFResource")
	ret0, _ := ret[0].(*schema.Resource)
	return ret0
}

// TFResource indicates an expected call of TFResource
func (mr *ResourceMockRecorder) TFResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TFResource", reflect.TypeOf((*Resource)(nil).TFResource))
}

// Type mocks base method
func (m *Resource) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *ResourceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*Resource)(nil).Type))
}
