package mobilecustomapplications

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_mobile_app_key_performance",
		"dynatrace_custom_app_enablement",
		"dynatrace_calculated_mobile_metric",
		"dynatrace_mobile_app_request_errors",
		"dynatrace_mobile_app_enablement",
		"dynatrace_mobile_application",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "mobilecustomapplications"
		})
	}
}
