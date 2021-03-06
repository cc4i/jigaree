/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JigareeSpec defines the desired state of Jigaree
type JigareeSpec struct {
	Air     AirSpec     `json:"air"`
	Gate    GateSpec    `json:"gate"`
	Website WebsiteSpec `json:"website"`
	Weather WeatherSpec `json:"weather"`
}

type GateSpec struct {
	ReplicaNum string `json:"replicaNumber,omitempty"`
}

type WebsiteSpec struct {
	ReplicaNum string `json:"replicaNumber,omitempty"`
}

type WeatherSpec struct {
	ReplicaNum string `json:"replicaNumber,omitempty"`
}

type AirSpec struct {
	ReplicaNum string `json:"replicaNumber,omitempty"`
}

// JigareeStatus defines the observed state of Jigaree
type JigareeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Jigaree is the Schema for the jigarees API
type Jigaree struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JigareeSpec   `json:"spec,omitempty"`
	Status JigareeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// JigareeList contains a list of Jigaree
type JigareeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Jigaree `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Jigaree{}, &JigareeList{})
}
