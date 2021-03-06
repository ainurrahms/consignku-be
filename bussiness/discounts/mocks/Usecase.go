// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	discounts "consignku/bussiness/discounts"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, discountsDomain
func (_m *Usecase) Delete(ctx context.Context, discountsDomain *discounts.Domain) (*discounts.Domain, error) {
	ret := _m.Called(ctx, discountsDomain)

	var r0 *discounts.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *discounts.Domain) *discounts.Domain); ok {
		r0 = rf(ctx, discountsDomain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discounts.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *discounts.Domain) error); ok {
		r1 = rf(ctx, discountsDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Usecase) GetAll(ctx context.Context) ([]discounts.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []discounts.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []discounts.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]discounts.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetByID(ctx context.Context, id int) (discounts.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 discounts.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) discounts.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(discounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, discountsDomain
func (_m *Usecase) Store(ctx context.Context, discountsDomain *discounts.Domain) error {
	ret := _m.Called(ctx, discountsDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discounts.Domain) error); ok {
		r0 = rf(ctx, discountsDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, discountsDomain
func (_m *Usecase) Update(ctx context.Context, discountsDomain *discounts.Domain) (*discounts.Domain, error) {
	ret := _m.Called(ctx, discountsDomain)

	var r0 *discounts.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *discounts.Domain) *discounts.Domain); ok {
		r0 = rf(ctx, discountsDomain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discounts.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *discounts.Domain) error); ok {
		r1 = rf(ctx, discountsDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
