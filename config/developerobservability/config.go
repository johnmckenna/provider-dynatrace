package developerobservability

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_devobs_agent_optin",
		"dynatrace_devobs_git_onprem",
		"dynatrace_devobs_data_masking",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "developerobservability"
		})
	}
}
