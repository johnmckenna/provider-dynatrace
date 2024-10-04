package usersettings

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_user_settings",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "usersettings"
		})
	}
}
