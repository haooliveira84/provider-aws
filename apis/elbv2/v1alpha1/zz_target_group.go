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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TargetGroupParameters defines the desired state of TargetGroup
type TargetGroupParameters struct {
	// Region is which region the TargetGroup will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Indicates whether health checks are enabled. If the target type is lambda,
	// health checks are disabled by default but can be enabled. If the target type
	// is instance, ip, or alb, health checks are always enabled and cannot be disabled.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty"`
	// The approximate amount of time, in seconds, between health checks of an individual
	// target. The range is 5-300. If the target group protocol is TCP, TLS, UDP,
	// TCP_UDP, HTTP or HTTPS, the default is 30 seconds. If the target group protocol
	// is GENEVE, the default is 10 seconds. If the target type is lambda, the default
	// is 35 seconds.
	HealthCheckIntervalSeconds *int64 `json:"healthCheckIntervalSeconds,omitempty"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the
	// format /package.service/method. The default is /Amazon Web Services.ALB/healthcheck.
	HealthCheckPath *string `json:"healthCheckPath,omitempty"`
	// The port the load balancer uses when performing health checks on targets.
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is
	// traffic-port, which is the port on which each target receives traffic from
	// the load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort *string `json:"healthCheckPort,omitempty"`
	// The protocol the load balancer uses when performing health checks on targets.
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers
	// and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported
	// for health checks if the protocol of the target group is HTTP or HTTPS. The
	// GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty"`
	// The amount of time, in seconds, during which no response from a target means
	// a failed health check. The range is 2–120 seconds. For target groups with
	// a protocol of HTTP, the default is 6 seconds. For target groups with a protocol
	// of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a
	// protocol of GENEVE, the default is 5 seconds. If the target type is lambda,
	// the default is 30 seconds.
	HealthCheckTimeoutSeconds *int64 `json:"healthCheckTimeoutSeconds,omitempty"`
	// The number of consecutive health check successes required before considering
	// a target healthy. The range is 2-10. If the target group protocol is TCP,
	// TCP_UDP, UDP, TLS, HTTP or HTTPS, the default is 5. For target groups with
	// a protocol of GENEVE, the default is 5. If the target type is lambda, the
	// default is 5.
	HealthyThresholdCount *int64 `json:"healthyThresholdCount,omitempty"`
	// The type of IP address used for this target group. The possible values are
	// ipv4 and ipv6. This is an optional parameter. If not specified, the IP address
	// type defaults to ipv4.
	IPAddressType *string `json:"ipAddressType,omitempty"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for
	// a successful response from a target. For target groups with a protocol of
	// TCP, TCP_UDP, UDP or TLS the range is 200-599. For target groups with a protocol
	// of HTTP or HTTPS, the range is 200-499. For target groups with a protocol
	// of GENEVE, the range is 200-399.
	Matcher *Matcher `json:"matcher,omitempty"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32
	// characters, must contain only alphanumeric characters or hyphens, and must
	// not begin or end with a hyphen.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The port on which the targets receive traffic. This port is used unless you
	// specify a port override when registering the target. If the target is a Lambda
	// function, this parameter does not apply. If the protocol is GENEVE, the supported
	// port is 6081.
	Port *int64 `json:"port,omitempty"`
	// The protocol to use for routing traffic to the targets. For Application Load
	// Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers,
	// the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers,
	// the supported protocol is GENEVE. A TCP_UDP listener must be associated with
	// a TCP_UDP target group. If the target is a Lambda function, this parameter
	// does not apply.
	Protocol *string `json:"protocol,omitempty"`
	// [HTTP/HTTPS protocol] The protocol version. Specify GRPC to send requests
	// to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2.
	// The default is HTTP1, which sends requests to targets using HTTP/1.1.
	ProtocolVersion *string `json:"protocolVersion,omitempty"`
	// The tags to assign to the target group.
	Tags []*Tag `json:"tags,omitempty"`
	// The type of target that you must specify when registering targets with this
	// target group. You can't specify targets for a target group using more than
	// one target type.
	//
	//    * instance - Register targets by instance ID. This is the default value.
	//
	//    * ip - Register targets by IP address. You can specify IP addresses from
	//    the subnets of the virtual private cloud (VPC) for the target group, the
	//    RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the
	//    RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP
	//    addresses.
	//
	//    * lambda - Register a single Lambda function as a target.
	//
	//    * alb - Register a single Application Load Balancer as a target.
	TargetType *string `json:"targetType,omitempty"`
	// The number of consecutive health check failures required before considering
	// a target unhealthy. The range is 2-10. If the target group protocol is TCP,
	// TCP_UDP, UDP, TLS, HTTP or HTTPS, the default is 2. For target groups with
	// a protocol of GENEVE, the default is 2. If the target type is lambda, the
	// default is 5.
	UnhealthyThresholdCount *int64 `json:"unhealthyThresholdCount,omitempty"`
	// The identifier of the virtual private cloud (VPC). If the target is a Lambda
	// function, this parameter does not apply. Otherwise, this parameter is required.
	VPCID                       *string `json:"vpcID,omitempty"`
	CustomTargetGroupParameters `json:",inline"`
}

// TargetGroupSpec defines the desired state of TargetGroup
type TargetGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TargetGroupParameters `json:"forProvider"`
}

// TargetGroupObservation defines the observed state of TargetGroup
type TargetGroupObservation struct {
	// Information about the target group.
	TargetGroups []*TargetGroup_SDK `json:"targetGroups,omitempty"`

	CustomTargetGroupObservation `json:",inline"`
}

// TargetGroupStatus defines the observed state of TargetGroup.
type TargetGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroup is the Schema for the TargetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetGroupSpec   `json:"spec"`
	Status            TargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroupList contains a list of TargetGroups
type TargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetGroup `json:"items"`
}

// Repository type metadata.
var (
	TargetGroupKind             = "TargetGroup"
	TargetGroupGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetGroupKind}.String()
	TargetGroupKindAPIVersion   = TargetGroupKind + "." + GroupVersion.String()
	TargetGroupGroupVersionKind = GroupVersion.WithKind(TargetGroupKind)
)

func init() {
	SchemeBuilder.Register(&TargetGroup{}, &TargetGroupList{})
}
