package credentials

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_aws_credentials",
		"dynatrace_aws_credentials",
		"dynatrace_aws_service",
		"dynatrace_credentials",
		"dynatrace_azure_credentials",
		"dynatrace_azure_service",
		"dynatrace_azure_credentials",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "credentials"
		})
	}
}
