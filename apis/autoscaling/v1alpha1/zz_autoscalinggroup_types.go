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

type AutoscalingGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type AutoscalingGroupParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityRebalance *bool `json:"capacityRebalance,omitempty" tf:"capacity_rebalance,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultCooldown *int64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCapacity *int64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	EnabledMetrics []*string `json:"enabledMetrics,omitempty" tf:"enabled_metrics,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDeleteWarmPool *bool `json:"forceDeleteWarmPool,omitempty" tf:"force_delete_warm_pool,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckGracePeriod *int64 `json:"healthCheckGracePeriod,omitempty" tf:"health_check_grace_period,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// +kubebuilder:validation:Optional
	InitialLifecycleHook []InitialLifecycleHookParameters `json:"initialLifecycleHook,omitempty" tf:"initial_lifecycle_hook,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRefresh []InstanceRefreshParameters `json:"instanceRefresh,omitempty" tf:"instance_refresh,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchConfiguration *string `json:"launchConfiguration,omitempty" tf:"launch_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancers []*string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`

	// +kubebuilder:validation:Optional
	MaxInstanceLifetime *int64 `json:"maxInstanceLifetime,omitempty" tf:"max_instance_lifetime,omitempty"`

	// +kubebuilder:validation:Required
	MaxSize *int64 `json:"maxSize" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsGranularity *string `json:"metricsGranularity,omitempty" tf:"metrics_granularity,omitempty"`

	// +kubebuilder:validation:Optional
	MinElbCapacity *int64 `json:"minElbCapacity,omitempty" tf:"min_elb_capacity,omitempty"`

	// +kubebuilder:validation:Required
	MinSize *int64 `json:"minSize" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	MixedInstancesPolicy []MixedInstancesPolicyParameters `json:"mixedInstancesPolicy,omitempty" tf:"mixed_instances_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// +kubebuilder:validation:Optional
	ProtectFromScaleIn *bool `json:"protectFromScaleIn,omitempty" tf:"protect_from_scale_in,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArn *string `json:"serviceLinkedRoleArn,omitempty" tf:"service_linked_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	SuspendedProcesses []*string `json:"suspendedProcesses,omitempty" tf:"suspended_processes,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1.LBTargetGroup
	// +kubebuilder:validation:Optional
	TargetGroupArns []*string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupArnsRefs []v1.Reference `json:"targetGroupArnsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetGroupArnsSelector *v1.Selector `json:"targetGroupArnsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	VpcZoneIdentifier []*string `json:"vpcZoneIdentifier,omitempty" tf:"vpc_zone_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	VpcZoneIdentifierRefs []v1.Reference `json:"vpcZoneIdentifierRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcZoneIdentifierSelector *v1.Selector `json:"vpcZoneIdentifierSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WaitForCapacityTimeout *string `json:"waitForCapacityTimeout,omitempty" tf:"wait_for_capacity_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForElbCapacity *int64 `json:"waitForElbCapacity,omitempty" tf:"wait_for_elb_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	WarmPool []WarmPoolParameters `json:"warmPool,omitempty" tf:"warm_pool,omitempty"`
}

type InitialLifecycleHookObservation struct {
}

type InitialLifecycleHookParameters struct {

	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// +kubebuilder:validation:Optional
	HeartbeatTimeout *int64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// +kubebuilder:validation:Required
	LifecycleTransition *string `json:"lifecycleTransition" tf:"lifecycle_transition,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type InstanceRefreshObservation struct {
}

type InstanceRefreshParameters struct {

	// +kubebuilder:validation:Optional
	Preferences []PreferencesParameters `json:"preferences,omitempty" tf:"preferences,omitempty"`

	// +kubebuilder:validation:Required
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Triggers []*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type InstancesDistributionObservation struct {
}

type InstancesDistributionParameters struct {

	// +kubebuilder:validation:Optional
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandBaseCapacity *int64 `json:"onDemandBaseCapacity,omitempty" tf:"on_demand_base_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemandPercentageAboveBaseCapacity *int64 `json:"onDemandPercentageAboveBaseCapacity,omitempty" tf:"on_demand_percentage_above_base_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	SpotAllocationStrategy *string `json:"spotAllocationStrategy,omitempty" tf:"spot_allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	SpotInstancePools *int64 `json:"spotInstancePools,omitempty" tf:"spot_instance_pools,omitempty"`

	// +kubebuilder:validation:Optional
	SpotMaxPrice *string `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MixedInstancesPolicyLaunchTemplateObservation struct {
}

type MixedInstancesPolicyLaunchTemplateParameters struct {

	// +kubebuilder:validation:Required
	LaunchTemplateSpecification []LaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification,omitempty"`

	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`
}

type MixedInstancesPolicyObservation struct {
}

type MixedInstancesPolicyParameters struct {

	// +kubebuilder:validation:Optional
	InstancesDistribution []InstancesDistributionParameters `json:"instancesDistribution,omitempty" tf:"instances_distribution,omitempty"`

	// +kubebuilder:validation:Required
	LaunchTemplate []MixedInstancesPolicyLaunchTemplateParameters `json:"launchTemplate" tf:"launch_template,omitempty"`
}

type OverrideLaunchTemplateSpecificationObservation struct {
}

type OverrideLaunchTemplateSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OverrideObservation struct {
}

type OverrideParameters struct {

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateSpecification []OverrideLaunchTemplateSpecificationParameters `json:"launchTemplateSpecification,omitempty" tf:"launch_template_specification,omitempty"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type PreferencesObservation struct {
}

type PreferencesParameters struct {

	// +kubebuilder:validation:Optional
	InstanceWarmup *string `json:"instanceWarmup,omitempty" tf:"instance_warmup,omitempty"`

	// +kubebuilder:validation:Optional
	MinHealthyPercentage *int64 `json:"minHealthyPercentage,omitempty" tf:"min_healthy_percentage,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	PropagateAtLaunch *bool `json:"propagateAtLaunch" tf:"propagate_at_launch,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type WarmPoolObservation struct {
}

type WarmPoolParameters struct {

	// +kubebuilder:validation:Optional
	MaxGroupPreparedCapacity *int64 `json:"maxGroupPreparedCapacity,omitempty" tf:"max_group_prepared_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	PoolState *string `json:"poolState,omitempty" tf:"pool_state,omitempty"`
}

// AutoscalingGroupSpec defines the desired state of AutoscalingGroup
type AutoscalingGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscalingGroupParameters `json:"forProvider"`
}

// AutoscalingGroupStatus defines the observed state of AutoscalingGroup.
type AutoscalingGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscalingGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroup is the Schema for the AutoscalingGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingGroupSpec   `json:"spec"`
	Status            AutoscalingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroupList contains a list of AutoscalingGroups
type AutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingGroup `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingGroupKind             = "AutoscalingGroup"
	AutoscalingGroupGroupKind        = schema.GroupKind{Group: Group, Kind: AutoscalingGroupKind}.String()
	AutoscalingGroupKindAPIVersion   = AutoscalingGroupKind + "." + GroupVersion.String()
	AutoscalingGroupGroupVersionKind = GroupVersion.WithKind(AutoscalingGroupKind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingGroup{}, &AutoscalingGroupList{})
}
