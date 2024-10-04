package deprecated

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_ddu_pool",
		"dynatrace_pg_anomalies",
		"dynatrace_cloudfoundry_credentials",
		"dynatrace_alerting_profile",
		"dynatrace_autotag",
		"dynatrace_dashboard",
		"dynatrace_maintenance_window",
		"dynatrace_database_anomalies",
		"dynatrace_application_anomalies",
		"dynatrace_host_anomalies",
		"dynatrace_management_zone",
		"dynatrace_notification",
		"dynatrace_alerting_profile",
		"dynatrace_custom_anomalies",
		"dynatrace_slo",
		"dynatrace_resource_attributes",
		"dynatrace_service_anomalies",
		"dynatrace_k8s_credentials",
		"dynatrace_span_events",
		"dynatrace_disk_anomalies",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "deprecated"
		})
	}
}
