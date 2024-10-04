package processmonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_process_monitoring",
		"dynatrace_process_availability",
		"dynatrace_process_visibility",
		"dynatrace_builtin_process_monitoring",
		"dynatrace_process_monitoring_rule",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "processmonitoring"
		})
	}
}
