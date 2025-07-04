// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	overtime "github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"
	mock "github.com/stretchr/testify/mock"
)

// OvertimeUsecase is an autogenerated mock type for the OvertimeUsecase type
type OvertimeUsecase struct {
	mock.Mock
}

// CreateOvertime provides a mock function with given fields: ctx, req
func (_m *OvertimeUsecase) CreateOvertime(ctx context.Context, req overtime.OvertimeRequest) (overtime.OvertimeResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateOvertime")
	}

	var r0 overtime.OvertimeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, overtime.OvertimeRequest) (overtime.OvertimeResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, overtime.OvertimeRequest) overtime.OvertimeResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(overtime.OvertimeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, overtime.OvertimeRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOvertimeUsecase creates a new instance of OvertimeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOvertimeUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OvertimeUsecase {
	mock := &OvertimeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
