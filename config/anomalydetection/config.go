package anomalydetection

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_service_anomalies_v2",
		"dynatrace_davis_anomaly_detectors",
		"dynatrace_metric_events",
		"dynatrace_disk_specific_anomalies_v2",
		"dynatrace_k8s_cluster_anomalies",
		"dynatrace_disk_anomalies_v2",
		"dynatrace_mobile_app_anomalies",
		"dynatrace_custom_app_crash_rate",
		"dynatrace_frequent_issues",
		"dynatrace_aws_anomalies",
		"dynatrace_web_app_anomalies",
		"dynatrace_k8s_workload_anomalies",
		"dynatrace_k8s_namespace_anomalies",
		"dynatrace_custom_app_anomalies",
		"dynatrace_host_anomalies_v2",
		"dynatrace_disk_anomaly_rules",
		"dynatrace_k8s_node_anomalies",
		"dynatrace_k8s_pvc_anomalies",
		"dynatrace_mobile_app_crash_rate",
		"dynatrace_vmware_anomalies",
		"dynatrace_disk_edge_anomaly_detectors",
		"dynatrace_database_anomalies_v2",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "anomalydetection"
		})
	}
}
