package browsermonitors

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_browser_monitor_performance",
		"dynatrace_browser_monitor_outage",
		"dynatrace_browser_monitor",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "browsermonitors"
		})
	}
}
