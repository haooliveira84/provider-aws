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

// EnvironmentParameters defines the desired state of Environment
type EnvironmentParameters struct {
	// Region is which region the Environment will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A list of key-value pairs containing the Apache Airflow configuration options
	// you want to attach to your environment. For more information, see Apache
	// Airflow configuration options (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html).
	AirflowConfigurationOptions map[string]*string `json:"airflowConfigurationOptions,omitempty"`
	// The Apache Airflow version for your environment. If no value is specified,
	// it defaults to the latest version. For more information, see Apache Airflow
	// versions on Amazon Managed Workflows for Apache Airflow (MWAA) (https://docs.aws.amazon.com/mwaa/latest/userguide/airflow-versions.html).
	//
	// Valid values: 1.10.12, 2.0.2, 2.2.2, 2.4.3, 2.5.1, 2.6.3, 2.7.2.
	AirflowVersion *string `json:"airflowVersion,omitempty"`
	// The relative path to the DAGs folder on your Amazon S3 bucket. For example,
	// dags. For more information, see Adding or updating DAGs (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html).
	// +kubebuilder:validation:Required
	DagS3Path *string `json:"dagS3Path"`
	// The environment class type. Valid values: mw1.small, mw1.medium, mw1.large.
	// For more information, see Amazon MWAA environment class (https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html).
	EnvironmentClass *string `json:"environmentClass,omitempty"`
	// Defines the Apache Airflow logs to send to CloudWatch Logs.
	LoggingConfiguration *LoggingConfigurationInput `json:"loggingConfiguration,omitempty"`
	// The maximum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify
	// in the MaxWorkers field. For example, 20. When there are no more tasks running,
	// and no more in the queue, MWAA disposes of the extra workers leaving the
	// one worker that is included with your environment, or the number you specify
	// in MinWorkers.
	MaxWorkers *int64 `json:"maxWorkers,omitempty"`
	// The minimum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify
	// in the MaxWorkers field. When there are no more tasks running, and no more
	// in the queue, MWAA disposes of the extra workers leaving the worker count
	// you specify in the MinWorkers field. For example, 2.
	MinWorkers *int64 `json:"minWorkers,omitempty"`
	// The version of the plugins.zip file on your Amazon S3 bucket. You must specify
	// a version each time a plugins.zip file is updated. For more information,
	// see How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html).
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion,omitempty"`
	// The relative path to the plugins.zip file on your Amazon S3 bucket. For example,
	// plugins.zip. If specified, then the plugins.zip version is required. For
	// more information, see Installing custom plugins (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html).
	PluginsS3Path *string `json:"pluginsS3Path,omitempty"`
	// The version of the requirements.txt file on your Amazon S3 bucket. You must
	// specify a version each time a requirements.txt file is updated. For more
	// information, see How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html).
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion,omitempty"`
	// The relative path to the requirements.txt file on your Amazon S3 bucket.
	// For example, requirements.txt. If specified, then a version is required.
	// For more information, see Installing Python dependencies (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html).
	RequirementsS3Path *string `json:"requirementsS3Path,omitempty"`
	// The number of Apache Airflow schedulers to run in your environment. Valid
	// values:
	//
	//    * v2 - Accepts between 2 to 5. Defaults to 2.
	//
	//    * v1 - Accepts 1.
	Schedulers *int64 `json:"schedulers,omitempty"`
	// The version of the startup shell script in your Amazon S3 bucket. You must
	// specify the version ID (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// that Amazon S3 assigns to the file every time you update the script.
	//
	// Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are
	// no more than 1,024 bytes long. The following is an example:
	//
	// 3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo
	//
	// For more information, see Using a startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html).
	StartupScriptS3ObjectVersion *string `json:"startupScriptS3ObjectVersion,omitempty"`
	// The relative path to the startup shell script in your Amazon S3 bucket. For
	// example, s3://mwaa-environment/startup.sh.
	//
	// Amazon MWAA runs the script as your environment starts, and before running
	// the Apache Airflow process. You can use this script to install dependencies,
	// modify Apache Airflow configuration options, and set environment variables.
	// For more information, see Using a startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html).
	StartupScriptS3Path *string `json:"startupScriptS3Path,omitempty"`
	// The key-value tag pairs you want to associate to your environment. For example,
	// "Environment": "Staging". For more information, see Tagging Amazon Web Services
	// resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags map[string]*string `json:"tags,omitempty"`
	// The Apache Airflow Web server access mode. For more information, see Apache
	// Airflow access modes (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html).
	WebserverAccessMode *string `json:"webserverAccessMode,omitempty"`
	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour
	// standard time to start weekly maintenance updates of your environment in
	// the following format: DAY:HH:MM. For example: TUE:03:30. You can specify
	// a start time in 30 minute increments only.
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart,omitempty"`
	CustomEnvironmentParameters  `json:",inline"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EnvironmentParameters `json:"forProvider"`
}

// EnvironmentObservation defines the observed state of Environment
type EnvironmentObservation struct {
	// The Amazon Resource Name (ARN) returned in the response for the environment.
	ARN *string `json:"arn,omitempty"`
	// The status of the last update on the environment.
	LastUpdate *LastUpdate `json:"lastUpdate,omitempty"`
	// The status of the Amazon MWAA environment. Valid values:
	//
	//    * CREATING - Indicates the request to create the environment is in progress.
	//
	//    * CREATING_SNAPSHOT - Indicates the request to update environment details,
	//    or upgrade the environment version, is in progress and Amazon MWAA is
	//    creating a storage volume snapshot of the Amazon RDS database cluster
	//    associated with the environment. A database snapshot is a backup created
	//    at a specific point in time. Amazon MWAA uses snapshots to recover environment
	//    metadata if the process to update or upgrade an environment fails.
	//
	//    * CREATE_FAILED - Indicates the request to create the environment failed,
	//    and the environment could not be created.
	//
	//    * AVAILABLE - Indicates the request was successful and the environment
	//    is ready to use.
	//
	//    * UPDATING - Indicates the request to update the environment is in progress.
	//
	//    * ROLLING_BACK - Indicates the request to update environment details,
	//    or upgrade the environment version, failed and Amazon MWAA is restoring
	//    the environment using the latest storage volume snapshot.
	//
	//    * DELETING - Indicates the request to delete the environment is in progress.
	//
	//    * DELETED - Indicates the request to delete the environment is complete,
	//    and the environment has been deleted.
	//
	//    * UNAVAILABLE - Indicates the request failed, but the environment was
	//    unable to rollback and is not in a stable state.
	//
	//    * UPDATE_FAILED - Indicates the request to update the environment failed,
	//    and the environment has rolled back successfully and is ready to use.
	//
	// We recommend reviewing our troubleshooting guide for a list of common errors
	// and their solutions. For more information, see Amazon MWAA troubleshooting
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/troubleshooting.html).
	Status *string `json:"status,omitempty"`

	CustomEnvironmentObservation `json:",inline"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvironmentSpec   `json:"spec"`
	Status            EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentKind             = "Environment"
	EnvironmentGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentKind}.String()
	EnvironmentKindAPIVersion   = EnvironmentKind + "." + GroupVersion.String()
	EnvironmentGroupVersionKind = GroupVersion.WithKind(EnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
