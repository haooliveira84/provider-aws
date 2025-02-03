//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEnvironmentObservation) DeepCopyInto(out *CustomEnvironmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEnvironmentObservation.
func (in *CustomEnvironmentObservation) DeepCopy() *CustomEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(CustomEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEnvironmentParameters) DeepCopyInto(out *CustomEnvironmentParameters) {
	*out = *in
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyRef != nil {
		in, out := &in.KMSKeyRef, &out.KMSKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeySelector != nil {
		in, out := &in.KMSKeySelector, &out.KMSKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceBucketARN != nil {
		in, out := &in.SourceBucketARN, &out.SourceBucketARN
		*out = new(string)
		**out = **in
	}
	if in.SourceBucketARNRef != nil {
		in, out := &in.SourceBucketARNRef, &out.SourceBucketARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceBucketARNSelector != nil {
		in, out := &in.SourceBucketARNSelector, &out.SourceBucketARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExecutionRoleARN != nil {
		in, out := &in.ExecutionRoleARN, &out.ExecutionRoleARN
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleARNRef != nil {
		in, out := &in.ExecutionRoleARNRef, &out.ExecutionRoleARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ExecutionRoleARNSelector != nil {
		in, out := &in.ExecutionRoleARNSelector, &out.ExecutionRoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.NetworkConfiguration.DeepCopyInto(&out.NetworkConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEnvironmentParameters.
func (in *CustomEnvironmentParameters) DeepCopy() *CustomEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(CustomEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomNetworkConfiguration) DeepCopyInto(out *CustomNetworkConfiguration) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomNetworkConfiguration.
func (in *CustomNetworkConfiguration) DeepCopy() *CustomNetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(CustomNetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension) DeepCopyInto(out *Dimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension.
func (in *Dimension) DeepCopy() *Dimension {
	if in == nil {
		return nil
	}
	out := new(Dimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentObservation) DeepCopyInto(out *EnvironmentObservation) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(LastUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	out.CustomEnvironmentObservation = in.CustomEnvironmentObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentObservation.
func (in *EnvironmentObservation) DeepCopy() *EnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentParameters) DeepCopyInto(out *EnvironmentParameters) {
	*out = *in
	if in.AirflowConfigurationOptions != nil {
		in, out := &in.AirflowConfigurationOptions, &out.AirflowConfigurationOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AirflowVersion != nil {
		in, out := &in.AirflowVersion, &out.AirflowVersion
		*out = new(string)
		**out = **in
	}
	if in.DagS3Path != nil {
		in, out := &in.DagS3Path, &out.DagS3Path
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentClass != nil {
		in, out := &in.EnvironmentClass, &out.EnvironmentClass
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = new(LoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxWorkers != nil {
		in, out := &in.MaxWorkers, &out.MaxWorkers
		*out = new(int64)
		**out = **in
	}
	if in.MinWorkers != nil {
		in, out := &in.MinWorkers, &out.MinWorkers
		*out = new(int64)
		**out = **in
	}
	if in.PluginsS3ObjectVersion != nil {
		in, out := &in.PluginsS3ObjectVersion, &out.PluginsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.PluginsS3Path != nil {
		in, out := &in.PluginsS3Path, &out.PluginsS3Path
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3ObjectVersion != nil {
		in, out := &in.RequirementsS3ObjectVersion, &out.RequirementsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3Path != nil {
		in, out := &in.RequirementsS3Path, &out.RequirementsS3Path
		*out = new(string)
		**out = **in
	}
	if in.Schedulers != nil {
		in, out := &in.Schedulers, &out.Schedulers
		*out = new(int64)
		**out = **in
	}
	if in.StartupScriptS3ObjectVersion != nil {
		in, out := &in.StartupScriptS3ObjectVersion, &out.StartupScriptS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.StartupScriptS3Path != nil {
		in, out := &in.StartupScriptS3Path, &out.StartupScriptS3Path
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WebserverAccessMode != nil {
		in, out := &in.WebserverAccessMode, &out.WebserverAccessMode
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindowStart != nil {
		in, out := &in.WeeklyMaintenanceWindowStart, &out.WeeklyMaintenanceWindowStart
		*out = new(string)
		**out = **in
	}
	in.CustomEnvironmentParameters.DeepCopyInto(&out.CustomEnvironmentParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentParameters.
func (in *EnvironmentParameters) DeepCopy() *EnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment_SDK) DeepCopyInto(out *Environment_SDK) {
	*out = *in
	if in.AirflowConfigurationOptions != nil {
		in, out := &in.AirflowConfigurationOptions, &out.AirflowConfigurationOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AirflowVersion != nil {
		in, out := &in.AirflowVersion, &out.AirflowVersion
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.DagS3Path != nil {
		in, out := &in.DagS3Path, &out.DagS3Path
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentClass != nil {
		in, out := &in.EnvironmentClass, &out.EnvironmentClass
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleARN != nil {
		in, out := &in.ExecutionRoleARN, &out.ExecutionRoleARN
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(LastUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = new(LoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxWorkers != nil {
		in, out := &in.MaxWorkers, &out.MaxWorkers
		*out = new(int64)
		**out = **in
	}
	if in.MinWorkers != nil {
		in, out := &in.MinWorkers, &out.MinWorkers
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PluginsS3ObjectVersion != nil {
		in, out := &in.PluginsS3ObjectVersion, &out.PluginsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.PluginsS3Path != nil {
		in, out := &in.PluginsS3Path, &out.PluginsS3Path
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3ObjectVersion != nil {
		in, out := &in.RequirementsS3ObjectVersion, &out.RequirementsS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.RequirementsS3Path != nil {
		in, out := &in.RequirementsS3Path, &out.RequirementsS3Path
		*out = new(string)
		**out = **in
	}
	if in.Schedulers != nil {
		in, out := &in.Schedulers, &out.Schedulers
		*out = new(int64)
		**out = **in
	}
	if in.ServiceRoleARN != nil {
		in, out := &in.ServiceRoleARN, &out.ServiceRoleARN
		*out = new(string)
		**out = **in
	}
	if in.SourceBucketARN != nil {
		in, out := &in.SourceBucketARN, &out.SourceBucketARN
		*out = new(string)
		**out = **in
	}
	if in.StartupScriptS3ObjectVersion != nil {
		in, out := &in.StartupScriptS3ObjectVersion, &out.StartupScriptS3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.StartupScriptS3Path != nil {
		in, out := &in.StartupScriptS3Path, &out.StartupScriptS3Path
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WebserverAccessMode != nil {
		in, out := &in.WebserverAccessMode, &out.WebserverAccessMode
		*out = new(string)
		**out = **in
	}
	if in.WebserverURL != nil {
		in, out := &in.WebserverURL, &out.WebserverURL
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindowStart != nil {
		in, out := &in.WeeklyMaintenanceWindowStart, &out.WeeklyMaintenanceWindowStart
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment_SDK.
func (in *Environment_SDK) DeepCopy() *Environment_SDK {
	if in == nil {
		return nil
	}
	out := new(Environment_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastUpdate) DeepCopyInto(out *LastUpdate) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(UpdateError)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastUpdate.
func (in *LastUpdate) DeepCopy() *LastUpdate {
	if in == nil {
		return nil
	}
	out := new(LastUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfiguration) DeepCopyInto(out *LoggingConfiguration) {
	*out = *in
	if in.DagProcessingLogs != nil {
		in, out := &in.DagProcessingLogs, &out.DagProcessingLogs
		*out = new(ModuleLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SchedulerLogs != nil {
		in, out := &in.SchedulerLogs, &out.SchedulerLogs
		*out = new(ModuleLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskLogs != nil {
		in, out := &in.TaskLogs, &out.TaskLogs
		*out = new(ModuleLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.WebserverLogs != nil {
		in, out := &in.WebserverLogs, &out.WebserverLogs
		*out = new(ModuleLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerLogs != nil {
		in, out := &in.WorkerLogs, &out.WorkerLogs
		*out = new(ModuleLoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfiguration.
func (in *LoggingConfiguration) DeepCopy() *LoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(LoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigurationInput) DeepCopyInto(out *LoggingConfigurationInput) {
	*out = *in
	if in.DagProcessingLogs != nil {
		in, out := &in.DagProcessingLogs, &out.DagProcessingLogs
		*out = new(ModuleLoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
	if in.SchedulerLogs != nil {
		in, out := &in.SchedulerLogs, &out.SchedulerLogs
		*out = new(ModuleLoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
	if in.TaskLogs != nil {
		in, out := &in.TaskLogs, &out.TaskLogs
		*out = new(ModuleLoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
	if in.WebserverLogs != nil {
		in, out := &in.WebserverLogs, &out.WebserverLogs
		*out = new(ModuleLoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerLogs != nil {
		in, out := &in.WorkerLogs, &out.WorkerLogs
		*out = new(ModuleLoggingConfigurationInput)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationInput.
func (in *LoggingConfigurationInput) DeepCopy() *LoggingConfigurationInput {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricDatum) DeepCopyInto(out *MetricDatum) {
	*out = *in
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricDatum.
func (in *MetricDatum) DeepCopy() *MetricDatum {
	if in == nil {
		return nil
	}
	out := new(MetricDatum)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleLoggingConfiguration) DeepCopyInto(out *ModuleLoggingConfiguration) {
	*out = *in
	if in.CloudWatchLogGroupARN != nil {
		in, out := &in.CloudWatchLogGroupARN, &out.CloudWatchLogGroupARN
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleLoggingConfiguration.
func (in *ModuleLoggingConfiguration) DeepCopy() *ModuleLoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ModuleLoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleLoggingConfigurationInput) DeepCopyInto(out *ModuleLoggingConfigurationInput) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleLoggingConfigurationInput.
func (in *ModuleLoggingConfigurationInput) DeepCopy() *ModuleLoggingConfigurationInput {
	if in == nil {
		return nil
	}
	out := new(ModuleLoggingConfigurationInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateError) DeepCopyInto(out *UpdateError) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateError.
func (in *UpdateError) DeepCopy() *UpdateError {
	if in == nil {
		return nil
	}
	out := new(UpdateError)
	in.DeepCopyInto(out)
	return out
}
