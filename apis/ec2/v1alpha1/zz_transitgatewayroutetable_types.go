/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayRouteTableObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DefaultAssociationRouteTable *bool `json:"defaultAssociationRouteTable,omitempty" tf:"default_association_route_table,omitempty"`

	DefaultPropagationRouteTable *bool `json:"defaultPropagationRouteTable,omitempty" tf:"default_propagation_route_table,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitGatewayRouteTableParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=TransitGateway
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteTableSpec defines the desired state of TransitGatewayRouteTable
type TransitGatewayRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteTableParameters `json:"forProvider"`
}

// TransitGatewayRouteTableStatus defines the observed state of TransitGatewayRouteTable.
type TransitGatewayRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTable is the Schema for the TransitGatewayRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTableSpec   `json:"spec"`
	Status            TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTableList contains a list of TransitGatewayRouteTables
type TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTable `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTableKind             = "TransitGatewayRouteTable"
	TransitGatewayRouteTableGroupKind        = schema.GroupKind{Group: Group, Kind: TransitGatewayRouteTableKind}.String()
	TransitGatewayRouteTableKindAPIVersion   = TransitGatewayRouteTableKind + "." + GroupVersion.String()
	TransitGatewayRouteTableGroupVersionKind = GroupVersion.WithKind(TransitGatewayRouteTableKind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTable{}, &TransitGatewayRouteTableList{})
}
