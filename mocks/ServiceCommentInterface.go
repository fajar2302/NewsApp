// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	comments "NEWSAPP/features/Comments"

	mock "github.com/stretchr/testify/mock"
)

// ServiceCommentInterface is an autogenerated mock type for the ServiceCommentInterface type
type ServiceCommentInterface struct {
	mock.Mock
}

// CreateNewComment provides a mock function with given fields: articlesid, comment
func (_m *ServiceCommentInterface) CreateNewComment(articlesid uint, comment comments.Comment) error {
	ret := _m.Called(articlesid, comment)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, comments.Comment) error); ok {
		r0 = rf(articlesid, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: commentID
func (_m *ServiceCommentInterface) DeleteComment(commentID uint) error {
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
func (_m *ServiceCommentInterface) GetAllComments() ([]comments.Comment, error) {
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

// NewServiceCommentInterface creates a new instance of ServiceCommentInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceCommentInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceCommentInterface {
	mock := &ServiceCommentInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
