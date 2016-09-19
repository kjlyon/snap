// + build legacy small medium large

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fixtures

import (
	"github.com/intelsdi-x/snap/core/serror"
	"github.com/intelsdi-x/snap/mgmt/tribe/agreement"
)

type mockTribe struct{}

type MockTribeManager struct{}

func (m *MockTribeManager) GetAgreement(name string) (*agreement.Agreement, serror.SnapError) {
	return nil, nil
}
func (m *MockTribeManager) GetAgreements() map[string]*agreement.Agreement { return nil }
func (m *MockTribeManager) AddAgreement(name string) serror.SnapError      { return nil }
func (m *MockTribeManager) RemoveAgreement(name string) serror.SnapError   { return nil }
func (m *MockTribeManager) JoinAgreement(agreementName, memberName string) serror.SnapError {
	return nil
}
func (m *MockTribeManager) LeaveAgreement(agreementName, memberName string) serror.SnapError {
	return nil
}
func (m *MockTribeManager) GetMembers() []string                    { return nil }
func (m *MockTribeManager) GetMember(name string) *agreement.Member { return nil }
