package httpmonitors

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_http_monitor",
		"dynatrace_http_monitor_outage",
		"dynatrace_http_monitor_cookies",
		"dynatrace_http_monitor_performance",
		"dynatrace_http_monitor_script",
		"dynatrace_http_monitor",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "httpmonitors"
		})
	}
}
