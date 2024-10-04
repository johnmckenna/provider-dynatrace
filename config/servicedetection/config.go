package servicedetection

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_unified_services_opentel",
		"dynatrace_management_zone_v2",
		"dynatrace_service_external_web_service",
		"dynatrace_management_zone_v2",
		"dynatrace_service_full_web_service",
		"dynatrace_management_zone_v2",
		"dynatrace_service_external_web_request",
		"dynatrace_unified_services_metrics",
		"dynatrace_management_zone_v2",
		"dynatrace_service_full_web_request",
		"dynatrace_custom_service",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "servicedetection"
		})
	}
}
