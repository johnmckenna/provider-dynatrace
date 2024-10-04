package tags

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_custom_tags",
		"dynatrace_autotag_v2",
		"dynatrace_autotag_v2",
		"dynatrace_autotag_v2",
		"dynatrace_autotag_rules",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "tags"
		})
	}
}
