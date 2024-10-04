package opentelemetryopentracing

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_opentelemetry_metrics",
		"dynatrace_attribute_allow_list",
		"dynatrace_attributes_preferences",
		"dynatrace_attribute_block_list",
		"dynatrace_attribute_masking",
		"dynatrace_span_context_propagation",
		"dynatrace_span_entry_point",
		"dynatrace_span_capture_rule",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "opentelemetryopentracing"
		})
	}
}
