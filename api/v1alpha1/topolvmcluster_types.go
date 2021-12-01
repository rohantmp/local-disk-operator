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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TopoLVMClusterSpec defines the desired state of TopoLVMCluster
type TopoLVMClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DeviceSelector is a set of rules that should match for a device to be included in this TopoLVMCluster
	// +Optional
	DeviceClasses []DeviceClass `json:"deviceClasses,omitempty"`
}

type DeviceClass struct {
	// Name of the class, the VG and possibly the storageclass.
	Name string `json:"name,omitempty"`
	// DeviceSelector is a set of rules that should match for a device to be included in this TopoLVMCluster
	// +optional
	DeviceSelector *DeviceSelector `json:"deviceSelector,omitempty"`
	// NodeSelector chooses nodes
	NodeSelector *corev1.NodeSelector `json:"nodeSelector,omitempty"`
	// Config for this deviceClass, lvm settings are a field here
	// +optional
	Config *DeviceClassConfig `json:"config,omitempty"`
}

// DeviceSelector matches if all the rules match.
// (ruleType1 AND ruleType2 AND ... AND ruleTypeN-1 AND ruleTypeN)
// Each characteristic is represented either by a ruleType, a ruleType is either a list of OR'd rules or a single rule
type DeviceSelector struct {
	// MaxMatchesPerNode to allow a use-cases where not all matches on a node are consumed.
	// +optional
	MaxMatchesPerNode *int `json:"maxMatchesPerNode,omitempty"`
	// Devices is the list of devices that should be used for automatic detection.
	// This would be one of the types supported by the local-storage operator. Currently,
	// the supported types are: disk, part. If the list is empty only `disk` types will be selected
	// +optional
	DeviceTypes []DeviceType `json:"deviceTypes,omitempty"`
	// DeviceMechanicalProperty denotes whether Rotational or NonRotational disks should be used.
	// by default, it selects both
	// +optional
	DeviceMechanicalProperties []DeviceMechanicalProperty `json:"deviceMechanicalProperties,omitempty"`
	// MinSize is the minimum size of the device which needs to be included. Defaults to `1Gi` if empty
	// +optional
	MinSize *resource.Quantity `json:"minSize,omitempty"`
	// MaxSize is the maximum size of the device which needs to be included
	// +optional
	MaxSize *resource.Quantity `json:"maxSize,omitempty"`
	// Models is a list of device models. If not empty, the device's model as outputted by lsblk needs
	// to contain at least one of these strings.
	// +optional
	Models []string `json:"models,omitempty"`
	// Vendors is a list of device vendors. If not empty, the device's model as outputted by lsblk needs
	// to contain at least one of these strings.
	// +optional
	Vendors []string `json:"vendors,omitempty"`
}

type DeviceClassConfig struct {
	LVMConfig *LVMConfig `json:"lvmConfig,omitempty"`
}

type LVMConfig struct {
	thinProvision bool `json:"thinProvision,omitempty"`
}

// TopoLVMClusterStatus defines the observed state of TopoLVMCluster
type TopoLVMClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TopoLVMCluster is the Schema for the topolvmclusters API
type TopoLVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TopoLVMClusterSpec   `json:"spec,omitempty"`
	Status TopoLVMClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TopoLVMClusterList contains a list of TopoLVMCluster
type TopoLVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopoLVMCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TopoLVMCluster{}, &TopoLVMClusterList{})
}

// DeviceType is the types that will be supported by the LSO.
type DeviceType string

const (
	// RawDisk represents a device-type of block disk
	RawDisk DeviceType = "disk"
	// Partition represents a device-type of partition
	Partition DeviceType = "part"
	// Loop type device
	Loop DeviceType = "loop"
)

// DeviceMechanicalProperty holds the device's mechanical spec. It can be rotational or nonRotational
type DeviceMechanicalProperty string

// The mechanical properties of the devices
const (
	// Rotational refers to magnetic disks
	Rotational DeviceMechanicalProperty = "Rotational"
	// NonRotational refers to ssds
	NonRotational DeviceMechanicalProperty = "NonRotational"
)
