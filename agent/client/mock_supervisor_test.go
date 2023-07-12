// Code generated by mockery v2.32.0. DO NOT EDIT.

package client

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"

	agentlocalpb "github.com/percona/pmm/api/agentlocalpb"
	agentpb "github.com/percona/pmm/api/agentpb"
)

// mockSupervisor is an autogenerated mock type for the supervisor type
type mockSupervisor struct {
	mock.Mock
}

// AgentLogByID provides a mock function with given fields: _a0
func (_m *mockSupervisor) AgentLogByID(_a0 string) ([]string, uint) {
	ret := _m.Called(_a0)

	var r0 []string
	var r1 uint
	if rf, ok := ret.Get(0).(func(string) ([]string, uint)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) uint); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(uint)
	}

	return r0, r1
}

// AgentsList provides a mock function with given fields:
func (_m *mockSupervisor) AgentsList() []*agentlocalpb.AgentInfo {
	ret := _m.Called()

	var r0 []*agentlocalpb.AgentInfo
	if rf, ok := ret.Get(0).(func() []*agentlocalpb.AgentInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*agentlocalpb.AgentInfo)
		}
	}

	return r0
}

// Changes provides a mock function with given fields:
func (_m *mockSupervisor) Changes() <-chan *agentpb.StateChangedRequest {
	ret := _m.Called()

	var r0 <-chan *agentpb.StateChangedRequest
	if rf, ok := ret.Get(0).(func() <-chan *agentpb.StateChangedRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *agentpb.StateChangedRequest)
		}
	}

	return r0
}

// ClearChangesChannel provides a mock function with given fields:
func (_m *mockSupervisor) ClearChangesChannel() {
	_m.Called()
}

// Collect provides a mock function with given fields: _a0
func (_m *mockSupervisor) Collect(_a0 chan<- prometheus.Metric) {
	_m.Called(_a0)
}

// Describe provides a mock function with given fields: _a0
func (_m *mockSupervisor) Describe(_a0 chan<- *prometheus.Desc) {
	_m.Called(_a0)
}

// QANRequests provides a mock function with given fields:
func (_m *mockSupervisor) QANRequests() <-chan *agentpb.QANCollectRequest {
	ret := _m.Called()

	var r0 <-chan *agentpb.QANCollectRequest
	if rf, ok := ret.Get(0).(func() <-chan *agentpb.QANCollectRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *agentpb.QANCollectRequest)
		}
	}

	return r0
}

// RestartAgents provides a mock function with given fields:
func (_m *mockSupervisor) RestartAgents() {
	_m.Called()
}

// SetState provides a mock function with given fields: _a0
func (_m *mockSupervisor) SetState(_a0 *agentpb.SetStateRequest) {
	_m.Called(_a0)
}

// newMockSupervisor creates a new instance of mockSupervisor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSupervisor(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockSupervisor {
	mock := &mockSupervisor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
