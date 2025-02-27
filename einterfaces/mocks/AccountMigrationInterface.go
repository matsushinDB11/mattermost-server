// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// AccountMigrationInterface is an autogenerated mock type for the AccountMigrationInterface type
type AccountMigrationInterface struct {
	mock.Mock
}

// MigrateToLdap provides a mock function with given fields: fromAuthService, forignUserFieldNameToMatch, force, dryRun
func (_m *AccountMigrationInterface) MigrateToLdap(fromAuthService string, forignUserFieldNameToMatch string, force bool, dryRun bool) *model.AppError {
	ret := _m.Called(fromAuthService, forignUserFieldNameToMatch, force, dryRun)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string, bool, bool) *model.AppError); ok {
		r0 = rf(fromAuthService, forignUserFieldNameToMatch, force, dryRun)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MigrateToSaml provides a mock function with given fields: fromAuthService, usersMap, auto, dryRun
func (_m *AccountMigrationInterface) MigrateToSaml(fromAuthService string, usersMap map[string]string, auto bool, dryRun bool) *model.AppError {
	ret := _m.Called(fromAuthService, usersMap, auto, dryRun)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, map[string]string, bool, bool) *model.AppError); ok {
		r0 = rf(fromAuthService, usersMap, auto, dryRun)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}
