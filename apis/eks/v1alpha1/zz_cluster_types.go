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

type CertificateAuthorityObservation struct {
	Data *string `json:"data,omitempty" tf:"data,omitempty"`
}

type CertificateAuthorityParameters struct {
}

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CertificateAuthority []CertificateAuthorityObservation `json:"certificateAuthority,omitempty" tf:"certificate_authority,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	EnabledClusterLogTypes []*string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// +kubebuilder:validation:Optional
	KubernetesNetworkConfig []KubernetesNetworkConfigParameters `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Required
	VpcConfig []VpcConfigParameters `json:"vpcConfig" tf:"vpc_config,omitempty"`
}

type EncryptionConfigObservation struct {
}

type EncryptionConfigParameters struct {

	// +kubebuilder:validation:Required
	Provider []ProviderParameters `json:"provider" tf:"provider,omitempty"`

	// +kubebuilder:validation:Required
	Resources []*string `json:"resources" tf:"resources,omitempty"`
}

type IdentityObservation struct {
	Oidc []OidcObservation `json:"oidc,omitempty" tf:"oidc,omitempty"`
}

type IdentityParameters struct {
}

type KubernetesNetworkConfigObservation struct {
}

type KubernetesNetworkConfigParameters struct {

	// +kubebuilder:validation:Optional
	ServiceIPv4Cidr *string `json:"serviceIpv4Cidr,omitempty" tf:"service_ipv4_cidr,omitempty"`
}

type OidcObservation struct {
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type OidcParameters struct {
}

type ProviderObservation struct {
}

type ProviderParameters struct {

	// +kubebuilder:validation:Required
	KeyArn *string `json:"keyArn" tf:"key_arn,omitempty"`
}

type VpcConfigObservation struct {
	ClusterSecurityGroupID *string `json:"clusterSecurityGroupId,omitempty" tf:"cluster_security_group_id,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Optional
	EndpointPrivateAccess *bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointPublicAccess *bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access,omitempty"`

	// +kubebuilder:validation:Optional
	PublicAccessCidrs []*string `json:"publicAccessCidrs,omitempty" tf:"public_access_cidrs,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	ClusterKind             = "Cluster"
	ClusterGroupKind        = schema.GroupKind{Group: Group, Kind: ClusterKind}.String()
	ClusterKindAPIVersion   = ClusterKind + "." + GroupVersion.String()
	ClusterGroupVersionKind = GroupVersion.WithKind(ClusterKind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
