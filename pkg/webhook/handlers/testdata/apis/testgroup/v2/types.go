/*
Copyright 2020 The cert-manager Authors.

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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:storageversion
// TestType in v2 is identical to v1, except TestFieldPtr has been renamed to TestFieldPtrAlt
type TestType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// TestField is used in tests.
	// Validation doesn't allow this to be set to the value of TestFieldValueNotAllowed.
	TestField string `json:"testField"`
	// +optional
	TestFieldPtrAlt *string `json:"testFieldPtrAlt,omitempty"`

	// TestFieldImmutable cannot be changed after being set to a non-zero value
	TestFieldImmutable string `json:"testFieldImmutable"`
}

const (
	DisallowedTestFieldValue = "not-allowed-in-v2"
)
