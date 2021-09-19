// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/jonatan5524/side-projects-manager/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// ParentDirectoryRepository is an autogenerated mock type for the ParentDirectoryRepository type
type ParentDirectoryRepository struct {
	mock.Mock
}

// Put provides a mock function with given fields: _a0
func (_m *ParentDirectoryRepository) Put(_a0 model.ParentDirectory) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(model.ParentDirectory) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.ParentDirectory) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}