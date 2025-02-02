// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProcessGroupDetectionFlagsInitParameters struct {

	// (Boolean) In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	// In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	AddNodeJsScriptName *bool `json:"addNodeJsScriptName,omitempty" tf:"add_node_js_script_name,omitempty"`

	// (Boolean) Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	// Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	AutoDetectCassandraClusters *bool `json:"autoDetectCassandraClusters,omitempty" tf:"auto_detect_cassandra_clusters,omitempty"`

	// (Boolean) Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	// Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	AutoDetectSpringBoot *bool `json:"autoDetectSpringBoot,omitempty" tf:"auto_detect_spring_boot,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectTibcoContainerEditionEngines *bool `json:"autoDetectTibcoContainerEditionEngines,omitempty" tf:"auto_detect_tibco_container_edition_engines,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectTibcoEngines *bool `json:"autoDetectTibcoEngines,omitempty" tf:"auto_detect_tibco_engines,omitempty"`

	// (Boolean) Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	// Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	AutoDetectWebMethodsIntegrationServer *bool `json:"autoDetectWebMethodsIntegrationServer,omitempty" tf:"auto_detect_web_methods_integration_server,omitempty"`

	// (Boolean) Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	// Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	AutoDetectWebSphereLibertyApplication *bool `json:"autoDetectWebSphereLibertyApplication,omitempty" tf:"auto_detect_web_sphere_liberty_application,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	// Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	GroupIbmmqbyInstanceName *bool `json:"groupIbmmqbyInstanceName,omitempty" tf:"group_ibmmqby_instance_name,omitempty"`

	// D[Server:] is not set.
	// Enabling this flag will detect the JBoss server name from the system property jboss.server.name=<server-name>, only if -D[Server:<server-name>] is not set.
	IdentifyJbossServerBySystemProperty *bool `json:"identifyJbossServerBySystemProperty,omitempty" tf:"identify_jboss_server_by_system_property,omitempty"`

	// existing processes as new processes.
	// To determine the unique identity of each detected process, and to generate a unique name for each detected process, Dynatrace evaluates the name of the directory that each process binary is contained within. For application containers like Tomcat and JBoss, Dynatrace evaluates important directories like CATALINA_HOME and JBOSS_HOME for this information. In some automated deployment scenarios such directory names are updated automatically with new version numbers, build numbers, dates, or GUIDs. Enable this setting to ensure that automated directory name changes don't result in Dynatrace registering pre-existing processes as new processes.
	IgnoreUniqueIdentifiers *bool `json:"ignoreUniqueIdentifiers,omitempty" tf:"ignore_unique_identifiers,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	// Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	ShortLivedProcessesMonitoring *bool `json:"shortLivedProcessesMonitoring,omitempty" tf:"short_lived_processes_monitoring,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	// Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	SplitOracleDatabasePg *bool `json:"splitOracleDatabasePg,omitempty" tf:"split_oracle_database_pg,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	// Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	SplitOracleListenerPg *bool `json:"splitOracleListenerPg,omitempty" tf:"split_oracle_listener_pg,omitempty"`

	// (Boolean) By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	// By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	UseCatalinaBase *bool `json:"useCatalinaBase,omitempty" tf:"use_catalina_base,omitempty"`

	// group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	// By default, Dynatrace uses image names as identifiers for individual process groups, with one process-group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	UseDockerContainerName *bool `json:"useDockerContainerName,omitempty" tf:"use_docker_container_name,omitempty"`
}

type ProcessGroupDetectionFlagsObservation struct {

	// (Boolean) In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	// In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	AddNodeJsScriptName *bool `json:"addNodeJsScriptName,omitempty" tf:"add_node_js_script_name,omitempty"`

	// (Boolean) Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	// Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	AutoDetectCassandraClusters *bool `json:"autoDetectCassandraClusters,omitempty" tf:"auto_detect_cassandra_clusters,omitempty"`

	// (Boolean) Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	// Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	AutoDetectSpringBoot *bool `json:"autoDetectSpringBoot,omitempty" tf:"auto_detect_spring_boot,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectTibcoContainerEditionEngines *bool `json:"autoDetectTibcoContainerEditionEngines,omitempty" tf:"auto_detect_tibco_container_edition_engines,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	AutoDetectTibcoEngines *bool `json:"autoDetectTibcoEngines,omitempty" tf:"auto_detect_tibco_engines,omitempty"`

	// (Boolean) Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	// Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	AutoDetectWebMethodsIntegrationServer *bool `json:"autoDetectWebMethodsIntegrationServer,omitempty" tf:"auto_detect_web_methods_integration_server,omitempty"`

	// (Boolean) Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	// Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	AutoDetectWebSphereLibertyApplication *bool `json:"autoDetectWebSphereLibertyApplication,omitempty" tf:"auto_detect_web_sphere_liberty_application,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	// Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	GroupIbmmqbyInstanceName *bool `json:"groupIbmmqbyInstanceName,omitempty" tf:"group_ibmmqby_instance_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// D[Server:] is not set.
	// Enabling this flag will detect the JBoss server name from the system property jboss.server.name=<server-name>, only if -D[Server:<server-name>] is not set.
	IdentifyJbossServerBySystemProperty *bool `json:"identifyJbossServerBySystemProperty,omitempty" tf:"identify_jboss_server_by_system_property,omitempty"`

	// existing processes as new processes.
	// To determine the unique identity of each detected process, and to generate a unique name for each detected process, Dynatrace evaluates the name of the directory that each process binary is contained within. For application containers like Tomcat and JBoss, Dynatrace evaluates important directories like CATALINA_HOME and JBOSS_HOME for this information. In some automated deployment scenarios such directory names are updated automatically with new version numbers, build numbers, dates, or GUIDs. Enable this setting to ensure that automated directory name changes don't result in Dynatrace registering pre-existing processes as new processes.
	IgnoreUniqueIdentifiers *bool `json:"ignoreUniqueIdentifiers,omitempty" tf:"ignore_unique_identifiers,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	// Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	ShortLivedProcessesMonitoring *bool `json:"shortLivedProcessesMonitoring,omitempty" tf:"short_lived_processes_monitoring,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	// Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	SplitOracleDatabasePg *bool `json:"splitOracleDatabasePg,omitempty" tf:"split_oracle_database_pg,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	// Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	SplitOracleListenerPg *bool `json:"splitOracleListenerPg,omitempty" tf:"split_oracle_listener_pg,omitempty"`

	// (Boolean) By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	// By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	UseCatalinaBase *bool `json:"useCatalinaBase,omitempty" tf:"use_catalina_base,omitempty"`

	// group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	// By default, Dynatrace uses image names as identifiers for individual process groups, with one process-group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	UseDockerContainerName *bool `json:"useDockerContainerName,omitempty" tf:"use_docker_container_name,omitempty"`
}

type ProcessGroupDetectionFlagsParameters struct {

	// (Boolean) In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	// In older versions, Node.js applications were distinguished based on their directory name, omitting the script name. Changing this setting may change the general handling of Node.js process groups. Leave unchanged if in doubt.
	// +kubebuilder:validation:Optional
	AddNodeJsScriptName *bool `json:"addNodeJsScriptName,omitempty" tf:"add_node_js_script_name,omitempty"`

	// (Boolean) Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	// Enabling this flag will detect separate Cassandra process groups based on the configured Cassandra cluster name.
	// +kubebuilder:validation:Optional
	AutoDetectCassandraClusters *bool `json:"autoDetectCassandraClusters,omitempty" tf:"auto_detect_cassandra_clusters,omitempty"`

	// (Boolean) Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	// Enabling this flag will detect Spring Boot process groups based on command line and applications' configuration files.
	// +kubebuilder:validation:Optional
	AutoDetectSpringBoot *bool `json:"autoDetectSpringBoot,omitempty" tf:"auto_detect_spring_boot,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// +kubebuilder:validation:Optional
	AutoDetectTibcoContainerEditionEngines *bool `json:"autoDetectTibcoContainerEditionEngines,omitempty" tf:"auto_detect_tibco_container_edition_engines,omitempty"`

	// (Boolean) Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// Enabling this flag will detect separate TIBCO BusinessWorks process groups per engine property file.
	// +kubebuilder:validation:Optional
	AutoDetectTibcoEngines *bool `json:"autoDetectTibcoEngines,omitempty" tf:"auto_detect_tibco_engines,omitempty"`

	// (Boolean) Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	// Enabling this flag will detect webMethods Integration Server including specific properties like install root and product name.
	// +kubebuilder:validation:Optional
	AutoDetectWebMethodsIntegrationServer *bool `json:"autoDetectWebMethodsIntegrationServer,omitempty" tf:"auto_detect_web_methods_integration_server,omitempty"`

	// (Boolean) Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	// Enabling this flag will detect separate WebSphere Liberty process groups based on java command line.
	// +kubebuilder:validation:Optional
	AutoDetectWebSphereLibertyApplication *bool `json:"autoDetectWebSphereLibertyApplication,omitempty" tf:"auto_detect_web_sphere_liberty_application,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	// Enable to group and separately analyze the processes of each IBM MQ Queue manager instance. Each process group receives a unique name based on the queue manager instance name.
	// +kubebuilder:validation:Optional
	GroupIbmmqbyInstanceName *bool `json:"groupIbmmqbyInstanceName,omitempty" tf:"group_ibmmqby_instance_name,omitempty"`

	// D[Server:] is not set.
	// Enabling this flag will detect the JBoss server name from the system property jboss.server.name=<server-name>, only if -D[Server:<server-name>] is not set.
	// +kubebuilder:validation:Optional
	IdentifyJbossServerBySystemProperty *bool `json:"identifyJbossServerBySystemProperty,omitempty" tf:"identify_jboss_server_by_system_property,omitempty"`

	// existing processes as new processes.
	// To determine the unique identity of each detected process, and to generate a unique name for each detected process, Dynatrace evaluates the name of the directory that each process binary is contained within. For application containers like Tomcat and JBoss, Dynatrace evaluates important directories like CATALINA_HOME and JBOSS_HOME for this information. In some automated deployment scenarios such directory names are updated automatically with new version numbers, build numbers, dates, or GUIDs. Enable this setting to ensure that automated directory name changes don't result in Dynatrace registering pre-existing processes as new processes.
	// +kubebuilder:validation:Optional
	IgnoreUniqueIdentifiers *bool `json:"ignoreUniqueIdentifiers,omitempty" tf:"ignore_unique_identifiers,omitempty"`

	// (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// (Boolean) Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	// Enable to monitor CPU and memory usage of short lived processes, otherwise being lost by traditional monitoring. Disabling this flag blocks passing data to cluster only, it does not stop data collection and has no effect on performance.
	// +kubebuilder:validation:Optional
	ShortLivedProcessesMonitoring *bool `json:"shortLivedProcessesMonitoring,omitempty" tf:"short_lived_processes_monitoring,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	// Enable to group and separately analyze the processes of each Oracle DB. Each process group receives a unique name based on the Oracle DB SID.
	// +kubebuilder:validation:Optional
	SplitOracleDatabasePg *bool `json:"splitOracleDatabasePg,omitempty" tf:"split_oracle_database_pg,omitempty"`

	// (Boolean) Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	// Enable to group and separately analyze the processes of each Oracle Listener. Each process group receives a unique name based on the Oracle Listener name.
	// +kubebuilder:validation:Optional
	SplitOracleListenerPg *bool `json:"splitOracleListenerPg,omitempty" tf:"split_oracle_listener_pg,omitempty"`

	// (Boolean) By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	// By default, Tomcat clusters are identified and named based on the CATALINA_HOME directory name. This setting results in the use of the CATALINA_BASE directory name to identify multiple Tomcat nodes within each Tomcat cluster. If this setting is not enabled, each CATALINA_HOME+CATALINA_BASE combination will be considered a separate Tomcat cluster. In other words, Tomcat clusters can't have multiple nodes on a single host.
	// +kubebuilder:validation:Optional
	UseCatalinaBase *bool `json:"useCatalinaBase,omitempty" tf:"use_catalina_base,omitempty"`

	// group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	// By default, Dynatrace uses image names as identifiers for individual process groups, with one process-group instance per host. Normally Docker container names can't serve as stable identifiers of process group instances because they are variable and auto-generated. You can however manually assign proper container names to their Docker instances. Such manually-assigned container names can serve as reliable process-group instance identifiers. This flag instructs Dynatrace to use Docker-provided names to distinguish between multiple instances of the same image. If this flag is not applied and you run multiple containers of the same image on the same host, the resulting processes will be consolidated into a single process view. Use this flag with caution!
	// +kubebuilder:validation:Optional
	UseDockerContainerName *bool `json:"useDockerContainerName,omitempty" tf:"use_docker_container_name,omitempty"`
}

// ProcessGroupDetectionFlagsSpec defines the desired state of ProcessGroupDetectionFlags
type ProcessGroupDetectionFlagsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProcessGroupDetectionFlagsParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProcessGroupDetectionFlagsInitParameters `json:"initProvider,omitempty"`
}

// ProcessGroupDetectionFlagsStatus defines the observed state of ProcessGroupDetectionFlags.
type ProcessGroupDetectionFlagsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProcessGroupDetectionFlagsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProcessGroupDetectionFlags is the Schema for the ProcessGroupDetectionFlagss API. The resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type ProcessGroupDetectionFlags struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addNodeJsScriptName) || (has(self.initProvider) && has(self.initProvider.addNodeJsScriptName))",message="spec.forProvider.addNodeJsScriptName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectCassandraClusters) || (has(self.initProvider) && has(self.initProvider.autoDetectCassandraClusters))",message="spec.forProvider.autoDetectCassandraClusters is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectSpringBoot) || (has(self.initProvider) && has(self.initProvider.autoDetectSpringBoot))",message="spec.forProvider.autoDetectSpringBoot is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectTibcoContainerEditionEngines) || (has(self.initProvider) && has(self.initProvider.autoDetectTibcoContainerEditionEngines))",message="spec.forProvider.autoDetectTibcoContainerEditionEngines is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectTibcoEngines) || (has(self.initProvider) && has(self.initProvider.autoDetectTibcoEngines))",message="spec.forProvider.autoDetectTibcoEngines is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectWebMethodsIntegrationServer) || (has(self.initProvider) && has(self.initProvider.autoDetectWebMethodsIntegrationServer))",message="spec.forProvider.autoDetectWebMethodsIntegrationServer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoDetectWebSphereLibertyApplication) || (has(self.initProvider) && has(self.initProvider.autoDetectWebSphereLibertyApplication))",message="spec.forProvider.autoDetectWebSphereLibertyApplication is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupIbmmqbyInstanceName) || (has(self.initProvider) && has(self.initProvider.groupIbmmqbyInstanceName))",message="spec.forProvider.groupIbmmqbyInstanceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identifyJbossServerBySystemProperty) || (has(self.initProvider) && has(self.initProvider.identifyJbossServerBySystemProperty))",message="spec.forProvider.identifyJbossServerBySystemProperty is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ignoreUniqueIdentifiers) || (has(self.initProvider) && has(self.initProvider.ignoreUniqueIdentifiers))",message="spec.forProvider.ignoreUniqueIdentifiers is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.shortLivedProcessesMonitoring) || (has(self.initProvider) && has(self.initProvider.shortLivedProcessesMonitoring))",message="spec.forProvider.shortLivedProcessesMonitoring is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.splitOracleDatabasePg) || (has(self.initProvider) && has(self.initProvider.splitOracleDatabasePg))",message="spec.forProvider.splitOracleDatabasePg is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.splitOracleListenerPg) || (has(self.initProvider) && has(self.initProvider.splitOracleListenerPg))",message="spec.forProvider.splitOracleListenerPg is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.useCatalinaBase) || (has(self.initProvider) && has(self.initProvider.useCatalinaBase))",message="spec.forProvider.useCatalinaBase is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.useDockerContainerName) || (has(self.initProvider) && has(self.initProvider.useDockerContainerName))",message="spec.forProvider.useDockerContainerName is a required parameter"
	Spec   ProcessGroupDetectionFlagsSpec   `json:"spec"`
	Status ProcessGroupDetectionFlagsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProcessGroupDetectionFlagsList contains a list of ProcessGroupDetectionFlagss
type ProcessGroupDetectionFlagsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProcessGroupDetectionFlags `json:"items"`
}

// Repository type metadata.
var (
	ProcessGroupDetectionFlags_Kind             = "ProcessGroupDetectionFlags"
	ProcessGroupDetectionFlags_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProcessGroupDetectionFlags_Kind}.String()
	ProcessGroupDetectionFlags_KindAPIVersion   = ProcessGroupDetectionFlags_Kind + "." + CRDGroupVersion.String()
	ProcessGroupDetectionFlags_GroupVersionKind = CRDGroupVersion.WithKind(ProcessGroupDetectionFlags_Kind)
)

func init() {
	SchemeBuilder.Register(&ProcessGroupDetectionFlags{}, &ProcessGroupDetectionFlagsList{})
}
