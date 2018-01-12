// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: RefreshTokenGrantStorage)

// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/spotxchange/fosite"
)

// Mock of RefreshTokenGrantStorage interface
type MockRefreshTokenGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockRefreshTokenGrantStorageRecorder
}

// Recorder for MockRefreshTokenGrantStorage (not exported)
type _MockRefreshTokenGrantStorageRecorder struct {
	mock *MockRefreshTokenGrantStorage
}

func NewMockRefreshTokenGrantStorage(ctrl *gomock.Controller) *MockRefreshTokenGrantStorage {
	mock := &MockRefreshTokenGrantStorage{ctrl: ctrl}
	mock.recorder = &_MockRefreshTokenGrantStorageRecorder{mock}
	return mock
}

func (_m *MockRefreshTokenGrantStorage) EXPECT() *_MockRefreshTokenGrantStorageRecorder {
	return _m.recorder
}

func (_m *MockRefreshTokenGrantStorage) CreateRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRefreshTokenGrantStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockRefreshTokenGrantStorage) DeleteRefreshTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRefreshTokenGrantStorageRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0, arg1)
}

func (_m *MockRefreshTokenGrantStorage) GetRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRefreshTokenGrantStorageRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockRefreshTokenGrantStorage) PersistRefreshTokenGrantSession(_param0 context.Context, _param1 string, _param2 string, _param3 string, _param4 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "PersistRefreshTokenGrantSession", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRefreshTokenGrantStorageRecorder) PersistRefreshTokenGrantSession(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PersistRefreshTokenGrantSession", arg0, arg1, arg2, arg3, arg4)
}
