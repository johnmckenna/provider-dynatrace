package appengine

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_infraops_app_feature_flags",
		"dynatrace_discovery_feature_flags",
		"dynatrace_infraops_app_settings",
		"dynatrace_db_app_feature_flags",
		"dynatrace_discovery_default_rules",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "appengine"
		})
	}
}
