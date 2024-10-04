package synthetic

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_calculated_synthetic_metric",
		"dynatrace_synthetic_location",
		"dynatrace_synthetic_availability",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "synthetic"
		})
	}
}
