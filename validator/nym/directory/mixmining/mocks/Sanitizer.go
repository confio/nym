// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym-directory/models"
	mock "github.com/stretchr/testify/mock"
)

// Sanitizer is an autogenerated mock type for the Sanitizer type
type Sanitizer struct {
	mock.Mock
}

// Sanitize provides a mock function with given fields: input
func (_m *Sanitizer) Sanitize(input models.MixStatus) models.MixStatus {
	ret := _m.Called(input)

	var r0 models.MixStatus
	if rf, ok := ret.Get(0).(func(models.MixStatus) models.MixStatus); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(models.MixStatus)
	}

	return r0
}