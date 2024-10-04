package businessevents

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_business_events_metrics",
		"dynatrace_business_events_oneagent_outgoing",
		"dynatrace_business_events_buckets",
		"dynatrace_business_events_security_context",
		"dynatrace_business_events_oneagent",
		"dynatrace_business_events_processing",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "businessevents"
		})
	}
}
