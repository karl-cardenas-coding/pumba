// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import filters "github.com/docker/docker/api/types/filters"
import mock "github.com/stretchr/testify/mock"
import network "github.com/docker/docker/api/types/network"
import types "github.com/docker/docker/api/types"

// NetworkAPIClient is an autogenerated mock type for the NetworkAPIClient type
type NetworkAPIClient struct {
	mock.Mock
}

// NetworkConnect provides a mock function with given fields: ctx, _a1, container, config
func (_m *NetworkAPIClient) NetworkConnect(ctx context.Context, _a1 string, container string, config *network.EndpointSettings) error {
	ret := _m.Called(ctx, _a1, container, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *network.EndpointSettings) error); ok {
		r0 = rf(ctx, _a1, container, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworkCreate provides a mock function with given fields: ctx, name, options
func (_m *NetworkAPIClient) NetworkCreate(ctx context.Context, name string, options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	ret := _m.Called(ctx, name, options)

	var r0 types.NetworkCreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NetworkCreate) types.NetworkCreateResponse); ok {
		r0 = rf(ctx, name, options)
	} else {
		r0 = ret.Get(0).(types.NetworkCreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NetworkCreate) error); ok {
		r1 = rf(ctx, name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkDisconnect provides a mock function with given fields: ctx, _a1, container, force
func (_m *NetworkAPIClient) NetworkDisconnect(ctx context.Context, _a1 string, container string, force bool) error {
	ret := _m.Called(ctx, _a1, container, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, _a1, container, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworkInspect provides a mock function with given fields: ctx, _a1, options
func (_m *NetworkAPIClient) NetworkInspect(ctx context.Context, _a1 string, options types.NetworkInspectOptions) (types.NetworkResource, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NetworkInspectOptions) types.NetworkResource); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.NetworkResource)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NetworkInspectOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkInspectWithRaw provides a mock function with given fields: ctx, _a1, options
func (_m *NetworkAPIClient) NetworkInspectWithRaw(ctx context.Context, _a1 string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NetworkInspectOptions) types.NetworkResource); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.NetworkResource)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NetworkInspectOptions) []byte); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, types.NetworkInspectOptions) error); ok {
		r2 = rf(ctx, _a1, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NetworkList provides a mock function with given fields: ctx, options
func (_m *NetworkAPIClient) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.NetworkResource
	if rf, ok := ret.Get(0).(func(context.Context, types.NetworkListOptions) []types.NetworkResource); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NetworkResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.NetworkListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkRemove provides a mock function with given fields: ctx, _a1
func (_m *NetworkAPIClient) NetworkRemove(ctx context.Context, _a1 string) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworksPrune provides a mock function with given fields: ctx, pruneFilter
func (_m *NetworkAPIClient) NetworksPrune(ctx context.Context, pruneFilter filters.Args) (types.NetworksPruneReport, error) {
	ret := _m.Called(ctx, pruneFilter)

	var r0 types.NetworksPruneReport
	if rf, ok := ret.Get(0).(func(context.Context, filters.Args) types.NetworksPruneReport); ok {
		r0 = rf(ctx, pruneFilter)
	} else {
		r0 = ret.Get(0).(types.NetworksPruneReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filters.Args) error); ok {
		r1 = rf(ctx, pruneFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}