// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	models "github.com/caraml-dev/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// EnvironmentService is an autogenerated mock type for the EnvironmentService type
type EnvironmentService struct {
	mock.Mock
}

// GetDefaultEnvironment provides a mock function with given fields:
func (_m *EnvironmentService) GetDefaultEnvironment() (*models.Environment, error) {
	ret := _m.Called()

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func() *models.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultPredictionJobEnvironment provides a mock function with given fields:
func (_m *EnvironmentService) GetDefaultPredictionJobEnvironment() (*models.Environment, error) {
	ret := _m.Called()

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func() *models.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnvironment provides a mock function with given fields: name
func (_m *EnvironmentService) GetEnvironment(name string) (*models.Environment, error) {
	ret := _m.Called(name)

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func(string) *models.Environment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEnvironments provides a mock function with given fields: name
func (_m *EnvironmentService) ListEnvironments(name string) ([]*models.Environment, error) {
	ret := _m.Called(name)

	var r0 []*models.Environment
	if rf, ok := ret.Get(0).(func(string) []*models.Environment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: env
func (_m *EnvironmentService) Save(env *models.Environment) (*models.Environment, error) {
	ret := _m.Called(env)

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func(*models.Environment) *models.Environment); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Environment) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
