package platform

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_grail_metrics_allowlist",
		"dynatrace_generic_setting",
		"dynatrace_grail_metrics_allowall",
		"dynatrace_platform_bucket",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "platform"
		})
	}
}
