// Code generated by mockery v2.32.0. DO NOT EDIT.

package management

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAgentsStateUpdater is an autogenerated mock type for the agentsStateUpdater type
type mockAgentsStateUpdater struct {
	mock.Mock
}

// RequestStateUpdate provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockAgentsStateUpdater) RequestStateUpdate(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}

// newMockAgentsStateUpdater creates a new instance of mockAgentsStateUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAgentsStateUpdater(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAgentsStateUpdater {
	mock := &mockAgentsStateUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
