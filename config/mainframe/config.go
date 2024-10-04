package mainframe

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_ibm_mq_filters",
		"dynatrace_queue_sharing_groups",
		"dynatrace_queue_manager",
		"dynatrace_mainframe_transaction_monitoring",
		"dynatrace_transaction_start_filters",
		"dynatrace_ims_bridges",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "mainframe"
		})
	}
}
