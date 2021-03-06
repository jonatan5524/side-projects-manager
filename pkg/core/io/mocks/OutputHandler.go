// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fmt "fmt"

	core "github.com/jonatan5524/side-projects-manager/pkg/core/io"

	mock "github.com/stretchr/testify/mock"
)

// OutputHandler is an autogenerated mock type for the OutputHandler type
type OutputHandler struct {
	mock.Mock
}

// PrintObject provides a mock function with given fields: _a0
func (_m *OutputHandler) PrintObject(_a0 fmt.Stringer) {
	_m.Called(_a0)
}

// PrintString provides a mock function with given fields: _a0
func (_m *OutputHandler) PrintString(_a0 string) {
	_m.Called(_a0)
}

// PrintTable provides a mock function with given fields: _a0
func (_m *OutputHandler) PrintTable(_a0 []core.Tabler) {
	_m.Called(_a0)
}
