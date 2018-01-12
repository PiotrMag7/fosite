// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/openid (interfaces: OpenIDConnectTokenStrategy)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/spotxchange/fosite"
)

// Mock of OpenIDConnectTokenStrategy interface
type MockOpenIDConnectTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockOpenIDConnectTokenStrategyRecorder
}

// Recorder for MockOpenIDConnectTokenStrategy (not exported)
type _MockOpenIDConnectTokenStrategyRecorder struct {
	mock *MockOpenIDConnectTokenStrategy
}

func NewMockOpenIDConnectTokenStrategy(ctrl *gomock.Controller) *MockOpenIDConnectTokenStrategy {
	mock := &MockOpenIDConnectTokenStrategy{ctrl: ctrl}
	mock.recorder = &_MockOpenIDConnectTokenStrategyRecorder{mock}
	return mock
}

func (_m *MockOpenIDConnectTokenStrategy) EXPECT() *_MockOpenIDConnectTokenStrategyRecorder {
	return _m.recorder
}

func (_m *MockOpenIDConnectTokenStrategy) GenerateIDToken(_param0 context.Context, _param1 fosite.Requester) (string, error) {
	ret := _m.ctrl.Call(_m, "GenerateIDToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOpenIDConnectTokenStrategyRecorder) GenerateIDToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateIDToken", arg0, arg1)
}
