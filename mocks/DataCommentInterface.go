// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	comments "NEWSAPP/features/Comments"

	mock "github.com/stretchr/testify/mock"
)

// DataCommentInterface is an autogenerated mock type for the DataCommentInterface type
type DataCommentInterface struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: comment
func (_m *DataCommentInterface) CreateComment(comment comments.Comment) error {
	ret := _m.Called(comment)

	if len(ret) == 0 {
		panic("no return value specified for CreateComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(comments.Comment) error); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: commentID
func (_m *DataCommentInterface) DeleteComment(commentID uint) error {
	ret := _m.Called(commentID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllComments provides a mock function with given fields:
func (_m *DataCommentInterface) GetAllComments() ([]comments.Comment, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllComments")
	}

	var r0 []comments.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]comments.Comment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []comments.Comment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDataCommentInterface creates a new instance of DataCommentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataCommentInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataCommentInterface {
	mock := &DataCommentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
