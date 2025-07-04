// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/alimasyhur/go-payroll-service/internal/app/entity"
	mock "github.com/stretchr/testify/mock"
)

// AttendancePeriod is an autogenerated mock type for the AttendancePeriod type
type AttendancePeriod struct {
	mock.Mock
}

// CreateAttendancePeriod provides a mock function with given fields: ctx, payload
func (_m *AttendancePeriod) CreateAttendancePeriod(ctx context.Context, payload entity.AttendancePeriod) (entity.AttendancePeriod, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateAttendancePeriod")
	}

	var r0 entity.AttendancePeriod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.AttendancePeriod) (entity.AttendancePeriod, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.AttendancePeriod) entity.AttendancePeriod); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(entity.AttendancePeriod)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.AttendancePeriod) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByDaterange provides a mock function with given fields: ctx, startDate, endDate
func (_m *AttendancePeriod) GetOneByDaterange(ctx context.Context, startDate string, endDate string) (entity.AttendancePeriod, error) {
	ret := _m.Called(ctx, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetOneByDaterange")
	}

	var r0 entity.AttendancePeriod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entity.AttendancePeriod, error)); ok {
		return rf(ctx, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entity.AttendancePeriod); ok {
		r0 = rf(ctx, startDate, endDate)
	} else {
		r0 = ret.Get(0).(entity.AttendancePeriod)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByUUID provides a mock function with given fields: ctx, uuid
func (_m *AttendancePeriod) GetOneByUUID(ctx context.Context, uuid string) (entity.AttendancePeriod, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetOneByUUID")
	}

	var r0 entity.AttendancePeriod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.AttendancePeriod, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.AttendancePeriod); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(entity.AttendancePeriod)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneClosedByDate provides a mock function with given fields: ctx, date
func (_m *AttendancePeriod) GetOneClosedByDate(ctx context.Context, date string) (entity.AttendancePeriod, error) {
	ret := _m.Called(ctx, date)

	if len(ret) == 0 {
		panic("no return value specified for GetOneClosedByDate")
	}

	var r0 entity.AttendancePeriod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.AttendancePeriod, error)); ok {
		return rf(ctx, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.AttendancePeriod); ok {
		r0 = rf(ctx, date)
	} else {
		r0 = ret.Get(0).(entity.AttendancePeriod)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAttendancePeriod provides a mock function with given fields: ctx, data
func (_m *AttendancePeriod) UpdateAttendancePeriod(ctx context.Context, data entity.AttendancePeriod) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAttendancePeriod")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.AttendancePeriod) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAttendancePeriod creates a new instance of AttendancePeriod. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAttendancePeriod(t interface {
	mock.TestingT
	Cleanup(func())
}) *AttendancePeriod {
	mock := &AttendancePeriod{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
