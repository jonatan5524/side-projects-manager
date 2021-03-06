// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/jonatan5524/side-projects-manager/pkg/model"
	util "github.com/jonatan5524/side-projects-manager/pkg/util/io"
	mock "github.com/stretchr/testify/mock"
)

// ParentDirectoryConstructor is an autogenerated mock type for the ParentDirectoryConstructor type
type ParentDirectoryConstructor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *ParentDirectoryConstructor) Execute(_a0 string, _a1 util.DirectoryGetter) (model.ParentDirectory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.ParentDirectory
	if rf, ok := ret.Get(0).(func(string, util.DirectoryGetter) model.ParentDirectory); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.ParentDirectory)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, util.DirectoryGetter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
