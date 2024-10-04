package cloudplatforms

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_k8s_monitoring",
		"dynatrace_kubernetes_enrichment",
		"dynatrace_kubernetes",
		"dynatrace_cloud_foundry",
		"dynatrace_kubernetes_app",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "cloudplatforms"
		})
	}
}
