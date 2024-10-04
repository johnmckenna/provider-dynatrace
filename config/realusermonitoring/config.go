package realusermonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_user_experience_score",
		"dynatrace_user_session_metrics",
		"dynatrace_geolocation",
		"dynatrace_process_group_rum",
		"dynatrace_rum_ip_determination",
		"dynatrace_rum_ip_locations",
		"dynatrace_user_action_metrics",
		"dynatrace_rum_host_headers",
		"dynatrace_usability_analytics",
		"dynatrace_usability_analytics",
		"dynatrace_rum_overload_prevention",
		"dynatrace_rum_provider_breakdown",
		"dynatrace_rum_advanced_correlation",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "realusermonitoring"
		})
	}
}
