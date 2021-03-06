// Code generated by counterfeiter. DO NOT EDIT.
package permissionsfakes

import (
	"context"
	"sync"

	"github.com/rleszilm/permissions"
)

type FakeRoleProvider struct {
	RolesStub        func(context.Context, permissions.UserID, permissions.GroupID) ([]permissions.Role, error)
	rolesMutex       sync.RWMutex
	rolesArgsForCall []struct {
		arg1 context.Context
		arg2 permissions.UserID
		arg3 permissions.GroupID
	}
	rolesReturns struct {
		result1 []permissions.Role
		result2 error
	}
	rolesReturnsOnCall map[int]struct {
		result1 []permissions.Role
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoleProvider) Roles(arg1 context.Context, arg2 permissions.UserID, arg3 permissions.GroupID) ([]permissions.Role, error) {
	fake.rolesMutex.Lock()
	ret, specificReturn := fake.rolesReturnsOnCall[len(fake.rolesArgsForCall)]
	fake.rolesArgsForCall = append(fake.rolesArgsForCall, struct {
		arg1 context.Context
		arg2 permissions.UserID
		arg3 permissions.GroupID
	}{arg1, arg2, arg3})
	fake.recordInvocation("Roles", []interface{}{arg1, arg2, arg3})
	fake.rolesMutex.Unlock()
	if fake.RolesStub != nil {
		return fake.RolesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.rolesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRoleProvider) RolesCallCount() int {
	fake.rolesMutex.RLock()
	defer fake.rolesMutex.RUnlock()
	return len(fake.rolesArgsForCall)
}

func (fake *FakeRoleProvider) RolesCalls(stub func(context.Context, permissions.UserID, permissions.GroupID) ([]permissions.Role, error)) {
	fake.rolesMutex.Lock()
	defer fake.rolesMutex.Unlock()
	fake.RolesStub = stub
}

func (fake *FakeRoleProvider) RolesArgsForCall(i int) (context.Context, permissions.UserID, permissions.GroupID) {
	fake.rolesMutex.RLock()
	defer fake.rolesMutex.RUnlock()
	argsForCall := fake.rolesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoleProvider) RolesReturns(result1 []permissions.Role, result2 error) {
	fake.rolesMutex.Lock()
	defer fake.rolesMutex.Unlock()
	fake.RolesStub = nil
	fake.rolesReturns = struct {
		result1 []permissions.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleProvider) RolesReturnsOnCall(i int, result1 []permissions.Role, result2 error) {
	fake.rolesMutex.Lock()
	defer fake.rolesMutex.Unlock()
	fake.RolesStub = nil
	if fake.rolesReturnsOnCall == nil {
		fake.rolesReturnsOnCall = make(map[int]struct {
			result1 []permissions.Role
			result2 error
		})
	}
	fake.rolesReturnsOnCall[i] = struct {
		result1 []permissions.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeRoleProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.rolesMutex.RLock()
	defer fake.rolesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoleProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ permissions.RoleProvider = new(FakeRoleProvider)
