// Code generated by MockGen. DO NOT EDIT.
// Source: client/aws_device_farm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	devicefarm "github.com/aws/aws-sdk-go-v2/service/devicefarm"
	gomock "github.com/golang/mock/gomock"
)

// MockProviderDeviceFarmInterface is a mock of ProviderDeviceFarmInterface interface.
type MockProviderDeviceFarmInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProviderDeviceFarmInterfaceMockRecorder
}

// MockProviderDeviceFarmInterfaceMockRecorder is the mock recorder for MockProviderDeviceFarmInterface.
type MockProviderDeviceFarmInterfaceMockRecorder struct {
	mock *MockProviderDeviceFarmInterface
}

// NewMockProviderDeviceFarmInterface creates a new mock instance.
func NewMockProviderDeviceFarmInterface(ctrl *gomock.Controller) *MockProviderDeviceFarmInterface {
	mock := &MockProviderDeviceFarmInterface{ctrl: ctrl}
	mock.recorder = &MockProviderDeviceFarmInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderDeviceFarmInterface) EXPECT() *MockProviderDeviceFarmInterfaceMockRecorder {
	return m.recorder
}

// CreateUpload mocks base method.
func (m *MockProviderDeviceFarmInterface) CreateUpload(ctx context.Context, params *devicefarm.CreateUploadInput, optFns ...func(*devicefarm.Options)) (*devicefarm.CreateUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUpload", varargs...)
	ret0, _ := ret[0].(*devicefarm.CreateUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpload indicates an expected call of CreateUpload.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) CreateUpload(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpload", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).CreateUpload), varargs...)
}

// GetRun mocks base method.
func (m *MockProviderDeviceFarmInterface) GetRun(ctx context.Context, params *devicefarm.GetRunInput, optFns ...func(*devicefarm.Options)) (*devicefarm.GetRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRun", varargs...)
	ret0, _ := ret[0].(*devicefarm.GetRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) GetRun(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).GetRun), varargs...)
}

// GetUpload mocks base method.
func (m *MockProviderDeviceFarmInterface) GetUpload(ctx context.Context, params *devicefarm.GetUploadInput, optFns ...func(*devicefarm.Options)) (*devicefarm.GetUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUpload", varargs...)
	ret0, _ := ret[0].(*devicefarm.GetUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpload indicates an expected call of GetUpload.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) GetUpload(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpload", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).GetUpload), varargs...)
}

// ListArtifacts mocks base method.
func (m *MockProviderDeviceFarmInterface) ListArtifacts(ctx context.Context, params *devicefarm.ListArtifactsInput, optFns ...func(*devicefarm.Options)) (*devicefarm.ListArtifactsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListArtifacts", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListArtifactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArtifacts indicates an expected call of ListArtifacts.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ListArtifacts(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArtifacts", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ListArtifacts), varargs...)
}

// ListDevicePools mocks base method.
func (m *MockProviderDeviceFarmInterface) ListDevicePools(arg0 context.Context, arg1 *devicefarm.ListDevicePoolsInput, arg2 ...func(*devicefarm.Options)) (*devicefarm.ListDevicePoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDevicePools", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListDevicePoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDevicePools indicates an expected call of ListDevicePools.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ListDevicePools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevicePools", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ListDevicePools), varargs...)
}

// ListJobs mocks base method.
func (m *MockProviderDeviceFarmInterface) ListJobs(arg0 context.Context, arg1 *devicefarm.ListJobsInput, arg2 ...func(*devicefarm.Options)) (*devicefarm.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobs", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ListJobs), varargs...)
}

// ListProjects mocks base method.
func (m *MockProviderDeviceFarmInterface) ListProjects(arg0 context.Context, arg1 *devicefarm.ListProjectsInput, arg2 ...func(*devicefarm.Options)) (*devicefarm.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ListProjects), varargs...)
}

// ListUploads mocks base method.
func (m *MockProviderDeviceFarmInterface) ListUploads(arg0 context.Context, arg1 *devicefarm.ListUploadsInput, arg2 ...func(*devicefarm.Options)) (*devicefarm.ListUploadsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUploads", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListUploadsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUploads indicates an expected call of ListUploads.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ListUploads(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUploads", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ListUploads), varargs...)
}

// ScheduleRun mocks base method.
func (m *MockProviderDeviceFarmInterface) ScheduleRun(ctx context.Context, params *devicefarm.ScheduleRunInput, optFns ...func(*devicefarm.Options)) (*devicefarm.ScheduleRunOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleRun", varargs...)
	ret0, _ := ret[0].(*devicefarm.ScheduleRunOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleRun indicates an expected call of ScheduleRun.
func (mr *MockProviderDeviceFarmInterfaceMockRecorder) ScheduleRun(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleRun", reflect.TypeOf((*MockProviderDeviceFarmInterface)(nil).ScheduleRun), varargs...)
}

// MockListProjectsPaginator is a mock of ListProjectsPaginator interface.
type MockListProjectsPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockListProjectsPaginatorMockRecorder
}

// MockListProjectsPaginatorMockRecorder is the mock recorder for MockListProjectsPaginator.
type MockListProjectsPaginatorMockRecorder struct {
	mock *MockListProjectsPaginator
}

// NewMockListProjectsPaginator creates a new mock instance.
func NewMockListProjectsPaginator(ctrl *gomock.Controller) *MockListProjectsPaginator {
	mock := &MockListProjectsPaginator{ctrl: ctrl}
	mock.recorder = &MockListProjectsPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListProjectsPaginator) EXPECT() *MockListProjectsPaginatorMockRecorder {
	return m.recorder
}

// HasMorePages mocks base method.
func (m *MockListProjectsPaginator) HasMorePages() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMorePages")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMorePages indicates an expected call of HasMorePages.
func (mr *MockListProjectsPaginatorMockRecorder) HasMorePages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMorePages", reflect.TypeOf((*MockListProjectsPaginator)(nil).HasMorePages))
}

// NextPage mocks base method.
func (m *MockListProjectsPaginator) NextPage(arg0 context.Context, arg1 ...func(*devicefarm.Options)) (*devicefarm.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NextPage", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPage indicates an expected call of NextPage.
func (mr *MockListProjectsPaginatorMockRecorder) NextPage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPage", reflect.TypeOf((*MockListProjectsPaginator)(nil).NextPage), varargs...)
}

// MockListUploadsPaginator is a mock of ListUploadsPaginator interface.
type MockListUploadsPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockListUploadsPaginatorMockRecorder
}

// MockListUploadsPaginatorMockRecorder is the mock recorder for MockListUploadsPaginator.
type MockListUploadsPaginatorMockRecorder struct {
	mock *MockListUploadsPaginator
}

// NewMockListUploadsPaginator creates a new mock instance.
func NewMockListUploadsPaginator(ctrl *gomock.Controller) *MockListUploadsPaginator {
	mock := &MockListUploadsPaginator{ctrl: ctrl}
	mock.recorder = &MockListUploadsPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListUploadsPaginator) EXPECT() *MockListUploadsPaginatorMockRecorder {
	return m.recorder
}

// HasMorePages mocks base method.
func (m *MockListUploadsPaginator) HasMorePages() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMorePages")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMorePages indicates an expected call of HasMorePages.
func (mr *MockListUploadsPaginatorMockRecorder) HasMorePages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMorePages", reflect.TypeOf((*MockListUploadsPaginator)(nil).HasMorePages))
}

// NextPage mocks base method.
func (m *MockListUploadsPaginator) NextPage(arg0 context.Context, arg1 ...func(*devicefarm.Options)) (*devicefarm.ListUploadsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NextPage", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListUploadsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPage indicates an expected call of NextPage.
func (mr *MockListUploadsPaginatorMockRecorder) NextPage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPage", reflect.TypeOf((*MockListUploadsPaginator)(nil).NextPage), varargs...)
}

// MockListDevicePoolPaginator is a mock of ListDevicePoolPaginator interface.
type MockListDevicePoolPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockListDevicePoolPaginatorMockRecorder
}

// MockListDevicePoolPaginatorMockRecorder is the mock recorder for MockListDevicePoolPaginator.
type MockListDevicePoolPaginatorMockRecorder struct {
	mock *MockListDevicePoolPaginator
}

// NewMockListDevicePoolPaginator creates a new mock instance.
func NewMockListDevicePoolPaginator(ctrl *gomock.Controller) *MockListDevicePoolPaginator {
	mock := &MockListDevicePoolPaginator{ctrl: ctrl}
	mock.recorder = &MockListDevicePoolPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListDevicePoolPaginator) EXPECT() *MockListDevicePoolPaginatorMockRecorder {
	return m.recorder
}

// HasMorePages mocks base method.
func (m *MockListDevicePoolPaginator) HasMorePages() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMorePages")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMorePages indicates an expected call of HasMorePages.
func (mr *MockListDevicePoolPaginatorMockRecorder) HasMorePages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMorePages", reflect.TypeOf((*MockListDevicePoolPaginator)(nil).HasMorePages))
}

// NextPage mocks base method.
func (m *MockListDevicePoolPaginator) NextPage(arg0 context.Context, arg1 ...func(*devicefarm.Options)) (*devicefarm.ListDevicePoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NextPage", varargs...)
	ret0, _ := ret[0].(*devicefarm.ListDevicePoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPage indicates an expected call of NextPage.
func (mr *MockListDevicePoolPaginatorMockRecorder) NextPage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPage", reflect.TypeOf((*MockListDevicePoolPaginator)(nil).NextPage), varargs...)
}
