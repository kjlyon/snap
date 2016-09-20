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

//type mockTribe struct{}

var mockTribeAgreement *agreement.Agreement
var mockTribeMember *agreement.Member

func init() {
	mockTribeAgreement = agreement.New("MyNewTribeNode")
	//var mlN *memberlist.Node
	//mockTribeMember = agreement.NewMember(mlN)

}

type MockTribeManager struct{}

func (m *MockTribeManager) GetAgreement(name string) (*agreement.Agreement, serror.SnapError) {
	return mockTribeAgreement, nil
}
func (m *MockTribeManager) GetAgreements() map[string]*agreement.Agreement {
	return map[string]*agreement.Agreement{"Agree1": mockTribeAgreement, "Agree2": mockTribeAgreement}
}
func (m *MockTribeManager) AddAgreement(name string) serror.SnapError    { return nil }
func (m *MockTribeManager) RemoveAgreement(name string) serror.SnapError { return nil }
func (m *MockTribeManager) JoinAgreement(agreementName, memberName string) serror.SnapError {
	return nil
}
func (m *MockTribeManager) LeaveAgreement(agreementName, memberName string) serror.SnapError {
	return nil
}
func (m *MockTribeManager) GetMembers() []string                    { return []string{"one", "two", "three"} }
func (m *MockTribeManager) GetMember(name string) *agreement.Member { return &agreement.Member{} }

const (
	GET_TRIBE_AGREEMENTS = `{
  "meta": {
    "code": 200,
    "message": "Tribe agreements retrieved",
    "type": "tribe_agreement_list_returned",
    "version": 1
  },
  "body": {
    "agreements": {
      "Agree1": {
        "name": "MyNewTribeNode",
        "plugin_agreement": {},
        "task_agreement": {}
      },
      "Agree2": {
        "name": "MyNewTribeNode",
        "plugin_agreement": {},
        "task_agreement": {}
      }
    }
  }
}`

	GET_TRIBE_AGREEMENTS_NAME = `{
  "meta": {
    "code": 200,
    "message": "Tribe agreement returned",
    "type": "tribe_agreement_returned",
    "version": 1
  },
  "body": {
    "agreement": {
      "name": "MyNewTribeNode",
      "plugin_agreement": {},
      "task_agreement": {}
    }
  }
}`

	GET_TRIBE_MEMBERS = `{
  "meta": {
    "code": 200,
    "message": "Tribe members retrieved",
    "type": "tribe_member_list_returned",
    "version": 1
  },
  "body": {
    "members": [
      "one",
      "two",
      "three"
    ]
  }
}`

	GET_TRIBE_MEMBER_NAME = `{
  "meta": {
    "code": 200,
    "message": "Tribe member details retrieved",
    "type": "tribe_member_details_returned",
    "version": 1
  },
  "body": {
    "name": "",
    "plugin_agreement": "",
    "tags": null,
    "task_agreements": null
  }
}`
)
