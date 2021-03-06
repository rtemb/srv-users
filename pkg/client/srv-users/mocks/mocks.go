// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	srv_users "github.com/rtemb/srv-users/pkg/client/srv-users"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SrvUsersMock struct {
	AddRoleStub        func(context.Context, *srv_users.AddRoleRequest, ...grpc.CallOption) (*emptypb.Empty, error)
	addRoleMutex       sync.RWMutex
	addRoleArgsForCall []struct {
		arg1 context.Context
		arg2 *srv_users.AddRoleRequest
		arg3 []grpc.CallOption
	}
	addRoleReturns struct {
		result1 *emptypb.Empty
		result2 error
	}
	addRoleReturnsOnCall map[int]struct {
		result1 *emptypb.Empty
		result2 error
	}
	AuthStub        func(context.Context, *srv_users.AuthRequest, ...grpc.CallOption) (*srv_users.AuthResponse, error)
	authMutex       sync.RWMutex
	authArgsForCall []struct {
		arg1 context.Context
		arg2 *srv_users.AuthRequest
		arg3 []grpc.CallOption
	}
	authReturns struct {
		result1 *srv_users.AuthResponse
		result2 error
	}
	authReturnsOnCall map[int]struct {
		result1 *srv_users.AuthResponse
		result2 error
	}
	CreateUserStub        func(context.Context, *srv_users.CreateUserRequest, ...grpc.CallOption) (*srv_users.CreateUserResponse, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 context.Context
		arg2 *srv_users.CreateUserRequest
		arg3 []grpc.CallOption
	}
	createUserReturns struct {
		result1 *srv_users.CreateUserResponse
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 *srv_users.CreateUserResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SrvUsersMock) AddRole(arg1 context.Context, arg2 *srv_users.AddRoleRequest, arg3 ...grpc.CallOption) (*emptypb.Empty, error) {
	fake.addRoleMutex.Lock()
	ret, specificReturn := fake.addRoleReturnsOnCall[len(fake.addRoleArgsForCall)]
	fake.addRoleArgsForCall = append(fake.addRoleArgsForCall, struct {
		arg1 context.Context
		arg2 *srv_users.AddRoleRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddRole", []interface{}{arg1, arg2, arg3})
	fake.addRoleMutex.Unlock()
	if fake.AddRoleStub != nil {
		return fake.AddRoleStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addRoleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SrvUsersMock) AddRoleCallCount() int {
	fake.addRoleMutex.RLock()
	defer fake.addRoleMutex.RUnlock()
	return len(fake.addRoleArgsForCall)
}

func (fake *SrvUsersMock) AddRoleCalls(stub func(context.Context, *srv_users.AddRoleRequest, ...grpc.CallOption) (*emptypb.Empty, error)) {
	fake.addRoleMutex.Lock()
	defer fake.addRoleMutex.Unlock()
	fake.AddRoleStub = stub
}

func (fake *SrvUsersMock) AddRoleArgsForCall(i int) (context.Context, *srv_users.AddRoleRequest, []grpc.CallOption) {
	fake.addRoleMutex.RLock()
	defer fake.addRoleMutex.RUnlock()
	argsForCall := fake.addRoleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SrvUsersMock) AddRoleReturns(result1 *emptypb.Empty, result2 error) {
	fake.addRoleMutex.Lock()
	defer fake.addRoleMutex.Unlock()
	fake.AddRoleStub = nil
	fake.addRoleReturns = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) AddRoleReturnsOnCall(i int, result1 *emptypb.Empty, result2 error) {
	fake.addRoleMutex.Lock()
	defer fake.addRoleMutex.Unlock()
	fake.AddRoleStub = nil
	if fake.addRoleReturnsOnCall == nil {
		fake.addRoleReturnsOnCall = make(map[int]struct {
			result1 *emptypb.Empty
			result2 error
		})
	}
	fake.addRoleReturnsOnCall[i] = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) Auth(arg1 context.Context, arg2 *srv_users.AuthRequest, arg3 ...grpc.CallOption) (*srv_users.AuthResponse, error) {
	fake.authMutex.Lock()
	ret, specificReturn := fake.authReturnsOnCall[len(fake.authArgsForCall)]
	fake.authArgsForCall = append(fake.authArgsForCall, struct {
		arg1 context.Context
		arg2 *srv_users.AuthRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Auth", []interface{}{arg1, arg2, arg3})
	fake.authMutex.Unlock()
	if fake.AuthStub != nil {
		return fake.AuthStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.authReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SrvUsersMock) AuthCallCount() int {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	return len(fake.authArgsForCall)
}

func (fake *SrvUsersMock) AuthCalls(stub func(context.Context, *srv_users.AuthRequest, ...grpc.CallOption) (*srv_users.AuthResponse, error)) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = stub
}

func (fake *SrvUsersMock) AuthArgsForCall(i int) (context.Context, *srv_users.AuthRequest, []grpc.CallOption) {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	argsForCall := fake.authArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SrvUsersMock) AuthReturns(result1 *srv_users.AuthResponse, result2 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	fake.authReturns = struct {
		result1 *srv_users.AuthResponse
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) AuthReturnsOnCall(i int, result1 *srv_users.AuthResponse, result2 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	if fake.authReturnsOnCall == nil {
		fake.authReturnsOnCall = make(map[int]struct {
			result1 *srv_users.AuthResponse
			result2 error
		})
	}
	fake.authReturnsOnCall[i] = struct {
		result1 *srv_users.AuthResponse
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) CreateUser(arg1 context.Context, arg2 *srv_users.CreateUserRequest, arg3 ...grpc.CallOption) (*srv_users.CreateUserResponse, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 context.Context
		arg2 *srv_users.CreateUserRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SrvUsersMock) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *SrvUsersMock) CreateUserCalls(stub func(context.Context, *srv_users.CreateUserRequest, ...grpc.CallOption) (*srv_users.CreateUserResponse, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *SrvUsersMock) CreateUserArgsForCall(i int) (context.Context, *srv_users.CreateUserRequest, []grpc.CallOption) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SrvUsersMock) CreateUserReturns(result1 *srv_users.CreateUserResponse, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 *srv_users.CreateUserResponse
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) CreateUserReturnsOnCall(i int, result1 *srv_users.CreateUserResponse, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 *srv_users.CreateUserResponse
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 *srv_users.CreateUserResponse
		result2 error
	}{result1, result2}
}

func (fake *SrvUsersMock) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addRoleMutex.RLock()
	defer fake.addRoleMutex.RUnlock()
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SrvUsersMock) recordInvocation(key string, args []interface{}) {
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

var _ srv_users.UsersServiceClient = new(SrvUsersMock)
