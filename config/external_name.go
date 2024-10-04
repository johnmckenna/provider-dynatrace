/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Access Tokens

	"dynatrace_activegate_token": config.IdentifierFromProvider,
	"dynatrace_ag_token":         config.IdentifierFromProvider,
	"dynatrace_api_token":        config.IdentifierFromProvider,
	"dynatrace_token_settings":   config.IdentifierFromProvider,

	// Alerting

	"dynatrace_alerting":            config.IdentifierFromProvider,
	"dynatrace_connectivity_alerts": config.IdentifierFromProvider,
	"dynatrace_maintenance":         config.IdentifierFromProvider,

	// Anomaly Detection

	"dynatrace_aws_anomalies":               config.IdentifierFromProvider,
	"dynatrace_custom_app_anomalies":        config.IdentifierFromProvider,
	"dynatrace_custom_app_crash_rate":       config.IdentifierFromProvider,
	"dynatrace_database_anomalies_v2":       config.IdentifierFromProvider,
	"dynatrace_davis_anomaly_detectors":     config.IdentifierFromProvider,
	"dynatrace_disk_anomalies_v2":           config.IdentifierFromProvider,
	"dynatrace_disk_anomaly_rules":          config.IdentifierFromProvider,
	"dynatrace_disk_edge_anomaly_detectors": config.IdentifierFromProvider,
	"dynatrace_disk_specific_anomalies_v2":  config.IdentifierFromProvider,
	"dynatrace_frequent_issues":             config.IdentifierFromProvider,
	"dynatrace_host_anomalies_v2":           config.IdentifierFromProvider,
	"dynatrace_k8s_cluster_anomalies":       config.IdentifierFromProvider,
	"dynatrace_k8s_namespace_anomalies":     config.IdentifierFromProvider,
	"dynatrace_k8s_node_anomalies":          config.IdentifierFromProvider,
	"dynatrace_k8s_pvc_anomalies":           config.IdentifierFromProvider,
	"dynatrace_k8s_workload_anomalies":      config.IdentifierFromProvider,
	"dynatrace_metric_events":               config.IdentifierFromProvider,
	"dynatrace_mobile_app_anomalies":        config.IdentifierFromProvider,
	"dynatrace_mobile_app_crash_rate":       config.IdentifierFromProvider,
	"dynatrace_service_anomalies_v2":        config.IdentifierFromProvider,
	"dynatrace_vmware_anomalies":            config.IdentifierFromProvider,
	"dynatrace_web_app_anomalies":           config.IdentifierFromProvider,

	// AppEngine

	"dynatrace_db_app_feature_flags":       config.IdentifierFromProvider,
	"dynatrace_discovery_default_rules":    config.IdentifierFromProvider,
	"dynatrace_discovery_feature_flags":    config.IdentifierFromProvider,
	"dynatrace_infraops_app_feature_flags": config.IdentifierFromProvider,
	"dynatrace_infraops_app_settings":      config.IdentifierFromProvider,

	// Application Security

	"dynatrace_appsec_notification":       config.IdentifierFromProvider,
	"dynatrace_attack_alerting":           config.IdentifierFromProvider,
	"dynatrace_attack_allowlist":          config.IdentifierFromProvider,
	"dynatrace_attack_rules":              config.IdentifierFromProvider,
	"dynatrace_attack_settings":           config.IdentifierFromProvider,
	"dynatrace_vulnerability_alerting":    config.IdentifierFromProvider,
	"dynatrace_vulnerability_code":        config.IdentifierFromProvider,
	"dynatrace_vulnerability_settings":    config.IdentifierFromProvider,
	"dynatrace_vulnerability_third_party": config.IdentifierFromProvider,

	// Automation

	"dynatrace_automation_business_calendar":        config.IdentifierFromProvider,
	"dynatrace_automation_scheduling_rule":          config.IdentifierFromProvider,
	"dynatrace_automation_workflow_aws_connections": config.IdentifierFromProvider,
	"dynatrace_automation_workflow":                 config.IdentifierFromProvider,
	"dynatrace_automation_workflow_jira":            config.IdentifierFromProvider,
	"dynatrace_automation_workflow_k8s_connections": config.IdentifierFromProvider,
	"dynatrace_automation_workflow_slack":           config.IdentifierFromProvider,
	"dynatrace_site_reliability_guardian":           config.IdentifierFromProvider,

	// Browser Monitors

	"dynatrace_browser_monitor":             config.IdentifierFromProvider,
	"dynatrace_browser_monitor_outage":      config.IdentifierFromProvider,
	"dynatrace_browser_monitor_performance": config.IdentifierFromProvider,

	// Business Events

	"dynatrace_business_events_buckets":           config.IdentifierFromProvider,
	"dynatrace_business_events_metrics":           config.IdentifierFromProvider,
	"dynatrace_business_events_oneagent":          config.IdentifierFromProvider,
	"dynatrace_business_events_oneagent_outgoing": config.IdentifierFromProvider,
	"dynatrace_business_events_processing":        config.IdentifierFromProvider,
	"dynatrace_business_events_security_context":  config.IdentifierFromProvider,

	// Cloud Platforms

	"dynatrace_cloud_foundry":         config.IdentifierFromProvider,
	"dynatrace_k8s_monitoring":        config.IdentifierFromProvider,
	"dynatrace_kubernetes_app":        config.IdentifierFromProvider,
	"dynatrace_kubernetes":            config.IdentifierFromProvider,
	"dynatrace_kubernetes_enrichment": config.IdentifierFromProvider,

	// Cluster Management

	"dynatrace_managed_backup":           config.IdentifierFromProvider,
	"dynatrace_managed_internet_proxy":   config.IdentifierFromProvider,
	"dynatrace_managed_network_zones":    config.IdentifierFromProvider,
	"dynatrace_managed_public_endpoints": config.IdentifierFromProvider,
	"dynatrace_managed_remote_access":    config.IdentifierFromProvider,
	"dynatrace_managed_smtp":             config.IdentifierFromProvider,

	// Containers

	"dynatrace_container_builtin_rule": config.IdentifierFromProvider,
	"dynatrace_container_registry":     config.IdentifierFromProvider,
	"dynatrace_container_rule":         config.IdentifierFromProvider,
	"dynatrace_container_technology":   config.IdentifierFromProvider,

	// Credentials

	"dynatrace_aws_credentials":   config.IdentifierFromProvider,
	"dynatrace_aws_service":       config.IdentifierFromProvider,
	"dynatrace_azure_credentials": config.IdentifierFromProvider,
	"dynatrace_azure_service":     config.IdentifierFromProvider,
	"dynatrace_credentials":       config.IdentifierFromProvider,

	// Dashboards

	"dynatrace_dashboards_allowlist": config.IdentifierFromProvider,
	"dynatrace_dashboards_general":   config.IdentifierFromProvider,
	"dynatrace_dashboards_presets":   config.IdentifierFromProvider,
	"dynatrace_json_dashboard_base":  config.IdentifierFromProvider,
	"dynatrace_json_dashboard":       config.IdentifierFromProvider,

	// Deprecated

	"dynatrace_alerting_profile":         config.IdentifierFromProvider,
	"dynatrace_application_anomalies":    config.IdentifierFromProvider,
	"dynatrace_autotag":                  config.IdentifierFromProvider,
	"dynatrace_cloudfoundry_credentials": config.IdentifierFromProvider,
	"dynatrace_custom_anomalies":         config.IdentifierFromProvider,
	"dynatrace_dashboard":                config.IdentifierFromProvider,
	"dynatrace_database_anomalies":       config.IdentifierFromProvider,
	"dynatrace_ddu_pool":                 config.IdentifierFromProvider,
	"dynatrace_disk_anomalies":           config.IdentifierFromProvider,
	"dynatrace_host_anomalies":           config.IdentifierFromProvider,
	"dynatrace_k8s_credentials":          config.IdentifierFromProvider,
	"dynatrace_maintenance_window":       config.IdentifierFromProvider,
	"dynatrace_management_zone":          config.IdentifierFromProvider,
	"dynatrace_notification":             config.IdentifierFromProvider,
	"dynatrace_pg_anomalies":             config.IdentifierFromProvider,
	"dynatrace_resource_attributes":      config.IdentifierFromProvider,
	"dynatrace_service_anomalies":        config.IdentifierFromProvider,
	"dynatrace_slo":                      config.IdentifierFromProvider,
	"dynatrace_span_events":              config.IdentifierFromProvider,

	// Developer Observability

	"dynatrace_devobs_agent_optin":  config.IdentifierFromProvider,
	"dynatrace_devobs_data_masking": config.IdentifierFromProvider,
	"dynatrace_devobs_git_onprem":   config.IdentifierFromProvider,

	// Documents

	"dynatrace_direct_shares": config.IdentifierFromProvider,
	"dynatrace_document":      config.IdentifierFromProvider,

	// Environment Settings

	"dynatrace_app_monitoring":             config.IdentifierFromProvider,
	"dynatrace_audit_log":                  config.IdentifierFromProvider,
	"dynatrace_data_privacy":               config.IdentifierFromProvider,
	"dynatrace_disk_options":               config.IdentifierFromProvider,
	"dynatrace_eula_settings":              config.IdentifierFromProvider,
	"dynatrace_hub_permissions":            config.IdentifierFromProvider,
	"dynatrace_hub_subscriptions":          config.IdentifierFromProvider,
	"dynatrace_ip_address_masking":         config.IdentifierFromProvider,
	"dynatrace_limit_outbound_connections": config.IdentifierFromProvider,
	"dynatrace_oneagent_default_mode":      config.IdentifierFromProvider,
	"dynatrace_oneagent_default_version":   config.IdentifierFromProvider,
	"dynatrace_oneagent_features":          config.IdentifierFromProvider,
	"dynatrace_oneagent_side_masking":      config.IdentifierFromProvider,

	// Extensions

	"dynatrace_extension_execution_controller": config.IdentifierFromProvider,
	"dynatrace_extension_execution_remote":     config.IdentifierFromProvider,
	"dynatrace_hub_extension_active_version":   config.IdentifierFromProvider,
	"dynatrace_hub_extension_config":           config.IdentifierFromProvider,

	// Failure Detection

	"dynatrace_failure_detection_parameters": config.IdentifierFromProvider,
	"dynatrace_failure_detection_rules":      config.IdentifierFromProvider,
	"dynatrace_service_failure":              config.IdentifierFromProvider,
	"dynatrace_service_http_failure":         config.IdentifierFromProvider,

	// Host Monitoring

	"dynatrace_aix_extension":                 config.IdentifierFromProvider,
	"dynatrace_crashdump_analytics":           config.IdentifierFromProvider,
	"dynatrace_disk_analytics":                config.IdentifierFromProvider,
	"dynatrace_ebpf_service_discovery":        config.IdentifierFromProvider,
	"dynatrace_host_monitoring_advanced":      config.IdentifierFromProvider,
	"dynatrace_host_monitoring":               config.IdentifierFromProvider,
	"dynatrace_host_monitoring_mode":          config.IdentifierFromProvider,
	"dynatrace_host_process_group_monitoring": config.IdentifierFromProvider,
	"dynatrace_nettracer":                     config.IdentifierFromProvider,
	"dynatrace_network_traffic":               config.IdentifierFromProvider,
	"dynatrace_os_services":                   config.IdentifierFromProvider,

	// HTTP Monitors

	"dynatrace_http_monitor":             config.IdentifierFromProvider,
	"dynatrace_http_monitor_cookies":     config.IdentifierFromProvider,
	"dynatrace_http_monitor_outage":      config.IdentifierFromProvider,
	"dynatrace_http_monitor_performance": config.IdentifierFromProvider,
	"dynatrace_http_monitor_script":      config.IdentifierFromProvider,

	// IAM

	"dynatrace_iam_group":              config.IdentifierFromProvider,
	"dynatrace_iam_permission":         config.IdentifierFromProvider,
	"dynatrace_iam_policy_bindings":    config.IdentifierFromProvider,
	"dynatrace_iam_policy_bindings_v2": config.IdentifierFromProvider,
	"dynatrace_iam_policy":             config.IdentifierFromProvider,
	"dynatrace_iam_user":               config.IdentifierFromProvider,
	"dynatrace_mgmz_permission":        config.IdentifierFromProvider,
	"dynatrace_policy_bindings":        config.IdentifierFromProvider,
	"dynatrace_policy":                 config.IdentifierFromProvider,
	"dynatrace_user":                   config.IdentifierFromProvider,
	"dynatrace_user_group":             config.IdentifierFromProvider,

	// Incubator

	"dynatrace_golden_state": config.IdentifierFromProvider,

	// Integrations

	"dynatrace_issue_tracking":      config.IdentifierFromProvider,
	"dynatrace_remote_environments": config.IdentifierFromProvider,

	// Log Monitoring

	"dynatrace_log_buckets":                config.IdentifierFromProvider,
	"dynatrace_log_custom_attribute":       config.IdentifierFromProvider,
	"dynatrace_log_custom_source":          config.IdentifierFromProvider,
	"dynatrace_log_debug_settings":         config.IdentifierFromProvider,
	"dynatrace_log_events":                 config.IdentifierFromProvider,
	"dynatrace_log_metrics":                config.IdentifierFromProvider,
	"dynatrace_log_oneagent":               config.IdentifierFromProvider,
	"dynatrace_log_processing":             config.IdentifierFromProvider,
	"dynatrace_log_security_context":       config.IdentifierFromProvider,
	"dynatrace_log_sensitive_data_masking": config.IdentifierFromProvider,
	"dynatrace_log_storage":                config.IdentifierFromProvider,
	"dynatrace_log_timestamp":              config.IdentifierFromProvider,

	// Mainframe

	"dynatrace_ibm_mq_filters":                   config.IdentifierFromProvider,
	"dynatrace_ims_bridges":                      config.IdentifierFromProvider,
	"dynatrace_mainframe_transaction_monitoring": config.IdentifierFromProvider,
	"dynatrace_queue_manager":                    config.IdentifierFromProvider,
	"dynatrace_queue_sharing_groups":             config.IdentifierFromProvider,
	"dynatrace_transaction_start_filters":        config.IdentifierFromProvider,

	// Management Zones

	"dynatrace_management_zone_v2": config.IdentifierFromProvider,

	// Metrics

	"dynatrace_custom_units":      config.IdentifierFromProvider,
	"dynatrace_histogram_metrics": config.IdentifierFromProvider,
	"dynatrace_metric_metadata":   config.IdentifierFromProvider,
	"dynatrace_metric_query":      config.IdentifierFromProvider,

	// Mobile & Custom Applications

	"dynatrace_calculated_mobile_metric":   config.IdentifierFromProvider,
	"dynatrace_custom_app_enablement":      config.IdentifierFromProvider,
	"dynatrace_mobile_app_enablement":      config.IdentifierFromProvider,
	"dynatrace_mobile_app_key_performance": config.IdentifierFromProvider,
	"dynatrace_mobile_application":         config.IdentifierFromProvider,
	"dynatrace_mobile_app_request_errors":  config.IdentifierFromProvider,

	// Monitored Entities

	"dynatrace_custom_device": config.IdentifierFromProvider,

	// Monitored Technologies

	"dynatrace_monitored_technologies_apache":      config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_dotnet":      config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_go":          config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_iis":         config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_java":        config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_nginx":       config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_nodejs":      config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_opentracing": config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_php":         config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_varnish":     config.IdentifierFromProvider,
	"dynatrace_monitored_technologies_wsmb":        config.IdentifierFromProvider,

	// Network Availability Monitors

	"dynatrace_network_monitor":        config.IdentifierFromProvider,
	"dynatrace_network_monitor_outage": config.IdentifierFromProvider,

	// Notifications

	"dynatrace_ansible_tower_notification": config.IdentifierFromProvider,
	"dynatrace_email_notification":         config.IdentifierFromProvider,
	"dynatrace_jira_notification":          config.IdentifierFromProvider,
	"dynatrace_mobile_notifications":       config.IdentifierFromProvider,
	"dynatrace_ops_genie_notification":     config.IdentifierFromProvider,
	"dynatrace_pager_duty_notification":    config.IdentifierFromProvider,
	"dynatrace_service_now_notification":   config.IdentifierFromProvider,
	"dynatrace_slack_notification":         config.IdentifierFromProvider,
	"dynatrace_trello_notification":        config.IdentifierFromProvider,
	"dynatrace_victor_ops_notification":    config.IdentifierFromProvider,
	"dynatrace_webhook_notification":       config.IdentifierFromProvider,
	"dynatrace_xmatters_notification":      config.IdentifierFromProvider,

	// OpenTelemetry & OpenTracing

	"dynatrace_attribute_allow_list":     config.IdentifierFromProvider,
	"dynatrace_attribute_block_list":     config.IdentifierFromProvider,
	"dynatrace_attribute_masking":        config.IdentifierFromProvider,
	"dynatrace_attributes_preferences":   config.IdentifierFromProvider,
	"dynatrace_opentelemetry_metrics":    config.IdentifierFromProvider,
	"dynatrace_span_capture_rule":        config.IdentifierFromProvider,
	"dynatrace_span_context_propagation": config.IdentifierFromProvider,
	"dynatrace_span_entry_point":         config.IdentifierFromProvider,

	// Ownership

	"dynatrace_ownership_config": config.IdentifierFromProvider,
	"dynatrace_ownership_teams":  config.IdentifierFromProvider,

	// Platform

	"dynatrace_generic_setting":         config.IdentifierFromProvider,
	"dynatrace_grail_metrics_allowall":  config.IdentifierFromProvider,
	"dynatrace_grail_metrics_allowlist": config.IdentifierFromProvider,
	"dynatrace_platform_bucket":         config.IdentifierFromProvider,

	// Process Group Monitoring

	"dynatrace_cloudapp_workloaddetection":     config.IdentifierFromProvider,
	"dynatrace_declarative_grouping":           config.IdentifierFromProvider,
	"dynatrace_pg_alerting":                    config.IdentifierFromProvider,
	"dynatrace_process_group_detection":        config.IdentifierFromProvider,
	"dynatrace_process_group_detection_flags":  config.IdentifierFromProvider,
	"dynatrace_process_group_monitoring":       config.IdentifierFromProvider,
	"dynatrace_processgroup_naming":            config.IdentifierFromProvider,
	"dynatrace_process_group_simple_detection": config.IdentifierFromProvider,

	// Process Monitoring

	"dynatrace_builtin_process_monitoring": config.IdentifierFromProvider,
	"dynatrace_process_availability":       config.IdentifierFromProvider,
	"dynatrace_process_monitoring":         config.IdentifierFromProvider,
	"dynatrace_process_monitoring_rule":    config.IdentifierFromProvider,
	"dynatrace_process_visibility":         config.IdentifierFromProvider,

	// Real User Monitoring

	"dynatrace_geolocation":              config.IdentifierFromProvider,
	"dynatrace_process_group_rum":        config.IdentifierFromProvider,
	"dynatrace_rum_advanced_correlation": config.IdentifierFromProvider,
	"dynatrace_rum_host_headers":         config.IdentifierFromProvider,
	"dynatrace_rum_ip_determination":     config.IdentifierFromProvider,
	"dynatrace_rum_ip_locations":         config.IdentifierFromProvider,
	"dynatrace_rum_overload_prevention":  config.IdentifierFromProvider,
	"dynatrace_rum_provider_breakdown":   config.IdentifierFromProvider,
	"dynatrace_usability_analytics":      config.IdentifierFromProvider,
	"dynatrace_user_action_metrics":      config.IdentifierFromProvider,
	"dynatrace_user_experience_score":    config.IdentifierFromProvider,
	"dynatrace_user_session_metrics":     config.IdentifierFromProvider,

	// Service Detection

	"dynatrace_custom_service":               config.IdentifierFromProvider,
	"dynatrace_service_external_web_request": config.IdentifierFromProvider,
	"dynatrace_service_external_web_service": config.IdentifierFromProvider,
	"dynatrace_service_full_web_request":     config.IdentifierFromProvider,
	"dynatrace_service_full_web_service":     config.IdentifierFromProvider,
	"dynatrace_unified_services_metrics":     config.IdentifierFromProvider,
	"dynatrace_unified_services_opentel":     config.IdentifierFromProvider,

	// Service-level Objective

	"dynatrace_slo_normalization": config.IdentifierFromProvider,
	"dynatrace_slo_v2":            config.IdentifierFromProvider,

	// Service Monitoring

	"dynatrace_api_detection":             config.IdentifierFromProvider,
	"dynatrace_calculated_service_metric": config.IdentifierFromProvider,
	"dynatrace_muted_requests":            config.IdentifierFromProvider,
	"dynatrace_request_attribute":         config.IdentifierFromProvider,
	"dynatrace_service_naming":            config.IdentifierFromProvider,
	"dynatrace_url_based_sampling":        config.IdentifierFromProvider,

	// Service Settings

	"dynatrace_davis_copilot": config.IdentifierFromProvider,

	// Session Replay

	"dynatrace_session_replay_resource_capture": config.IdentifierFromProvider,
	"dynatrace_session_replay_web_privacy":      config.IdentifierFromProvider,

	// Synthetic

	"dynatrace_calculated_synthetic_metric": config.IdentifierFromProvider,
	"dynatrace_synthetic_availability":      config.IdentifierFromProvider,
	"dynatrace_synthetic_location":          config.IdentifierFromProvider,

	// Tags

	"dynatrace_autotag_rules": config.IdentifierFromProvider,
	"dynatrace_autotag_v2":    config.IdentifierFromProvider,
	"dynatrace_custom_tags":   config.IdentifierFromProvider,

	// Topology Model

	"dynatrace_generic_relationships":  config.IdentifierFromProvider,
	"dynatrace_generic_types":          config.IdentifierFromProvider,
	"dynatrace_grail_security_context": config.IdentifierFromProvider,

	// Updates

	"dynatrace_activegate_updates": config.IdentifierFromProvider,
	"dynatrace_oneagent_updates":   config.IdentifierFromProvider,
	"dynatrace_update_windows":     config.IdentifierFromProvider,

	// User Settings

	"dynatrace_user_settings": config.IdentifierFromProvider,

	// Virtualization

	"dynatrace_vmware": config.IdentifierFromProvider,

	// Web Applications

	"dynatrace_application_detection_rule_v2":    config.IdentifierFromProvider,
	"dynatrace_calculated_web_metric":            config.IdentifierFromProvider,
	"dynatrace_key_user_action":                  config.IdentifierFromProvider,
	"dynatrace_web_app_beacon_endpoint":          config.IdentifierFromProvider,
	"dynatrace_web_app_beacon_origins":           config.IdentifierFromProvider,
	"dynatrace_web_app_custom_config_properties": config.IdentifierFromProvider,
	"dynatrace_web_app_custom_errors":            config.IdentifierFromProvider,
	"dynatrace_web_app_custom_injection":         config.IdentifierFromProvider,
	"dynatrace_web_app_enablement":               config.IdentifierFromProvider,
	"dynatrace_web_app_injection_cookie":         config.IdentifierFromProvider,
	"dynatrace_web_app_javascript_updates":       config.IdentifierFromProvider,
	"dynatrace_web_app_javascript_version":       config.IdentifierFromProvider,
	"dynatrace_web_app_key_performance_custom":   config.IdentifierFromProvider,
	"dynatrace_web_app_key_performance_load":     config.IdentifierFromProvider,
	"dynatrace_web_app_key_performance_xhr":      config.IdentifierFromProvider,
	"dynatrace_web_application":                  config.IdentifierFromProvider,
	"dynatrace_web_app_request_errors":           config.IdentifierFromProvider,
	"dynatrace_web_app_resource_cleanup":         config.IdentifierFromProvider,
	"dynatrace_web_app_resource_types":           config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
