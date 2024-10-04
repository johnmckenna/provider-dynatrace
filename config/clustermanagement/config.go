package clustermanagement

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_managed_public_endpoints",
		"dynatrace_managed_remote_access",
		"dynatrace_managed_network_zones",
		"dynatrace_managed_backup",
		"dynatrace_managed_smtp",
		"dynatrace_managed_internet_proxy",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "clustermanagement"
		})
	}
}
