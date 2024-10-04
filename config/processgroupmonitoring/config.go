package processgroupmonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	resources := []string{
		"dynatrace_process_group_detection",
		"dynatrace_pg_alerting",
		"dynatrace_cloudapp_workloaddetection",
		"dynatrace_process_group_monitoring",
		"dynatrace_process_group_simple_detection",
		"dynatrace_declarative_grouping",
		"dynatrace_processgroup_naming",
		"dynatrace_process_group_detection_flags",
	}

	for _, resource := range resources {
		p.AddResourceConfigurator(resource, func(r *config.Resource) {
			r.ShortGroup = "processgroupmonitoring"
		})
	}
}
