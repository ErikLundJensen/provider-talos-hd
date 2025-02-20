/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// serverParameters are the configurable fields of a server.
type serverParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// serverObservation are the observable fields of a server.
type serverObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A serverSpec defines the desired state of a server.
type serverSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       serverParameters `json:"forProvider"`
}

// A serverStatus represents the observed state of a server.
type serverStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          serverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A server is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,taloshd}
type server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   serverSpec   `json:"spec"`
	Status serverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// serverList contains a list of server
type serverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []server `json:"items"`
}

// server type metadata.
var (
	serverKind             = reflect.TypeOf(server{}).Name()
	serverGroupKind        = schema.GroupKind{Group: Group, Kind: serverKind}.String()
	serverKindAPIVersion   = serverKind + "." + SchemeGroupVersion.String()
	serverGroupVersionKind = SchemeGroupVersion.WithKind(serverKind)
)

func init() {
	SchemeBuilder.Register(&server{}, &serverList{})
}
