package dashboards

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard_base",
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard_base",
		"dynatrace_json_dashboard",
		"dynatrace_json_dashboard_base",
		"dynatrace_json_dashboard",
		"dynatrace_dashboards_allowlist",
		"dynatrace_dashboards_general",
		"dynatrace_dashboards_presets",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "dashboards"
		})
	}
}
