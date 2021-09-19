// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// DirectoryGetter is an autogenerated mock type for the DirectoryGetter type
type DirectoryGetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: path
func (_m *DirectoryGetter) Execute(path string) (fs.FileInfo, error) {
	ret := _m.Called(path)

	var r0 fs.FileInfo
	if rf, ok := ret.Get(0).(func(string) fs.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
