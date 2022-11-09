// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	database "kek-backend/internal/alert/database"

	mock "github.com/stretchr/testify/mock"

	model "kek-backend/internal/alert/model"
)

// AlertDB is an autogenerated mock type for the AlertDB type
type AlertDB struct {
	mock.Mock
}

// DeleteAlertBySlug provides a mock function with given fields: ctx, accountId, slug
func (_m *AlertDB) DeleteAlertBySlug(ctx context.Context, accountId uint, slug string) error {
	ret := _m.Called(ctx, accountId, slug)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) error); ok {
		r0 = rf(ctx, accountId, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCommentById provides a mock function with given fields: ctx, accountId, slug, id
func (_m *AlertDB) DeleteCommentById(ctx context.Context, accountId uint, slug string, id uint) error {
	ret := _m.Called(ctx, accountId, slug, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, uint) error); ok {
		r0 = rf(ctx, accountId, slug, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComments provides a mock function with given fields: ctx, accountId, slug
func (_m *AlertDB) DeleteComments(ctx context.Context, accountId uint, slug string) (int64, error) {
	ret := _m.Called(ctx, accountId, slug)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) int64); ok {
		r0 = rf(ctx, accountId, slug)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, accountId, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAlertBySlug provides a mock function with given fields: ctx, slug
func (_m *AlertDB) FindAlertBySlug(ctx context.Context, slug string) (*model.Alert, error) {
	ret := _m.Called(ctx, slug)

	var r0 *model.Alert
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Alert); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAlerts provides a mock function with given fields: ctx, criteria
func (_m *AlertDB) FindAlerts(ctx context.Context, criteria database.IterateAlertCriteria) ([]*model.Alert, int64, error) {
	ret := _m.Called(ctx, criteria)

	var r0 []*model.Alert
	if rf, ok := ret.Get(0).(func(context.Context, database.IterateAlertCriteria) []*model.Alert); ok {
		r0 = rf(ctx, criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Alert)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, database.IterateAlertCriteria) int64); ok {
		r1 = rf(ctx, criteria)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, database.IterateAlertCriteria) error); ok {
		r2 = rf(ctx, criteria)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RunInTx provides a mock function with given fields: ctx, f
func (_m *AlertDB) RunInTx(ctx context.Context, f func(context.Context) error) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAlert provides a mock function with given fields: ctx, alert
func (_m *AlertDB) SaveAlert(ctx context.Context, alert *model.Alert) error {
	ret := _m.Called(ctx, alert)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Alert) error); ok {
		r0 = rf(ctx, alert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
