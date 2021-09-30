/*
Copyright 2021.

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

// ScrambledSpec defines the desired state of Scrambled
type ScrambledSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// CleanWord The word to scramble.
	CleanWord string `json:"clean-word,omitempty"`
}

// ScrambledStatus defines the observed state of Scrambled
type ScrambledStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ScrambledWord string `json:"scrambled-word,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Scrambled is the Schema for the scrambleds API
type Scrambled struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScrambledSpec   `json:"spec,omitempty"`
	Status ScrambledStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScrambledList contains a list of Scrambled
type ScrambledList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scrambled `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Scrambled{}, &ScrambledList{})
}
