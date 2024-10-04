package logmonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_log_storage",
		"dynatrace_log_security_context",
		"dynatrace_log_debug_settings",
		"dynatrace_log_events",
		"dynatrace_log_sensitive_data_masking",
		"dynatrace_log_sensitive_data_masking",
		"dynatrace_log_custom_attribute",
		"dynatrace_log_buckets",
		"dynatrace_log_custom_source",
		"dynatrace_log_processing",
		"dynatrace_log_metrics",
		"dynatrace_log_oneagent",
		"dynatrace_log_timestamp",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "logmonitoring"
		})
	}
}
