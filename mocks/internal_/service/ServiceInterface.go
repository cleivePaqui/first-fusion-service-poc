// Code generated by mockery v2.30.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

type ServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceInterface) EXPECT() *ServiceInterface_Expecter {
	return &ServiceInterface_Expecter{mock: &_m.Mock}
}

// IntegrationPointsHealthCheck provides a mock function with given fields:
func (_m *ServiceInterface) IntegrationPointsHealthCheck() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceInterface_IntegrationPointsHealthCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IntegrationPointsHealthCheck'
type ServiceInterface_IntegrationPointsHealthCheck_Call struct {
	*mock.Call
}

// IntegrationPointsHealthCheck is a helper method to define mock.On call
func (_e *ServiceInterface_Expecter) IntegrationPointsHealthCheck() *ServiceInterface_IntegrationPointsHealthCheck_Call {
	return &ServiceInterface_IntegrationPointsHealthCheck_Call{Call: _e.mock.On("IntegrationPointsHealthCheck")}
}

func (_c *ServiceInterface_IntegrationPointsHealthCheck_Call) Run(run func()) *ServiceInterface_IntegrationPointsHealthCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceInterface_IntegrationPointsHealthCheck_Call) Return(_a0 error) *ServiceInterface_IntegrationPointsHealthCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceInterface_IntegrationPointsHealthCheck_Call) RunAndReturn(run func() error) *ServiceInterface_IntegrationPointsHealthCheck_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}