// Code generated by MockGen. DO NOT EDIT.
// Source: internal/mempool/actions.go

// Package mempool is a generated GoMock package.
package mempool

import (
	gomock "github.com/golang/mock/gomock"
	chainindex "github.com/olympus-protocol/ogen/internal/chainindex"
	state "github.com/olympus-protocol/ogen/internal/state"
	primitives "github.com/olympus-protocol/ogen/pkg/primitives"
	reflect "reflect"
)

// MockActionMempool is a mock of ActionMempool interface
type MockActionMempool struct {
	ctrl     *gomock.Controller
	recorder *MockActionMempoolMockRecorder
}

// MockActionMempoolMockRecorder is the mock recorder for MockActionMempool
type MockActionMempoolMockRecorder struct {
	mock *MockActionMempool
}

// NewMockActionMempool creates a new mock instance
func NewMockActionMempool(ctrl *gomock.Controller) *MockActionMempool {
	mock := &MockActionMempool{ctrl: ctrl}
	mock.recorder = &MockActionMempoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionMempool) EXPECT() *MockActionMempoolMockRecorder {
	return m.recorder
}

// NotifyIllegalVotes mocks base method
func (m *MockActionMempool) NotifyIllegalVotes(slashing *primitives.VoteSlashing) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyIllegalVotes", slashing)
}

// NotifyIllegalVotes indicates an expected call of NotifyIllegalVotes
func (mr *MockActionMempoolMockRecorder) NotifyIllegalVotes(slashing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyIllegalVotes", reflect.TypeOf((*MockActionMempool)(nil).NotifyIllegalVotes), slashing)
}

// NewTip mocks base method
func (m *MockActionMempool) NewTip(arg0 *chainindex.BlockRow, arg1 *primitives.Block, arg2 state.State, arg3 []*primitives.EpochReceipt) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewTip", arg0, arg1, arg2, arg3)
}

// NewTip indicates an expected call of NewTip
func (mr *MockActionMempoolMockRecorder) NewTip(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTip", reflect.TypeOf((*MockActionMempool)(nil).NewTip), arg0, arg1, arg2, arg3)
}

// ProposerSlashingConditionViolated mocks base method
func (m *MockActionMempool) ProposerSlashingConditionViolated(slashing *primitives.ProposerSlashing) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProposerSlashingConditionViolated", slashing)
}

// ProposerSlashingConditionViolated indicates an expected call of ProposerSlashingConditionViolated
func (mr *MockActionMempoolMockRecorder) ProposerSlashingConditionViolated(slashing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposerSlashingConditionViolated", reflect.TypeOf((*MockActionMempool)(nil).ProposerSlashingConditionViolated), slashing)
}

// AddDeposit mocks base method
func (m *MockActionMempool) AddDeposit(deposit *primitives.Deposit, state state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeposit", deposit, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDeposit indicates an expected call of AddDeposit
func (mr *MockActionMempoolMockRecorder) AddDeposit(deposit, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeposit", reflect.TypeOf((*MockActionMempool)(nil).AddDeposit), deposit, state)
}

// GetDeposits mocks base method
func (m *MockActionMempool) GetDeposits(num int, withState state.State) ([]*primitives.Deposit, state.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeposits", num, withState)
	ret0, _ := ret[0].([]*primitives.Deposit)
	ret1, _ := ret[1].(state.State)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeposits indicates an expected call of GetDeposits
func (mr *MockActionMempoolMockRecorder) GetDeposits(num, withState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeposits", reflect.TypeOf((*MockActionMempool)(nil).GetDeposits), num, withState)
}

// RemoveByBlock mocks base method
func (m *MockActionMempool) RemoveByBlock(b *primitives.Block, tipState state.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveByBlock", b, tipState)
}

// RemoveByBlock indicates an expected call of RemoveByBlock
func (mr *MockActionMempoolMockRecorder) RemoveByBlock(b, tipState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByBlock", reflect.TypeOf((*MockActionMempool)(nil).RemoveByBlock), b, tipState)
}

// AddGovernanceVote mocks base method
func (m *MockActionMempool) AddGovernanceVote(vote *primitives.GovernanceVote, state state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGovernanceVote", vote, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGovernanceVote indicates an expected call of AddGovernanceVote
func (mr *MockActionMempoolMockRecorder) AddGovernanceVote(vote, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGovernanceVote", reflect.TypeOf((*MockActionMempool)(nil).AddGovernanceVote), vote, state)
}

// AddExit mocks base method
func (m *MockActionMempool) AddExit(exit *primitives.Exit, state state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExit", exit, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddExit indicates an expected call of AddExit
func (mr *MockActionMempoolMockRecorder) AddExit(exit, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExit", reflect.TypeOf((*MockActionMempool)(nil).AddExit), exit, state)
}

// GetProposerSlashings mocks base method
func (m *MockActionMempool) GetProposerSlashings(num int, state state.State) ([]*primitives.ProposerSlashing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerSlashings", num, state)
	ret0, _ := ret[0].([]*primitives.ProposerSlashing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposerSlashings indicates an expected call of GetProposerSlashings
func (mr *MockActionMempoolMockRecorder) GetProposerSlashings(num, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerSlashings", reflect.TypeOf((*MockActionMempool)(nil).GetProposerSlashings), num, state)
}

// GetExits mocks base method
func (m *MockActionMempool) GetExits(num int, state state.State) ([]*primitives.Exit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExits", num, state)
	ret0, _ := ret[0].([]*primitives.Exit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExits indicates an expected call of GetExits
func (mr *MockActionMempoolMockRecorder) GetExits(num, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExits", reflect.TypeOf((*MockActionMempool)(nil).GetExits), num, state)
}

// GetVoteSlashings mocks base method
func (m *MockActionMempool) GetVoteSlashings(num int, state state.State) ([]*primitives.VoteSlashing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoteSlashings", num, state)
	ret0, _ := ret[0].([]*primitives.VoteSlashing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoteSlashings indicates an expected call of GetVoteSlashings
func (mr *MockActionMempoolMockRecorder) GetVoteSlashings(num, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoteSlashings", reflect.TypeOf((*MockActionMempool)(nil).GetVoteSlashings), num, state)
}

// GetRANDAOSlashings mocks base method
func (m *MockActionMempool) GetRANDAOSlashings(num int, state state.State) ([]*primitives.RANDAOSlashing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRANDAOSlashings", num, state)
	ret0, _ := ret[0].([]*primitives.RANDAOSlashing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRANDAOSlashings indicates an expected call of GetRANDAOSlashings
func (mr *MockActionMempoolMockRecorder) GetRANDAOSlashings(num, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRANDAOSlashings", reflect.TypeOf((*MockActionMempool)(nil).GetRANDAOSlashings), num, state)
}

// GetGovernanceVotes mocks base method
func (m *MockActionMempool) GetGovernanceVotes(num int, state state.State) ([]*primitives.GovernanceVote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGovernanceVotes", num, state)
	ret0, _ := ret[0].([]*primitives.GovernanceVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGovernanceVotes indicates an expected call of GetGovernanceVotes
func (mr *MockActionMempoolMockRecorder) GetGovernanceVotes(num, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGovernanceVotes", reflect.TypeOf((*MockActionMempool)(nil).GetGovernanceVotes), num, state)
}