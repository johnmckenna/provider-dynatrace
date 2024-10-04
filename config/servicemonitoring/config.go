package servicemonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_url_based_sampling",
		"dynatrace_service_naming",
		"dynatrace_api_detection",
		"dynatrace_request_attribute",
		"dynatrace_muted_requests",
		"dynatrace_calculated_service_metric",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "servicemonitoring"
		})
	}
}
