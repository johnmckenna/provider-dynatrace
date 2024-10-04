package monitoredtechnologies

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_monitored_technologies_opentracing",
		"dynatrace_monitored_technologies_php",
		"dynatrace_monitored_technologies_go",
		"dynatrace_monitored_technologies_wsmb",
		"dynatrace_monitored_technologies_nodejs",
		"dynatrace_monitored_technologies_varnish",
		"dynatrace_monitored_technologies_iis",
		"dynatrace_monitored_technologies_nginx",
		"dynatrace_monitored_technologies_apache",
		"dynatrace_monitored_technologies_java",
		"dynatrace_monitored_technologies_dotnet",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "monitoredtechnologies"
		})
	}
}
