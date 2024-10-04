package accesstokens

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_activegate_token",
		"dynatrace_ag_token",
		"dynatrace_token_settings",
		"dynatrace_api_token",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "accesstokens"
		})
	}
}