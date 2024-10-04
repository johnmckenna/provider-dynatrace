package metrics

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_histogram_metrics",
		"dynatrace_custom_units",
		"dynatrace_metric_query",
		"dynatrace_metric_metadata",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "metrics"
		})
	}
}
