/*
Copyright 2023.

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
	"net/url"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QuayMirrorLocationSpec defines the desired state of QuayMirrorLocation
type QuayMirrorLocationSpec struct {
	// QuayHostname is the hostname of the Quay registry.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Quay hostname"
	// +kubebuilder:validation:Required
	QuayHostname string `json:"quayHostname"`

	// QuayOrganization is the organization within the Quay registry to mirror.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Quay organization"
	// +kubebuilder:validation:Required
	QuayOrganization string `json:"quayOrganization"`

	// QuayCredentialsSecret refers to the Secret containing credentials to communicate with the Quay registry.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Quay Credentials secret",xDescriptors={"urn:alm:descriptor:io.kubernetes:Secret"}
	// +kubebuilder:validation:Required
	QuayCredentialsSecret *SecretRef `json:"quayCredentialsSecret"`

	// InsecureRegistry refers to whether to skip TLS verification to the Quay registry.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Insecure Registry"
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	InsecureRegistry bool `json:"insecureRegistry,omitempty"`
}

// QuayMirrorLocationStatus defines the observed state of QuayMirrorLocation
type QuayMirrorLocationStatus struct {

	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Conditions",xDescriptors={"urn:alm:descriptor:io.kubernetes.conditions"}
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// RepositoryCount is the number of repositories in the Quay organization.
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Repository count"
	// +kubebuilder:validation:Optional
	RepositoryCount int `json:"repositoryCount,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster
//+operator-sdk:csv:customresourcedefinitions:displayName="Quay Mirror Location"

// QuayMirrorLocation is the Schema for the quaymirrorlocations API
type QuayMirrorLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuayMirrorLocationSpec   `json:"spec,omitempty"`
	Status QuayMirrorLocationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// QuayMirrorLocationList contains a list of QuayMirrorLocation
type QuayMirrorLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuayMirrorLocation `json:"items"`
}

// SecretRef represents a reference to an item within a Secret
type SecretRef struct {

	// Name represents the name of the secret
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Name of the secret",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Namespace represents the namespace containing the secret
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace containing the secret",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`

	// Key represents the specific key to reference from the secret
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Key within the secret",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	// +kubebuilder:validation:Optional
	Key string `json:"key,omitempty"`
}

func (q *QuayMirrorLocation) GetConditions() []metav1.Condition {
	return q.Status.Conditions
}

func (q *QuayMirrorLocation) SetConditions(conditions []metav1.Condition) {
	q.Status.Conditions = conditions
}

func init() {
	SchemeBuilder.Register(&QuayMirrorLocation{}, &QuayMirrorLocationList{})
}

func (qi *QuayMirrorLocation) GetRegistryHostname() (string, error) {
	quayURL, err := url.Parse(qi.Spec.QuayHostname)

	if err != nil {
		return "", err
	}

	return quayURL.Host, nil
}
