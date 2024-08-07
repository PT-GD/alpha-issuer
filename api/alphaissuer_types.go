/*
Copyright 2024.

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

package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	issuer "github.com/cert-manager/issuer-lib/api/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlphaIssuerSpec defines the desired state of AlphaIssuer
type AlphaIssuerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AlphaIssuer. Edit alphaissuer_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// AlphaIssuerStatus defines the observed state of AlphaIssuer
type AlphaIssuerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message"
// +kubebuilder:printcolumn:name="LastTransition",type="string",type="date",JSONPath=".status.conditions[?(@.type==\"Ready\")].lastTransitionTime"
// +kubebuilder:printcolumn:name="ObservedGeneration",type="integer",JSONPath=".status.conditions[?(@.type==\"Ready\")].observedGeneration"
// +kubebuilder:printcolumn:name="Generation",type="integer",JSONPath=".metadata.generation"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// AlphaIssuer is the Schema for the alphaissuers API
type AlphaIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlphaIssuerSpec     `json:"spec,omitempty"`
	Status issuer.IssuerStatus `json:"status,omitempty"`
}

func (vi *AlphaIssuer) GetStatus() *issuer.IssuerStatus {
	return &vi.Status
}

func (vi *AlphaIssuer) GetIssuerTypeIdentifier() string {
	return "alphaissuers.certmanager.alpha-issuer.io"
}

var _ issuer.Issuer = &AlphaIssuer{}

// +kubebuilder:object:root=true

// AlphaIssuerList contains a list of AlphaIssuer
type AlphaIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlphaIssuer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlphaIssuer{}, &AlphaIssuerList{})
}
