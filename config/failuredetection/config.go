package failuredetection

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_failure_detection_rules",
		"dynatrace_failure_detection_parameters",
		"dynatrace_service_http_failure",
		"dynatrace_failure_detection_parameters",
		"dynatrace_service_failure",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "failuredetection"
		})
	}
}
