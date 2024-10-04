package extensions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_hub_extension_config",
		"dynatrace_extension_execution_remote",
		"dynatrace_hub_extension_active_version",
		"dynatrace_extension_execution_controller",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "extensions"
		})
	}
}
