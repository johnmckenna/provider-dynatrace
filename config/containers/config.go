package containers

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_container_builtin_rule",
		"dynatrace_container_technology",
		"dynatrace_container_rule",
		"dynatrace_container_registry",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "containers"
		})
	}
}
